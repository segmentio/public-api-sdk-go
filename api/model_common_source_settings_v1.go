/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 54.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the CommonSourceSettingsV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonSourceSettingsV1{}

// CommonSourceSettingsV1 struct for CommonSourceSettingsV1
type CommonSourceSettingsV1 struct {
	Track    *TrackSourceSettingsV1    `json:"track,omitempty"`
	Identify *IdentifySourceSettingsV1 `json:"identify,omitempty"`
	Group    *GroupSourceSettingsV1    `json:"group,omitempty"`
	// SourceId to forward violations to.
	ForwardingViolationsTo *string `json:"forwardingViolationsTo,omitempty"`
	// SourceId to forward blocked events to.
	ForwardingBlockedEventsTo *string `json:"forwardingBlockedEventsTo,omitempty"`
}

// NewCommonSourceSettingsV1 instantiates a new CommonSourceSettingsV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonSourceSettingsV1() *CommonSourceSettingsV1 {
	this := CommonSourceSettingsV1{}
	return &this
}

// NewCommonSourceSettingsV1WithDefaults instantiates a new CommonSourceSettingsV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonSourceSettingsV1WithDefaults() *CommonSourceSettingsV1 {
	this := CommonSourceSettingsV1{}
	return &this
}

// GetTrack returns the Track field value if set, zero value otherwise.
func (o *CommonSourceSettingsV1) GetTrack() TrackSourceSettingsV1 {
	if o == nil || IsNil(o.Track) {
		var ret TrackSourceSettingsV1
		return ret
	}
	return *o.Track
}

// GetTrackOk returns a tuple with the Track field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonSourceSettingsV1) GetTrackOk() (*TrackSourceSettingsV1, bool) {
	if o == nil || IsNil(o.Track) {
		return nil, false
	}
	return o.Track, true
}

// HasTrack returns a boolean if a field has been set.
func (o *CommonSourceSettingsV1) HasTrack() bool {
	if o != nil && !IsNil(o.Track) {
		return true
	}

	return false
}

// SetTrack gets a reference to the given TrackSourceSettingsV1 and assigns it to the Track field.
func (o *CommonSourceSettingsV1) SetTrack(v TrackSourceSettingsV1) {
	o.Track = &v
}

// GetIdentify returns the Identify field value if set, zero value otherwise.
func (o *CommonSourceSettingsV1) GetIdentify() IdentifySourceSettingsV1 {
	if o == nil || IsNil(o.Identify) {
		var ret IdentifySourceSettingsV1
		return ret
	}
	return *o.Identify
}

// GetIdentifyOk returns a tuple with the Identify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonSourceSettingsV1) GetIdentifyOk() (*IdentifySourceSettingsV1, bool) {
	if o == nil || IsNil(o.Identify) {
		return nil, false
	}
	return o.Identify, true
}

// HasIdentify returns a boolean if a field has been set.
func (o *CommonSourceSettingsV1) HasIdentify() bool {
	if o != nil && !IsNil(o.Identify) {
		return true
	}

	return false
}

// SetIdentify gets a reference to the given IdentifySourceSettingsV1 and assigns it to the Identify field.
func (o *CommonSourceSettingsV1) SetIdentify(v IdentifySourceSettingsV1) {
	o.Identify = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *CommonSourceSettingsV1) GetGroup() GroupSourceSettingsV1 {
	if o == nil || IsNil(o.Group) {
		var ret GroupSourceSettingsV1
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonSourceSettingsV1) GetGroupOk() (*GroupSourceSettingsV1, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *CommonSourceSettingsV1) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given GroupSourceSettingsV1 and assigns it to the Group field.
func (o *CommonSourceSettingsV1) SetGroup(v GroupSourceSettingsV1) {
	o.Group = &v
}

// GetForwardingViolationsTo returns the ForwardingViolationsTo field value if set, zero value otherwise.
func (o *CommonSourceSettingsV1) GetForwardingViolationsTo() string {
	if o == nil || IsNil(o.ForwardingViolationsTo) {
		var ret string
		return ret
	}
	return *o.ForwardingViolationsTo
}

// GetForwardingViolationsToOk returns a tuple with the ForwardingViolationsTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonSourceSettingsV1) GetForwardingViolationsToOk() (*string, bool) {
	if o == nil || IsNil(o.ForwardingViolationsTo) {
		return nil, false
	}
	return o.ForwardingViolationsTo, true
}

// HasForwardingViolationsTo returns a boolean if a field has been set.
func (o *CommonSourceSettingsV1) HasForwardingViolationsTo() bool {
	if o != nil && !IsNil(o.ForwardingViolationsTo) {
		return true
	}

	return false
}

// SetForwardingViolationsTo gets a reference to the given string and assigns it to the ForwardingViolationsTo field.
func (o *CommonSourceSettingsV1) SetForwardingViolationsTo(v string) {
	o.ForwardingViolationsTo = &v
}

// GetForwardingBlockedEventsTo returns the ForwardingBlockedEventsTo field value if set, zero value otherwise.
func (o *CommonSourceSettingsV1) GetForwardingBlockedEventsTo() string {
	if o == nil || IsNil(o.ForwardingBlockedEventsTo) {
		var ret string
		return ret
	}
	return *o.ForwardingBlockedEventsTo
}

// GetForwardingBlockedEventsToOk returns a tuple with the ForwardingBlockedEventsTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonSourceSettingsV1) GetForwardingBlockedEventsToOk() (*string, bool) {
	if o == nil || IsNil(o.ForwardingBlockedEventsTo) {
		return nil, false
	}
	return o.ForwardingBlockedEventsTo, true
}

// HasForwardingBlockedEventsTo returns a boolean if a field has been set.
func (o *CommonSourceSettingsV1) HasForwardingBlockedEventsTo() bool {
	if o != nil && !IsNil(o.ForwardingBlockedEventsTo) {
		return true
	}

	return false
}

// SetForwardingBlockedEventsTo gets a reference to the given string and assigns it to the ForwardingBlockedEventsTo field.
func (o *CommonSourceSettingsV1) SetForwardingBlockedEventsTo(v string) {
	o.ForwardingBlockedEventsTo = &v
}

func (o CommonSourceSettingsV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonSourceSettingsV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Track) {
		toSerialize["track"] = o.Track
	}
	if !IsNil(o.Identify) {
		toSerialize["identify"] = o.Identify
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.ForwardingViolationsTo) {
		toSerialize["forwardingViolationsTo"] = o.ForwardingViolationsTo
	}
	if !IsNil(o.ForwardingBlockedEventsTo) {
		toSerialize["forwardingBlockedEventsTo"] = o.ForwardingBlockedEventsTo
	}
	return toSerialize, nil
}

type NullableCommonSourceSettingsV1 struct {
	value *CommonSourceSettingsV1
	isSet bool
}

func (v NullableCommonSourceSettingsV1) Get() *CommonSourceSettingsV1 {
	return v.value
}

func (v *NullableCommonSourceSettingsV1) Set(val *CommonSourceSettingsV1) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonSourceSettingsV1) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonSourceSettingsV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonSourceSettingsV1(
	val *CommonSourceSettingsV1,
) *NullableCommonSourceSettingsV1 {
	return &NullableCommonSourceSettingsV1{value: val, isSet: true}
}

func (v NullableCommonSourceSettingsV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonSourceSettingsV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
