package env

import "os"

type Env struct {
	BTC BTCEnv
}

type BTCEnv struct {
	Host string
	User string
	Pass string
}

func Load() Env {
	return Env{
		BTC: BTCEnv{
			Host: os.Getenv("CFLOW_BTC_HOST"),
			User: os.Getenv("CFLOW_BTC_USER"),
			Pass: os.Getenv("CFLOW_BTC_PASS"),
		},
	}
}
