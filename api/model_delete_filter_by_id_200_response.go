/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 55.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the DeleteFilterById200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteFilterById200Response{}

// DeleteFilterById200Response struct for DeleteFilterById200Response
type DeleteFilterById200Response struct {
	Data *DeleteFilterByIdOutput `json:"data,omitempty"`
}

// NewDeleteFilterById200Response instantiates a new DeleteFilterById200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteFilterById200Response() *DeleteFilterById200Response {
	this := DeleteFilterById200Response{}
	return &this
}

// NewDeleteFilterById200ResponseWithDefaults instantiates a new DeleteFilterById200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteFilterById200ResponseWithDefaults() *DeleteFilterById200Response {
	this := DeleteFilterById200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DeleteFilterById200Response) GetData() DeleteFilterByIdOutput {
	if o == nil || IsNil(o.Data) {
		var ret DeleteFilterByIdOutput
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteFilterById200Response) GetDataOk() (*DeleteFilterByIdOutput, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DeleteFilterById200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given DeleteFilterByIdOutput and assigns it to the Data field.
func (o *DeleteFilterById200Response) SetData(v DeleteFilterByIdOutput) {
	o.Data = &v
}

func (o DeleteFilterById200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteFilterById200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableDeleteFilterById200Response struct {
	value *DeleteFilterById200Response
	isSet bool
}

func (v NullableDeleteFilterById200Response) Get() *DeleteFilterById200Response {
	return v.value
}

func (v *NullableDeleteFilterById200Response) Set(val *DeleteFilterById200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteFilterById200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteFilterById200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteFilterById200Response(
	val *DeleteFilterById200Response,
) *NullableDeleteFilterById200Response {
	return &NullableDeleteFilterById200Response{value: val, isSet: true}
}

func (v NullableDeleteFilterById200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteFilterById200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
