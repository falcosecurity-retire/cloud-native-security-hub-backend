package usecases_test

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/event"
	mock_event "github.com/falcosecurity/cloud-native-security-hub/pkg/event/mock_resource"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource/mock_resource"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/falcosecurity/cloud-native-security-hub/test/fixtures/resources"

	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/usecases"
)

var _ = Describe("RetrieveOneResource use case", func() {
	var (
		mockCtrl            *gomock.Controller
		mockUpdater         *mock_resource.MockUpdater
		mockEventDispatcher *mock_event.MockDispatcher
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockUpdater = mock_resource.NewMockUpdater(mockCtrl)
		mockEventDispatcher = mock_event.NewMockDispatcher(mockCtrl)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	It("returns one resource", func() {
		mockEventDispatcher.EXPECT().
			Dispatch(gomock.Eq(&event.RetrievedResource{
				ResourceID: "apache",
				Updater:    mockUpdater,
			})).
			Do(func(event event.Interface) {
				event.Handle()
			}).
			Times(1)

		mockUpdater.EXPECT().
			IncrementDownloadCountFor(gomock.Eq("apache")).
			Times(1).
			Return(nil)

		useCase := usecases.RetrieveOneResource{
			ResourceRepository: NewResourceRepository(),
			ResourceID:         "apache",
			EventDispatcher:    mockEventDispatcher,
			Updater:            mockUpdater,
		}

		result, _ := useCase.Execute()

		Expect(result).To(Equal(resources.Apache()))
	})

	Context("when resource does not exist", func() {
		It("returns resource not found error", func() {
			mockEventDispatcher.EXPECT().
				Dispatch(gomock.Eq(&event.RetrievedResource{
					ResourceID: "notFound",
					Updater:    mockUpdater,
				})).
				Do(func(ev event.Interface) {
					ev.Handle()
				}).
				Times(1)

			mockUpdater.EXPECT().
				IncrementDownloadCountFor(gomock.Eq("notFound")).
				Times(1).
				Return(resource.ErrResourceNotFound)

			useCase := usecases.RetrieveOneResource{
				ResourceRepository: NewResourceRepository(),
				ResourceID:         "notFound",
				EventDispatcher:    mockEventDispatcher,
				Updater:            mockUpdater,
			}

			retrieved, err := useCase.Execute()

			Expect(retrieved).To(BeNil())
			Expect(err).To(MatchError(resource.ErrResourceNotFound))
		})
	})
})
