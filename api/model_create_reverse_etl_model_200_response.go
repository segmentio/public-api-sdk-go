/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 38.4.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the CreateReverseEtlModel200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateReverseEtlModel200Response{}

// CreateReverseEtlModel200Response struct for CreateReverseEtlModel200Response
type CreateReverseEtlModel200Response struct {
	Data *CreateReverseEtlModelOutput `json:"data,omitempty"`
}

// NewCreateReverseEtlModel200Response instantiates a new CreateReverseEtlModel200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateReverseEtlModel200Response() *CreateReverseEtlModel200Response {
	this := CreateReverseEtlModel200Response{}
	return &this
}

// NewCreateReverseEtlModel200ResponseWithDefaults instantiates a new CreateReverseEtlModel200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateReverseEtlModel200ResponseWithDefaults() *CreateReverseEtlModel200Response {
	this := CreateReverseEtlModel200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateReverseEtlModel200Response) GetData() CreateReverseEtlModelOutput {
	if o == nil || IsNil(o.Data) {
		var ret CreateReverseEtlModelOutput
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateReverseEtlModel200Response) GetDataOk() (*CreateReverseEtlModelOutput, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateReverseEtlModel200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CreateReverseEtlModelOutput and assigns it to the Data field.
func (o *CreateReverseEtlModel200Response) SetData(v CreateReverseEtlModelOutput) {
	o.Data = &v
}

func (o CreateReverseEtlModel200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateReverseEtlModel200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCreateReverseEtlModel200Response struct {
	value *CreateReverseEtlModel200Response
	isSet bool
}

func (v NullableCreateReverseEtlModel200Response) Get() *CreateReverseEtlModel200Response {
	return v.value
}

func (v *NullableCreateReverseEtlModel200Response) Set(val *CreateReverseEtlModel200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateReverseEtlModel200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateReverseEtlModel200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateReverseEtlModel200Response(
	val *CreateReverseEtlModel200Response,
) *NullableCreateReverseEtlModel200Response {
	return &NullableCreateReverseEtlModel200Response{value: val, isSet: true}
}

func (v NullableCreateReverseEtlModel200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateReverseEtlModel200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
