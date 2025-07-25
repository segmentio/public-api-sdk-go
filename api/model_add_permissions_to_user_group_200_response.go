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

// checks if the AddPermissionsToUserGroup200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddPermissionsToUserGroup200Response{}

// AddPermissionsToUserGroup200Response struct for AddPermissionsToUserGroup200Response
type AddPermissionsToUserGroup200Response struct {
	Data *AddPermissionsToUserGroupV1Output `json:"data,omitempty"`
}

// NewAddPermissionsToUserGroup200Response instantiates a new AddPermissionsToUserGroup200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddPermissionsToUserGroup200Response() *AddPermissionsToUserGroup200Response {
	this := AddPermissionsToUserGroup200Response{}
	return &this
}

// NewAddPermissionsToUserGroup200ResponseWithDefaults instantiates a new AddPermissionsToUserGroup200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddPermissionsToUserGroup200ResponseWithDefaults() *AddPermissionsToUserGroup200Response {
	this := AddPermissionsToUserGroup200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AddPermissionsToUserGroup200Response) GetData() AddPermissionsToUserGroupV1Output {
	if o == nil || IsNil(o.Data) {
		var ret AddPermissionsToUserGroupV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPermissionsToUserGroup200Response) GetDataOk() (*AddPermissionsToUserGroupV1Output, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AddPermissionsToUserGroup200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AddPermissionsToUserGroupV1Output and assigns it to the Data field.
func (o *AddPermissionsToUserGroup200Response) SetData(v AddPermissionsToUserGroupV1Output) {
	o.Data = &v
}

func (o AddPermissionsToUserGroup200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddPermissionsToUserGroup200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAddPermissionsToUserGroup200Response struct {
	value *AddPermissionsToUserGroup200Response
	isSet bool
}

func (v NullableAddPermissionsToUserGroup200Response) Get() *AddPermissionsToUserGroup200Response {
	return v.value
}

func (v *NullableAddPermissionsToUserGroup200Response) Set(
	val *AddPermissionsToUserGroup200Response,
) {
	v.value = val
	v.isSet = true
}

func (v NullableAddPermissionsToUserGroup200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAddPermissionsToUserGroup200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddPermissionsToUserGroup200Response(
	val *AddPermissionsToUserGroup200Response,
) *NullableAddPermissionsToUserGroup200Response {
	return &NullableAddPermissionsToUserGroup200Response{value: val, isSet: true}
}

func (v NullableAddPermissionsToUserGroup200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddPermissionsToUserGroup200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
