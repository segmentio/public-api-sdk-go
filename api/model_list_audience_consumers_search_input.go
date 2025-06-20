/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.6.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ListAudienceConsumersSearchInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAudienceConsumersSearchInput{}

// ListAudienceConsumersSearchInput Search criteria input for list audience consumers.
type ListAudienceConsumersSearchInput struct {
	// Field to filter by.
	Type string `json:"type"`
	// Text to match the field value.
	Query string `json:"query"`
}

// NewListAudienceConsumersSearchInput instantiates a new ListAudienceConsumersSearchInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAudienceConsumersSearchInput(
	type_ string,
	query string,
) *ListAudienceConsumersSearchInput {
	this := ListAudienceConsumersSearchInput{}
	this.Type = type_
	this.Query = query
	return &this
}

// NewListAudienceConsumersSearchInputWithDefaults instantiates a new ListAudienceConsumersSearchInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAudienceConsumersSearchInputWithDefaults() *ListAudienceConsumersSearchInput {
	this := ListAudienceConsumersSearchInput{}
	return &this
}

// GetType returns the Type field value
func (o *ListAudienceConsumersSearchInput) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ListAudienceConsumersSearchInput) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ListAudienceConsumersSearchInput) SetType(v string) {
	o.Type = v
}

// GetQuery returns the Query field value
func (o *ListAudienceConsumersSearchInput) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *ListAudienceConsumersSearchInput) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *ListAudienceConsumersSearchInput) SetQuery(v string) {
	o.Query = v
}

func (o ListAudienceConsumersSearchInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListAudienceConsumersSearchInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["query"] = o.Query
	return toSerialize, nil
}

type NullableListAudienceConsumersSearchInput struct {
	value *ListAudienceConsumersSearchInput
	isSet bool
}

func (v NullableListAudienceConsumersSearchInput) Get() *ListAudienceConsumersSearchInput {
	return v.value
}

func (v *NullableListAudienceConsumersSearchInput) Set(val *ListAudienceConsumersSearchInput) {
	v.value = val
	v.isSet = true
}

func (v NullableListAudienceConsumersSearchInput) IsSet() bool {
	return v.isSet
}

func (v *NullableListAudienceConsumersSearchInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAudienceConsumersSearchInput(
	val *ListAudienceConsumersSearchInput,
) *NullableListAudienceConsumersSearchInput {
	return &NullableListAudienceConsumersSearchInput{value: val, isSet: true}
}

func (v NullableListAudienceConsumersSearchInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAudienceConsumersSearchInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
