/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 57.4.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ListReverseETLSyncStatusesFromModelAndSubscriptionId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListReverseETLSyncStatusesFromModelAndSubscriptionId200Response{}

// ListReverseETLSyncStatusesFromModelAndSubscriptionId200Response struct for ListReverseETLSyncStatusesFromModelAndSubscriptionId200Response
type ListReverseETLSyncStatusesFromModelAndSubscriptionId200Response struct {
	Data *ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput `json:"data,omitempty"`
}

// NewListReverseETLSyncStatusesFromModelAndSubscriptionId200Response instantiates a new ListReverseETLSyncStatusesFromModelAndSubscriptionId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListReverseETLSyncStatusesFromModelAndSubscriptionId200Response() *ListReverseETLSyncStatusesFromModelAndSubscriptionId200Response {
	this := ListReverseETLSyncStatusesFromModelAndSubscriptionId200Response{}
	return &this
}

// NewListReverseETLSyncStatusesFromModelAndSubscriptionId200ResponseWithDefaults instantiates a new ListReverseETLSyncStatusesFromModelAndSubscriptionId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListReverseETLSyncStatusesFromModelAndSubscriptionId200ResponseWithDefaults() *ListReverseETLSyncStatusesFromModelAndSubscriptionId200Response {
	this := ListReverseETLSyncStatusesFromModelAndSubscriptionId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListReverseETLSyncStatusesFromModelAndSubscriptionId200Response) GetData() ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput {
	if o == nil || IsNil(o.Data) {
		var ret ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListReverseETLSyncStatusesFromModelAndSubscriptionId200Response) GetDataOk() (*ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListReverseETLSyncStatusesFromModelAndSubscriptionId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput and assigns it to the Data field.
func (o *ListReverseETLSyncStatusesFromModelAndSubscriptionId200Response) SetData(
	v ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput,
) {
	o.Data = &v
}

func (o ListReverseETLSyncStatusesFromModelAndSubscriptionId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListReverseETLSyncStatusesFromModelAndSubscriptionId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableListReverseETLSyncStatusesFromModelAndSubscriptionId200Response struct {
	value *ListReverseETLSyncStatusesFromModelAndSubscriptionId200Response
	isSet bool
}

func (v NullableListReverseETLSyncStatusesFromModelAndSubscriptionId200Response) Get() *ListReverseETLSyncStatusesFromModelAndSubscriptionId200Response {
	return v.value
}

func (v *NullableListReverseETLSyncStatusesFromModelAndSubscriptionId200Response) Set(
	val *ListReverseETLSyncStatusesFromModelAndSubscriptionId200Response,
) {
	v.value = val
	v.isSet = true
}

func (v NullableListReverseETLSyncStatusesFromModelAndSubscriptionId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListReverseETLSyncStatusesFromModelAndSubscriptionId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListReverseETLSyncStatusesFromModelAndSubscriptionId200Response(
	val *ListReverseETLSyncStatusesFromModelAndSubscriptionId200Response,
) *NullableListReverseETLSyncStatusesFromModelAndSubscriptionId200Response {
	return &NullableListReverseETLSyncStatusesFromModelAndSubscriptionId200Response{
		value: val,
		isSet: true,
	}
}

func (v NullableListReverseETLSyncStatusesFromModelAndSubscriptionId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListReverseETLSyncStatusesFromModelAndSubscriptionId200Response) UnmarshalJSON(
	src []byte,
) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
