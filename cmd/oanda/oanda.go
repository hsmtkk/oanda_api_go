package main

import (
	"log"

	"github.com/hsmtkk/oanda_api_go/cmd/oanda/accounts"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	command := &cobra.Command{
		Use: "oanda",
	}
	command.AddCommand(accounts.Command)
	command.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
	viper.BindPFlag("verbose", command.PersistentFlags().Lookup("verbose"))
	if err := command.Execute(); err != nil {
		log.Fatal(err)
	}
}
