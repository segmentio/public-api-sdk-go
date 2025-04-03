/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.0.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the AudienceDefinitionBeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AudienceDefinitionBeta{}

// AudienceDefinitionBeta Defines an audience definition.
type AudienceDefinitionBeta struct {
	// The query language string defining the audience segmentation criteria.
	Query string `json:"query"`
	// The underlying data type being segmented for this audience.  Possible values: users, accounts.
	Type string `json:"type"`
}

// NewAudienceDefinitionBeta instantiates a new AudienceDefinitionBeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAudienceDefinitionBeta(query string, type_ string) *AudienceDefinitionBeta {
	this := AudienceDefinitionBeta{}
	this.Query = query
	this.Type = type_
	return &this
}

// NewAudienceDefinitionBetaWithDefaults instantiates a new AudienceDefinitionBeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAudienceDefinitionBetaWithDefaults() *AudienceDefinitionBeta {
	this := AudienceDefinitionBeta{}
	return &this
}

// GetQuery returns the Query field value
func (o *AudienceDefinitionBeta) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *AudienceDefinitionBeta) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *AudienceDefinitionBeta) SetQuery(v string) {
	o.Query = v
}

// GetType returns the Type field value
func (o *AudienceDefinitionBeta) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AudienceDefinitionBeta) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AudienceDefinitionBeta) SetType(v string) {
	o.Type = v
}

func (o AudienceDefinitionBeta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AudienceDefinitionBeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["query"] = o.Query
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableAudienceDefinitionBeta struct {
	value *AudienceDefinitionBeta
	isSet bool
}

func (v NullableAudienceDefinitionBeta) Get() *AudienceDefinitionBeta {
	return v.value
}

func (v *NullableAudienceDefinitionBeta) Set(val *AudienceDefinitionBeta) {
	v.value = val
	v.isSet = true
}

func (v NullableAudienceDefinitionBeta) IsSet() bool {
	return v.isSet
}

func (v *NullableAudienceDefinitionBeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAudienceDefinitionBeta(
	val *AudienceDefinitionBeta,
) *NullableAudienceDefinitionBeta {
	return &NullableAudienceDefinitionBeta{value: val, isSet: true}
}

func (v NullableAudienceDefinitionBeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAudienceDefinitionBeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
