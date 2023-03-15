/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.3.2
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Settings1 The output of Source settings.
type Settings1 struct {
	Track    *Track    `json:"track,omitempty"`
	Identify *Identify `json:"identify,omitempty"`
	Group    *Group    `json:"group,omitempty"`
	// SourceId to forward violations to.
	ForwardingViolationsTo *string `json:"forwardingViolationsTo,omitempty"`
	// SourceId to forward blocked events to.
	ForwardingBlockedEventsTo *string `json:"forwardingBlockedEventsTo,omitempty"`
}

// NewSettings1 instantiates a new Settings1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettings1() *Settings1 {
	this := Settings1{}
	return &this
}

// NewSettings1WithDefaults instantiates a new Settings1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettings1WithDefaults() *Settings1 {
	this := Settings1{}
	return &this
}

// GetTrack returns the Track field value if set, zero value otherwise.
func (o *Settings1) GetTrack() Track {
	if o == nil || o.Track == nil {
		var ret Track
		return ret
	}
	return *o.Track
}

// GetTrackOk returns a tuple with the Track field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Settings1) GetTrackOk() (*Track, bool) {
	if o == nil || o.Track == nil {
		return nil, false
	}
	return o.Track, true
}

// HasTrack returns a boolean if a field has been set.
func (o *Settings1) HasTrack() bool {
	if o != nil && o.Track != nil {
		return true
	}

	return false
}

// SetTrack gets a reference to the given Track and assigns it to the Track field.
func (o *Settings1) SetTrack(v Track) {
	o.Track = &v
}

// GetIdentify returns the Identify field value if set, zero value otherwise.
func (o *Settings1) GetIdentify() Identify {
	if o == nil || o.Identify == nil {
		var ret Identify
		return ret
	}
	return *o.Identify
}

// GetIdentifyOk returns a tuple with the Identify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Settings1) GetIdentifyOk() (*Identify, bool) {
	if o == nil || o.Identify == nil {
		return nil, false
	}
	return o.Identify, true
}

// HasIdentify returns a boolean if a field has been set.
func (o *Settings1) HasIdentify() bool {
	if o != nil && o.Identify != nil {
		return true
	}

	return false
}

// SetIdentify gets a reference to the given Identify and assigns it to the Identify field.
func (o *Settings1) SetIdentify(v Identify) {
	o.Identify = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *Settings1) GetGroup() Group {
	if o == nil || o.Group == nil {
		var ret Group
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Settings1) GetGroupOk() (*Group, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *Settings1) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given Group and assigns it to the Group field.
func (o *Settings1) SetGroup(v Group) {
	o.Group = &v
}

// GetForwardingViolationsTo returns the ForwardingViolationsTo field value if set, zero value otherwise.
func (o *Settings1) GetForwardingViolationsTo() string {
	if o == nil || o.ForwardingViolationsTo == nil {
		var ret string
		return ret
	}
	return *o.ForwardingViolationsTo
}

// GetForwardingViolationsToOk returns a tuple with the ForwardingViolationsTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Settings1) GetForwardingViolationsToOk() (*string, bool) {
	if o == nil || o.ForwardingViolationsTo == nil {
		return nil, false
	}
	return o.ForwardingViolationsTo, true
}

// HasForwardingViolationsTo returns a boolean if a field has been set.
func (o *Settings1) HasForwardingViolationsTo() bool {
	if o != nil && o.ForwardingViolationsTo != nil {
		return true
	}

	return false
}

// SetForwardingViolationsTo gets a reference to the given string and assigns it to the ForwardingViolationsTo field.
func (o *Settings1) SetForwardingViolationsTo(v string) {
	o.ForwardingViolationsTo = &v
}

// GetForwardingBlockedEventsTo returns the ForwardingBlockedEventsTo field value if set, zero value otherwise.
func (o *Settings1) GetForwardingBlockedEventsTo() string {
	if o == nil || o.ForwardingBlockedEventsTo == nil {
		var ret string
		return ret
	}
	return *o.ForwardingBlockedEventsTo
}

// GetForwardingBlockedEventsToOk returns a tuple with the ForwardingBlockedEventsTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Settings1) GetForwardingBlockedEventsToOk() (*string, bool) {
	if o == nil || o.ForwardingBlockedEventsTo == nil {
		return nil, false
	}
	return o.ForwardingBlockedEventsTo, true
}

// HasForwardingBlockedEventsTo returns a boolean if a field has been set.
func (o *Settings1) HasForwardingBlockedEventsTo() bool {
	if o != nil && o.ForwardingBlockedEventsTo != nil {
		return true
	}

	return false
}

// SetForwardingBlockedEventsTo gets a reference to the given string and assigns it to the ForwardingBlockedEventsTo field.
func (o *Settings1) SetForwardingBlockedEventsTo(v string) {
	o.ForwardingBlockedEventsTo = &v
}

func (o Settings1) MarshalJSON() ([]byte, error) {
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

type NullableSettings1 struct {
	value *Settings1
	isSet bool
}

func (v NullableSettings1) Get() *Settings1 {
	return v.value
}

func (v *NullableSettings1) Set(val *Settings1) {
	v.value = val
	v.isSet = true
}

func (v NullableSettings1) IsSet() bool {
	return v.isSet
}

func (v *NullableSettings1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettings1(val *Settings1) *NullableSettings1 {
	return &NullableSettings1{value: val, isSet: true}
}

func (v NullableSettings1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettings1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
