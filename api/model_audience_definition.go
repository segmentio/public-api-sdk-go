/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 44.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the AudienceDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AudienceDefinition{}

// AudienceDefinition Defines an audience definition.
type AudienceDefinition struct {
	// The query language string defining the audience segmentation criteria.
	Query string `json:"query"`
	// The underlying data type being segmented for this audience.  Possible values: users, accounts.
	Type string `json:"type"`
}

// NewAudienceDefinition instantiates a new AudienceDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAudienceDefinition(query string, type_ string) *AudienceDefinition {
	this := AudienceDefinition{}
	this.Query = query
	this.Type = type_
	return &this
}

// NewAudienceDefinitionWithDefaults instantiates a new AudienceDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAudienceDefinitionWithDefaults() *AudienceDefinition {
	this := AudienceDefinition{}
	return &this
}

// GetQuery returns the Query field value
func (o *AudienceDefinition) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *AudienceDefinition) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *AudienceDefinition) SetQuery(v string) {
	o.Query = v
}

// GetType returns the Type field value
func (o *AudienceDefinition) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AudienceDefinition) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AudienceDefinition) SetType(v string) {
	o.Type = v
}

func (o AudienceDefinition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AudienceDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["query"] = o.Query
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableAudienceDefinition struct {
	value *AudienceDefinition
	isSet bool
}

func (v NullableAudienceDefinition) Get() *AudienceDefinition {
	return v.value
}

func (v *NullableAudienceDefinition) Set(val *AudienceDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableAudienceDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableAudienceDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAudienceDefinition(val *AudienceDefinition) *NullableAudienceDefinition {
	return &NullableAudienceDefinition{value: val, isSet: true}
}

func (v NullableAudienceDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAudienceDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
