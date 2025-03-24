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

// checks if the GetComputedTrait200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetComputedTrait200Response{}

// GetComputedTrait200Response struct for GetComputedTrait200Response
type GetComputedTrait200Response struct {
	Data *GetComputedTraitAlphaOutput `json:"data,omitempty"`
}

// NewGetComputedTrait200Response instantiates a new GetComputedTrait200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetComputedTrait200Response() *GetComputedTrait200Response {
	this := GetComputedTrait200Response{}
	return &this
}

// NewGetComputedTrait200ResponseWithDefaults instantiates a new GetComputedTrait200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetComputedTrait200ResponseWithDefaults() *GetComputedTrait200Response {
	this := GetComputedTrait200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetComputedTrait200Response) GetData() GetComputedTraitAlphaOutput {
	if o == nil || IsNil(o.Data) {
		var ret GetComputedTraitAlphaOutput
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetComputedTrait200Response) GetDataOk() (*GetComputedTraitAlphaOutput, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetComputedTrait200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GetComputedTraitAlphaOutput and assigns it to the Data field.
func (o *GetComputedTrait200Response) SetData(v GetComputedTraitAlphaOutput) {
	o.Data = &v
}

func (o GetComputedTrait200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetComputedTrait200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetComputedTrait200Response struct {
	value *GetComputedTrait200Response
	isSet bool
}

func (v NullableGetComputedTrait200Response) Get() *GetComputedTrait200Response {
	return v.value
}

func (v *NullableGetComputedTrait200Response) Set(val *GetComputedTrait200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetComputedTrait200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetComputedTrait200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetComputedTrait200Response(
	val *GetComputedTrait200Response,
) *NullableGetComputedTrait200Response {
	return &NullableGetComputedTrait200Response{value: val, isSet: true}
}

func (v NullableGetComputedTrait200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetComputedTrait200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
