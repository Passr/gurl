package requestbuilder_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRequestbuilder(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Requestbuilder Suite")
}
