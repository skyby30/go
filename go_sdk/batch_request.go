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

// Request for batch operations.
type BatchRequest struct {

	// If the specified parameter is CONTINUE, the remaining APIs in the batch continue to be executed.
	ActionOnFailure string `json:"action_on_failure,omitempty" bson:"action_on_failure,omitempty"`

	// A list of API requests in the batch.
	ApiRequestList []ApiRequest `json:"api_request_list,omitempty" bson:"api_request_list,omitempty"`

	// The current API version.
	ApiVersion string `json:"api_version,omitempty" bson:"api_version,omitempty"`

	// The order of execution of the APIs in the batch. Can be either Sequential (default value) or Parallel.
	ExecutionOrder string `json:"execution_order,omitempty" bson:"execution_order,omitempty"`
}
