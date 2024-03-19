/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 45.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GetDestinationMetadataV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDestinationMetadataV1Output{}

// GetDestinationMetadataV1Output Returns a Destination catalog item.
type GetDestinationMetadataV1Output struct {
	DestinationMetadata DestinationMetadataV1 `json:"destinationMetadata"`
}

// NewGetDestinationMetadataV1Output instantiates a new GetDestinationMetadataV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDestinationMetadataV1Output(
	destinationMetadata DestinationMetadataV1,
) *GetDestinationMetadataV1Output {
	this := GetDestinationMetadataV1Output{}
	this.DestinationMetadata = destinationMetadata
	return &this
}

// NewGetDestinationMetadataV1OutputWithDefaults instantiates a new GetDestinationMetadataV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDestinationMetadataV1OutputWithDefaults() *GetDestinationMetadataV1Output {
	this := GetDestinationMetadataV1Output{}
	return &this
}

// GetDestinationMetadata returns the DestinationMetadata field value
func (o *GetDestinationMetadataV1Output) GetDestinationMetadata() DestinationMetadataV1 {
	if o == nil {
		var ret DestinationMetadataV1
		return ret
	}

	return o.DestinationMetadata
}

// GetDestinationMetadataOk returns a tuple with the DestinationMetadata field value
// and a boolean to check if the value has been set.
func (o *GetDestinationMetadataV1Output) GetDestinationMetadataOk() (*DestinationMetadataV1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationMetadata, true
}

// SetDestinationMetadata sets field value
func (o *GetDestinationMetadataV1Output) SetDestinationMetadata(v DestinationMetadataV1) {
	o.DestinationMetadata = v
}

func (o GetDestinationMetadataV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDestinationMetadataV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["destinationMetadata"] = o.DestinationMetadata
	return toSerialize, nil
}

type NullableGetDestinationMetadataV1Output struct {
	value *GetDestinationMetadataV1Output
	isSet bool
}

func (v NullableGetDestinationMetadataV1Output) Get() *GetDestinationMetadataV1Output {
	return v.value
}

func (v *NullableGetDestinationMetadataV1Output) Set(val *GetDestinationMetadataV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDestinationMetadataV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDestinationMetadataV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDestinationMetadataV1Output(
	val *GetDestinationMetadataV1Output,
) *NullableGetDestinationMetadataV1Output {
	return &NullableGetDestinationMetadataV1Output{value: val, isSet: true}
}

func (v NullableGetDestinationMetadataV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDestinationMetadataV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
