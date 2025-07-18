/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.13.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ListRegulationsFromSource200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListRegulationsFromSource200Response{}

// ListRegulationsFromSource200Response struct for ListRegulationsFromSource200Response
type ListRegulationsFromSource200Response struct {
	Data *ListRegulationsFromSourceV1Output `json:"data,omitempty"`
}

// NewListRegulationsFromSource200Response instantiates a new ListRegulationsFromSource200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListRegulationsFromSource200Response() *ListRegulationsFromSource200Response {
	this := ListRegulationsFromSource200Response{}
	return &this
}

// NewListRegulationsFromSource200ResponseWithDefaults instantiates a new ListRegulationsFromSource200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListRegulationsFromSource200ResponseWithDefaults() *ListRegulationsFromSource200Response {
	this := ListRegulationsFromSource200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListRegulationsFromSource200Response) GetData() ListRegulationsFromSourceV1Output {
	if o == nil || IsNil(o.Data) {
		var ret ListRegulationsFromSourceV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRegulationsFromSource200Response) GetDataOk() (*ListRegulationsFromSourceV1Output, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListRegulationsFromSource200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ListRegulationsFromSourceV1Output and assigns it to the Data field.
func (o *ListRegulationsFromSource200Response) SetData(v ListRegulationsFromSourceV1Output) {
	o.Data = &v
}

func (o ListRegulationsFromSource200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListRegulationsFromSource200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableListRegulationsFromSource200Response struct {
	value *ListRegulationsFromSource200Response
	isSet bool
}

func (v NullableListRegulationsFromSource200Response) Get() *ListRegulationsFromSource200Response {
	return v.value
}

func (v *NullableListRegulationsFromSource200Response) Set(
	val *ListRegulationsFromSource200Response,
) {
	v.value = val
	v.isSet = true
}

func (v NullableListRegulationsFromSource200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListRegulationsFromSource200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRegulationsFromSource200Response(
	val *ListRegulationsFromSource200Response,
) *NullableListRegulationsFromSource200Response {
	return &NullableListRegulationsFromSource200Response{value: val, isSet: true}
}

func (v NullableListRegulationsFromSource200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRegulationsFromSource200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
