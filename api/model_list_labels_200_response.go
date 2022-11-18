/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 33.0.2
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ListLabels200Response struct for ListLabels200Response
type ListLabels200Response struct {
	Data *ListLabelsV1Output `json:"data,omitempty"`
}

// NewListLabels200Response instantiates a new ListLabels200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListLabels200Response() *ListLabels200Response {
	this := ListLabels200Response{}
	return &this
}

// NewListLabels200ResponseWithDefaults instantiates a new ListLabels200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListLabels200ResponseWithDefaults() *ListLabels200Response {
	this := ListLabels200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListLabels200Response) GetData() ListLabelsV1Output {
	if o == nil || o.Data == nil {
		var ret ListLabelsV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLabels200Response) GetDataOk() (*ListLabelsV1Output, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListLabels200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ListLabelsV1Output and assigns it to the Data field.
func (o *ListLabels200Response) SetData(v ListLabelsV1Output) {
	o.Data = &v
}

func (o ListLabels200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListLabels200Response struct {
	value *ListLabels200Response
	isSet bool
}

func (v NullableListLabels200Response) Get() *ListLabels200Response {
	return v.value
}

func (v *NullableListLabels200Response) Set(val *ListLabels200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListLabels200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListLabels200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListLabels200Response(val *ListLabels200Response) *NullableListLabels200Response {
	return &NullableListLabels200Response{value: val, isSet: true}
}

func (v NullableListLabels200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListLabels200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


