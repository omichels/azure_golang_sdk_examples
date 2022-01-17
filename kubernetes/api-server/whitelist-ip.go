package api_server

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2021-10-01/containerservice"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/auth"
	"github.com/Azure/go-autorest/autorest"
	"github.com/omichels/azure_golang_sdk_examples/internal/pkg"
	"log"
)

func azureAuthFromEnv() autorest.Authorizer {
	auth, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}
	return auth
}


func WhitelistIPOnAPIServer(ip string, subscription string, resourcegroup string, aksName string) (Error error) {

	var authorizer autorest.Authorizer
	authorizer = azureAuthFromEnv()

	aksClient := containerservice.NewManagedClustersClient(subscription)
	aksClient.Authorizer = authorizer
	managedCluster, err := aksClient.Get(context.Background(), resourcegroup, aksName)
	if err != nil {
		log.Fatal(err)
		return err
	}
	updateNeeded := false
	if managedCluster.APIServerAccessProfile != nil && managedCluster.APIServerAccessProfile.AuthorizedIPRanges != nil {
		ipRanges := *managedCluster.APIServerAccessProfile.AuthorizedIPRanges
		if !pkg.Contains(ipRanges, ip) {
			ipRanges = append(ipRanges, ip)
			updateNeeded = true
			*managedCluster.APIServerAccessProfile.AuthorizedIPRanges = ipRanges
		}
	} else {
		updateNeeded = true
		var newProfile containerservice.ManagedClusterAPIServerAccessProfile
		newRange := make([]string, 0)
		newRange = append(newRange, ip)
		newProfile.AuthorizedIPRanges = &newRange
		managedCluster.APIServerAccessProfile = &newProfile
	}
	if updateNeeded {
		_, err := aksClient.CreateOrUpdate(context.Background(), resourcegroup, aksName, managedCluster)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	return nil
}