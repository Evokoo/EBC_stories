package quest01_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test01(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "$ Suite")
}