package vendor_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestVendor(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Vendor Suite")
}
