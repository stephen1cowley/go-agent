package server

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/sashabaranov/go-openai"
)

var editAppJSResp ArgsAppJS
var editAppCSSResp ArgsAppCSS
var newFileResp ArgsCreateFile
var libsResp ArgsLibraries

func AppJSTool() {

	// Start off by cleaning the React App source code
	cmd := exec.Command("shell_script/onStartup.sh")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(output))

	apiKey := os.Getenv("OPEN_AI_API_KEY")
	client := openai.NewClient(apiKey)
	messages := make([]openai.ChatCompletionMessage, 0)

	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: "You are a helpful software engineer. Currently we are working on a fresh React App boilerplate. You are able to change App.js and App.css. You are able to create new JavaScript files to assist you in creating the application, ensure these are correctly imported into App.js. You are able to import external libraries which you should utilise for advanced app functionality.",
	})

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Conversation")
	fmt.Println("---------------------")
	myTools := []openai.Tool{AppJSEdit, AppCSSEdit, NewJsonFile, ImportLibraries}

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

		// jsonData, err := json.MarshalIndent(resp, "", "    ")
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		// fmt.Println(string(jsonData))

		content := resp.Choices[0].Message.Content
		fmt.Println(content)

		tool_calls := resp.Choices[0].Message.ToolCalls
		if len(tool_calls) != 0 {
			fmt.Println("Now making any tool calls ...")
		}

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
			case "new_js_file_func":
				fmt.Println("Creating new JS file ...")
				// fmt.Println(val.Function.Arguments)
				json.Unmarshal([]byte(val.Function.Arguments), &newFileResp)
				CreateJSFile(
					newFileResp,
				)
			case "libraries_func":
				fmt.Println("Importing libraries ...")
				// fmt.Println(val.Function.Arguments)
				json.Unmarshal([]byte(val.Function.Arguments), &libsResp)
				InstallLibraries(
					libsResp,
				)
			}
		}
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: content,
		})
	}
}
