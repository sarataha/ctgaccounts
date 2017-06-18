package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func Testctgaccounts(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ctgaccounts Suite")
}
