/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 50.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the RuleInputV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleInputV1{}

// RuleInputV1 Represents a rule to add to a Tracking Plan.
type RuleInputV1 struct {
	// The type for this Tracking Plan rule.
	Type string `json:"type"`
	// Key to this rule (free-form string like 'Button clicked').
	Key *string `json:"key,omitempty"`
	// JSON Schema of this rule.
	JsonSchema interface{} `json:"jsonSchema"`
	// Version of this rule.
	Version float32 `json:"version"`
}

// NewRuleInputV1 instantiates a new RuleInputV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleInputV1(type_ string, jsonSchema interface{}, version float32) *RuleInputV1 {
	this := RuleInputV1{}
	this.Type = type_
	this.JsonSchema = jsonSchema
	this.Version = version
	return &this
}

// NewRuleInputV1WithDefaults instantiates a new RuleInputV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleInputV1WithDefaults() *RuleInputV1 {
	this := RuleInputV1{}
	return &this
}

// GetType returns the Type field value
func (o *RuleInputV1) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RuleInputV1) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RuleInputV1) SetType(v string) {
	o.Type = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *RuleInputV1) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleInputV1) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *RuleInputV1) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *RuleInputV1) SetKey(v string) {
	o.Key = &v
}

// GetJsonSchema returns the JsonSchema field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *RuleInputV1) GetJsonSchema() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.JsonSchema
}

// GetJsonSchemaOk returns a tuple with the JsonSchema field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RuleInputV1) GetJsonSchemaOk() (*interface{}, bool) {
	if o == nil || IsNil(o.JsonSchema) {
		return nil, false
	}
	return &o.JsonSchema, true
}

// SetJsonSchema sets field value
func (o *RuleInputV1) SetJsonSchema(v interface{}) {
	o.JsonSchema = v
}

// GetVersion returns the Version field value
func (o *RuleInputV1) GetVersion() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *RuleInputV1) GetVersionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *RuleInputV1) SetVersion(v float32) {
	o.Version = v
}

func (o RuleInputV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuleInputV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if o.JsonSchema != nil {
		toSerialize["jsonSchema"] = o.JsonSchema
	}
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

type NullableRuleInputV1 struct {
	value *RuleInputV1
	isSet bool
}

func (v NullableRuleInputV1) Get() *RuleInputV1 {
	return v.value
}

func (v *NullableRuleInputV1) Set(val *RuleInputV1) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleInputV1) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleInputV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleInputV1(val *RuleInputV1) *NullableRuleInputV1 {
	return &NullableRuleInputV1{value: val, isSet: true}
}

func (v NullableRuleInputV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleInputV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
