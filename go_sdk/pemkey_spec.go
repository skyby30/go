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

// SSL key
type PemkeySpec struct {
	CaChain string `json:"ca_chain,omitempty" bson:"ca_chain,omitempty"`

	Cert string `json:"cert,omitempty" bson:"cert,omitempty"`

	Key string `json:"key,omitempty" bson:"key,omitempty"`

	Name string `json:"name,omitempty" bson:"name,omitempty"`

	// SSL key type. Key types with RSA_2048, ECDSA_256 and ECDSA_384 are supported for key generation and importing.
	Type_ string `json:"type,omitempty" bson:"type,omitempty"`
}
