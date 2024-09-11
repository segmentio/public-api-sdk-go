/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 54.3.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the CreateSource201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSource201Response{}

// CreateSource201Response struct for CreateSource201Response
type CreateSource201Response struct {
	Data *CreateSourceV1Output `json:"data,omitempty"`
}

// NewCreateSource201Response instantiates a new CreateSource201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSource201Response() *CreateSource201Response {
	this := CreateSource201Response{}
	return &this
}

// NewCreateSource201ResponseWithDefaults instantiates a new CreateSource201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSource201ResponseWithDefaults() *CreateSource201Response {
	this := CreateSource201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateSource201Response) GetData() CreateSourceV1Output {
	if o == nil || IsNil(o.Data) {
		var ret CreateSourceV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSource201Response) GetDataOk() (*CreateSourceV1Output, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateSource201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CreateSourceV1Output and assigns it to the Data field.
func (o *CreateSource201Response) SetData(v CreateSourceV1Output) {
	o.Data = &v
}

func (o CreateSource201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSource201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCreateSource201Response struct {
	value *CreateSource201Response
	isSet bool
}

func (v NullableCreateSource201Response) Get() *CreateSource201Response {
	return v.value
}

func (v *NullableCreateSource201Response) Set(val *CreateSource201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSource201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSource201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSource201Response(
	val *CreateSource201Response,
) *NullableCreateSource201Response {
	return &NullableCreateSource201Response{value: val, isSet: true}
}

func (v NullableCreateSource201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSource201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
