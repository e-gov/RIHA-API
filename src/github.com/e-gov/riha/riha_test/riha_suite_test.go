package riha_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFox(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RIHA Suite")
}
