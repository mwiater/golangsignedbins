# golangsignedbins

Signing a binary in Go (or any other language) typically involves creating a hash of the binary file and then encrypting this hash with a private key. The encrypted hash constitutes the digital signature.

Here's a step-by-step guide and explanation on how to sign a Go binary: [Golang: Verifying Application Integrity by Signing Binaries](https://medium.com/@matt.wiater/golang-verifying-integrity-by-signing-binaries-9b4497d5d761)

## Setup Heartbeat Test Application

### 1. Build the heartbeat binary

**Execute before compiling:**

`go run heartbeat/heartbeat.go`

```
2023-11-19T23:31:58Z Heartbeat...
2023-11-19T23:32:03Z Heartbeat...
2023-11-19T23:32:08Z Heartbeat...
2023-11-19T23:32:13Z Heartbeat...
2023-11-19T23:32:18Z Heartbeat...
2023-11-19T23:32:23Z Heartbeat...
```

**Build:**

`go build -o heartbeat/bin/heartbeat ./heartbeat`

**Execute binary:**

`./heartbeat/bin/heartbeat`

## Signing the Binary

### 1. Generate a Private/Public Key Pair

**Generate a RSA private/public key pair using OpenSSL:**

```bash
openssl genpkey -algorithm RSA -out ./keys/private_key.pem
openssl rsa -pubout -in ./keys/private_key.pem -out ./keys/public_key.pem
```

This will create a `./keys/private_key.pem` and `./keys/public_key.pem` file.

### 2. Sign the Binary

Run: `go run signer/signer.go`

This will sign the binary by creating and saving the signature file: `./signatures/heartbeat.sig`

### 3. Verify and Run the Binary

Run: `go run runner/runner.go`

This will verify the signed binary using the `./signatures/heartbeat.sig` file and run the signed and verified binary.

### 3. Tamper and Fail

To check for an invalid binary, I've included a file to modify the binary after signing. To see the results, after signing the binary, add some extra data to the binary and try running it again:

1. Run: `go run signer/signer.go`
2. Run: `go run tamper/tamper.go`
3. Run: `go run runner/runner.go`

This will now fail with the error message: `[Error] verify signature: crypto/rsa: verification error`

_To bring the signed binary back to a valid state, just run this again:_ `go run signer/signer.go`

## Testing

`go test ./common/common.go ./common/common_test.go`

```
ok      command-line-arguments  0.038s
```
