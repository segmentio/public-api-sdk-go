/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 49.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the CreateProfilesWarehouse201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateProfilesWarehouse201Response{}

// CreateProfilesWarehouse201Response struct for CreateProfilesWarehouse201Response
type CreateProfilesWarehouse201Response struct {
	Data *CreateProfilesWarehouseAlphaOutput `json:"data,omitempty"`
}

// NewCreateProfilesWarehouse201Response instantiates a new CreateProfilesWarehouse201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProfilesWarehouse201Response() *CreateProfilesWarehouse201Response {
	this := CreateProfilesWarehouse201Response{}
	return &this
}

// NewCreateProfilesWarehouse201ResponseWithDefaults instantiates a new CreateProfilesWarehouse201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProfilesWarehouse201ResponseWithDefaults() *CreateProfilesWarehouse201Response {
	this := CreateProfilesWarehouse201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateProfilesWarehouse201Response) GetData() CreateProfilesWarehouseAlphaOutput {
	if o == nil || IsNil(o.Data) {
		var ret CreateProfilesWarehouseAlphaOutput
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProfilesWarehouse201Response) GetDataOk() (*CreateProfilesWarehouseAlphaOutput, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateProfilesWarehouse201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CreateProfilesWarehouseAlphaOutput and assigns it to the Data field.
func (o *CreateProfilesWarehouse201Response) SetData(v CreateProfilesWarehouseAlphaOutput) {
	o.Data = &v
}

func (o CreateProfilesWarehouse201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateProfilesWarehouse201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCreateProfilesWarehouse201Response struct {
	value *CreateProfilesWarehouse201Response
	isSet bool
}

func (v NullableCreateProfilesWarehouse201Response) Get() *CreateProfilesWarehouse201Response {
	return v.value
}

func (v *NullableCreateProfilesWarehouse201Response) Set(val *CreateProfilesWarehouse201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProfilesWarehouse201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProfilesWarehouse201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProfilesWarehouse201Response(
	val *CreateProfilesWarehouse201Response,
) *NullableCreateProfilesWarehouse201Response {
	return &NullableCreateProfilesWarehouse201Response{value: val, isSet: true}
}

func (v NullableCreateProfilesWarehouse201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProfilesWarehouse201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
