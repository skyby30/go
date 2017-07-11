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

// Source spec that describes an entity like a VM, Image that should be translated into an AppBlueprint
type AppBlueprintRenderInput struct {

	// A description for the rendered app blueprint
	Description string `json:"description,omitempty" bson:"description,omitempty"`

	InputSpec AppBlueprintRenderInputInputSpec `json:"input_spec,omitempty" bson:"input_spec,omitempty"`

	InputType string `json:"input_type,omitempty" bson:"input_type,omitempty"`

	// Name of the rendered app blueprint
	Name string `json:"name,omitempty" bson:"name,omitempty"`
}