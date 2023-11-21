# golangsignedbins

Signing a binary in Go (or any other language) typically involves creating a hash of the binary file and then encrypting this hash with a private key. The encrypted hash constitutes the digital signature. Here's a step-by-step guide on how to sign a Go binary:

## Build the heartbeat binary

Test:

`go run heartbeat/heartbeat.go`

```
2023-11-19T23:31:58Z Heartbeat...
2023-11-19T23:32:03Z Heartbeat...
2023-11-19T23:32:08Z Heartbeat...
2023-11-19T23:32:13Z Heartbeat...
2023-11-19T23:32:18Z Heartbeat...
2023-11-19T23:32:23Z Heartbeat...
```

Build:

`go build -o heartbeat/bin/heartbeat ./heartbeat`

Test binary:

`./heartbeat/bin/heartbeat`


### 1. Generate a Private/Public Key Pair

First, you need a RSA private/public key pair. You can generate this using OpenSSL:

```bash
openssl genpkey -algorithm RSA -out ./keys/private_key.pem
openssl rsa -pubout -in ./keys/private_key.pem -out ./keys/public_key.pem
```

This will create a `./keys/private_key.pem` and `./keys/public_key.pem` file.

### 2. Sign the Binary

Run: `go run signer/signer.go`

This will sign the binary and save the file: `./signatures/heartbeat.sig`

### Integrating with Your Go Application

In your Go application, as outlined in the previous response, you'll load the `public_key.pem` and use it to verify the `./signatures/heartbeat.sig` against the binary file before executing it. This ensures that the binary hasn't been tampered with since it was signed.

## Testing

`go test ./...`

Or

`gotestsum --format testname`
