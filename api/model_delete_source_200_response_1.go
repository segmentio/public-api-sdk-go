/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 55.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the DeleteSource200Response1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteSource200Response1{}

// DeleteSource200Response1 struct for DeleteSource200Response1
type DeleteSource200Response1 struct {
	Data *DeleteSourceAlphaOutput `json:"data,omitempty"`
}

// NewDeleteSource200Response1 instantiates a new DeleteSource200Response1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteSource200Response1() *DeleteSource200Response1 {
	this := DeleteSource200Response1{}
	return &this
}

// NewDeleteSource200Response1WithDefaults instantiates a new DeleteSource200Response1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteSource200Response1WithDefaults() *DeleteSource200Response1 {
	this := DeleteSource200Response1{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DeleteSource200Response1) GetData() DeleteSourceAlphaOutput {
	if o == nil || IsNil(o.Data) {
		var ret DeleteSourceAlphaOutput
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteSource200Response1) GetDataOk() (*DeleteSourceAlphaOutput, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DeleteSource200Response1) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given DeleteSourceAlphaOutput and assigns it to the Data field.
func (o *DeleteSource200Response1) SetData(v DeleteSourceAlphaOutput) {
	o.Data = &v
}

func (o DeleteSource200Response1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteSource200Response1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableDeleteSource200Response1 struct {
	value *DeleteSource200Response1
	isSet bool
}

func (v NullableDeleteSource200Response1) Get() *DeleteSource200Response1 {
	return v.value
}

func (v *NullableDeleteSource200Response1) Set(val *DeleteSource200Response1) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteSource200Response1) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteSource200Response1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteSource200Response1(
	val *DeleteSource200Response1,
) *NullableDeleteSource200Response1 {
	return &NullableDeleteSource200Response1{value: val, isSet: true}
}

func (v NullableDeleteSource200Response1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteSource200Response1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
