package account

type Account struct {
	balance int64
}

func Open(balance int64) *Account {
	if balance < 0 {
		return nil
	}

	return &Account{balance: balance}
}

func (a *Account) Close() (int64, bool) {
	balance := a.balance
	a = nil
	return balance, true
}

func (a *Account) Balance() (int64, bool) {
	if a.balance < 0 {
		return a.balance, false
	}

	return a.balance, true
}
