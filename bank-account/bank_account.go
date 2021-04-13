package account

type Account struct {
	balance int64
}

func Open(balance int64) *Account {
	return &Account{balance: balance}
}
