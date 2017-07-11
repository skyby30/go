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

// The status of a REST API call; only used when there is a failure to report.
type ProjectStatus struct {
	ApiVersion string `json:"api_version,omitempty" bson:"api_version,omitempty"`

	// The HTTP error code
	Code int64 `json:"code,omitempty" bson:"code,omitempty"`

	// The kind name
	Kind string `json:"kind,omitempty" bson:"kind,omitempty"`

	MessageList []MessageResource `json:"message_list,omitempty" bson:"message_list,omitempty"`

	State string `json:"state,omitempty" bson:"state,omitempty"`
}
