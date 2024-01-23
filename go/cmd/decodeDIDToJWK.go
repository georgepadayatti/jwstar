package cmd

import (
	"fmt"
	"log"

	"github.com/georgepadayatti/jwstar/go/did"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// decodeDIDToJWKConfig structure to hold the configuration values
type decodeDIDToJWKConfig struct {
	DID string `mapstructure:"did"`
}

func decodeDIDToJWKCmd() *cobra.Command {
	var verifyJWSCmd = &cobra.Command{
		Use:   "decode-did",
		Short: "Decode DID",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {

			// Set the configuration file path
			viper.SetConfigFile("config.yaml")

			// Read the configuration file
			if err := viper.ReadInConfig(); err != nil {
				log.Fatalf("Error reading config file: %s\n", err)
			}

			// Initialize a Config structure to hold the configuration values
			var config decodeDIDToJWKConfig

			// Unmarshal the configuration into the Config structure
			if err := viper.Unmarshal(&config); err != nil {
				log.Fatalf("Error unmarshalling config: %s\n", err)
			}

			var didkey did.DIDKey
			if err := didkey.DecodeToBytes(config.DID); err != nil {
				log.Fatalf("Error decoding did:key to bytes: %s\n", err)
			}

			fmt.Println(didkey.ToJSON())
		},
	}

	return verifyJWSCmd
}
