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

// Virtual Machine NIC Status.
type VmNicOutputStatus struct {

	// IP endpoints for the adapter. Currently, IPv4 addresses are supported.
	IpEndpointList []IpAddress `json:"ip_endpoint_list,omitempty" bson:"ip_endpoint_list,omitempty"`

	// The MAC address for the adapter.
	MacAddress string `json:"mac_address,omitempty" bson:"mac_address,omitempty"`

	NetworkFunctionChainReference NetworkFunctionChainReference `json:"network_function_chain_reference,omitempty" bson:"network_function_chain_reference,omitempty"`

	// The type of this Network function NIC. Defaults to INGRESS.
	NetworkFunctionNicType string `json:"network_function_nic_type,omitempty" bson:"network_function_nic_type,omitempty"`

	// The type of this NIC. Defaults to NORMAL_NIC.
	NicType string `json:"nic_type,omitempty" bson:"nic_type,omitempty"`

	SubnetReference Reference `json:"subnet_reference,omitempty" bson:"subnet_reference,omitempty"`
}
