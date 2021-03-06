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

// Get Entities Response.
type GroupsGetEntitiesResponse struct {
	EntityType string `json:"entity_type,omitempty" bson:"entity_type,omitempty"`

	FilteredEntityCount int64 `json:"filtered_entity_count,omitempty" bson:"filtered_entity_count,omitempty"`

	GroupResults []GroupsGroupResult `json:"group_results,omitempty" bson:"group_results,omitempty"`

	TotalEntityCount int64 `json:"total_entity_count,omitempty" bson:"total_entity_count,omitempty"`

	TotalGroupCount int64 `json:"total_group_count,omitempty" bson:"total_group_count,omitempty"`
}
