package base62_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBase62(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Base62 Suite")
}
