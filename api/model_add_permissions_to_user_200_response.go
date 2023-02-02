/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AddPermissionsToUser200Response struct for AddPermissionsToUser200Response
type AddPermissionsToUser200Response struct {
	Data *AddPermissionsToUserV1Output `json:"data,omitempty"`
}

// NewAddPermissionsToUser200Response instantiates a new AddPermissionsToUser200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddPermissionsToUser200Response() *AddPermissionsToUser200Response {
	this := AddPermissionsToUser200Response{}
	return &this
}

// NewAddPermissionsToUser200ResponseWithDefaults instantiates a new AddPermissionsToUser200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddPermissionsToUser200ResponseWithDefaults() *AddPermissionsToUser200Response {
	this := AddPermissionsToUser200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AddPermissionsToUser200Response) GetData() AddPermissionsToUserV1Output {
	if o == nil || o.Data == nil {
		var ret AddPermissionsToUserV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPermissionsToUser200Response) GetDataOk() (*AddPermissionsToUserV1Output, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AddPermissionsToUser200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given AddPermissionsToUserV1Output and assigns it to the Data field.
func (o *AddPermissionsToUser200Response) SetData(v AddPermissionsToUserV1Output) {
	o.Data = &v
}

func (o AddPermissionsToUser200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableAddPermissionsToUser200Response struct {
	value *AddPermissionsToUser200Response
	isSet bool
}

func (v NullableAddPermissionsToUser200Response) Get() *AddPermissionsToUser200Response {
	return v.value
}

func (v *NullableAddPermissionsToUser200Response) Set(val *AddPermissionsToUser200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAddPermissionsToUser200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAddPermissionsToUser200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddPermissionsToUser200Response(
	val *AddPermissionsToUser200Response,
) *NullableAddPermissionsToUser200Response {
	return &NullableAddPermissionsToUser200Response{value: val, isSet: true}
}

func (v NullableAddPermissionsToUser200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddPermissionsToUser200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
