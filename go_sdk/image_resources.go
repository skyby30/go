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

// Describes the image spec resources object.
type ImageResources struct {

	// Checksum of the image. The checksum is used for image validation if the image has a source specified. For images that do not have their source specified the checksum is generated by the image service.
	Checksum Checksum `json:"checksum,omitempty" bson:"checksum,omitempty"`

	// The type of image.
	ImageType string `json:"image_type,omitempty" bson:"image_type,omitempty"`

	// The source URI points at the location of a the source image which is used to create/update image.
	SourceUri string `json:"source_uri,omitempty" bson:"source_uri,omitempty"`

	// The image version
	Version ImageVersion `json:"version,omitempty" bson:"version,omitempty"`
}