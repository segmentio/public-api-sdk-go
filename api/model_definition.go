/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 42.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the Definition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Definition{}

// Definition Query language definition and type.
type Definition struct {
	// The query language string defining the computed trait aggregation criteria.
	Query string `json:"query"`
	// The underlying data type being aggregated for this computed trait.  Possible values: users, accounts.
	Type string `json:"type"`
}

// NewDefinition instantiates a new Definition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefinition(query string, type_ string) *Definition {
	this := Definition{}
	this.Query = query
	this.Type = type_
	return &this
}

// NewDefinitionWithDefaults instantiates a new Definition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefinitionWithDefaults() *Definition {
	this := Definition{}
	return &this
}

// GetQuery returns the Query field value
func (o *Definition) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *Definition) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *Definition) SetQuery(v string) {
	o.Query = v
}

// GetType returns the Type field value
func (o *Definition) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Definition) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Definition) SetType(v string) {
	o.Type = v
}

func (o Definition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Definition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["query"] = o.Query
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableDefinition struct {
	value *Definition
	isSet bool
}

func (v NullableDefinition) Get() *Definition {
	return v.value
}

func (v *NullableDefinition) Set(val *Definition) {
	v.value = val
	v.isSet = true
}

func (v NullableDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefinition(val *Definition) *NullableDefinition {
	return &NullableDefinition{value: val, isSet: true}
}

func (v NullableDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
