package resource_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
)

var _ = Describe("ResourceID", func() {
	It("slugifies the name", func() {
		one := resource.NewResourceID("a name", "FalcoRules")
		other := resource.NewResourceID("a-name", "FalcoRules")

		Expect(one).To(Equal(other))
	})

	Context("when comparing ResourceID's", func() {
		It("considers the name of the security resource", func() {
			one := resource.NewResourceID("Nginx", "FalcoRules")
			other := resource.NewResourceID("Apache", "FalcoRules")

			Expect(one).NotTo(Equal(other))

		})

		It("considers the kind of the security resource", func() {
			one := resource.NewResourceID("Nginx", "FalcoRules")
			other := resource.NewResourceID("Nginx", "OpenPolicyAgentPolicies")

			Expect(one).NotTo(Equal(other))
		})
	})
})
