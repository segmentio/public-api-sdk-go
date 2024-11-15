/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 57.0.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the CreateFilterForSpace200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFilterForSpace200Response{}

// CreateFilterForSpace200Response struct for CreateFilterForSpace200Response
type CreateFilterForSpace200Response struct {
	Data *CreateFilterForSpaceOutput `json:"data,omitempty"`
}

// NewCreateFilterForSpace200Response instantiates a new CreateFilterForSpace200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFilterForSpace200Response() *CreateFilterForSpace200Response {
	this := CreateFilterForSpace200Response{}
	return &this
}

// NewCreateFilterForSpace200ResponseWithDefaults instantiates a new CreateFilterForSpace200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFilterForSpace200ResponseWithDefaults() *CreateFilterForSpace200Response {
	this := CreateFilterForSpace200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateFilterForSpace200Response) GetData() CreateFilterForSpaceOutput {
	if o == nil || IsNil(o.Data) {
		var ret CreateFilterForSpaceOutput
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFilterForSpace200Response) GetDataOk() (*CreateFilterForSpaceOutput, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateFilterForSpace200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CreateFilterForSpaceOutput and assigns it to the Data field.
func (o *CreateFilterForSpace200Response) SetData(v CreateFilterForSpaceOutput) {
	o.Data = &v
}

func (o CreateFilterForSpace200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateFilterForSpace200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCreateFilterForSpace200Response struct {
	value *CreateFilterForSpace200Response
	isSet bool
}

func (v NullableCreateFilterForSpace200Response) Get() *CreateFilterForSpace200Response {
	return v.value
}

func (v *NullableCreateFilterForSpace200Response) Set(val *CreateFilterForSpace200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFilterForSpace200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFilterForSpace200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFilterForSpace200Response(
	val *CreateFilterForSpace200Response,
) *NullableCreateFilterForSpace200Response {
	return &NullableCreateFilterForSpace200Response{value: val, isSet: true}
}

func (v NullableCreateFilterForSpace200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFilterForSpace200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
