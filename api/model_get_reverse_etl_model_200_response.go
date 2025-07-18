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

// checks if the GetReverseEtlModel200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetReverseEtlModel200Response{}

// GetReverseEtlModel200Response struct for GetReverseEtlModel200Response
type GetReverseEtlModel200Response struct {
	Data *GetReverseEtlModelOutput `json:"data,omitempty"`
}

// NewGetReverseEtlModel200Response instantiates a new GetReverseEtlModel200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetReverseEtlModel200Response() *GetReverseEtlModel200Response {
	this := GetReverseEtlModel200Response{}
	return &this
}

// NewGetReverseEtlModel200ResponseWithDefaults instantiates a new GetReverseEtlModel200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetReverseEtlModel200ResponseWithDefaults() *GetReverseEtlModel200Response {
	this := GetReverseEtlModel200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetReverseEtlModel200Response) GetData() GetReverseEtlModelOutput {
	if o == nil || IsNil(o.Data) {
		var ret GetReverseEtlModelOutput
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReverseEtlModel200Response) GetDataOk() (*GetReverseEtlModelOutput, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetReverseEtlModel200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GetReverseEtlModelOutput and assigns it to the Data field.
func (o *GetReverseEtlModel200Response) SetData(v GetReverseEtlModelOutput) {
	o.Data = &v
}

func (o GetReverseEtlModel200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetReverseEtlModel200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetReverseEtlModel200Response struct {
	value *GetReverseEtlModel200Response
	isSet bool
}

func (v NullableGetReverseEtlModel200Response) Get() *GetReverseEtlModel200Response {
	return v.value
}

func (v *NullableGetReverseEtlModel200Response) Set(val *GetReverseEtlModel200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetReverseEtlModel200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetReverseEtlModel200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetReverseEtlModel200Response(
	val *GetReverseEtlModel200Response,
) *NullableGetReverseEtlModel200Response {
	return &NullableGetReverseEtlModel200Response{value: val, isSet: true}
}

func (v NullableGetReverseEtlModel200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetReverseEtlModel200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
