package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/spf13/cobra"
)

func rootCommands() *cobra.Command {
	apiKey := ""

	rootCmd := &cobra.Command{
		Use:   "chatgpt",
		Short: "Chat with ChatGPT in console.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return initializeConfig(cmd)
		},
		Run: func(cmd *cobra.Command, args []string) {

			ctx := context.Background()
			client := gpt3.NewClient(apiKey)

			scanner := bufio.NewScanner(os.Stdin)
			quit := false

			currentUser, err := user.Current()
			if err != nil {
				log.Fatalf(err.Error())
			}

			username := currentUser.Name

			for !quit {
				fmt.Printf("Hello %s, how can I help you today? (type quit to exit): ", username)

				if !scanner.Scan() {
					break
				}

				question := scanner.Text()
				switch question {
				case "quit":
					quit = true

				default:
					GetResponse(client, ctx, question)
				}
			}
		},
	}
	rootCmd.Flags().StringVarP(&apiKey, "apiKey", "k", "NONE", "This is the OpenAI ChatGPT AI Key available at 'https://beta.openai.com/account/api-keys'")

	return rootCmd
}
