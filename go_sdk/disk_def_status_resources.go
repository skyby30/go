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

// Disk resources.
type DiskDefStatusResources struct {

	// Disk feature flags - 'CanAddAsNewDisk': Flag to indicate if this disk can be added as    new disk. - 'CanAddAsOldDisk': Flag to indicate if the disk can be added as    old disk. - 'BootDisk': Flag to indicate if its a boot disk. - 'OnlyBootDisk': Flag to indicate if the disk is boot only and    no disk operation to be run on it. - 'SelfEncryptingEnabled': Flag to indicate if the disk has self    encryption enabled. - 'PasswordProtected': Flag to indicate if the disk is password    protected.
	EnabledFeaturesList []string `json:"enabled_features_list,omitempty" bson:"enabled_features_list,omitempty"`

	// Firmware version.
	FirmwareVersion string `json:"firmware_version,omitempty" bson:"firmware_version,omitempty"`

	HostReference Reference `json:"host_reference,omitempty" bson:"host_reference,omitempty"`

	// Disk model.
	Model string `json:"model,omitempty" bson:"model,omitempty"`

	// Mount path.
	MountPath string `json:"mount_path,omitempty" bson:"mount_path,omitempty"`

	// Disk serial number.
	SerialNumber string `json:"serial_number,omitempty" bson:"serial_number,omitempty"`

	// Disk size in Mib.
	SizeMib int64 `json:"size_mib,omitempty" bson:"size_mib,omitempty"`

	// Disk location in a node.
	SlotNumber int64 `json:"slot_number,omitempty" bson:"slot_number,omitempty"`

	// Array of disk states - DataMigrationInitiated: Data Migration Initiated. - MarkedForRemovalButNotDetachable: Marked for removal, data    migration is in progress. - ReadyToDetach: Flag to indicate the disk is detachable. - DataMigrated: Flag to indicate if data migration is completed for    this disk. - MarkedForRemoval: Flag to indicate if the disk is marked for    removal. - Online: Flag to indicate if the disk is online. - Bad: Flag to indicate if the disk is bad. - Mounted: Flag to indicate if the disk is mounted. - UnderDiagnosis: Flag to indicate if the disk is under diagnosis.
	StateList []string `json:"state_list,omitempty" bson:"state_list,omitempty"`

	// Storage pool uuid.
	StoragePoolUuid string `json:"storage_pool_uuid,omitempty" bson:"storage_pool_uuid,omitempty"`

	// Storage tier type.
	StorageTierType string `json:"storage_tier_type,omitempty" bson:"storage_tier_type,omitempty"`

	// Disk vendor.
	Vendor string `json:"vendor,omitempty" bson:"vendor,omitempty"`
}
