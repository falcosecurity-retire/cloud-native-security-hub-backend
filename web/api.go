package web

import (
	"encoding/json"
	"fmt"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/usecases"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type HandlerRepository interface {
	notFound() http.HandlerFunc
	retrieveAllResourcesHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params)
	retrieveOneResourcesHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	retrieveFalcoRulesForHelmChartHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	retrieveAllVendorsHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params)
	retrieveOneVendorsHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	retrieveAllResourcesFromVendorHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type handlerRepository struct {
	factory usecases.Factory
	logger  *log.Logger
}

func NewHandlerRepository(logger *log.Logger) HandlerRepository {
	return &handlerRepository{
		factory: usecases.NewFactory(),
		logger:  logger,
	}
}

func (h *handlerRepository) notFound() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		h.logRequest(request, 404)
		http.NotFound(writer, request)
	}

}
func (h *handlerRepository) logRequest(request *http.Request, statusCode int) {
	if h.logger == nil {
		return
	}

	line := fmt.Sprintf("%d [%s] %s %s", statusCode, request.RemoteAddr, request.Method, request.URL)
	h.logger.Println(line)
}

func (h *handlerRepository) retrieveAllResourcesHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	useCase := h.factory.NewRetrieveAllResourcesUseCase()
	resources, err := useCase.Execute()
	if err != nil {
		h.logRequest(request, 500)
		writer.WriteHeader(500)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.Header().Set("Content-Type", "application/json")

	h.logRequest(request, 200)
	json.NewEncoder(writer).Encode(resources)
}

func (h *handlerRepository) retrieveOneResourcesHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	useCase := h.factory.NewRetrieveOneResourceUseCase(params.ByName("resource"))
	resources, err := useCase.Execute()
	if err != nil {
		writer.WriteHeader(500)
		h.logRequest(request, 500)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	h.logRequest(request, 200)
	json.NewEncoder(writer).Encode(resources)
}

func (h *handlerRepository) retrieveFalcoRulesForHelmChartHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	useCase := h.factory.NewRetrieveFalcoRulesForHelmChartUseCase(params.ByName("resource"))
	content, err := useCase.Execute()
	if err != nil {
		writer.WriteHeader(500)
		h.logRequest(request, 500)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.Header().Set("Content-Type", "application/x-yaml")
	h.logRequest(request, 200)
	writer.Write(content)
}

func (h *handlerRepository) retrieveAllVendorsHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	useCase := h.factory.NewRetrieveAllVendorsUseCase()
	resources, err := useCase.Execute()
	if err != nil {
		writer.WriteHeader(500)
		h.logRequest(request, 500)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	h.logRequest(request, 200)
	json.NewEncoder(writer).Encode(resources)
}

func (h *handlerRepository) retrieveOneVendorsHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	useCase := h.factory.NewRetrieveOneVendorUseCase(params.ByName("vendor"))
	resources, err := useCase.Execute()
	if err != nil {
		writer.WriteHeader(500)
		h.logRequest(request, 500)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	h.logRequest(request, 200)
	json.NewEncoder(writer).Encode(resources)
}

func (h *handlerRepository) retrieveAllResourcesFromVendorHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	useCase := h.factory.NewRetrieveAllResourcesFromVendorUseCase(params.ByName("vendor"))
	resources, err := useCase.Execute()
	if err != nil {
		writer.WriteHeader(500)
		h.logRequest(request, 500)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	h.logRequest(request, 200)
	json.NewEncoder(writer).Encode(resources)
}
