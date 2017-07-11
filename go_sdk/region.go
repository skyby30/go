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

// A contiguous set of offsets.
type Region struct {

	// The length of the region in bytes
	Length int64 `json:"length,omitempty" bson:"length,omitempty"`

	// The byte offset indicating the start of the region
	Offset int64 `json:"offset,omitempty" bson:"offset,omitempty"`

	// The type of the region
	Type_ string `json:"type,omitempty" bson:"type,omitempty"`
}
