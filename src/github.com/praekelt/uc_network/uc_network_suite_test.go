package uc_network_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestUcNetwork(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UcNetwork Suite")
}
