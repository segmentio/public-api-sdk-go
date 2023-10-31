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

// ListAudiences200Response struct for ListAudiences200Response
type ListAudiences200Response struct {
	Data *ListAudiencesAlphaOutput `json:"data,omitempty"`
}

// NewListAudiences200Response instantiates a new ListAudiences200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAudiences200Response() *ListAudiences200Response {
	this := ListAudiences200Response{}
	return &this
}

// NewListAudiences200ResponseWithDefaults instantiates a new ListAudiences200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAudiences200ResponseWithDefaults() *ListAudiences200Response {
	this := ListAudiences200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListAudiences200Response) GetData() ListAudiencesAlphaOutput {
	if o == nil || o.Data == nil {
		var ret ListAudiencesAlphaOutput
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAudiences200Response) GetDataOk() (*ListAudiencesAlphaOutput, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListAudiences200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ListAudiencesAlphaOutput and assigns it to the Data field.
func (o *ListAudiences200Response) SetData(v ListAudiencesAlphaOutput) {
	o.Data = &v
}

func (o ListAudiences200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListAudiences200Response struct {
	value *ListAudiences200Response
	isSet bool
}

func (v NullableListAudiences200Response) Get() *ListAudiences200Response {
	return v.value
}

func (v *NullableListAudiences200Response) Set(val *ListAudiences200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListAudiences200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListAudiences200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAudiences200Response(
	val *ListAudiences200Response,
) *NullableListAudiences200Response {
	return &NullableListAudiences200Response{value: val, isSet: true}
}

func (v NullableListAudiences200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAudiences200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
