package cmd

import (
	"errors"
	"log"

	"github.com/georgepadayatti/jwstar/go/jwk"
	"github.com/georgepadayatti/jwstar/go/jws"
	"github.com/spf13/cobra"
)

func generateJWSCmd() *cobra.Command {

	var claimsString string

	var generateJWSCmd = &cobra.Command{
		Use:   "generate-jws",
		Short: "Generate JWS",
		Args:  cobra.ExactArgs(0),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if claimsString == "" {
				return errors.New("--claims is required")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			var jwkObj jwk.JWK
			jwsObj := jws.JWS{Key: jwkObj, Claims: claimsString}
			err := jwsObj.Generate()
			if err != nil {
				log.Fatalf("Error occured while generating JWS: %v", err)
			}
		},
	}

	generateJWSCmd.Flags().StringVarP(&claimsString, "claims", "c", "", "JSON representation of claims")
	generateJWSCmd.MarkPersistentFlagRequired("claims")

	return generateJWSCmd
}
