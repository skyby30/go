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

type RoleApi struct {
	Configuration Configuration
}

func NewRoleApi() *RoleApi {
	configuration := NewConfiguration()
	return &RoleApi{
		Configuration: *configuration,
	}
}

func NewRoleApiWithBasePath(basePath string) *RoleApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &RoleApi{
		Configuration: *configuration,
	}
}

/**
 * List the roles.
 * Get roles. Supported filters are - uuid - name
 *
 * @param getEntitiesRequest
 * @return *RoleListIntentResponse
 */
func (a RoleApi) RolesListPost(getEntitiesRequest RoleListMetadata) (*RoleListIntentResponse, *APIResponse, error) {

	var httpMethod = "Post"
	// create path and map variables
	path := a.Configuration.BasePath + "/roles/list"

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

	var successPayload = new(RoleListIntentResponse)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Create a role.
 * \&quot;A role is a collection of permissions defined for one or more kinds. A kind represents the type of an entity (such as VM). Roles are defined by users who have permission to create roles and assign roles to projects. All users in a project inherit the role.\&quot;
 *
 * @param body
 * @return *RoleIntentResponse
 */
func (a RoleApi) RolesPost(body RoleIntentInput) (*RoleIntentResponse, *APIResponse, error) {

	var httpMethod = "Post"
	// create path and map variables
	path := a.Configuration.BasePath + "/roles"

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

	var successPayload = new(RoleIntentResponse)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Delete a role.
 * Delete a role.
 *
 * @param uuid The UUID of the entity
 * @return void
 */
func (a RoleApi) RolesUuidDelete(uuid string) (*APIResponse, error) {

	var httpMethod = "Delete"
	// create path and map variables
	path := a.Configuration.BasePath + "/roles/{uuid}"
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
 * Get a role.
 * Get a role.
 *
 * @param uuid The UUID of the entity
 * @return *RoleIntentResponse
 */
func (a RoleApi) RolesUuidGet(uuid string) (*RoleIntentResponse, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/roles/{uuid}"
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
	var successPayload = new(RoleIntentResponse)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Update a role.
 * Update a role.
 *
 * @param uuid The UUID of the entity
 * @param body
 * @return *RoleIntentResponse
 */
func (a RoleApi) RolesUuidPut(uuid string, body RoleIntentInput) (*RoleIntentResponse, *APIResponse, error) {

	var httpMethod = "Put"
	// create path and map variables
	path := a.Configuration.BasePath + "/roles/{uuid}"
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

	var successPayload = new(RoleIntentResponse)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}
