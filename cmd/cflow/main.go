package main

import (
	"log"

	"github.com/mlloc/cflow/internal/btc"
	"github.com/mlloc/cflow/internal/env"
)

func main() {
	vars := env.Load()

	client, err := btc.NewClient(vars)
	if err != nil {
		log.Fatal(err)
	}

	info, err := client.GetBlockchainInfo()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("chain: %s", info.Chain)
	log.Printf("blocks: %d", info.Blocks)
}