package web

import (
	"cloud-native-visibility-hub/pkg/usecases"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var factory = usecases.NewFactory()

func retrieveAllResourcesHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	useCase := factory.NewRetrieveAllResourcesUseCase()

	resources, _ := useCase.Execute()
	resourcesAsJSON, _ := json.Marshal(resources)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(resourcesAsJSON)
}
