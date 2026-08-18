package main

import (
	"log"

	"github.com/mlloc/cflow/internal/btc"
	"github.com/mlloc/cflow/internal/env"
)

func main() {
	vars := env.Load()

	_, err := btc.NewClient(vars)
	if err != nil {
		log.Fatal(err)
	}
}