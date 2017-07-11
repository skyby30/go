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

// Input for the changed regions query.
type ChangedRegionsQuery struct {

	// The absolute offset in bytes up to which to query for the changed regions. Note that the interval specified by the start_offset together with the end_offset is right half-open. If the end_offset is not specified, the portion from the start_offset till the end of the file will be included in the query.
	EndOffset int64 `json:"end_offset,omitempty" bson:"end_offset,omitempty"`

	// Absolute path of a file within a snapshot that must be used as the reference in the computation of the changed regions. If this path is not specified, then the changed regions will not be computed. Instead, the sparse and the non-sparse regions of the file specified in snapshot_file_path will be returned.
	ReferenceSnapshotFilePath string `json:"reference_snapshot_file_path,omitempty" bson:"reference_snapshot_file_path,omitempty"`

	// Absolute path of a file within a snapshot of an entity such as a virtual machine, a volume group, or a protection domain.
	SnapshotFilePath string `json:"snapshot_file_path,omitempty" bson:"snapshot_file_path,omitempty"`

	// The absolute offset in bytes from where to query for the changed regions.
	StartOffset int64 `json:"start_offset,omitempty" bson:"start_offset,omitempty"`
}