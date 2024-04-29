/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 49.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the CreateInsertFunctionInstance200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateInsertFunctionInstance200Response{}

// CreateInsertFunctionInstance200Response struct for CreateInsertFunctionInstance200Response
type CreateInsertFunctionInstance200Response struct {
	Data *CreateInsertFunctionInstanceAlphaOutput `json:"data,omitempty"`
}

// NewCreateInsertFunctionInstance200Response instantiates a new CreateInsertFunctionInstance200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateInsertFunctionInstance200Response() *CreateInsertFunctionInstance200Response {
	this := CreateInsertFunctionInstance200Response{}
	return &this
}

// NewCreateInsertFunctionInstance200ResponseWithDefaults instantiates a new CreateInsertFunctionInstance200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateInsertFunctionInstance200ResponseWithDefaults() *CreateInsertFunctionInstance200Response {
	this := CreateInsertFunctionInstance200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateInsertFunctionInstance200Response) GetData() CreateInsertFunctionInstanceAlphaOutput {
	if o == nil || IsNil(o.Data) {
		var ret CreateInsertFunctionInstanceAlphaOutput
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInsertFunctionInstance200Response) GetDataOk() (*CreateInsertFunctionInstanceAlphaOutput, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateInsertFunctionInstance200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CreateInsertFunctionInstanceAlphaOutput and assigns it to the Data field.
func (o *CreateInsertFunctionInstance200Response) SetData(
	v CreateInsertFunctionInstanceAlphaOutput,
) {
	o.Data = &v
}

func (o CreateInsertFunctionInstance200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateInsertFunctionInstance200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCreateInsertFunctionInstance200Response struct {
	value *CreateInsertFunctionInstance200Response
	isSet bool
}

func (v NullableCreateInsertFunctionInstance200Response) Get() *CreateInsertFunctionInstance200Response {
	return v.value
}

func (v *NullableCreateInsertFunctionInstance200Response) Set(
	val *CreateInsertFunctionInstance200Response,
) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateInsertFunctionInstance200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateInsertFunctionInstance200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateInsertFunctionInstance200Response(
	val *CreateInsertFunctionInstance200Response,
) *NullableCreateInsertFunctionInstance200Response {
	return &NullableCreateInsertFunctionInstance200Response{value: val, isSet: true}
}

func (v NullableCreateInsertFunctionInstance200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateInsertFunctionInstance200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
