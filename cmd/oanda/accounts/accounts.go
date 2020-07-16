package accounts

import (
	"fmt"
	"log"

	"github.com/hsmtkk/oanda_api_go/cmd/oanda/util"
	"github.com/hsmtkk/oanda_api_go/pkg/accounts/detail"
	"github.com/hsmtkk/oanda_api_go/pkg/accounts/list"
	"github.com/hsmtkk/oanda_api_go/pkg/env"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use: "accounts",
}

func init() {
	detailCommand := &cobra.Command{
		Use:  "detail",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			accountID := args[0]
			token, err := util.GetToken()
			if err != nil {
				log.Fatal(err)
			}
			logger := util.GetLogger()
			getter, err := detail.New(env.FxTradePractice, token, logger)
			if err != nil {
				log.Fatal(err)
			}
			detail, err := getter.Detail(accountID)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(detail)
		},
	}
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
	Command.AddCommand(detailCommand)
	Command.AddCommand(listCommand)
}
