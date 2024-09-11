/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 54.3.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GetSourceMetadataV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSourceMetadataV1Output{}

// GetSourceMetadataV1Output Returns the Source catalog item looked up by id.
type GetSourceMetadataV1Output struct {
	SourceMetadata SourceMetadataV1 `json:"sourceMetadata"`
}

// NewGetSourceMetadataV1Output instantiates a new GetSourceMetadataV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSourceMetadataV1Output(sourceMetadata SourceMetadataV1) *GetSourceMetadataV1Output {
	this := GetSourceMetadataV1Output{}
	this.SourceMetadata = sourceMetadata
	return &this
}

// NewGetSourceMetadataV1OutputWithDefaults instantiates a new GetSourceMetadataV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSourceMetadataV1OutputWithDefaults() *GetSourceMetadataV1Output {
	this := GetSourceMetadataV1Output{}
	return &this
}

// GetSourceMetadata returns the SourceMetadata field value
func (o *GetSourceMetadataV1Output) GetSourceMetadata() SourceMetadataV1 {
	if o == nil {
		var ret SourceMetadataV1
		return ret
	}

	return o.SourceMetadata
}

// GetSourceMetadataOk returns a tuple with the SourceMetadata field value
// and a boolean to check if the value has been set.
func (o *GetSourceMetadataV1Output) GetSourceMetadataOk() (*SourceMetadataV1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceMetadata, true
}

// SetSourceMetadata sets field value
func (o *GetSourceMetadataV1Output) SetSourceMetadata(v SourceMetadataV1) {
	o.SourceMetadata = v
}

func (o GetSourceMetadataV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSourceMetadataV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sourceMetadata"] = o.SourceMetadata
	return toSerialize, nil
}

type NullableGetSourceMetadataV1Output struct {
	value *GetSourceMetadataV1Output
	isSet bool
}

func (v NullableGetSourceMetadataV1Output) Get() *GetSourceMetadataV1Output {
	return v.value
}

func (v *NullableGetSourceMetadataV1Output) Set(val *GetSourceMetadataV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSourceMetadataV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSourceMetadataV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSourceMetadataV1Output(
	val *GetSourceMetadataV1Output,
) *NullableGetSourceMetadataV1Output {
	return &NullableGetSourceMetadataV1Output{value: val, isSet: true}
}

func (v NullableGetSourceMetadataV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSourceMetadataV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
