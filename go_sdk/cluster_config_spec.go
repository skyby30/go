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

// Cluster Configuration.
type ClusterConfigSpec struct {

	// List of valid ssh keys for the cluster.
	AuthorizedPublicKeyList []PublicKey `json:"authorized_public_key_list,omitempty" bson:"authorized_public_key_list,omitempty"`

	CertificationSigningInfo CertificationSigningInfo `json:"certification_signing_info,omitempty" bson:"certification_signing_info,omitempty"`

	ClientAuth ClientAuth `json:"client_auth,omitempty" bson:"client_auth,omitempty"`

	// Array of enabled features.
	EnabledFeatureList []string `json:"enabled_feature_list,omitempty" bson:"enabled_feature_list,omitempty"`

	// Cluster encryption status.
	EncryptionStatus string `json:"encryption_status,omitempty" bson:"encryption_status,omitempty"`

	// Cluster supported redundancy factor. Default is 2.
	RedundancyFactor int64 `json:"redundancy_factor,omitempty" bson:"redundancy_factor,omitempty"`

	// Map of software on the cluster with software type as the key.
	SoftwareMap map[string]ClusterSoftware `json:"software_map,omitempty" bson:"software_map,omitempty"`

	// Verbosity level settings for populating support information. - 'Nothing': Send nothing - 'Basic': Send basic information - skip core dump and hypervisor            stats information - 'BasicPlusCoreDump': Send basic and core dump information - 'All': Send all information
	SupportedInformationVerbosity string `json:"supported_information_verbosity,omitempty" bson:"supported_information_verbosity,omitempty"`

	// Zone name used in value of TZ environment variable.
	Timezone string `json:"timezone,omitempty" bson:"timezone,omitempty"`
}
