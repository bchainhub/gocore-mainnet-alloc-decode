 # Description
Simple script to decode [Gocore](https://github.com/core-coin/go-core)'s built-in mainnet genesis allocation in human-readable format.

Genesis allocation content is an RLP-encoded list of (address, balance) tuples.

# Installation

Install [Golang](https://go.dev/doc/install) and [Git](https://github.com/git-guides/install-git)

Clone this repo to your local environment

```bash
git clone git@github.com:error2215/gocore-mainnet-alloc-decode.git
```

# Usage
Go inside install directory
```bash
cd gocore-mainnet-alloc-decode
```

Run script with command 

```bash
go run main.go
```

It will return a list of preallocated addresses with their balances. All external dependencies and data are taken from official Gocore repository - https://github.com/core-coin/go-core