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

type DepositEvent struct {
	Event
	Amount int
}

type PayEvent struct {
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

func NewDepositEvent(id string, amt int) DepositEvent {
	event := new(DepositEvent)
	event.Type = "DepositEvent"
	event.AccId = id
	event.Amount = amt
	return *event
}

func NewPayEvent(id string, amt int) PayEvent {
	event := new(PayEvent)
	event.Type = "PayEvent"
	event.AccId = id
	event.Amount = amt
	return *event
}
