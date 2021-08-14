package account

import "sync"

type Account struct {
	balance int64
	open    bool
	mu      sync.Mutex
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}

	return &Account{balance: initialDeposit, open: true}
}

func (a *Account) Balance() (balance int64, ok bool) {
	return a.balance, a.open
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	return a.safeChange(func() (int64, bool) {
		if a.balance+amount < 0 {
			return a.balance, false
		}

		a.balance += amount
		return a.balance, true
	})
}

func (a *Account) Close() (payout int64, ok bool) {
	return a.safeChange(func() (int64, bool) {
		b := a.balance
		a.open = false
		a.balance = 0
		return b, true
	})
}

func (a *Account) safeChange(operation func() (int64, bool)) (amount int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.open {
		return 0, false
	}

	return operation()
}