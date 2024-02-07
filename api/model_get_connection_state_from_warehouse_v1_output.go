/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 40.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GetConnectionStateFromWarehouseV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetConnectionStateFromWarehouseV1Output{}

// GetConnectionStateFromWarehouseV1Output Returns the status of a Warehouse connection settings after an attempt to connect to it.
type GetConnectionStateFromWarehouseV1Output struct {
	// Represents the status for the current connection settings.
	ConnectionState string `json:"connectionState"`
}

// NewGetConnectionStateFromWarehouseV1Output instantiates a new GetConnectionStateFromWarehouseV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetConnectionStateFromWarehouseV1Output(
	connectionState string,
) *GetConnectionStateFromWarehouseV1Output {
	this := GetConnectionStateFromWarehouseV1Output{}
	this.ConnectionState = connectionState
	return &this
}

// NewGetConnectionStateFromWarehouseV1OutputWithDefaults instantiates a new GetConnectionStateFromWarehouseV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetConnectionStateFromWarehouseV1OutputWithDefaults() *GetConnectionStateFromWarehouseV1Output {
	this := GetConnectionStateFromWarehouseV1Output{}
	return &this
}

// GetConnectionState returns the ConnectionState field value
func (o *GetConnectionStateFromWarehouseV1Output) GetConnectionState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionState
}

// GetConnectionStateOk returns a tuple with the ConnectionState field value
// and a boolean to check if the value has been set.
func (o *GetConnectionStateFromWarehouseV1Output) GetConnectionStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionState, true
}

// SetConnectionState sets field value
func (o *GetConnectionStateFromWarehouseV1Output) SetConnectionState(v string) {
	o.ConnectionState = v
}

func (o GetConnectionStateFromWarehouseV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetConnectionStateFromWarehouseV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connectionState"] = o.ConnectionState
	return toSerialize, nil
}

type NullableGetConnectionStateFromWarehouseV1Output struct {
	value *GetConnectionStateFromWarehouseV1Output
	isSet bool
}

func (v NullableGetConnectionStateFromWarehouseV1Output) Get() *GetConnectionStateFromWarehouseV1Output {
	return v.value
}

func (v *NullableGetConnectionStateFromWarehouseV1Output) Set(
	val *GetConnectionStateFromWarehouseV1Output,
) {
	v.value = val
	v.isSet = true
}

func (v NullableGetConnectionStateFromWarehouseV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableGetConnectionStateFromWarehouseV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetConnectionStateFromWarehouseV1Output(
	val *GetConnectionStateFromWarehouseV1Output,
) *NullableGetConnectionStateFromWarehouseV1Output {
	return &NullableGetConnectionStateFromWarehouseV1Output{value: val, isSet: true}
}

func (v NullableGetConnectionStateFromWarehouseV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetConnectionStateFromWarehouseV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
