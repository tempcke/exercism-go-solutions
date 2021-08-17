package account

import "sync"

// Account is a bank account
type Account struct {
	mu      sync.RWMutex
	balance int64
	open    bool
}

// Open a new Account
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}

	return &Account{balance: initialDeposit, open: true}
}

// Balance of the Account
func (a *Account) Balance() (balance int64, ok bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.balance, a.open
}

// Deposit into the Account, negative amount does a withdrawal
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	return a.safeWrite(func() (int64, bool) {
		if a.balance+amount < 0 {
			return a.balance, false
		}

		a.balance += amount
		return a.balance, true
	})
}

// Close the Account
func (a *Account) Close() (payout int64, ok bool) {
	return a.safeWrite(func() (int64, bool) {
		b := a.balance
		a.open = false
		a.balance = 0
		return b, true
	})
}

func (a *Account) safeWrite(operation func() (int64, bool)) (amount int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.open {
		return 0, false
	}

	return operation()
}
