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

var _ = Describe("RetrieveOneResourceByVersion use case", func() {
	var (
		mockCtrl         *gomock.Controller
		mockUpdater      *mock_resource.MockUpdater
		mockEventHandler *mock_event.MockHandler
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockUpdater = mock_resource.NewMockUpdater(mockCtrl)
		mockEventHandler = mock_event.NewMockHandler(mockCtrl)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	It("returns one resource", func() {
		mockUpdater.EXPECT().
			IncrementDownloadCountFor(gomock.Eq("apache")).
			Times(1).
			Return(nil)

		mockEventHandler.EXPECT().
			HandleEvent(gomock.Eq(&event.RetrievedResource{
				ResourceID: "apache",
				Updater:    mockUpdater,
			})).
			Do(func(ev event.Interface) {
				ev.Handle()
			}).
			Times(1)

		useCase := usecases.RetrieveOneResourceByVersion{
			ResourceRepository: newResourceRepositoryWithVersions(),
			ResourceID:         "apache",
			Version:            "1.0.1",
			Updater:            mockUpdater,
			EventHandler:       mockEventHandler,
		}

		result, _ := useCase.Execute()

		apacheWithSpecificVersion := resources.Apache()
		apacheWithSpecificVersion.Version = "1.0.1"
		Expect(result).To(Equal(apacheWithSpecificVersion))
	})

	Context("when version does not exist", func() {
		It("returns an error", func() {
			mockUpdater.EXPECT().
				IncrementDownloadCountFor(gomock.Eq("apache")).
				Times(1).
				Return(resource.ErrResourceNotFound)

			mockEventHandler.EXPECT().
				HandleEvent(gomock.Eq(&event.RetrievedResource{
					ResourceID: "apache",
					Updater:    mockUpdater,
				})).
				Do(func(ev event.Interface) {
					ev.Handle()
				}).
				Times(1)

			useCase := usecases.RetrieveOneResourceByVersion{
				ResourceRepository: newResourceRepositoryWithVersions(),
				ResourceID:         "apache",
				Version:            "2.0.0",
				EventHandler:       mockEventHandler,
				Updater:            mockUpdater,
			}

			result, err := useCase.Execute()

			Expect(result).To(BeNil())
			Expect(err).To(MatchError(resource.ErrResourceNotFound))
		})
	})
})

func newResourceRepositoryWithVersions() resource.Repository {
	apache := resources.Apache()
	apache.Version = "1.0.1"

	return resource.NewMemoryRepository(
		[]*resource.Resource{
			resources.Apache(),
			apache,
		},
	)
}
