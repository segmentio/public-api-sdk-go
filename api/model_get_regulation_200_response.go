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

// checks if the GetRegulation200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRegulation200Response{}

// GetRegulation200Response struct for GetRegulation200Response
type GetRegulation200Response struct {
	Data *GetRegulationV1Output `json:"data,omitempty"`
}

// NewGetRegulation200Response instantiates a new GetRegulation200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRegulation200Response() *GetRegulation200Response {
	this := GetRegulation200Response{}
	return &this
}

// NewGetRegulation200ResponseWithDefaults instantiates a new GetRegulation200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRegulation200ResponseWithDefaults() *GetRegulation200Response {
	this := GetRegulation200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetRegulation200Response) GetData() GetRegulationV1Output {
	if o == nil || IsNil(o.Data) {
		var ret GetRegulationV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRegulation200Response) GetDataOk() (*GetRegulationV1Output, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetRegulation200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GetRegulationV1Output and assigns it to the Data field.
func (o *GetRegulation200Response) SetData(v GetRegulationV1Output) {
	o.Data = &v
}

func (o GetRegulation200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRegulation200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetRegulation200Response struct {
	value *GetRegulation200Response
	isSet bool
}

func (v NullableGetRegulation200Response) Get() *GetRegulation200Response {
	return v.value
}

func (v *NullableGetRegulation200Response) Set(val *GetRegulation200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRegulation200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRegulation200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRegulation200Response(
	val *GetRegulation200Response,
) *NullableGetRegulation200Response {
	return &NullableGetRegulation200Response{value: val, isSet: true}
}

func (v NullableGetRegulation200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRegulation200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
