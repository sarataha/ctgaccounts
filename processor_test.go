package main_test

import (
	. "ctgaccounts"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Event processor", func() {
	var (
		name = "Sara Taha"
	)

	Describe("NewCreateAccountEvent", func() {
		It("can create an account", func() {
			acc, err := NewCreateAccountEvent(name).Process()

			Expect(err).To(BeNil())
			Expect(acc.Name).To(Equal(name))
			Expect(acc.Balance).To(Equal(0))
		})
	})

	Describe("NewPaymentEvent", func() {
		It("can do payment to an account", func() {
			acc, _ := NewCreateAccountEvent(name).Process()

			acc, _ = NewDepositEvent(acc.Id, 20).Process()
			acc, err := NewDepositEvent(acc.Id, 25).Process()

			Expect(err).To(BeNil())
			Expect(acc.Balance).To(Equal(45))
		})
	})
})
