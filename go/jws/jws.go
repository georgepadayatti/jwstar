package jws

import (
	"encoding/json"
	"fmt"

	"github.com/georgepadayatti/jwstar/go/jwk"
	"github.com/go-jose/go-jose/v3"
)

// JWS represents a JSON Web Signature (JWS)
type JWS struct {
	Claims    string  `json:"-"`
	Key       jwk.JWK `json:"-"`
	Signature string  `json:"-"`
}

// Generate generate JSON web signature (JWS)
func (obj *JWS) Generate() error {
	// Generate EC key pair
	privateKey := obj.Key.GenerateECKey()
	// Serialise to JWK string
	jwkString := obj.Key.ToJSON()
	// Create signer
	signer, err := jose.NewSigner(jose.SigningKey{
		Algorithm: jose.ES256,
		Key:       privateKey,
	}, &jose.SignerOptions{
		ExtraHeaders: map[jose.HeaderKey]interface{}{"jwk": jwk.FromJSON(jwkString), "typ": "JWT"},
	})
	if err != nil {
		return err
	}
	// Sign claims
	jws, err := signer.Sign([]byte(obj.Claims))
	if err != nil {
		return err
	}
	// Serialise to JWS string
	jwsString, err := jws.CompactSerialize()
	if err != nil {
		return err
	}

	res, err := json.Marshal(map[string]string{"jwk": jwkString, "jws": jwsString})
	if err != nil {
		return err
	}

	// Print public key JWK and JWS to terminal
	fmt.Println(string(res))

	return nil
}

// Verify verify JSON web signature (JWS)
func (obj *JWS) Verify() error {
	// Create EC public key
	pubKey := obj.Key.ToECPublicKey()
	// Deserialise to JWS
	jws, err := jose.ParseSigned(obj.Signature)
	if err != nil {
		return err
	}
	// Verify signature and return decoded payload
	payload, err := jws.Verify(pubKey)
	if err != nil {
		return err
	}
	// Print decoded payload
	fmt.Println(string(payload))

	return nil
}
