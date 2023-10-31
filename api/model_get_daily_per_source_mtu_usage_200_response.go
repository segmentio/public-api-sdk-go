/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 37.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GetDailyPerSourceMTUUsage200Response struct for GetDailyPerSourceMTUUsage200Response
type GetDailyPerSourceMTUUsage200Response struct {
	Data *GetDailyPerSourceMTUUsageV1Output `json:"data,omitempty"`
}

// NewGetDailyPerSourceMTUUsage200Response instantiates a new GetDailyPerSourceMTUUsage200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDailyPerSourceMTUUsage200Response() *GetDailyPerSourceMTUUsage200Response {
	this := GetDailyPerSourceMTUUsage200Response{}
	return &this
}

// NewGetDailyPerSourceMTUUsage200ResponseWithDefaults instantiates a new GetDailyPerSourceMTUUsage200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDailyPerSourceMTUUsage200ResponseWithDefaults() *GetDailyPerSourceMTUUsage200Response {
	this := GetDailyPerSourceMTUUsage200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetDailyPerSourceMTUUsage200Response) GetData() GetDailyPerSourceMTUUsageV1Output {
	if o == nil || o.Data == nil {
		var ret GetDailyPerSourceMTUUsageV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDailyPerSourceMTUUsage200Response) GetDataOk() (*GetDailyPerSourceMTUUsageV1Output, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetDailyPerSourceMTUUsage200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GetDailyPerSourceMTUUsageV1Output and assigns it to the Data field.
func (o *GetDailyPerSourceMTUUsage200Response) SetData(v GetDailyPerSourceMTUUsageV1Output) {
	o.Data = &v
}

func (o GetDailyPerSourceMTUUsage200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetDailyPerSourceMTUUsage200Response struct {
	value *GetDailyPerSourceMTUUsage200Response
	isSet bool
}

func (v NullableGetDailyPerSourceMTUUsage200Response) Get() *GetDailyPerSourceMTUUsage200Response {
	return v.value
}

func (v *NullableGetDailyPerSourceMTUUsage200Response) Set(
	val *GetDailyPerSourceMTUUsage200Response,
) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDailyPerSourceMTUUsage200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDailyPerSourceMTUUsage200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDailyPerSourceMTUUsage200Response(
	val *GetDailyPerSourceMTUUsage200Response,
) *NullableGetDailyPerSourceMTUUsage200Response {
	return &NullableGetDailyPerSourceMTUUsage200Response{value: val, isSet: true}
}

func (v NullableGetDailyPerSourceMTUUsage200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDailyPerSourceMTUUsage200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
