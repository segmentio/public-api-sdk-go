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

// checks if the PreviewDestinationFilterV1Input type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PreviewDestinationFilterV1Input{}

// PreviewDestinationFilterV1Input Input of the Destination filter to preview. For guidance on using FQL, see the Segment documentation site.
type PreviewDestinationFilterV1Input struct {
	Filter PreviewDestinationFilterV1 `json:"filter"`
	// The JSON payload to apply the filter to.
	Payload map[string]interface{} `json:"payload"`
}

// NewPreviewDestinationFilterV1Input instantiates a new PreviewDestinationFilterV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreviewDestinationFilterV1Input(
	filter PreviewDestinationFilterV1,
	payload map[string]interface{},
) *PreviewDestinationFilterV1Input {
	this := PreviewDestinationFilterV1Input{}
	this.Filter = filter
	this.Payload = payload
	return &this
}

// NewPreviewDestinationFilterV1InputWithDefaults instantiates a new PreviewDestinationFilterV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreviewDestinationFilterV1InputWithDefaults() *PreviewDestinationFilterV1Input {
	this := PreviewDestinationFilterV1Input{}
	return &this
}

// GetFilter returns the Filter field value
func (o *PreviewDestinationFilterV1Input) GetFilter() PreviewDestinationFilterV1 {
	if o == nil {
		var ret PreviewDestinationFilterV1
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *PreviewDestinationFilterV1Input) GetFilterOk() (*PreviewDestinationFilterV1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value
func (o *PreviewDestinationFilterV1Input) SetFilter(v PreviewDestinationFilterV1) {
	o.Filter = v
}

// GetPayload returns the Payload field value
func (o *PreviewDestinationFilterV1Input) GetPayload() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *PreviewDestinationFilterV1Input) GetPayloadOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Payload, true
}

// SetPayload sets field value
func (o *PreviewDestinationFilterV1Input) SetPayload(v map[string]interface{}) {
	o.Payload = v
}

func (o PreviewDestinationFilterV1Input) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PreviewDestinationFilterV1Input) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["filter"] = o.Filter
	toSerialize["payload"] = o.Payload
	return toSerialize, nil
}

type NullablePreviewDestinationFilterV1Input struct {
	value *PreviewDestinationFilterV1Input
	isSet bool
}

func (v NullablePreviewDestinationFilterV1Input) Get() *PreviewDestinationFilterV1Input {
	return v.value
}

func (v *NullablePreviewDestinationFilterV1Input) Set(val *PreviewDestinationFilterV1Input) {
	v.value = val
	v.isSet = true
}

func (v NullablePreviewDestinationFilterV1Input) IsSet() bool {
	return v.isSet
}

func (v *NullablePreviewDestinationFilterV1Input) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreviewDestinationFilterV1Input(
	val *PreviewDestinationFilterV1Input,
) *NullablePreviewDestinationFilterV1Input {
	return &NullablePreviewDestinationFilterV1Input{value: val, isSet: true}
}

func (v NullablePreviewDestinationFilterV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreviewDestinationFilterV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
