package main

import (
	"errors"
)

func (e CreateEvent) Process() (*CinemaAccount, error) {
	return updateAccount(e.AccId, map[string]interface{}{
		"Id":      e.AccId,
		"Name":    e.AccName,
		"Balance": "0",
	})
}

func (e PaymentEvent) Process() (*CinemaAccount, error) {
	if acc, err := FetchAccount(e.AccId); err != nil {
		return nil, err
	} else {
		newBalance := acc.Balance + e.Amount
		return updateAccount(e.AccId, map[string]interface{}{
			"Balance": newBalance,
		})
	}
}
