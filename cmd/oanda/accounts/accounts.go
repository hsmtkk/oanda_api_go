package accounts

import (
	"fmt"
	"log"

	"github.com/hsmtkk/oanda_api_go/cmd/oanda/util"
	"github.com/hsmtkk/oanda_api_go/pkg/accounts/list"
	"github.com/hsmtkk/oanda_api_go/pkg/env"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use: "accounts",
}

func init() {
	listCommand := &cobra.Command{
		Use: "list",
		Run: func(cmd *cobra.Command, args []string) {
			token, err := util.GetToken()
			if err != nil {
				log.Fatal(err)
			}
			logger := util.GetLogger()
			getter, err := list.New(env.FxTradePractice, token, logger)
			if err != nil {
				log.Fatal(err)
			}
			ids, err := getter.List()
			if err != nil {
				log.Fatal(err)
			}
			for _, id := range ids {
				fmt.Println(id)
			}
		},
	}
	Command.AddCommand(listCommand)
}
