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

	Describe("NewDepositEvent", func() {
		It("can deposit to an account", func() {
			acc, _ := NewCreateAccountEvent(name).Process()

			acc, _ = NewDepositEvent(acc.Id, 20).Process()
			acc, err := NewDepositEvent(acc.Id, 25).Process()

			Expect(err).To(BeNil())
			Expect(acc.Balance).To(Equal(45))
		})
	})

	Describe("NewPayEvent", func() {
		var acc *ctgAccount

		BeforeEach(func() {
			acc, _ = NewCreateAccountEvent(name).Process()
			acc, _ = NewDepositEvent(acc.Id, 20).Process()
		})

		It("can pay money from an account with sufficient balance", func() {
			acc, err := NewPayEvent(acc.Id, 30).Process()

			Expect(err).NotTo(BeNil())
			Expect(acc).To(BeNil())
		})

		It("can pay money from an account with sufficient balance", func() {
			acc, _ = NewPayEvent(acc.Id, 5).Process()
			acc, _ = NewPayEvent(acc.Id, 2).Process()

			Expect(acc.Balance).To(Equal(13))
		})
	})
})
