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

// Cluster status definition.  A Nutanix cluster is comprised of three or more Nutanix nodes. Each node in the cluster contains memory, CPU, RAM, and storage (SSD/HDD). Each node in the cluster runs standard hypervisor such as VMware vSphere, Microsoft Hyper-V, or AHV. A Controller VM (CVM) runs on each node in the cluster. The CVM enables each node to share local storage from all nodes in the cluster.
type ClusterDefStatus struct {
	MessageList []MessageResource `json:"message_list,omitempty" bson:"message_list,omitempty"`

	// Cluster name.
	Name string `json:"name,omitempty" bson:"name,omitempty"`

	Resources ClusterDefStatusResources `json:"resources,omitempty" bson:"resources,omitempty"`

	// The state of the cluster entity.
	State string `json:"state,omitempty" bson:"state,omitempty"`
}
