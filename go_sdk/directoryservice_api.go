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

type DirectoryserviceApi struct {
	Configuration Configuration
}

func NewDirectoryserviceApi() *DirectoryserviceApi {
	configuration := NewConfiguration()
	return &DirectoryserviceApi{
		Configuration: *configuration,
	}
}

func NewDirectoryserviceApiWithBasePath(basePath string) *DirectoryserviceApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &DirectoryserviceApi{
		Configuration: *configuration,
	}
}

/**
 * Get a list of directory service configurations
 * This operation gets a list of directory service configurations, allowing for sorting, filtering, and pagination. Note: Entities that have not been created successfully are not listed.
 *
 * @param getEntitiesRequest
 * @return *DirectoryServiceListIntentResponse
 */
func (a DirectoryserviceApi) DirectoryServicesListPost(getEntitiesRequest DirectoryServiceListMetadata) (*DirectoryServiceListIntentResponse, *APIResponse, error) {

	var httpMethod = "Post"
	// create path and map variables
	path := a.Configuration.BasePath + "/directory_services/list"

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

	var successPayload = new(DirectoryServiceListIntentResponse)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Create a directory service configuration
 * This operation submits a request to create a directory service configuration based on the input parameters.
 *
 * @param body
 * @return *DirectoryServiceIntentResponse
 */
func (a DirectoryserviceApi) DirectoryServicesPost(body DirectoryServiceIntentInput) (*DirectoryServiceIntentResponse, *APIResponse, error) {

	var httpMethod = "Post"
	// create path and map variables
	path := a.Configuration.BasePath + "/directory_services"

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

	var successPayload = new(DirectoryServiceIntentResponse)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Delete a directory service configuration
 * This operation submits a request to delete a directory service configuration.
 *
 * @param uuid The UUID of the entity
 * @return void
 */
func (a DirectoryserviceApi) DirectoryServicesUuidDelete(uuid string) (*APIResponse, error) {

	var httpMethod = "Delete"
	// create path and map variables
	path := a.Configuration.BasePath + "/directory_services/{uuid}"
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
 * Get a directory service configuration
 * This operation gets a directory service configuration.
 *
 * @param uuid The UUID of the entity
 * @return *DirectoryServiceIntentResponse
 */
func (a DirectoryserviceApi) DirectoryServicesUuidGet(uuid string) (*DirectoryServiceIntentResponse, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/directory_services/{uuid}"
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
	var successPayload = new(DirectoryServiceIntentResponse)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Update a directory service configuration
 * This operation submits a request to update a directory service configuration based on the input parameters.
 *
 * @param uuid The UUID of the entity
 * @param body
 * @return *DirectoryServiceIntentResponse
 */
func (a DirectoryserviceApi) DirectoryServicesUuidPut(uuid string, body DirectoryServiceIntentInput) (*DirectoryServiceIntentResponse, *APIResponse, error) {

	var httpMethod = "Put"
	// create path and map variables
	path := a.Configuration.BasePath + "/directory_services/{uuid}"
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

	var successPayload = new(DirectoryServiceIntentResponse)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Search user or group in the directory service.
 * Retrieves a user or a group available in the directory service based on the UUID specified.
 *
 * @param uuid The UUID of the entity
 * @param body
 * @return *DirectoryServiceSearchResponse
 */
func (a DirectoryserviceApi) DirectoryServicesUuidSearchPost(uuid string, body DirectoryServiceSearchMetadata) (*DirectoryServiceSearchResponse, *APIResponse, error) {

	var httpMethod = "Post"
	// create path and map variables
	path := a.Configuration.BasePath + "/directory_services/{uuid}/search"
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

	var successPayload = new(DirectoryServiceSearchResponse)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}
