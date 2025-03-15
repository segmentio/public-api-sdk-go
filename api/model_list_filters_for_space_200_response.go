/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.0.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ListFiltersForSpace200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListFiltersForSpace200Response{}

// ListFiltersForSpace200Response struct for ListFiltersForSpace200Response
type ListFiltersForSpace200Response struct {
	Data *ListFiltersForSpaceOutput `json:"data,omitempty"`
}

// NewListFiltersForSpace200Response instantiates a new ListFiltersForSpace200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListFiltersForSpace200Response() *ListFiltersForSpace200Response {
	this := ListFiltersForSpace200Response{}
	return &this
}

// NewListFiltersForSpace200ResponseWithDefaults instantiates a new ListFiltersForSpace200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListFiltersForSpace200ResponseWithDefaults() *ListFiltersForSpace200Response {
	this := ListFiltersForSpace200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListFiltersForSpace200Response) GetData() ListFiltersForSpaceOutput {
	if o == nil || IsNil(o.Data) {
		var ret ListFiltersForSpaceOutput
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFiltersForSpace200Response) GetDataOk() (*ListFiltersForSpaceOutput, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListFiltersForSpace200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ListFiltersForSpaceOutput and assigns it to the Data field.
func (o *ListFiltersForSpace200Response) SetData(v ListFiltersForSpaceOutput) {
	o.Data = &v
}

func (o ListFiltersForSpace200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListFiltersForSpace200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableListFiltersForSpace200Response struct {
	value *ListFiltersForSpace200Response
	isSet bool
}

func (v NullableListFiltersForSpace200Response) Get() *ListFiltersForSpace200Response {
	return v.value
}

func (v *NullableListFiltersForSpace200Response) Set(val *ListFiltersForSpace200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListFiltersForSpace200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListFiltersForSpace200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListFiltersForSpace200Response(
	val *ListFiltersForSpace200Response,
) *NullableListFiltersForSpace200Response {
	return &NullableListFiltersForSpace200Response{value: val, isSet: true}
}

func (v NullableListFiltersForSpace200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListFiltersForSpace200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
