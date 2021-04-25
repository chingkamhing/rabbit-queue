package main

import (
	"log"

	"github.com/spf13/cobra"
)

// service root cli command settings
var rootCmd = &cobra.Command{
	Use:   "",
	Short: "RabbitMQ simple queue",
	Run: func(cmd *cobra.Command, args []string) {
		// default command: print usage
		cmd.Usage()
	},
}

// send cli command settings
var sendCmd = &cobra.Command{
	Use:   "send [message]",
	Short: "Send RabbitMQ message",
	Args:  cobra.ExactArgs(1),
	Run:   runSend,
}

// receive cli command settings
var receiveCmd = &cobra.Command{
	Use:   "receive",
	Short: "Receive RabbitMQ message",
	Args:  cobra.ExactArgs(0),
	Run:   runReceive,
}

var scheme string
var host string
var port int
var username string
var password string

const queueName = "hello"
const consumerName = ""

func init() {
	rootCmd.PersistentFlags().StringVar(&scheme, "scheme", "amqp", "AMQP scheme: amqp, amqps")
	rootCmd.PersistentFlags().StringVar(&host, "host", "localhost", "AMQP host name")
	rootCmd.PersistentFlags().IntVar(&port, "port", 5672, "AMQP port number")
	rootCmd.PersistentFlags().StringVar(&username, "username", "", "AMQP username")
	rootCmd.PersistentFlags().StringVar(&password, "password", "", "AMQP password")
	rootCmd.AddCommand(sendCmd)
	rootCmd.AddCommand(receiveCmd)
}

func failOnError(err error, msg string) {
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		log.Println(err)
	}
}
