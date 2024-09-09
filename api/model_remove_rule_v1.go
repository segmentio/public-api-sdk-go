/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 54.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the RemoveRuleV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoveRuleV1{}

// RemoveRuleV1 Represents the parameters needed to identify a rule on the backend-side.
type RemoveRuleV1 struct {
	// The type for this Tracking Plan rule.
	Type string `json:"type"`
	// Key to this rule (free-form string like 'Button clicked').
	Key *string `json:"key,omitempty"`
	// Version of this rule.
	Version float32 `json:"version"`
}

// NewRemoveRuleV1 instantiates a new RemoveRuleV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveRuleV1(type_ string, version float32) *RemoveRuleV1 {
	this := RemoveRuleV1{}
	this.Type = type_
	this.Version = version
	return &this
}

// NewRemoveRuleV1WithDefaults instantiates a new RemoveRuleV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveRuleV1WithDefaults() *RemoveRuleV1 {
	this := RemoveRuleV1{}
	return &this
}

// GetType returns the Type field value
func (o *RemoveRuleV1) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RemoveRuleV1) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RemoveRuleV1) SetType(v string) {
	o.Type = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *RemoveRuleV1) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoveRuleV1) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *RemoveRuleV1) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *RemoveRuleV1) SetKey(v string) {
	o.Key = &v
}

// GetVersion returns the Version field value
func (o *RemoveRuleV1) GetVersion() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *RemoveRuleV1) GetVersionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *RemoveRuleV1) SetVersion(v float32) {
	o.Version = v
}

func (o RemoveRuleV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoveRuleV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

type NullableRemoveRuleV1 struct {
	value *RemoveRuleV1
	isSet bool
}

func (v NullableRemoveRuleV1) Get() *RemoveRuleV1 {
	return v.value
}

func (v *NullableRemoveRuleV1) Set(val *RemoveRuleV1) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveRuleV1) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveRuleV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveRuleV1(val *RemoveRuleV1) *NullableRemoveRuleV1 {
	return &NullableRemoveRuleV1{value: val, isSet: true}
}

func (v NullableRemoveRuleV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveRuleV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
