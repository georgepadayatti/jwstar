package jws

import (
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
	}, nil)
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
	// Print public key JWK and JWS to terminal
	fmt.Printf("Public key JWK: \n\n%v\n\n", jwkString)
	fmt.Printf("JWS: \n\n%v\n", jwsString)

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
	fmt.Println("Signature verified.")
	fmt.Printf("\nPayload: \n\n%v\n", string(payload))

	return nil
}
