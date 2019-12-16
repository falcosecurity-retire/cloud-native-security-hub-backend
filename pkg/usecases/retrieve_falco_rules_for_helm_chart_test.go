package usecases_test

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/event"
	mock_event "github.com/falcosecurity/cloud-native-security-hub/pkg/event/mock_resource"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource/mock_resource"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/falcosecurity/cloud-native-security-hub/pkg/usecases"
)

var _ = Describe("RetrieveFalcoRulesForHelmChart use case", func() {
	var (
		mockCtrl           *gomock.Controller
		mockUpdater        *mock_resource.MockUpdater
		newEventDispatcher *mock_event.MockDispatcher
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockUpdater = mock_resource.NewMockUpdater(mockCtrl)
		newEventDispatcher = mock_event.NewMockDispatcher(mockCtrl)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	It("returns the rules for being used with the Helm chart", func() {
		newEventDispatcher.EXPECT().
			Dispatch(gomock.Eq(&event.RetrievedResource{
				ResourceID: "apache",
				Updater:    mockUpdater,
			})).
			Times(1).
			Do(func(ev event.Interface) {
				ev.Handle()
			})

		mockUpdater.EXPECT().
			IncrementDownloadCountFor(gomock.Eq("apache")).
			Times(1).
			Return(nil)

		useCase := usecases.RetrieveFalcoRulesForHelmChart{
			ResourceRepository: NewResourceRepository(),
			EventDispatcher:    newEventDispatcher,
			Updater:            mockUpdater,
			ResourceID:         "apache",
		}

		result, _ := useCase.Execute()

		expected := `customRules:
  rules-apache.yaml: |
    - macro: apache_consider_syscalls
      condition: (evt.num < 0)
`
		Expect(expected).To(Equal(string(result)))
	})

	Context("when resource doesn't exist", func() {
		It("it returns a resource not found error", func() {
			useCase := usecases.RetrieveFalcoRulesForHelmChart{
				ResourceRepository: NewResourceRepository(),
				ResourceID:         "notFound",
			}

			_, err := useCase.Execute()

			Expect(err).To(HaveOccurred())
		})

	})
})
