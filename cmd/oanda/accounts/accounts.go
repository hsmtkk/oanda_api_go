package accounts

import (
	"fmt"
	"log"

	"github.com/hsmtkk/oanda_api_go/cmd/oanda/util"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use: "accounts",
	Run: func(cmd *cobra.Command, args []string) {
		executor, err := util.NewExecutor()
		if err != nil {
			log.Fatal(err)
		}
		insts, err := executor.Accounts()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(insts)
	},
}
