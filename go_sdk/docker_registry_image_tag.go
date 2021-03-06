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

// Docker Hub tagged image
type DockerRegistryImageTag struct {

	// Last modified date in RFC 3339
	ModifiedDate time.Time `json:"modified_date,omitempty" bson:"modified_date,omitempty"`

	// Image tag name
	Name string `json:"name,omitempty" bson:"name,omitempty"`

	// Size of the image in MiB
	SizeMib int64 `json:"size_mib,omitempty" bson:"size_mib,omitempty"`
}
