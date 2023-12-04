package response_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestResponse(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Response Suite")
}
