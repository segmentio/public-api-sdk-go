/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 38.5.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the DestinationMetadataComponentV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DestinationMetadataComponentV1{}

// DestinationMetadataComponentV1 Represents a component this Destination provides.
type DestinationMetadataComponentV1 struct {
	// The component type.
	Type string `json:"type"`
	// Link to the repository hosting the code for this component.
	Code string `json:"code"`
	// The owner of this component. Either 'SEGMENT' or 'PARTNER'.
	Owner *string `json:"owner,omitempty"`
}

// NewDestinationMetadataComponentV1 instantiates a new DestinationMetadataComponentV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestinationMetadataComponentV1(type_ string, code string) *DestinationMetadataComponentV1 {
	this := DestinationMetadataComponentV1{}
	this.Type = type_
	this.Code = code
	return &this
}

// NewDestinationMetadataComponentV1WithDefaults instantiates a new DestinationMetadataComponentV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestinationMetadataComponentV1WithDefaults() *DestinationMetadataComponentV1 {
	this := DestinationMetadataComponentV1{}
	return &this
}

// GetType returns the Type field value
func (o *DestinationMetadataComponentV1) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadataComponentV1) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DestinationMetadataComponentV1) SetType(v string) {
	o.Type = v
}

// GetCode returns the Code field value
func (o *DestinationMetadataComponentV1) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadataComponentV1) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *DestinationMetadataComponentV1) SetCode(v string) {
	o.Code = v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *DestinationMetadataComponentV1) GetOwner() string {
	if o == nil || IsNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationMetadataComponentV1) GetOwnerOk() (*string, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *DestinationMetadataComponentV1) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *DestinationMetadataComponentV1) SetOwner(v string) {
	o.Owner = &v
}

func (o DestinationMetadataComponentV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DestinationMetadataComponentV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["code"] = o.Code
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	return toSerialize, nil
}

type NullableDestinationMetadataComponentV1 struct {
	value *DestinationMetadataComponentV1
	isSet bool
}

func (v NullableDestinationMetadataComponentV1) Get() *DestinationMetadataComponentV1 {
	return v.value
}

func (v *NullableDestinationMetadataComponentV1) Set(val *DestinationMetadataComponentV1) {
	v.value = val
	v.isSet = true
}

func (v NullableDestinationMetadataComponentV1) IsSet() bool {
	return v.isSet
}

func (v *NullableDestinationMetadataComponentV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestinationMetadataComponentV1(
	val *DestinationMetadataComponentV1,
) *NullableDestinationMetadataComponentV1 {
	return &NullableDestinationMetadataComponentV1{value: val, isSet: true}
}

func (v NullableDestinationMetadataComponentV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDestinationMetadataComponentV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
