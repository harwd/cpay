package pay

type Payment struct {
	Id       string
	Currency string
	Amount   int
	Address  string
	Status   string
	TXID     string
}