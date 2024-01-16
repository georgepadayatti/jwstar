import json

from jwcrypto import jwk, jwt


def generate_secp256r1_jwk():
    # Generate EC key pair for secp256r1 curve
    key = jwk.JWK.generate(kty="EC", crv="P-256")
    return key


def sign_jws(payload, private_key):
    # Generate JSON web signature
    key = jwk.JWK(**private_key)
    token = jwt.JWT(header={"alg": "ES256"}, claims=payload)
    token.make_signed_token(key)
    return token.serialize()


if __name__ == "__main__":
    private_key = generate_secp256r1_jwk()

    payload = {"data": "Hello, JWS!"}

    jws_token = sign_jws(payload, private_key.export(as_dict=True))

    print(
        f"Public key JWK:\n{json.dumps(private_key.export(private_key=False, as_dict=True), indent=2)}\n"
    )
    print(f"JWS:\n{jws_token}")
