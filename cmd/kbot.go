/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v3"
)

var (
	TeleToken  = os.Getenv("token")
	apiVersion = "v1.0.0"
)

// kbotCmd represents the kbot command
var kbotCmd = &cobra.Command{
	Use:     "kbot",
	Aliases: []string{"start"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("kbot %s has started", apiVersion)
		if TeleToken == "" {
			panic("Teletoken is empty")
		}

		kbot, err := telebot.NewBot(
			telebot.Settings{
				URL:    "",
				Token:  TeleToken,
				Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
			},
		)

		commands := []telebot.Command{
			{Text: "/help", Description: "List available commands"},
			{Text: "/hello", Description: "Say hello"},
		}
		kbot.SetCommands(commands)

		if err != nil {
			log.Fatalf("Please check TeleToken env variable. %s", err)
			return
		}

		kbot.Handle(telebot.OnText, func(c telebot.Context) error {
			text := c.Text()
			log.Printf("payload %s, text %s\n", c.Message().Payload, text)
			switch text {

			case "/hello":
				return c.Send(fmt.Sprintf("Hello I am kbot %s", apiVersion))
			case "/help":
				res := ""
				for _, c := range commands {
					res += fmt.Sprintf("%s - %s\n", c.Text, c.Description)
				}
				return c.Send(res)
			case "/time":
				return c.Send(time.Now().Format("2006-01-02T15:04:05"))
			case "/timestamp":
				return c.Send("timestamp ~~~~: " + strconv.Itoa(time.Now().Second()))
			default:
				log.Println("unknown case: ", text)
				return c.Send(fmt.Sprintf("Unknown command %s", text))
			}
		})


		kbot.Start()
	},
}

func init() {
	rootCmd.AddCommand(kbotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kbotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kbotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
