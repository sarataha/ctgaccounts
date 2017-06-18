package main_test

import (
	. "ctgaccounts"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Event", func() {
	var (
		accId  = "AB928323-2837232832-3283232-232"
		amount = 2000
	)

	Describe("NewCreateAccountEvent", func() {
		It("can create a create account event", func() {
			name := "Sara Taha"

			event := NewCreateAccountEvent(name)

			Expect(event.AccName).To(Equal(name))
			Expect(event.AccId).NotTo(BeNil())
			Expect(event.Type).To(Equal("CreateEvent"))
		})
	})

	Describe("NewPaymentEvent", func() {
		It("can create a payment event", func() {
			event := NewDepositEvent(accId, amount)

			Expect(event.AccId).To(Equal(accId))
			Expect(event.Amount).To(Equal(amount))
			Expect(event.Type).To(Equal("PaymentEvent"))
		})
	})
})
