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

var editHtmlResp ArgsHtml

func HtmlTool() {
	apiKey := os.Getenv("OPEN_AI_API_KEY")
	client := openai.NewClient(apiKey)
	messages := make([]openai.ChatCompletionMessage, 0)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Conversation")
	fmt.Println("---------------------")
	myTools := []openai.Tool{HtmlEdit, RunTheServer}

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
				Model:    openai.GPT4o,
				Messages: messages,
				Tools:    myTools,
				// ToolChoice: "required",
			},
		)

		cancel()

		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			continue
		}

		jsonData, err := json.MarshalIndent(resp, "", "    ")
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(jsonData))
		content := resp.Choices[0].Message.Content
		tool_calls := resp.Choices[0].Message.ToolCalls

		for _, val := range tool_calls {
			switch val.Function.Name {
			case "html_edit_func":
				fmt.Println("Editting the HTML code...")
				fmt.Println(val.Function.Arguments)
				json.Unmarshal([]byte(val.Function.Arguments), &editHtmlResp)
				EditWebsite(
					editHtmlResp.HtmlCode,
				)
			case "run_server_func":
				fmt.Println("Now running the server...")
				RunServer()
			}
		}

		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: content,
		})
		fmt.Println(content)
	}
}
