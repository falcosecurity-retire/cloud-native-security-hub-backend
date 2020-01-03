package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/usecases"
)

type HandlerRepository interface {
	notFound() http.HandlerFunc
	healthCheckHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params)

	retrieveAllResourcesHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params)

	retrieveOneResourcesHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	retrieveFalcoRulesForHelmChartHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params)

	retrieveOneResourceByVersionHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	retrieveFalcoRulesForHelmChartByVersionHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params)

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
	json.NewEncoder(writer).Encode(collectionToDTO(resources))
}

func collectionToDTO(resources []*resource.Resource) []*resource.ResourceDTO {
	var result []*resource.ResourceDTO

	for _, current := range resources {
		result = append(result, resource.NewResourceDTO(current))
	}

	return result
}

func (h *handlerRepository) retrieveOneResourcesHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	kind, err := parseKind(params.ByName("kind"))
	if err != nil {
		writer.WriteHeader(404)
		h.logRequest(request, 404)
		writer.Write([]byte(err.Error()))
		return
	}
	useCase := h.factory.NewRetrieveOneResourceUseCase(params.ByName("resource"), kind)
	resources, err := useCase.Execute()
	if err != nil {
		writer.WriteHeader(500)
		h.logRequest(request, 500)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	h.logRequest(request, 200)
	json.NewEncoder(writer).Encode(resource.NewResourceDTO(resources))
}

func parseKind(slug string) (string, error) {
	switch slug {
	case "falco-rules":
		return resource.FalcoRules, nil
	case "open-policy-agent-policies":
		return resource.OpenPolicyAgentPolicies, nil
	default:
		return "", errors.New(fmt.Sprintf("%s is not a valid kind type", slug))
	}
}

func (h *handlerRepository) retrieveFalcoRulesForHelmChartHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	if params.ByName("kind") != "falco-rules" {
		writer.WriteHeader(400)
		h.logRequest(request, 400)
		writer.Write([]byte(errors.New(fmt.Sprintf("This operation is only allowed for falco-rules")).Error()))
		return
	}

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

func (h *handlerRepository) retrieveOneResourceByVersionHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	kind, err := parseKind(params.ByName("kind"))
	if err != nil {
		writer.WriteHeader(404)
		h.logRequest(request, 404)
		writer.Write([]byte(err.Error()))
		return
	}
	useCase := h.factory.NewRetrieveOneResourceByVersionUseCase(params.ByName("resource"), kind, params.ByName("version"))
	resources, err := useCase.Execute()
	if err != nil {
		writer.WriteHeader(500)
		h.logRequest(request, 500)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	h.logRequest(request, 200)
	json.NewEncoder(writer).Encode(resource.NewResourceDTO(resources))
}

func (h *handlerRepository) retrieveFalcoRulesForHelmChartByVersionHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	if params.ByName("kind") != "falco-rules" {
		writer.WriteHeader(400)
		h.logRequest(request, 400)
		writer.Write([]byte(errors.New(fmt.Sprintf("This operation is only allowed for falco-rules")).Error()))
		return
	}

	useCase := h.factory.NewRetrieveFalcoRulesForHelmChartByVersionUseCase(params.ByName("resource"), params.ByName("version"))
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
	vendors, err := useCase.Execute()
	if err != nil {
		writer.WriteHeader(500)
		h.logRequest(request, 500)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	h.logRequest(request, 200)
	json.NewEncoder(writer).Encode(vendors)
}

func (h *handlerRepository) retrieveOneVendorsHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	useCase := h.factory.NewRetrieveOneVendorUseCase(params.ByName("vendor"))
	vendor, err := useCase.Execute()
	if err != nil {
		writer.WriteHeader(500)
		h.logRequest(request, 500)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	h.logRequest(request, 200)
	json.NewEncoder(writer).Encode(vendor)
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
	json.NewEncoder(writer).Encode(collectionToDTO(resources))
}

func (h *handlerRepository) healthCheckHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	h.logRequest(request, 200)
	writer.Header().Set("Content-Type", "text/plain")
	writer.Write([]byte("OK"))
}
