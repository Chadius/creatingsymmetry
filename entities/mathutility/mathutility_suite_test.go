package mathutility_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMathutility(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mathutility Suite")
}
