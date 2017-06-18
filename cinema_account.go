package main

import (
	"errors"
	"strconv"
)

type CinemaAccount struct {
	Id      string
	Name    string
	Balance int
}

func FetchAccount(id string) (*CinemaAccount, error) {
	cmd := Redis.HGetAll(id)
	if err := cmd.Err(); err != nil {
		return nil, err
	}

	data := cmd.Val()
	if len(data) == 0 {
		return nil, nil
	} else {
		return ToAccount(data)
	}
}

func updateAccount(id string, data map[string]interface{}) (*CinemaAccount, error) {
	cmd := Redis.HMSet(id, data)

	if err := cmd.Err(); err != nil {
		return nil, err
	} else {
		return FetchAccount(id)
	}
}

func ToAccount(m map[string]string) (*CinemaAccount, error) {
	balance, err := strconv.Atoi(m["Balance"])
	if err != nil {
		return nil, err
	}

	if _, ok := m["Id"]; !ok {
		return nil, errors.New("Missing account ID")
	}

	return &CinemaAccount{
		Id:      m["Id"],
		Name:    m["Name"],
		Balance: balance,
	}, nil
}
