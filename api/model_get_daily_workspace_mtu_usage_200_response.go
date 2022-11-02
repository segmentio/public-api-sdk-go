/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 32.0.2
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GetDailyWorkspaceMTUUsage200Response struct for GetDailyWorkspaceMTUUsage200Response
type GetDailyWorkspaceMTUUsage200Response struct {
	Data *GetDailyWorkspaceMTUUsageV1Output `json:"data,omitempty"`
}

// NewGetDailyWorkspaceMTUUsage200Response instantiates a new GetDailyWorkspaceMTUUsage200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDailyWorkspaceMTUUsage200Response() *GetDailyWorkspaceMTUUsage200Response {
	this := GetDailyWorkspaceMTUUsage200Response{}
	return &this
}

// NewGetDailyWorkspaceMTUUsage200ResponseWithDefaults instantiates a new GetDailyWorkspaceMTUUsage200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDailyWorkspaceMTUUsage200ResponseWithDefaults() *GetDailyWorkspaceMTUUsage200Response {
	this := GetDailyWorkspaceMTUUsage200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetDailyWorkspaceMTUUsage200Response) GetData() GetDailyWorkspaceMTUUsageV1Output {
	if o == nil || o.Data == nil {
		var ret GetDailyWorkspaceMTUUsageV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDailyWorkspaceMTUUsage200Response) GetDataOk() (*GetDailyWorkspaceMTUUsageV1Output, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetDailyWorkspaceMTUUsage200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GetDailyWorkspaceMTUUsageV1Output and assigns it to the Data field.
func (o *GetDailyWorkspaceMTUUsage200Response) SetData(v GetDailyWorkspaceMTUUsageV1Output) {
	o.Data = &v
}

func (o GetDailyWorkspaceMTUUsage200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetDailyWorkspaceMTUUsage200Response struct {
	value *GetDailyWorkspaceMTUUsage200Response
	isSet bool
}

func (v NullableGetDailyWorkspaceMTUUsage200Response) Get() *GetDailyWorkspaceMTUUsage200Response {
	return v.value
}

func (v *NullableGetDailyWorkspaceMTUUsage200Response) Set(val *GetDailyWorkspaceMTUUsage200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDailyWorkspaceMTUUsage200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDailyWorkspaceMTUUsage200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDailyWorkspaceMTUUsage200Response(val *GetDailyWorkspaceMTUUsage200Response) *NullableGetDailyWorkspaceMTUUsage200Response {
	return &NullableGetDailyWorkspaceMTUUsage200Response{value: val, isSet: true}
}

func (v NullableGetDailyWorkspaceMTUUsage200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDailyWorkspaceMTUUsage200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


