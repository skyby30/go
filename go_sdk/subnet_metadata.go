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
	"time"
)

// The subnet kind metadata
type SubnetMetadata struct {

	// Categories for the subnet
	Categories map[string]string `json:"categories,omitempty" bson:"categories,omitempty"`

	// Time when subnet was created
	CreationTime time.Time `json:"creation_time,omitempty" bson:"creation_time,omitempty"`

	// Monotonically increasing number
	EntityVersion int64 `json:"entity_version,omitempty" bson:"entity_version,omitempty"`

	// The kind name
	Kind string `json:"kind,omitempty" bson:"kind,omitempty"`

	// Time when subnet was last updated
	LastUpdateTime time.Time `json:"last_update_time,omitempty" bson:"last_update_time,omitempty"`

	// subnet name
	Name string `json:"name,omitempty" bson:"name,omitempty"`

	OwnerReference UserReference `json:"owner_reference,omitempty" bson:"owner_reference,omitempty"`

	// subnet UUID
	Uuid string `json:"uuid,omitempty" bson:"uuid,omitempty"`
}
