package rpsls_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRpsls(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RPSLS Suite")
}
