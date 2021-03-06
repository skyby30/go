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

// Defines the snapshot schedule
type SnapshotSchedule struct {

	// Duration of the event. If set, an event of duration duration_usecs will repeat as per the recurrence defined in interval_type
	DurationSecs int64 `json:"duration_secs,omitempty" bson:"duration_secs,omitempty"`

	// End time of the snapshot schedule
	EndTime int64 `json:"end_time,omitempty" bson:"end_time,omitempty"`

	// Multiple of interval_type
	IntervalMultiple int64 `json:"interval_multiple,omitempty" bson:"interval_multiple,omitempty"`

	// Type of schedule interval
	IntervalType string `json:"interval_type,omitempty" bson:"interval_type,omitempty"`

	// Whether the snapshot schedule is suspended
	IsSuspended bool `json:"is_suspended,omitempty" bson:"is_suspended,omitempty"`

	// Start time of the snapshot schedule
	StartTime int64 `json:"start_time,omitempty" bson:"start_time,omitempty"`
}
