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

// checks if the LabelV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LabelV1{}

// LabelV1 A label lets Workspace owners assign permissions to users, and grant these users access to groups.  A Workspace owner may use labels to grant users access to groups of resources.  When you add a label to a Source or Personas Spaces, any users granted access to that label gain access to those resources.  All Workspaces include labels for Dev (development) and Prod (production) environments. On top of those, Free and Team plan customers may create up to five labels.  Customers with the Enterprise pricing package may create an unlimited number of labels.
type LabelV1 struct {
	// The key that represents the name of this label.
	Key string `json:"key"`
	// The value associated with the key of this label.
	Value string `json:"value"`
	// An optional description of the purpose of this label.
	Description *string `json:"description,omitempty"`
}

// NewLabelV1 instantiates a new LabelV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabelV1(key string, value string) *LabelV1 {
	this := LabelV1{}
	this.Key = key
	this.Value = value
	return &this
}

// NewLabelV1WithDefaults instantiates a new LabelV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelV1WithDefaults() *LabelV1 {
	this := LabelV1{}
	return &this
}

// GetKey returns the Key field value
func (o *LabelV1) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *LabelV1) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *LabelV1) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *LabelV1) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *LabelV1) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *LabelV1) SetValue(v string) {
	o.Value = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LabelV1) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelV1) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LabelV1) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LabelV1) SetDescription(v string) {
	o.Description = &v
}

func (o LabelV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LabelV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["value"] = o.Value
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableLabelV1 struct {
	value *LabelV1
	isSet bool
}

func (v NullableLabelV1) Get() *LabelV1 {
	return v.value
}

func (v *NullableLabelV1) Set(val *LabelV1) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelV1) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelV1(val *LabelV1) *NullableLabelV1 {
	return &NullableLabelV1{value: val, isSet: true}
}

func (v NullableLabelV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabelV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
