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

// checks if the PreviewDestinationFilterV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PreviewDestinationFilterV1{}

// PreviewDestinationFilterV1 A simplified Destination filter that includes the if and actions for a DestinationFilterV1.
type PreviewDestinationFilterV1 struct {
	// A FQL statement which determines if the provided filter's actions will apply to the provided JSON payload. The literal string \"all\" will result in this filter to all events. For guidance on using FQL, see the Segment documentation site.
	If string `json:"if"`
	// The filtering action to take on events that match the \"if\" statement. Action types must be one of: \"drop\", \"allow_properties\", \"drop_properties\" or \"sample\".
	Actions []DestinationFilterActionV1 `json:"actions"`
}

// NewPreviewDestinationFilterV1 instantiates a new PreviewDestinationFilterV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreviewDestinationFilterV1(
	if_ string,
	actions []DestinationFilterActionV1,
) *PreviewDestinationFilterV1 {
	this := PreviewDestinationFilterV1{}
	this.If = if_
	this.Actions = actions
	return &this
}

// NewPreviewDestinationFilterV1WithDefaults instantiates a new PreviewDestinationFilterV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreviewDestinationFilterV1WithDefaults() *PreviewDestinationFilterV1 {
	this := PreviewDestinationFilterV1{}
	return &this
}

// GetIf returns the If field value
func (o *PreviewDestinationFilterV1) GetIf() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.If
}

// GetIfOk returns a tuple with the If field value
// and a boolean to check if the value has been set.
func (o *PreviewDestinationFilterV1) GetIfOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.If, true
}

// SetIf sets field value
func (o *PreviewDestinationFilterV1) SetIf(v string) {
	o.If = v
}

// GetActions returns the Actions field value
func (o *PreviewDestinationFilterV1) GetActions() []DestinationFilterActionV1 {
	if o == nil {
		var ret []DestinationFilterActionV1
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *PreviewDestinationFilterV1) GetActionsOk() ([]DestinationFilterActionV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Actions, true
}

// SetActions sets field value
func (o *PreviewDestinationFilterV1) SetActions(v []DestinationFilterActionV1) {
	o.Actions = v
}

func (o PreviewDestinationFilterV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PreviewDestinationFilterV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["if"] = o.If
	toSerialize["actions"] = o.Actions
	return toSerialize, nil
}

type NullablePreviewDestinationFilterV1 struct {
	value *PreviewDestinationFilterV1
	isSet bool
}

func (v NullablePreviewDestinationFilterV1) Get() *PreviewDestinationFilterV1 {
	return v.value
}

func (v *NullablePreviewDestinationFilterV1) Set(val *PreviewDestinationFilterV1) {
	v.value = val
	v.isSet = true
}

func (v NullablePreviewDestinationFilterV1) IsSet() bool {
	return v.isSet
}

func (v *NullablePreviewDestinationFilterV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreviewDestinationFilterV1(
	val *PreviewDestinationFilterV1,
) *NullablePreviewDestinationFilterV1 {
	return &NullablePreviewDestinationFilterV1{value: val, isSet: true}
}

func (v NullablePreviewDestinationFilterV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreviewDestinationFilterV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
