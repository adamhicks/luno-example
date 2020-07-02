# luno-example
Example use of the Luno API using the Go SDK

View the slides at https://talks.godoc.org/github.com/adamhicks/luno-example/content.slide

# Getting Started
- Clone this repository
- Install go https://golang.org/doc/install
- Create a Luno API key https://www.luno.com/wallet/security/api_keys
- Paste the API key and secret into `balances/balances.go` `lunoClient.SetAuth()`
- `go run balances/balances.go`