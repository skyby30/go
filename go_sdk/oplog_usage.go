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

// oplog disk usage.
type OplogUsage struct {

	// Oplog disk size used in percentage.
	OplogDiskPct float32 `json:"oplog_disk_pct,omitempty" bson:"oplog_disk_pct,omitempty"`

	// Size of oplog disk in bytes.
	OplogDiskSize int64 `json:"oplog_disk_size,omitempty" bson:"oplog_disk_size,omitempty"`
}
