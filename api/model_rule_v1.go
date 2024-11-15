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

// checks if the RuleV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleV1{}

// RuleV1 Represents a rule from a Tracking Plan.
type RuleV1 struct {
	// The type for this Tracking Plan rule.
	Type string `json:"type"`
	// Key to this rule (free-form string like 'Button clicked').
	Key *string `json:"key,omitempty"`
	// JSON Schema of this rule.
	JsonSchema interface{} `json:"jsonSchema"`
	// Version of this rule.
	Version float32 `json:"version"`
	// The timestamp of this rule's creation.
	CreatedAt *string `json:"createdAt,omitempty"`
	// The timestamp of this rule's last change.
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// The timestamp of this rule's deprecation.
	DeprecatedAt *string `json:"deprecatedAt,omitempty"`
}

// NewRuleV1 instantiates a new RuleV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleV1(type_ string, jsonSchema interface{}, version float32) *RuleV1 {
	this := RuleV1{}
	this.Type = type_
	this.JsonSchema = jsonSchema
	this.Version = version
	return &this
}

// NewRuleV1WithDefaults instantiates a new RuleV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleV1WithDefaults() *RuleV1 {
	this := RuleV1{}
	return &this
}

// GetType returns the Type field value
func (o *RuleV1) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RuleV1) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RuleV1) SetType(v string) {
	o.Type = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *RuleV1) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleV1) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *RuleV1) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *RuleV1) SetKey(v string) {
	o.Key = &v
}

// GetJsonSchema returns the JsonSchema field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *RuleV1) GetJsonSchema() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.JsonSchema
}

// GetJsonSchemaOk returns a tuple with the JsonSchema field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RuleV1) GetJsonSchemaOk() (*interface{}, bool) {
	if o == nil || IsNil(o.JsonSchema) {
		return nil, false
	}
	return &o.JsonSchema, true
}

// SetJsonSchema sets field value
func (o *RuleV1) SetJsonSchema(v interface{}) {
	o.JsonSchema = v
}

// GetVersion returns the Version field value
func (o *RuleV1) GetVersion() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *RuleV1) GetVersionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *RuleV1) SetVersion(v float32) {
	o.Version = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RuleV1) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleV1) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RuleV1) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *RuleV1) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *RuleV1) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleV1) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RuleV1) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *RuleV1) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeprecatedAt returns the DeprecatedAt field value if set, zero value otherwise.
func (o *RuleV1) GetDeprecatedAt() string {
	if o == nil || IsNil(o.DeprecatedAt) {
		var ret string
		return ret
	}
	return *o.DeprecatedAt
}

// GetDeprecatedAtOk returns a tuple with the DeprecatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleV1) GetDeprecatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeprecatedAt) {
		return nil, false
	}
	return o.DeprecatedAt, true
}

// HasDeprecatedAt returns a boolean if a field has been set.
func (o *RuleV1) HasDeprecatedAt() bool {
	if o != nil && !IsNil(o.DeprecatedAt) {
		return true
	}

	return false
}

// SetDeprecatedAt gets a reference to the given string and assigns it to the DeprecatedAt field.
func (o *RuleV1) SetDeprecatedAt(v string) {
	o.DeprecatedAt = &v
}

func (o RuleV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuleV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if o.JsonSchema != nil {
		toSerialize["jsonSchema"] = o.JsonSchema
	}
	toSerialize["version"] = o.Version
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.DeprecatedAt) {
		toSerialize["deprecatedAt"] = o.DeprecatedAt
	}
	return toSerialize, nil
}

type NullableRuleV1 struct {
	value *RuleV1
	isSet bool
}

func (v NullableRuleV1) Get() *RuleV1 {
	return v.value
}

func (v *NullableRuleV1) Set(val *RuleV1) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleV1) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleV1(val *RuleV1) *NullableRuleV1 {
	return &NullableRuleV1{value: val, isSet: true}
}

func (v NullableRuleV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
