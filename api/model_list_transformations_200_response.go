/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.1.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ListTransformations200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListTransformations200Response{}

// ListTransformations200Response struct for ListTransformations200Response
type ListTransformations200Response struct {
	Data *ListTransformationsV1Output `json:"data,omitempty"`
}

// NewListTransformations200Response instantiates a new ListTransformations200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransformations200Response() *ListTransformations200Response {
	this := ListTransformations200Response{}
	return &this
}

// NewListTransformations200ResponseWithDefaults instantiates a new ListTransformations200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransformations200ResponseWithDefaults() *ListTransformations200Response {
	this := ListTransformations200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListTransformations200Response) GetData() ListTransformationsV1Output {
	if o == nil || IsNil(o.Data) {
		var ret ListTransformationsV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTransformations200Response) GetDataOk() (*ListTransformationsV1Output, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListTransformations200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ListTransformationsV1Output and assigns it to the Data field.
func (o *ListTransformations200Response) SetData(v ListTransformationsV1Output) {
	o.Data = &v
}

func (o ListTransformations200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListTransformations200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableListTransformations200Response struct {
	value *ListTransformations200Response
	isSet bool
}

func (v NullableListTransformations200Response) Get() *ListTransformations200Response {
	return v.value
}

func (v *NullableListTransformations200Response) Set(val *ListTransformations200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransformations200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransformations200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransformations200Response(
	val *ListTransformations200Response,
) *NullableListTransformations200Response {
	return &NullableListTransformations200Response{value: val, isSet: true}
}

func (v NullableListTransformations200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransformations200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
