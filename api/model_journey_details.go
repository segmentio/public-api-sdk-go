/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.1.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the JourneyDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JourneyDetails{}

// JourneyDetails struct for JourneyDetails
type JourneyDetails struct {
	ContainerId   *string  `json:"containerId,omitempty"`
	VersionNumber *float32 `json:"versionNumber,omitempty"`
}

// NewJourneyDetails instantiates a new JourneyDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJourneyDetails() *JourneyDetails {
	this := JourneyDetails{}
	return &this
}

// NewJourneyDetailsWithDefaults instantiates a new JourneyDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJourneyDetailsWithDefaults() *JourneyDetails {
	this := JourneyDetails{}
	return &this
}

// GetContainerId returns the ContainerId field value if set, zero value otherwise.
func (o *JourneyDetails) GetContainerId() string {
	if o == nil || IsNil(o.ContainerId) {
		var ret string
		return ret
	}
	return *o.ContainerId
}

// GetContainerIdOk returns a tuple with the ContainerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JourneyDetails) GetContainerIdOk() (*string, bool) {
	if o == nil || IsNil(o.ContainerId) {
		return nil, false
	}
	return o.ContainerId, true
}

// HasContainerId returns a boolean if a field has been set.
func (o *JourneyDetails) HasContainerId() bool {
	if o != nil && !IsNil(o.ContainerId) {
		return true
	}

	return false
}

// SetContainerId gets a reference to the given string and assigns it to the ContainerId field.
func (o *JourneyDetails) SetContainerId(v string) {
	o.ContainerId = &v
}

// GetVersionNumber returns the VersionNumber field value if set, zero value otherwise.
func (o *JourneyDetails) GetVersionNumber() float32 {
	if o == nil || IsNil(o.VersionNumber) {
		var ret float32
		return ret
	}
	return *o.VersionNumber
}

// GetVersionNumberOk returns a tuple with the VersionNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JourneyDetails) GetVersionNumberOk() (*float32, bool) {
	if o == nil || IsNil(o.VersionNumber) {
		return nil, false
	}
	return o.VersionNumber, true
}

// HasVersionNumber returns a boolean if a field has been set.
func (o *JourneyDetails) HasVersionNumber() bool {
	if o != nil && !IsNil(o.VersionNumber) {
		return true
	}

	return false
}

// SetVersionNumber gets a reference to the given float32 and assigns it to the VersionNumber field.
func (o *JourneyDetails) SetVersionNumber(v float32) {
	o.VersionNumber = &v
}

func (o JourneyDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JourneyDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContainerId) {
		toSerialize["containerId"] = o.ContainerId
	}
	if !IsNil(o.VersionNumber) {
		toSerialize["versionNumber"] = o.VersionNumber
	}
	return toSerialize, nil
}

type NullableJourneyDetails struct {
	value *JourneyDetails
	isSet bool
}

func (v NullableJourneyDetails) Get() *JourneyDetails {
	return v.value
}

func (v *NullableJourneyDetails) Set(val *JourneyDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableJourneyDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableJourneyDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJourneyDetails(val *JourneyDetails) *NullableJourneyDetails {
	return &NullableJourneyDetails{value: val, isSet: true}
}

func (v NullableJourneyDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJourneyDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
