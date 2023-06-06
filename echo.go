package main

import (
	"EchoServer/networkListener"
	"EchoServer/networkSender"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

var Verbose bool
var NewLine bool

var rootCmd = &cobra.Command{
	Use:   "echo",
	Short: "An echo command.",
	Long:  "A basic echo command; however, it can also echo to an arbitrary IP and port.",
	Run: func(cmd *cobra.Command, args []string) {
		if Verbose {
			log.Println("Running the root command.")
		}

		fmt.Printf(strings.Join(cmd.Flags().Args(), " "))

		if NewLine == true {
			fmt.Printf("\n")
		}
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Set this flag to have more verbose output.")
}

func main() {
	if Verbose {
		log.Println("Starting up echo command...")
	}
	rootCmd.AddCommand(networkSender.NetworkEcho, networkListener.Listener)
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
