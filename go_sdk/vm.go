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

// VM Input Definition.
type Vm struct {

	// Reference to the cluster where this VM exists or needs to be migrated to. This is to support migration of a VM from one cluster to another cluster.
	ClusterReference ClusterReference `json:"cluster_reference,omitempty" bson:"cluster_reference,omitempty"`

	// A description or user annotation for the VM.
	Description string `json:"description,omitempty" bson:"description,omitempty"`

	// VM Name.
	Name string `json:"name,omitempty" bson:"name,omitempty"`

	Resources VmResources `json:"resources,omitempty" bson:"resources,omitempty"`
}
