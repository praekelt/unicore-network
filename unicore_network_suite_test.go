package unicore_network

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestUnicoreNetwork(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UnicoreNetwork Suite")
}
