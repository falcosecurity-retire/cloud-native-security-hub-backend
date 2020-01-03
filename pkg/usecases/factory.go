package usecases

import (
	"database/sql"
	"log"
	"os"

	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/vendor"
)

type Factory interface {
	NewRetrieveAllResourcesUseCase() *RetrieveAllResources
	NewRetrieveOneResourceUseCase() *RetrieveOneResource
	NewRetrieveOneResourceByVersionUseCase() *RetrieveOneResourceByVersion
	NewRetrieveFalcoRulesForHelmChartUseCase() *RetrieveFalcoRulesForHelmChart
	NewRetrieveFalcoRulesForHelmChartByVersionUseCase() *RetrieveFalcoRulesForHelmChartByVersion
	NewRetrieveAllVendorsUseCase() *RetrieveAllVendors
	NewRetrieveOneVendorUseCase() *RetrieveOneVendor
	NewRetrieveAllResourcesFromVendorUseCase() *RetrieveAllResourcesFromVendor

	NewResourcesRepository() resource.Repository
	NewVendorRepository() vendor.Repository
}

func NewFactory() Factory {
	factory := &factory{}
	factory.db = factory.newDB()
	factory.resourceRepository = factory.NewResourcesRepository()
	factory.vendorRepository = factory.NewVendorRepository()
	return factory
}

// TODO: Instantiate useCases only once
type factory struct {
	db                 *sql.DB
	vendorRepository   vendor.Repository
	resourceRepository resource.Repository
}

func (f *factory) NewRetrieveAllResourcesUseCase() *RetrieveAllResources {
	return &RetrieveAllResources{ResourceRepository: f.resourceRepository}
}

func (f *factory) NewRetrieveOneResourceUseCase() *RetrieveOneResource {
	return &RetrieveOneResource{ResourceRepository: f.resourceRepository}
}

func (f *factory) NewRetrieveOneResourceByVersionUseCase() *RetrieveOneResourceByVersion {
	return &RetrieveOneResourceByVersion{ResourceRepository: f.resourceRepository}
}

func (f *factory) NewRetrieveFalcoRulesForHelmChartUseCase() *RetrieveFalcoRulesForHelmChart {
	return &RetrieveFalcoRulesForHelmChart{ResourceRepository: f.resourceRepository}
}

func (f *factory) NewRetrieveFalcoRulesForHelmChartByVersionUseCase() *RetrieveFalcoRulesForHelmChartByVersion {
	return &RetrieveFalcoRulesForHelmChartByVersion{ResourceRepository: f.resourceRepository}
}

func (f *factory) NewRetrieveAllVendorsUseCase() *RetrieveAllVendors {
	return &RetrieveAllVendors{
		VendorRepository: f.vendorRepository,
	}
}

func (f *factory) NewRetrieveOneVendorUseCase() *RetrieveOneVendor {
	return &RetrieveOneVendor{VendorRepository: f.vendorRepository}
}

func (f *factory) NewRetrieveAllResourcesFromVendorUseCase() *RetrieveAllResourcesFromVendor {
	return &RetrieveAllResourcesFromVendor{
		VendorRepository:   f.vendorRepository,
		ResourceRepository: f.resourceRepository,
	}
}

func (f *factory) NewResourcesRepository() resource.Repository {
	return resource.NewPostgresRepository(f.db)
}

func (f *factory) NewVendorRepository() vendor.Repository {
	return vendor.NewPostgresRepository(f.db)
}

func (f *factory) newDB() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return db
}
