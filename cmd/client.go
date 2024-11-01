package cmd

import (
	"github.com/ComradeCaro/SimpleChat/client"
	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Run the SimpleChat client",
	Run: func(cmd *cobra.Command, args []string) {
		client.Run()
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
