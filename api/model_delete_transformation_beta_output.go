/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 55.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the DeleteTransformationBetaOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteTransformationBetaOutput{}

// DeleteTransformationBetaOutput The output of delete Transformation.
type DeleteTransformationBetaOutput struct {
	// The operation status.
	Status string `json:"status"`
}

// NewDeleteTransformationBetaOutput instantiates a new DeleteTransformationBetaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteTransformationBetaOutput(status string) *DeleteTransformationBetaOutput {
	this := DeleteTransformationBetaOutput{}
	this.Status = status
	return &this
}

// NewDeleteTransformationBetaOutputWithDefaults instantiates a new DeleteTransformationBetaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteTransformationBetaOutputWithDefaults() *DeleteTransformationBetaOutput {
	this := DeleteTransformationBetaOutput{}
	return &this
}

// GetStatus returns the Status field value
func (o *DeleteTransformationBetaOutput) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DeleteTransformationBetaOutput) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DeleteTransformationBetaOutput) SetStatus(v string) {
	o.Status = v
}

func (o DeleteTransformationBetaOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteTransformationBetaOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableDeleteTransformationBetaOutput struct {
	value *DeleteTransformationBetaOutput
	isSet bool
}

func (v NullableDeleteTransformationBetaOutput) Get() *DeleteTransformationBetaOutput {
	return v.value
}

func (v *NullableDeleteTransformationBetaOutput) Set(val *DeleteTransformationBetaOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteTransformationBetaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteTransformationBetaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteTransformationBetaOutput(
	val *DeleteTransformationBetaOutput,
) *NullableDeleteTransformationBetaOutput {
	return &NullableDeleteTransformationBetaOutput{value: val, isSet: true}
}

func (v NullableDeleteTransformationBetaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteTransformationBetaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
