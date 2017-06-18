package main

import (
	uuid "github.com/satori/go.uuid"
)

/* each event expressed as its own struct */

type Event struct {
	AccId string
	Type  string
}

type CreateEvent struct {
	Event
	AccName string
}

type PaymentEvent struct {
	Event
	Amount int
}

/* helper to create events */

func NewCreateAccountEvent(name string) CreateEvent {
	event := new(CreateEvent)
	event.Type = "CreateEvent"
	event.AccId = uuid.NewV4().String()
	event.AccName = name
	return *event
}

func NewPaymentEvent(id string, amt int) PaymentEvent {
	event := new(DepositEvent)
	event.Type = "PaymentEvent"
	event.AccId = id
	event.Amount = amt
	return *event
}
