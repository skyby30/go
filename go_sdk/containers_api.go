/*
 * Nutanix Intentful API
 *
 * Move programming from the user to the machine
 *
 * OpenAPI spec version: 3.0.0
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package nutanix

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

type ContainersApi struct {
	Configuration Configuration
}

func NewContainersApi() *ContainersApi {
	configuration := NewConfiguration()
	return &ContainersApi{
		Configuration: *configuration,
	}
}

func NewContainersApiWithBasePath(basePath string) *ContainersApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &ContainersApi{
		Configuration: *configuration,
	}
}

/**
 * Get a list of Containers
 * This operation gets a list of Containers, allowing for sorting, filtering, and pagination. Note: Entities that have not been created successfully are not listed.
 *
 * @param getEntitiesRequest
 * @return *ContainerListIntentResponse
 */
func (a ContainersApi) ContainersListPost(getEntitiesRequest ContainerListMetadata) (*ContainerListIntentResponse, *APIResponse, error) {

	var httpMethod = "Post"
	// create path and map variables
	path := a.Configuration.BasePath + "/containers/list"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication (basicAuth) required

	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != "" {
		headerParams["Authorization"] = "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	switch reflect.TypeOf(getEntitiesRequest) {
	case reflect.TypeOf(""):
		postBody = getEntitiesRequest
	default:
		postBody = &getEntitiesRequest
	}

	var successPayload = new(ContainerListIntentResponse)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Create a Container
 * This operation submits a request to create a Container based on the input parameters.
 *
 * @param body
 * @return *ContainerIntentResponse
 */
func (a ContainersApi) ContainersPost(body ContainerIntentInput) (*ContainerIntentResponse, *APIResponse, error) {

	var httpMethod = "Post"
	// create path and map variables
	path := a.Configuration.BasePath + "/containers"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication (basicAuth) required

	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != "" {
		headerParams["Authorization"] = "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	switch reflect.TypeOf(body) {
	case reflect.TypeOf(""):
		postBody = body
	default:
		postBody = &body
	}

	var successPayload = new(ContainerIntentResponse)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Delete a Container
 * This operation submits a request to delete a Container.
 *
 * @param uuid The UUID of the entity
 * @return void
 */
func (a ContainersApi) ContainersUuidDelete(uuid string) (*APIResponse, error) {

	var httpMethod = "Delete"
	// create path and map variables
	path := a.Configuration.BasePath + "/containers/{uuid}"
	path = strings.Replace(path, "{"+"uuid"+"}", fmt.Sprintf("%v", uuid), -1)

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication (basicAuth) required

	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != "" {
		headerParams["Authorization"] = "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}

	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return NewAPIResponse(httpResponse.RawResponse), err
	}

	return NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Get a Container
 * This operation gets a Container.
 *
 * @param uuid The UUID of the entity
 * @return *ContainerIntentResponse
 */
func (a ContainersApi) ContainersUuidGet(uuid string) (*ContainerIntentResponse, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/containers/{uuid}"
	path = strings.Replace(path, "{"+"uuid"+"}", fmt.Sprintf("%v", uuid), -1)

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication (basicAuth) required

	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != "" {
		headerParams["Authorization"] = "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(ContainerIntentResponse)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Update a Container
 * This operation submits a request to update a Container based on the input parameters. Only state change of container is allowed.
 *
 * @param uuid The UUID of the entity
 * @param body Intent Spec for Container.
 * @return *ContainerIntentResponse
 */
func (a ContainersApi) ContainersUuidPut(uuid string, body ContainerIntentInput) (*ContainerIntentResponse, *APIResponse, error) {

	var httpMethod = "Put"
	// create path and map variables
	path := a.Configuration.BasePath + "/containers/{uuid}"
	path = strings.Replace(path, "{"+"uuid"+"}", fmt.Sprintf("%v", uuid), -1)

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication (basicAuth) required

	// http basic authentication required
	if a.Configuration.Username != "" || a.Configuration.Password != "" {
		headerParams["Authorization"] = "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	switch reflect.TypeOf(body) {
	case reflect.TypeOf(""):
		postBody = body
	default:
		postBody = &body
	}

	var successPayload = new(ContainerIntentResponse)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}