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

// All API calls that return a list will have this metadata block
type BatchListMetadata struct {

	// The filter used for the results
	Filter string `json:"filter,omitempty" bson:"filter,omitempty"`

	// The kind name
	Kind string `json:"kind,omitempty" bson:"kind,omitempty"`

	// The number of records to retrieve relative to the offset
	Length int64 `json:"length,omitempty" bson:"length,omitempty"`

	// Offset from the start of the entity list
	Offset int64 `json:"offset,omitempty" bson:"offset,omitempty"`

	// The attribute to perform sort on
	SortAttribute string `json:"sort_attribute,omitempty" bson:"sort_attribute,omitempty"`

	// The sort order in which results are returned
	SortOrder string `json:"sort_order,omitempty" bson:"sort_order,omitempty"`

	// Total matches found for the kind
	TotalMatches int64 `json:"total_matches,omitempty" bson:"total_matches,omitempty"`
}
