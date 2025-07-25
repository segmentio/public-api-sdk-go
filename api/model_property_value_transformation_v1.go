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

// checks if the PropertyValueTransformationV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PropertyValueTransformationV1{}

// PropertyValueTransformationV1 struct for PropertyValueTransformationV1
type PropertyValueTransformationV1 struct {
	// The property paths. The maximum number of paths is 10.
	PropertyPaths []string `json:"propertyPaths"`
	// The new value of the property paths.
	PropertyValue string `json:"propertyValue"`
}

// NewPropertyValueTransformationV1 instantiates a new PropertyValueTransformationV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyValueTransformationV1(
	propertyPaths []string,
	propertyValue string,
) *PropertyValueTransformationV1 {
	this := PropertyValueTransformationV1{}
	this.PropertyPaths = propertyPaths
	this.PropertyValue = propertyValue
	return &this
}

// NewPropertyValueTransformationV1WithDefaults instantiates a new PropertyValueTransformationV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyValueTransformationV1WithDefaults() *PropertyValueTransformationV1 {
	this := PropertyValueTransformationV1{}
	return &this
}

// GetPropertyPaths returns the PropertyPaths field value
func (o *PropertyValueTransformationV1) GetPropertyPaths() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PropertyPaths
}

// GetPropertyPathsOk returns a tuple with the PropertyPaths field value
// and a boolean to check if the value has been set.
func (o *PropertyValueTransformationV1) GetPropertyPathsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PropertyPaths, true
}

// SetPropertyPaths sets field value
func (o *PropertyValueTransformationV1) SetPropertyPaths(v []string) {
	o.PropertyPaths = v
}

// GetPropertyValue returns the PropertyValue field value
func (o *PropertyValueTransformationV1) GetPropertyValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PropertyValue
}

// GetPropertyValueOk returns a tuple with the PropertyValue field value
// and a boolean to check if the value has been set.
func (o *PropertyValueTransformationV1) GetPropertyValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PropertyValue, true
}

// SetPropertyValue sets field value
func (o *PropertyValueTransformationV1) SetPropertyValue(v string) {
	o.PropertyValue = v
}

func (o PropertyValueTransformationV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PropertyValueTransformationV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["propertyPaths"] = o.PropertyPaths
	toSerialize["propertyValue"] = o.PropertyValue
	return toSerialize, nil
}

type NullablePropertyValueTransformationV1 struct {
	value *PropertyValueTransformationV1
	isSet bool
}

func (v NullablePropertyValueTransformationV1) Get() *PropertyValueTransformationV1 {
	return v.value
}

func (v *NullablePropertyValueTransformationV1) Set(val *PropertyValueTransformationV1) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyValueTransformationV1) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyValueTransformationV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyValueTransformationV1(
	val *PropertyValueTransformationV1,
) *NullablePropertyValueTransformationV1 {
	return &NullablePropertyValueTransformationV1{value: val, isSet: true}
}

func (v NullablePropertyValueTransformationV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyValueTransformationV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
