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

// Cluster build details.
type BuildInfo struct {

	// Build type, one of {dbg, opt, release}.
	BuildType string `json:"build_type,omitempty" bson:"build_type,omitempty"`

	// Date/time of the last commit.
	CommitDate time.Time `json:"commit_date,omitempty" bson:"commit_date,omitempty"`

	// Last Git commit id which the build is based on.
	CommitId string `json:"commit_id,omitempty" bson:"commit_id,omitempty"`

	// First 6 characters of the last Git commit id.
	ShortCommitId string `json:"short_commit_id,omitempty" bson:"short_commit_id,omitempty"`

	// Version string in format <code_name>-<version_numbers>-<branch_type>, i.e master, danube-4.5.0.2-stable
	Version string `json:"version,omitempty" bson:"version,omitempty"`
}