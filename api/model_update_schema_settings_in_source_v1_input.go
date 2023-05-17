/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.3.6
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UpdateSchemaSettingsInSourceV1Input Input to update a Source's settings.
type UpdateSchemaSettingsInSourceV1Input struct {
	Track    *Track    `json:"track,omitempty"`
	Identify *Identify `json:"identify,omitempty"`
	Group    *Group    `json:"group,omitempty"`
	// Source id to forward violations to.
	ForwardingViolationsTo *string `json:"forwardingViolationsTo,omitempty"`
	// Source id to forward blocked events to.
	ForwardingBlockedEventsTo *string `json:"forwardingBlockedEventsTo,omitempty"`
}

// NewUpdateSchemaSettingsInSourceV1Input instantiates a new UpdateSchemaSettingsInSourceV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSchemaSettingsInSourceV1Input() *UpdateSchemaSettingsInSourceV1Input {
	this := UpdateSchemaSettingsInSourceV1Input{}
	return &this
}

// NewUpdateSchemaSettingsInSourceV1InputWithDefaults instantiates a new UpdateSchemaSettingsInSourceV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSchemaSettingsInSourceV1InputWithDefaults() *UpdateSchemaSettingsInSourceV1Input {
	this := UpdateSchemaSettingsInSourceV1Input{}
	return &this
}

// GetTrack returns the Track field value if set, zero value otherwise.
func (o *UpdateSchemaSettingsInSourceV1Input) GetTrack() Track {
	if o == nil || o.Track == nil {
		var ret Track
		return ret
	}
	return *o.Track
}

// GetTrackOk returns a tuple with the Track field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSchemaSettingsInSourceV1Input) GetTrackOk() (*Track, bool) {
	if o == nil || o.Track == nil {
		return nil, false
	}
	return o.Track, true
}

// HasTrack returns a boolean if a field has been set.
func (o *UpdateSchemaSettingsInSourceV1Input) HasTrack() bool {
	if o != nil && o.Track != nil {
		return true
	}

	return false
}

// SetTrack gets a reference to the given Track and assigns it to the Track field.
func (o *UpdateSchemaSettingsInSourceV1Input) SetTrack(v Track) {
	o.Track = &v
}

// GetIdentify returns the Identify field value if set, zero value otherwise.
func (o *UpdateSchemaSettingsInSourceV1Input) GetIdentify() Identify {
	if o == nil || o.Identify == nil {
		var ret Identify
		return ret
	}
	return *o.Identify
}

// GetIdentifyOk returns a tuple with the Identify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSchemaSettingsInSourceV1Input) GetIdentifyOk() (*Identify, bool) {
	if o == nil || o.Identify == nil {
		return nil, false
	}
	return o.Identify, true
}

// HasIdentify returns a boolean if a field has been set.
func (o *UpdateSchemaSettingsInSourceV1Input) HasIdentify() bool {
	if o != nil && o.Identify != nil {
		return true
	}

	return false
}

// SetIdentify gets a reference to the given Identify and assigns it to the Identify field.
func (o *UpdateSchemaSettingsInSourceV1Input) SetIdentify(v Identify) {
	o.Identify = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *UpdateSchemaSettingsInSourceV1Input) GetGroup() Group {
	if o == nil || o.Group == nil {
		var ret Group
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSchemaSettingsInSourceV1Input) GetGroupOk() (*Group, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *UpdateSchemaSettingsInSourceV1Input) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given Group and assigns it to the Group field.
func (o *UpdateSchemaSettingsInSourceV1Input) SetGroup(v Group) {
	o.Group = &v
}

// GetForwardingViolationsTo returns the ForwardingViolationsTo field value if set, zero value otherwise.
func (o *UpdateSchemaSettingsInSourceV1Input) GetForwardingViolationsTo() string {
	if o == nil || o.ForwardingViolationsTo == nil {
		var ret string
		return ret
	}
	return *o.ForwardingViolationsTo
}

// GetForwardingViolationsToOk returns a tuple with the ForwardingViolationsTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSchemaSettingsInSourceV1Input) GetForwardingViolationsToOk() (*string, bool) {
	if o == nil || o.ForwardingViolationsTo == nil {
		return nil, false
	}
	return o.ForwardingViolationsTo, true
}

// HasForwardingViolationsTo returns a boolean if a field has been set.
func (o *UpdateSchemaSettingsInSourceV1Input) HasForwardingViolationsTo() bool {
	if o != nil && o.ForwardingViolationsTo != nil {
		return true
	}

	return false
}

// SetForwardingViolationsTo gets a reference to the given string and assigns it to the ForwardingViolationsTo field.
func (o *UpdateSchemaSettingsInSourceV1Input) SetForwardingViolationsTo(v string) {
	o.ForwardingViolationsTo = &v
}

// GetForwardingBlockedEventsTo returns the ForwardingBlockedEventsTo field value if set, zero value otherwise.
func (o *UpdateSchemaSettingsInSourceV1Input) GetForwardingBlockedEventsTo() string {
	if o == nil || o.ForwardingBlockedEventsTo == nil {
		var ret string
		return ret
	}
	return *o.ForwardingBlockedEventsTo
}

// GetForwardingBlockedEventsToOk returns a tuple with the ForwardingBlockedEventsTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSchemaSettingsInSourceV1Input) GetForwardingBlockedEventsToOk() (*string, bool) {
	if o == nil || o.ForwardingBlockedEventsTo == nil {
		return nil, false
	}
	return o.ForwardingBlockedEventsTo, true
}

// HasForwardingBlockedEventsTo returns a boolean if a field has been set.
func (o *UpdateSchemaSettingsInSourceV1Input) HasForwardingBlockedEventsTo() bool {
	if o != nil && o.ForwardingBlockedEventsTo != nil {
		return true
	}

	return false
}

// SetForwardingBlockedEventsTo gets a reference to the given string and assigns it to the ForwardingBlockedEventsTo field.
func (o *UpdateSchemaSettingsInSourceV1Input) SetForwardingBlockedEventsTo(v string) {
	o.ForwardingBlockedEventsTo = &v
}

func (o UpdateSchemaSettingsInSourceV1Input) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Track != nil {
		toSerialize["track"] = o.Track
	}
	if o.Identify != nil {
		toSerialize["identify"] = o.Identify
	}
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	if o.ForwardingViolationsTo != nil {
		toSerialize["forwardingViolationsTo"] = o.ForwardingViolationsTo
	}
	if o.ForwardingBlockedEventsTo != nil {
		toSerialize["forwardingBlockedEventsTo"] = o.ForwardingBlockedEventsTo
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateSchemaSettingsInSourceV1Input struct {
	value *UpdateSchemaSettingsInSourceV1Input
	isSet bool
}

func (v NullableUpdateSchemaSettingsInSourceV1Input) Get() *UpdateSchemaSettingsInSourceV1Input {
	return v.value
}

func (v *NullableUpdateSchemaSettingsInSourceV1Input) Set(
	val *UpdateSchemaSettingsInSourceV1Input,
) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSchemaSettingsInSourceV1Input) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSchemaSettingsInSourceV1Input) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSchemaSettingsInSourceV1Input(
	val *UpdateSchemaSettingsInSourceV1Input,
) *NullableUpdateSchemaSettingsInSourceV1Input {
	return &NullableUpdateSchemaSettingsInSourceV1Input{value: val, isSet: true}
}

func (v NullableUpdateSchemaSettingsInSourceV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSchemaSettingsInSourceV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
