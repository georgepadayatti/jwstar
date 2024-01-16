package cmd

import (
	"errors"
	"log"

	"github.com/georgepadayatti/jwstar/go/jwk"
	"github.com/georgepadayatti/jwstar/go/jws"
	"github.com/spf13/cobra"
)

func verifyJWSCmd() *cobra.Command {

	var jwkString string
	var jwsString string

	var verifyJWSCmd = &cobra.Command{
		Use:   "verify-jws",
		Short: "Verify JWS",
		Args:  cobra.ExactArgs(0),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if jwkString == "" || jwsString == "" {
				return errors.New("both --jwk and --jws is required")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			jwsObj := jws.JWS{Key: jwk.FromJSON(jwkString), Signature: jwsString}
			err := jwsObj.Verify()
			if err != nil {
				log.Fatalf("Error occured while verifying JWS: %v", err)
			}
		},
	}

	verifyJWSCmd.Flags().StringVarP(&jwkString, "jwk", "k", "", "JSON Web Key")
	verifyJWSCmd.Flags().StringVarP(&jwsString, "jws", "s", "", "JSON Web Signature")
	verifyJWSCmd.MarkPersistentFlagRequired("jwk")
	verifyJWSCmd.MarkPersistentFlagRequired("jws")

	return verifyJWSCmd
}
