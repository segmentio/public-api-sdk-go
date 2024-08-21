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

// checks if the GetReverseETLSyncStatusesBySubscriptionId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetReverseETLSyncStatusesBySubscriptionId200Response{}

// GetReverseETLSyncStatusesBySubscriptionId200Response struct for GetReverseETLSyncStatusesBySubscriptionId200Response
type GetReverseETLSyncStatusesBySubscriptionId200Response struct {
	Data *GetReverseETLSyncStatusesBySubscriptionIdOutput `json:"data,omitempty"`
}

// NewGetReverseETLSyncStatusesBySubscriptionId200Response instantiates a new GetReverseETLSyncStatusesBySubscriptionId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetReverseETLSyncStatusesBySubscriptionId200Response() *GetReverseETLSyncStatusesBySubscriptionId200Response {
	this := GetReverseETLSyncStatusesBySubscriptionId200Response{}
	return &this
}

// NewGetReverseETLSyncStatusesBySubscriptionId200ResponseWithDefaults instantiates a new GetReverseETLSyncStatusesBySubscriptionId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetReverseETLSyncStatusesBySubscriptionId200ResponseWithDefaults() *GetReverseETLSyncStatusesBySubscriptionId200Response {
	this := GetReverseETLSyncStatusesBySubscriptionId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetReverseETLSyncStatusesBySubscriptionId200Response) GetData() GetReverseETLSyncStatusesBySubscriptionIdOutput {
	if o == nil || IsNil(o.Data) {
		var ret GetReverseETLSyncStatusesBySubscriptionIdOutput
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReverseETLSyncStatusesBySubscriptionId200Response) GetDataOk() (*GetReverseETLSyncStatusesBySubscriptionIdOutput, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetReverseETLSyncStatusesBySubscriptionId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GetReverseETLSyncStatusesBySubscriptionIdOutput and assigns it to the Data field.
func (o *GetReverseETLSyncStatusesBySubscriptionId200Response) SetData(
	v GetReverseETLSyncStatusesBySubscriptionIdOutput,
) {
	o.Data = &v
}

func (o GetReverseETLSyncStatusesBySubscriptionId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetReverseETLSyncStatusesBySubscriptionId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetReverseETLSyncStatusesBySubscriptionId200Response struct {
	value *GetReverseETLSyncStatusesBySubscriptionId200Response
	isSet bool
}

func (v NullableGetReverseETLSyncStatusesBySubscriptionId200Response) Get() *GetReverseETLSyncStatusesBySubscriptionId200Response {
	return v.value
}

func (v *NullableGetReverseETLSyncStatusesBySubscriptionId200Response) Set(
	val *GetReverseETLSyncStatusesBySubscriptionId200Response,
) {
	v.value = val
	v.isSet = true
}

func (v NullableGetReverseETLSyncStatusesBySubscriptionId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetReverseETLSyncStatusesBySubscriptionId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetReverseETLSyncStatusesBySubscriptionId200Response(
	val *GetReverseETLSyncStatusesBySubscriptionId200Response,
) *NullableGetReverseETLSyncStatusesBySubscriptionId200Response {
	return &NullableGetReverseETLSyncStatusesBySubscriptionId200Response{value: val, isSet: true}
}

func (v NullableGetReverseETLSyncStatusesBySubscriptionId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetReverseETLSyncStatusesBySubscriptionId200Response) UnmarshalJSON(
	src []byte,
) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
