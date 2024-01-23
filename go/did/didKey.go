package did

import (
	"strings"

	"github.com/georgepadayatti/jwstar/go/jwk"
	"github.com/multiformats/go-multibase"
	"github.com/wealdtech/go-multicodec"
)

type DIDKey struct {
	decoded []byte
}

func (obj *DIDKey) DecodeToBytes(did string) error {
	_, dataWithCodec, err := multibase.Decode(strings.TrimPrefix(did, "did:key:"))
	if err != nil {
		return err
	}

	// Remove the codec
	dataWithoutCodec, _, err := multicodec.RemoveCodec(dataWithCodec)
	if err != nil {
		return err
	}

	obj.decoded = dataWithoutCodec

	return nil
}

func (obj *DIDKey) ToJSON() string {
	return string(obj.decoded)
}

func (obj *DIDKey) ToJWK() jwk.JWK {
	return jwk.FromJSON(string(obj.decoded))
}
