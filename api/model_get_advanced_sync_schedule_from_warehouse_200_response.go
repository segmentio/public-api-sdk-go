/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 53.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GetAdvancedSyncScheduleFromWarehouse200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAdvancedSyncScheduleFromWarehouse200Response{}

// GetAdvancedSyncScheduleFromWarehouse200Response struct for GetAdvancedSyncScheduleFromWarehouse200Response
type GetAdvancedSyncScheduleFromWarehouse200Response struct {
	Data *GetAdvancedSyncScheduleFromWarehouseV1Output `json:"data,omitempty"`
}

// NewGetAdvancedSyncScheduleFromWarehouse200Response instantiates a new GetAdvancedSyncScheduleFromWarehouse200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAdvancedSyncScheduleFromWarehouse200Response() *GetAdvancedSyncScheduleFromWarehouse200Response {
	this := GetAdvancedSyncScheduleFromWarehouse200Response{}
	return &this
}

// NewGetAdvancedSyncScheduleFromWarehouse200ResponseWithDefaults instantiates a new GetAdvancedSyncScheduleFromWarehouse200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAdvancedSyncScheduleFromWarehouse200ResponseWithDefaults() *GetAdvancedSyncScheduleFromWarehouse200Response {
	this := GetAdvancedSyncScheduleFromWarehouse200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetAdvancedSyncScheduleFromWarehouse200Response) GetData() GetAdvancedSyncScheduleFromWarehouseV1Output {
	if o == nil || IsNil(o.Data) {
		var ret GetAdvancedSyncScheduleFromWarehouseV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdvancedSyncScheduleFromWarehouse200Response) GetDataOk() (*GetAdvancedSyncScheduleFromWarehouseV1Output, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetAdvancedSyncScheduleFromWarehouse200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GetAdvancedSyncScheduleFromWarehouseV1Output and assigns it to the Data field.
func (o *GetAdvancedSyncScheduleFromWarehouse200Response) SetData(
	v GetAdvancedSyncScheduleFromWarehouseV1Output,
) {
	o.Data = &v
}

func (o GetAdvancedSyncScheduleFromWarehouse200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAdvancedSyncScheduleFromWarehouse200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetAdvancedSyncScheduleFromWarehouse200Response struct {
	value *GetAdvancedSyncScheduleFromWarehouse200Response
	isSet bool
}

func (v NullableGetAdvancedSyncScheduleFromWarehouse200Response) Get() *GetAdvancedSyncScheduleFromWarehouse200Response {
	return v.value
}

func (v *NullableGetAdvancedSyncScheduleFromWarehouse200Response) Set(
	val *GetAdvancedSyncScheduleFromWarehouse200Response,
) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAdvancedSyncScheduleFromWarehouse200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAdvancedSyncScheduleFromWarehouse200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAdvancedSyncScheduleFromWarehouse200Response(
	val *GetAdvancedSyncScheduleFromWarehouse200Response,
) *NullableGetAdvancedSyncScheduleFromWarehouse200Response {
	return &NullableGetAdvancedSyncScheduleFromWarehouse200Response{value: val, isSet: true}
}

func (v NullableGetAdvancedSyncScheduleFromWarehouse200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAdvancedSyncScheduleFromWarehouse200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
