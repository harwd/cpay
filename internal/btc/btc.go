package btc

import (
	"github.com/btcsuite/btcd/rpcclient"
)

type Client struct {
	rpc *rpcclient.Client
}