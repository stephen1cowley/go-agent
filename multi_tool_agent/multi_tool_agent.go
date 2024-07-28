package main

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

func main() {
	apiKey := os.Getenv("OPEN_AI_API_KEY")
	client := openai.NewClient(apiKey)
	messages := make([]openai.ChatCompletionMessage, 0)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Conversation")
	fmt.Println("---------------------")
	fmt.Println(Xk)
	myTools := []openai.Tool{MyTool1, MyTool2}

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: text,
		})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		resp, err := client.CreateChatCompletion(
			ctx,
			openai.ChatCompletionRequest{
				Model:    openai.GPT3Dot5Turbo,
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
		toolCall := resp.Choices[0].Message.ToolCalls
		fmt.Println(toolCall)
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: content,
		})
		fmt.Println(content)
	}
}
