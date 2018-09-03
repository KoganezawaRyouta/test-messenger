package cmd

import (
	"github.com/spf13/cobra"

)

var RootCmd = &cobra.Command{
	Use:   "root",
	Short: "hey",
	Long:  "hey",
	Run: func(cmd *cobra.Command, args []string) {
		// ...
	},
}

func init() {
	cobra.OnInitialize()
}
