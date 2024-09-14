package server

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/sashabaranov/go-openai"
)

var editAppJSResp ArgsAppJS
var editAppCSSResp ArgsAppCSS

func AppJSTool() {
	apiKey := os.Getenv("OPEN_AI_API_KEY")
	client := openai.NewClient(apiKey)
	messages := make([]openai.ChatCompletionMessage, 0)

	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: "You are a helpful software engineer. Currently we are working on a fresh React App boilerplate. You are able to change App.js and App.css only, and you are not able to import any external libraries.",
	})

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Conversation")
	fmt.Println("---------------------")
	myTools := []openai.Tool{AppJSEdit, AppCSSEdit}

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: text,
		})

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

		resp, err := client.CreateChatCompletion(
			ctx,
			openai.ChatCompletionRequest{
				Model:       openai.GPT4o,
				Messages:    messages,
				Tools:       myTools,
				Temperature: 0.0,
				// ToolChoice: "required",
			},
		)

		cancel()

		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			continue
		}

		// jsonData, err := json.MarshalIndent(resp, "", "    ")
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		// fmt.Println(string(jsonData))

		content := resp.Choices[0].Message.Content
		tool_calls := resp.Choices[0].Message.ToolCalls

		fmt.Println(content)

		// test

		if len(tool_calls) != 0 {
			fmt.Println("Now making any tool calls ...")
		}

		if len(messages) >= 6 {
			messages = messages[1:]
		}

		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: content,
		})

		fmt.Println(messages)

		for _, val := range tool_calls {
			switch val.Function.Name {
			case "app_js_edit_func":
				fmt.Println("Updating App.js ...")
				// fmt.Println(val.Function.Arguments)
				json.Unmarshal([]byte(val.Function.Arguments), &editAppJSResp)
				EditAppJS(
					editAppJSResp.AppJSCode,
				)
			case "app_css_edit_func":
				fmt.Println("Updating App.css ...")
				// fmt.Println(val.Function.Arguments)
				json.Unmarshal([]byte(val.Function.Arguments), &editAppCSSResp)
				EditAppCSS(
					editAppCSSResp.AppCSSCode,
				)
			}
		}

	}
}
