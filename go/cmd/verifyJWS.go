package cmd

import (
	"log"

	"github.com/georgepadayatti/jwstar/go/jwk"
	"github.com/georgepadayatti/jwstar/go/jws"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// verifyJWSConfig structure to hold the configuration values
type verifyJWSConfig struct {
	JWS string `mapstructure:"jws"`
	JWK string `mapstructure:"jwk"`
}

func verifyJWSCmd() *cobra.Command {
	var verifyJWSCmd = &cobra.Command{
		Use:   "verify-jws",
		Short: "Verify JWS",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {

			// Set the configuration file path
			viper.SetConfigFile("config.yaml")

			// Read the configuration file
			if err := viper.ReadInConfig(); err != nil {
				log.Fatalf("Error reading config file: %s\n", err)
			}

			// Initialize a Config structure to hold the configuration values
			var config verifyJWSConfig

			// Unmarshal the configuration into the Config structure
			if err := viper.Unmarshal(&config); err != nil {
				log.Fatalf("Error unmarshalling config: %s\n", err)
			}

			jwsObj := jws.JWS{Key: jwk.FromJSON(config.JWK), Signature: config.JWS}
			err := jwsObj.Verify()
			if err != nil {
				log.Fatalf("Error occured while verifying JWS: %v", err)
			}
		},
	}

	return verifyJWSCmd
}
