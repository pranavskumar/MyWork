package main

import (
	"context"
	"fmt"
	"os"

	openapiclient "github.com/pranavskumar/MySDK"
)

func main() {
	configuration := openapiclient.NewConfiguration()
	configuration.Host = "localhost:8000"
	configuration.Scheme = "http"
	configuration.DefaultHeader = map[string]string{
		"Content-Type": "application/json",
	}
	configuration.UserAgent = "OpenAPI-Generator/1.0.0/go"
	configuration.Debug = false
	configuration.Servers = openapiclient.ServerConfigurations{
		{
			URL:         "http://localhost:8080",
			Description: "No description provided",
		},
	}
	configuration.OperationServers = map[string]openapiclient.ServerConfigurations{}
	configuration.HTTPClient = nil

	CreateOriginPool(configuration)
	ListOriginPool(configuration)
	ReplaceOriginPool(configuration)
	GetOriginPool(configuration)
	DeleteOriginPool(configuration)
}

func CreateOriginPool(configuration *openapiclient.Configuration) {
	apiClient := openapiclient.NewAPIClient(configuration)
	metadataNamespace := "create"

	name := "originPool_1"
	description := "origin pool description"
	var threshold int64 = 10
	originpool := openapiclient.OriginPoolCreateRequest{
		Metadata: &openapiclient.SchemaObjectCreateMetaType{
			Name:        &name,
			Description: &description,
		},
		Spec: &openapiclient.ViewsoriginPoolCreateSpecType{
			AdvancedOptions: &openapiclient.OriginPoolOriginPoolAdvancedOptions{
				PanicThreshold: &threshold,
			},
		},
	}

	resp, r, err := apiClient.DefaultAPI.VesIoSchemaViewsOriginPoolAPICreate(context.Background(), metadataNamespace).OriginPoolCreateRequest(originpool).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.VesIoSchemaViewsOriginPoolAPICreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	fmt.Fprintf(os.Stdout, "%v\n", r)
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.VesIoSchemaViewsOriginPoolAPICreate`: %v\n", resp)
}

func ListOriginPool(configuration *openapiclient.Configuration) {
	apiClient := openapiclient.NewAPIClient(configuration)

	metadataNamespace := "list"
	resp, r, err := apiClient.DefaultAPI.VesIoSchemaViewsOriginPoolAPIList(context.Background(), metadataNamespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.VesIoSchemaViewsOriginPoolAPIList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	fmt.Fprintf(os.Stdout, "%v\n", r)
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.VesIoSchemaViewsOriginPoolAPIList`: %v\n", resp)
}

func GetOriginPool(configuration *openapiclient.Configuration) {
	apiClient := openapiclient.NewAPIClient(configuration)

	metadataNamespace := "get"
	originPoolName := "originPool_1"
	resp, r, err := apiClient.DefaultAPI.VesIoSchemaViewsOriginPoolAPIGet(context.Background(), metadataNamespace, originPoolName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.VesIoSchemaViewsOriginPoolAPIGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	fmt.Fprintf(os.Stdout, "%v\n", r)
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.VesIoSchemaViewsOriginPoolAPIGet`: %v\n", resp)
}

func ReplaceOriginPool(configuration *openapiclient.Configuration) {
	apiClient := openapiclient.NewAPIClient(configuration)
	metadataNamespace := "replace"
	originPoolName := "originPool1"
	name := "originPool_1"
	description := "origin pool description"
	var threshold int64 = 10
	originpool := openapiclient.OriginPoolReplaceRequest{
		Metadata: &openapiclient.SchemaObjectReplaceMetaType{
			Name:        &name,
			Description: &description,
		},
		Spec: &openapiclient.ViewsoriginPoolReplaceSpecType{
			AdvancedOptions: &openapiclient.OriginPoolOriginPoolAdvancedOptions{
				PanicThreshold: &threshold,
			},
		},
	}
	resp, r, err := apiClient.DefaultAPI.VesIoSchemaViewsOriginPoolAPIReplace(context.Background(), metadataNamespace, originPoolName).OriginPoolReplaceRequest(originpool).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.VesIoSchemaViewsOriginPoolAPIDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	fmt.Fprintf(os.Stdout, "%v\n", r)
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.VesIoSchemaViewsOriginPoolAPIDelete`: %v\n", resp)
}

func DeleteOriginPool(configuration *openapiclient.Configuration) {
	apiClient := openapiclient.NewAPIClient(configuration)

	metadataNamespace := "delete"
	originPoolName := "originPoolName"

	name := "originPool_1"
	originpool := openapiclient.OriginPoolDeleteRequest{
		Name: &name,
	}

	resp, r, err := apiClient.DefaultAPI.VesIoSchemaViewsOriginPoolAPIDelete(context.Background(), metadataNamespace, originPoolName).OriginPoolDeleteRequest(originpool).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.VesIoSchemaViewsOriginPoolAPIDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	fmt.Fprintf(os.Stdout, "%v\n", r)
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.VesIoSchemaViewsOriginPoolAPIDelete`: %v\n", resp)

}
