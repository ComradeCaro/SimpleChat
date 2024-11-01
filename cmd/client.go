package cmd

import (
	"github.com/ComradeCaro/SimpleChat/client"
	"github.com/spf13/cobra"
)

var address string

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Run the SimpleChat client",
	Run: func(cmd *cobra.Command, args []string) {
		client.Run(address)
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
	clientCmd.Flags().StringVarP(&address, "address", "a", "", "Address client will connect to. Format is ip:port")
}
