/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.3.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GetDailyWorkspaceAPICallsUsage200Response struct for GetDailyWorkspaceAPICallsUsage200Response
type GetDailyWorkspaceAPICallsUsage200Response struct {
	Data *GetDailyWorkspaceAPICallsUsageV1Output `json:"data,omitempty"`
}

// NewGetDailyWorkspaceAPICallsUsage200Response instantiates a new GetDailyWorkspaceAPICallsUsage200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDailyWorkspaceAPICallsUsage200Response() *GetDailyWorkspaceAPICallsUsage200Response {
	this := GetDailyWorkspaceAPICallsUsage200Response{}
	return &this
}

// NewGetDailyWorkspaceAPICallsUsage200ResponseWithDefaults instantiates a new GetDailyWorkspaceAPICallsUsage200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDailyWorkspaceAPICallsUsage200ResponseWithDefaults() *GetDailyWorkspaceAPICallsUsage200Response {
	this := GetDailyWorkspaceAPICallsUsage200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetDailyWorkspaceAPICallsUsage200Response) GetData() GetDailyWorkspaceAPICallsUsageV1Output {
	if o == nil || o.Data == nil {
		var ret GetDailyWorkspaceAPICallsUsageV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDailyWorkspaceAPICallsUsage200Response) GetDataOk() (*GetDailyWorkspaceAPICallsUsageV1Output, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetDailyWorkspaceAPICallsUsage200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GetDailyWorkspaceAPICallsUsageV1Output and assigns it to the Data field.
func (o *GetDailyWorkspaceAPICallsUsage200Response) SetData(
	v GetDailyWorkspaceAPICallsUsageV1Output,
) {
	o.Data = &v
}

func (o GetDailyWorkspaceAPICallsUsage200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetDailyWorkspaceAPICallsUsage200Response struct {
	value *GetDailyWorkspaceAPICallsUsage200Response
	isSet bool
}

func (v NullableGetDailyWorkspaceAPICallsUsage200Response) Get() *GetDailyWorkspaceAPICallsUsage200Response {
	return v.value
}

func (v *NullableGetDailyWorkspaceAPICallsUsage200Response) Set(
	val *GetDailyWorkspaceAPICallsUsage200Response,
) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDailyWorkspaceAPICallsUsage200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDailyWorkspaceAPICallsUsage200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDailyWorkspaceAPICallsUsage200Response(
	val *GetDailyWorkspaceAPICallsUsage200Response,
) *NullableGetDailyWorkspaceAPICallsUsage200Response {
	return &NullableGetDailyWorkspaceAPICallsUsage200Response{value: val, isSet: true}
}

func (v NullableGetDailyWorkspaceAPICallsUsage200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDailyWorkspaceAPICallsUsage200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
