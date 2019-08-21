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


func retrieveOneResourcesHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	useCase := factory.NewRetrieveOneResourceUseCase(params.ByName("hash"))
	resources, err := useCase.Execute()
	if err != nil {
		writer.WriteHeader(500)
		writer.Write([]byte(err.Error()))
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(resources)
}
func retrieveAllVendorsHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	useCase := factory.NewRetrieveAllVendorsUseCase()
	resources, err := useCase.Execute()
	if err != nil {
		writer.WriteHeader(500)
		writer.Write([]byte(err.Error()))
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(resources)
}

func retrieveOneVendorsHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	useCase := factory.NewRetrieveOneVendorUseCase(params.ByName("vendor"))
	resources, err := useCase.Execute()
	if err != nil {
		writer.WriteHeader(500)
		writer.Write([]byte(err.Error()))
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(resources)
}
