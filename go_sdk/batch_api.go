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
	"net/url"
	"reflect"
)

type BatchApi struct {
	Configuration Configuration
}

func NewBatchApi() *BatchApi {
	configuration := NewConfiguration()
	return &BatchApi{
		Configuration: *configuration,
	}
}

func NewBatchApiWithBasePath(basePath string) *BatchApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &BatchApi{
		Configuration: *configuration,
	}
}

/**
 * Submit a list of one or more intentful REST APIs to be processed
 * Batching allows for instructions for several operations to be sent using a single HTTP request. Depending on the batch parameters, the Nutanix v3 gateway processes each independent operation sequentially or in parallel. Once all operations in the batch have been completed, a consolidated response is returned and the HTTP connection is closed. The batch API takes an array of logical HTTP requests represented as JSON arrays. Each request comprises the following: - A method (corresponding to HTTP methods such as GET, PUT, and POST) - A relative URL (relative_url) - (Optional) A body (for POST and PUT requests). The batch API returns an array of logical HTTP responses represented as JSON arrays containing the following: - A status code - (Optional) A body represented as a JSON-encoded string
 *
 * @param intentList List of intent APIs
 * @return *BatchResponse
 */
func (a BatchApi) BatchPost(intentList BatchRequest) (*BatchResponse, *APIResponse, error) {

	var httpMethod = "Post"
	// create path and map variables
	path := a.Configuration.BasePath + "/batch"

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
	switch reflect.TypeOf(intentList) {
	case reflect.TypeOf(""):
		postBody = intentList
	default:
		postBody = &intentList
	}

	var successPayload = new(BatchResponse)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}
