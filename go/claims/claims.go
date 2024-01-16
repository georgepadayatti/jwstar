package claims

import (
	"encoding/json"
	"log"
)

// LoadClaimsFromString unmarshal claims to `map[string]interface{}`
func LoadClaimsFromString(claimsString string) map[string]interface{} {
	var claims map[string]interface{}
	err := json.Unmarshal([]byte(claimsString), &claims)
	if err != nil {
		log.Fatal("Failed to unmarshal claims:", err)
	}
	return claims
}
