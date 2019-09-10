package web

import (
	"encoding/json"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/usecases"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type HandlerRepository interface {
	retrieveAllResourcesHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params)
	retrieveOneResourcesHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	retrieveOneResourcesRawHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	retrieveAllVendorsHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params)
	retrieveOneVendorsHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	retrieveAllResourcesFromVendorHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type handlerRepository struct {
	factory usecases.Factory
}

func NewHandlerRepository() HandlerRepository {
	return &handlerRepository{
		factory: usecases.NewFactory(),
	}
}

func (h *handlerRepository) retrieveAllResourcesHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	useCase := h.factory.NewRetrieveAllResourcesUseCase()
	resources, err := useCase.Execute()
	if err != nil {
		writer.WriteHeader(500)
		writer.Write([]byte(err.Error()))
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(resources)
}

func (h *handlerRepository) retrieveOneResourcesHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	useCase := h.factory.NewRetrieveOneResourceUseCase(params.ByName("hash"))
	resources, err := useCase.Execute()
	if err != nil {
		writer.WriteHeader(500)
		writer.Write([]byte(err.Error()))
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(resources)
}


func (h *handlerRepository) retrieveOneResourcesRawHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	useCase := h.factory.NewRetrieveOneRawResourceUseCase(params.ByName("hash"))
	content, err := useCase.Execute()
	if err != nil {
		writer.WriteHeader(500)
		writer.Write([]byte(err.Error()))
	}
	writer.Header().Set("Content-Type", "application/x-yaml")
	writer.Write(content.Raw())
}


func (h *handlerRepository) retrieveAllVendorsHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	useCase := h.factory.NewRetrieveAllVendorsUseCase()
	resources, err := useCase.Execute()
	if err != nil {
		writer.WriteHeader(500)
		writer.Write([]byte(err.Error()))
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(resources)
}

func (h *handlerRepository) retrieveOneVendorsHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	useCase := h.factory.NewRetrieveOneVendorUseCase(params.ByName("vendor"))
	resources, err := useCase.Execute()
	if err != nil {
		writer.WriteHeader(500)
		writer.Write([]byte(err.Error()))
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(resources)
}

func (h *handlerRepository) retrieveAllResourcesFromVendorHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	useCase := h.factory.NewRetrieveAllResourcesFromVendorUseCase(params.ByName("vendor"))
	resources, err := useCase.Execute()
	if err != nil {
		writer.WriteHeader(500)
		writer.Write([]byte(err.Error()))
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(resources)
}
