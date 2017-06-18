package main

import (
	"errors"
)

func (e CreateEvent) Process() (*ctgAccount, error) {
	return updateAccount(e.AccId, map[string]interface{}{
		"Id":      e.AccId,
		"Name":    e.AccName,
		"Balance": "0",
	})
}

func (e DepositEvent) Process() (*ctgAccount, error) {
	if acc, err := FetchAccount(e.AccId); err != nil {
		return nil, err
	} else {
		newBalance := acc.Balance + e.Amount
		return updateAccount(e.AccId, map[string]interface{}{
			"Balance": newBalance,
		})
	}
}

func (e PayEvent) Process() (*ctgAccount, error) {
	if acc, err := FetchAccount(e.AccId); err != nil {
		return nil, err
	} else {
		if acc.Balance >= e.Amount {
			newBalance := acc.Balance - e.Amount
			return updateAccount(e.AccId, map[string]interface{}{
				"Balance": newBalance,
			})
		} else {
			return nil, errors.New("Insufficient amount")
		}
	}
}
