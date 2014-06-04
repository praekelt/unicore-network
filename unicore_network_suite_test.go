package unicore_network_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestUnicoreNetwork(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UnicoreNetwork Suite")
}
