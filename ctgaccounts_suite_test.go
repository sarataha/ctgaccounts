package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCinemaToGo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CinemaToGo Suite")
}
