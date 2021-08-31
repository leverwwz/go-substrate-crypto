# go-substrate-crypto
go sdk for palkdot ecology

usage

0. add "github.com/leverwwz/go-substrate-crypto" into go.mod file

    go get github.com/leverwwz/go-substrate-crypto
    
1. add a prefix for the target chain if the prefix not contained in the list.

    ss58.AddPrefix()

2. generate the private.

3. generate address against to the prefix of a chain.
    
4. sign the message(note,three type of signing way available)
    
    ed25519.Sign
    sr25519.Sign
    ecdsa.Sign

for more details, please refer to testing files in "go-substrate-crypto/test".