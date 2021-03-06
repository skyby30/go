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

// IP config.
type IpConfig struct {

	// Default gateway IP address.
	DefaultGatewayIp string `json:"default_gateway_ip,omitempty" bson:"default_gateway_ip,omitempty"`

	DhcpOptions DhcpOptions `json:"dhcp_options,omitempty" bson:"dhcp_options,omitempty"`

	DhcpServerAddress Address `json:"dhcp_server_address,omitempty" bson:"dhcp_server_address,omitempty"`

	PoolList []IpPool `json:"pool_list,omitempty" bson:"pool_list,omitempty"`

	PrefixLength int64 `json:"prefix_length,omitempty" bson:"prefix_length,omitempty"`

	// Subnet IP address.
	SubnetIp string `json:"subnet_ip,omitempty" bson:"subnet_ip,omitempty"`
}
