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
	resources, err := useCase.Execute()
	if err != nil {
		writer.WriteHeader(500)
		writer.Write([]byte(err.Error()))
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(resources)
}
