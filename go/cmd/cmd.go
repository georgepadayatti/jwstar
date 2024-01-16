package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "jwstar",
	Short: "Create and verify JWK, JWS",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(generateJWSCmd())
	rootCmd.AddCommand(verifyJWSCmd())
}
