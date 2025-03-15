/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.0.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the DeleteUserGroup200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteUserGroup200Response{}

// DeleteUserGroup200Response struct for DeleteUserGroup200Response
type DeleteUserGroup200Response struct {
	Data *DeleteUserGroupV1Output `json:"data,omitempty"`
}

// NewDeleteUserGroup200Response instantiates a new DeleteUserGroup200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteUserGroup200Response() *DeleteUserGroup200Response {
	this := DeleteUserGroup200Response{}
	return &this
}

// NewDeleteUserGroup200ResponseWithDefaults instantiates a new DeleteUserGroup200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteUserGroup200ResponseWithDefaults() *DeleteUserGroup200Response {
	this := DeleteUserGroup200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DeleteUserGroup200Response) GetData() DeleteUserGroupV1Output {
	if o == nil || IsNil(o.Data) {
		var ret DeleteUserGroupV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteUserGroup200Response) GetDataOk() (*DeleteUserGroupV1Output, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DeleteUserGroup200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given DeleteUserGroupV1Output and assigns it to the Data field.
func (o *DeleteUserGroup200Response) SetData(v DeleteUserGroupV1Output) {
	o.Data = &v
}

func (o DeleteUserGroup200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteUserGroup200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableDeleteUserGroup200Response struct {
	value *DeleteUserGroup200Response
	isSet bool
}

func (v NullableDeleteUserGroup200Response) Get() *DeleteUserGroup200Response {
	return v.value
}

func (v *NullableDeleteUserGroup200Response) Set(val *DeleteUserGroup200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteUserGroup200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteUserGroup200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteUserGroup200Response(
	val *DeleteUserGroup200Response,
) *NullableDeleteUserGroup200Response {
	return &NullableDeleteUserGroup200Response{value: val, isSet: true}
}

func (v NullableDeleteUserGroup200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteUserGroup200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
