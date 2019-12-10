package usecases_test

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource/mock_resource"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/usecases"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("IncrementDownloadCount use case", func() {
	var (
		mockCtrl    *gomock.Controller
		mockUpdater *mock_resource.MockUpdater
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockUpdater = mock_resource.NewMockUpdater(mockCtrl)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	It("is being called by the use case", func() {
		incrementDownloadCount := usecases.IncrementDownloadCount{
			Updater:    mockUpdater,
			ResourceID: "apache",
		}

		mockUpdater.
			EXPECT().
			IncrementDownloadCountFor(gomock.Eq("apache")).
			Times(1)

		err := incrementDownloadCount.Execute()

		Expect(err).To(BeNil())
	})
})
