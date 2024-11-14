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

// checks if the TraitDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TraitDefinition{}

// TraitDefinition struct for TraitDefinition
type TraitDefinition struct {
	// The underlying data type being aggregated for this computed trait.  Possible values: users, accounts.
	Type string `json:"type"`
	// The query language string defining the computed trait aggregation criteria. For guidance on using the query language, see the [Segment documentation site](https://segment.com/docs/api/public-api/query-language).
	Query string `json:"query"`
}

// NewTraitDefinition instantiates a new TraitDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTraitDefinition(type_ string, query string) *TraitDefinition {
	this := TraitDefinition{}
	this.Type = type_
	this.Query = query
	return &this
}

// NewTraitDefinitionWithDefaults instantiates a new TraitDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTraitDefinitionWithDefaults() *TraitDefinition {
	this := TraitDefinition{}
	return &this
}

// GetType returns the Type field value
func (o *TraitDefinition) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TraitDefinition) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TraitDefinition) SetType(v string) {
	o.Type = v
}

// GetQuery returns the Query field value
func (o *TraitDefinition) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *TraitDefinition) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *TraitDefinition) SetQuery(v string) {
	o.Query = v
}

func (o TraitDefinition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TraitDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["query"] = o.Query
	return toSerialize, nil
}

type NullableTraitDefinition struct {
	value *TraitDefinition
	isSet bool
}

func (v NullableTraitDefinition) Get() *TraitDefinition {
	return v.value
}

func (v *NullableTraitDefinition) Set(val *TraitDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableTraitDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableTraitDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraitDefinition(val *TraitDefinition) *NullableTraitDefinition {
	return &NullableTraitDefinition{value: val, isSet: true}
}

func (v NullableTraitDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraitDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
