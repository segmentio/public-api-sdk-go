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

// checks if the RestoreFunctionVersionAlphaInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestoreFunctionVersionAlphaInput{}

// RestoreFunctionVersionAlphaInput Restore a given Function version.
type RestoreFunctionVersionAlphaInput struct {
	// An identifier for this version.
	VersionId string `json:"versionId"`
}

// NewRestoreFunctionVersionAlphaInput instantiates a new RestoreFunctionVersionAlphaInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestoreFunctionVersionAlphaInput(versionId string) *RestoreFunctionVersionAlphaInput {
	this := RestoreFunctionVersionAlphaInput{}
	this.VersionId = versionId
	return &this
}

// NewRestoreFunctionVersionAlphaInputWithDefaults instantiates a new RestoreFunctionVersionAlphaInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestoreFunctionVersionAlphaInputWithDefaults() *RestoreFunctionVersionAlphaInput {
	this := RestoreFunctionVersionAlphaInput{}
	return &this
}

// GetVersionId returns the VersionId field value
func (o *RestoreFunctionVersionAlphaInput) GetVersionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
func (o *RestoreFunctionVersionAlphaInput) GetVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionId, true
}

// SetVersionId sets field value
func (o *RestoreFunctionVersionAlphaInput) SetVersionId(v string) {
	o.VersionId = v
}

func (o RestoreFunctionVersionAlphaInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestoreFunctionVersionAlphaInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["versionId"] = o.VersionId
	return toSerialize, nil
}

type NullableRestoreFunctionVersionAlphaInput struct {
	value *RestoreFunctionVersionAlphaInput
	isSet bool
}

func (v NullableRestoreFunctionVersionAlphaInput) Get() *RestoreFunctionVersionAlphaInput {
	return v.value
}

func (v *NullableRestoreFunctionVersionAlphaInput) Set(val *RestoreFunctionVersionAlphaInput) {
	v.value = val
	v.isSet = true
}

func (v NullableRestoreFunctionVersionAlphaInput) IsSet() bool {
	return v.isSet
}

func (v *NullableRestoreFunctionVersionAlphaInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestoreFunctionVersionAlphaInput(
	val *RestoreFunctionVersionAlphaInput,
) *NullableRestoreFunctionVersionAlphaInput {
	return &NullableRestoreFunctionVersionAlphaInput{value: val, isSet: true}
}

func (v NullableRestoreFunctionVersionAlphaInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestoreFunctionVersionAlphaInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
