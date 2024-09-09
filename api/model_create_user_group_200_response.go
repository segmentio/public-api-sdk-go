/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 54.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the CreateUserGroup200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateUserGroup200Response{}

// CreateUserGroup200Response struct for CreateUserGroup200Response
type CreateUserGroup200Response struct {
	Data *CreateUserGroupV1Output `json:"data,omitempty"`
}

// NewCreateUserGroup200Response instantiates a new CreateUserGroup200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUserGroup200Response() *CreateUserGroup200Response {
	this := CreateUserGroup200Response{}
	return &this
}

// NewCreateUserGroup200ResponseWithDefaults instantiates a new CreateUserGroup200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUserGroup200ResponseWithDefaults() *CreateUserGroup200Response {
	this := CreateUserGroup200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateUserGroup200Response) GetData() CreateUserGroupV1Output {
	if o == nil || IsNil(o.Data) {
		var ret CreateUserGroupV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserGroup200Response) GetDataOk() (*CreateUserGroupV1Output, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateUserGroup200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CreateUserGroupV1Output and assigns it to the Data field.
func (o *CreateUserGroup200Response) SetData(v CreateUserGroupV1Output) {
	o.Data = &v
}

func (o CreateUserGroup200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateUserGroup200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCreateUserGroup200Response struct {
	value *CreateUserGroup200Response
	isSet bool
}

func (v NullableCreateUserGroup200Response) Get() *CreateUserGroup200Response {
	return v.value
}

func (v *NullableCreateUserGroup200Response) Set(val *CreateUserGroup200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUserGroup200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUserGroup200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUserGroup200Response(
	val *CreateUserGroup200Response,
) *NullableCreateUserGroup200Response {
	return &NullableCreateUserGroup200Response{value: val, isSet: true}
}

func (v NullableCreateUserGroup200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUserGroup200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
