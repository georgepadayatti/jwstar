package cmd

import (
	"log"

	"github.com/georgepadayatti/jwstar/go/jwk"
	"github.com/georgepadayatti/jwstar/go/jws"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// generateJWSConfig structure to hold the configuration values
type generateJWSConfig struct {
	Claims string `mapstructure:"claims"`
}

func generateJWSCmd() *cobra.Command {

	var generateJWSCmd = &cobra.Command{
		Use:   "generate-jws",
		Short: "Generate JWS",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			// Set the configuration file path
			viper.SetConfigFile("config.yaml")

			// Read the configuration file
			if err := viper.ReadInConfig(); err != nil {
				log.Fatalf("Error reading config file: %s\n", err)
			}

			// Initialize a Config structure to hold the configuration values
			var config generateJWSConfig

			// Unmarshal the configuration into the Config structure
			if err := viper.Unmarshal(&config); err != nil {
				log.Fatalf("Error unmarshalling config: %s\n", err)
			}

			var jwkObj jwk.JWK
			jwsObj := jws.JWS{Key: jwkObj, Claims: config.Claims}
			err := jwsObj.Generate()
			if err != nil {
				log.Fatalf("Error occured while generating JWS: %v", err)
			}
		},
	}

	return generateJWSCmd
}
