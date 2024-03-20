/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 46.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the CreateLabel201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLabel201Response{}

// CreateLabel201Response struct for CreateLabel201Response
type CreateLabel201Response struct {
	Data *CreateLabelV1Output `json:"data,omitempty"`
}

// NewCreateLabel201Response instantiates a new CreateLabel201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLabel201Response() *CreateLabel201Response {
	this := CreateLabel201Response{}
	return &this
}

// NewCreateLabel201ResponseWithDefaults instantiates a new CreateLabel201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLabel201ResponseWithDefaults() *CreateLabel201Response {
	this := CreateLabel201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateLabel201Response) GetData() CreateLabelV1Output {
	if o == nil || IsNil(o.Data) {
		var ret CreateLabelV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLabel201Response) GetDataOk() (*CreateLabelV1Output, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateLabel201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CreateLabelV1Output and assigns it to the Data field.
func (o *CreateLabel201Response) SetData(v CreateLabelV1Output) {
	o.Data = &v
}

func (o CreateLabel201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateLabel201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCreateLabel201Response struct {
	value *CreateLabel201Response
	isSet bool
}

func (v NullableCreateLabel201Response) Get() *CreateLabel201Response {
	return v.value
}

func (v *NullableCreateLabel201Response) Set(val *CreateLabel201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLabel201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLabel201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLabel201Response(
	val *CreateLabel201Response,
) *NullableCreateLabel201Response {
	return &NullableCreateLabel201Response{value: val, isSet: true}
}

func (v NullableCreateLabel201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLabel201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
