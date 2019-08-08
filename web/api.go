package web

import (
	"cloud-native-visibility-hub/pkg/resource"
	"cloud-native-visibility-hub/pkg/usecases"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type dummyResourcesRepository struct {
}

func (resources *dummyResourcesRepository) All() ([]resource.Resource, error) {
	return []resource.Resource{
		{
			Name: "Falco profile for Nginx",
		},
		{
			Name: "Grafana Dashboard for Traefik",
		},
	}, nil
}

func retrieveAllResourcesHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	useCase := usecases.RetrieveAllResources{
		Resources: &dummyResourcesRepository{},
	}

	resources, _ := useCase.Execute()
	resourcesAsJSON, _ := json.Marshal(resources)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(resourcesAsJSON)
}
