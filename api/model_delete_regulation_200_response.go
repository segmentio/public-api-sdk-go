/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 39.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the DeleteRegulation200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteRegulation200Response{}

// DeleteRegulation200Response struct for DeleteRegulation200Response
type DeleteRegulation200Response struct {
	Data *DeleteRegulationV1Output `json:"data,omitempty"`
}

// NewDeleteRegulation200Response instantiates a new DeleteRegulation200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteRegulation200Response() *DeleteRegulation200Response {
	this := DeleteRegulation200Response{}
	return &this
}

// NewDeleteRegulation200ResponseWithDefaults instantiates a new DeleteRegulation200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteRegulation200ResponseWithDefaults() *DeleteRegulation200Response {
	this := DeleteRegulation200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DeleteRegulation200Response) GetData() DeleteRegulationV1Output {
	if o == nil || IsNil(o.Data) {
		var ret DeleteRegulationV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteRegulation200Response) GetDataOk() (*DeleteRegulationV1Output, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DeleteRegulation200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given DeleteRegulationV1Output and assigns it to the Data field.
func (o *DeleteRegulation200Response) SetData(v DeleteRegulationV1Output) {
	o.Data = &v
}

func (o DeleteRegulation200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteRegulation200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableDeleteRegulation200Response struct {
	value *DeleteRegulation200Response
	isSet bool
}

func (v NullableDeleteRegulation200Response) Get() *DeleteRegulation200Response {
	return v.value
}

func (v *NullableDeleteRegulation200Response) Set(val *DeleteRegulation200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteRegulation200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteRegulation200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteRegulation200Response(
	val *DeleteRegulation200Response,
) *NullableDeleteRegulation200Response {
	return &NullableDeleteRegulation200Response{value: val, isSet: true}
}

func (v NullableDeleteRegulation200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteRegulation200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
