/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 36.1.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// TrackSourceSettingsV1 struct for TrackSourceSettingsV1
type TrackSourceSettingsV1 struct {
	// Enable to allow unplanned track events.  Config API note: equal to `allowUnplannedTrackEvents`.
	AllowUnplannedEvents *bool `json:"allowUnplannedEvents,omitempty"`
	// Enable to allow unplanned track event properties.  Config API note: equal to `allowUnplannedTrackEventProperties`.
	AllowUnplannedEventProperties *bool `json:"allowUnplannedEventProperties,omitempty"`
	// Allow track event on violations.  Config API note: equal to `allowTrackEventOnViolations`.
	AllowEventOnViolations *bool `json:"allowEventOnViolations,omitempty"`
	// Enable to allow track properties on violations.  Config API note: equal to `allowTrackEventPropertiesOnViolations`.
	AllowPropertiesOnViolations *bool `json:"allowPropertiesOnViolations,omitempty"`
	// The common track event on violations.  Config API note: equal to `commonTrackEventOnViolations`.
	CommonEventOnViolations *string `json:"commonEventOnViolations,omitempty"`
}

// NewTrackSourceSettingsV1 instantiates a new TrackSourceSettingsV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrackSourceSettingsV1() *TrackSourceSettingsV1 {
	this := TrackSourceSettingsV1{}
	return &this
}

// NewTrackSourceSettingsV1WithDefaults instantiates a new TrackSourceSettingsV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackSourceSettingsV1WithDefaults() *TrackSourceSettingsV1 {
	this := TrackSourceSettingsV1{}
	return &this
}

// GetAllowUnplannedEvents returns the AllowUnplannedEvents field value if set, zero value otherwise.
func (o *TrackSourceSettingsV1) GetAllowUnplannedEvents() bool {
	if o == nil || o.AllowUnplannedEvents == nil {
		var ret bool
		return ret
	}
	return *o.AllowUnplannedEvents
}

// GetAllowUnplannedEventsOk returns a tuple with the AllowUnplannedEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackSourceSettingsV1) GetAllowUnplannedEventsOk() (*bool, bool) {
	if o == nil || o.AllowUnplannedEvents == nil {
		return nil, false
	}
	return o.AllowUnplannedEvents, true
}

// HasAllowUnplannedEvents returns a boolean if a field has been set.
func (o *TrackSourceSettingsV1) HasAllowUnplannedEvents() bool {
	if o != nil && o.AllowUnplannedEvents != nil {
		return true
	}

	return false
}

// SetAllowUnplannedEvents gets a reference to the given bool and assigns it to the AllowUnplannedEvents field.
func (o *TrackSourceSettingsV1) SetAllowUnplannedEvents(v bool) {
	o.AllowUnplannedEvents = &v
}

// GetAllowUnplannedEventProperties returns the AllowUnplannedEventProperties field value if set, zero value otherwise.
func (o *TrackSourceSettingsV1) GetAllowUnplannedEventProperties() bool {
	if o == nil || o.AllowUnplannedEventProperties == nil {
		var ret bool
		return ret
	}
	return *o.AllowUnplannedEventProperties
}

// GetAllowUnplannedEventPropertiesOk returns a tuple with the AllowUnplannedEventProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackSourceSettingsV1) GetAllowUnplannedEventPropertiesOk() (*bool, bool) {
	if o == nil || o.AllowUnplannedEventProperties == nil {
		return nil, false
	}
	return o.AllowUnplannedEventProperties, true
}

// HasAllowUnplannedEventProperties returns a boolean if a field has been set.
func (o *TrackSourceSettingsV1) HasAllowUnplannedEventProperties() bool {
	if o != nil && o.AllowUnplannedEventProperties != nil {
		return true
	}

	return false
}

// SetAllowUnplannedEventProperties gets a reference to the given bool and assigns it to the AllowUnplannedEventProperties field.
func (o *TrackSourceSettingsV1) SetAllowUnplannedEventProperties(v bool) {
	o.AllowUnplannedEventProperties = &v
}

// GetAllowEventOnViolations returns the AllowEventOnViolations field value if set, zero value otherwise.
func (o *TrackSourceSettingsV1) GetAllowEventOnViolations() bool {
	if o == nil || o.AllowEventOnViolations == nil {
		var ret bool
		return ret
	}
	return *o.AllowEventOnViolations
}

// GetAllowEventOnViolationsOk returns a tuple with the AllowEventOnViolations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackSourceSettingsV1) GetAllowEventOnViolationsOk() (*bool, bool) {
	if o == nil || o.AllowEventOnViolations == nil {
		return nil, false
	}
	return o.AllowEventOnViolations, true
}

// HasAllowEventOnViolations returns a boolean if a field has been set.
func (o *TrackSourceSettingsV1) HasAllowEventOnViolations() bool {
	if o != nil && o.AllowEventOnViolations != nil {
		return true
	}

	return false
}

// SetAllowEventOnViolations gets a reference to the given bool and assigns it to the AllowEventOnViolations field.
func (o *TrackSourceSettingsV1) SetAllowEventOnViolations(v bool) {
	o.AllowEventOnViolations = &v
}

// GetAllowPropertiesOnViolations returns the AllowPropertiesOnViolations field value if set, zero value otherwise.
func (o *TrackSourceSettingsV1) GetAllowPropertiesOnViolations() bool {
	if o == nil || o.AllowPropertiesOnViolations == nil {
		var ret bool
		return ret
	}
	return *o.AllowPropertiesOnViolations
}

// GetAllowPropertiesOnViolationsOk returns a tuple with the AllowPropertiesOnViolations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackSourceSettingsV1) GetAllowPropertiesOnViolationsOk() (*bool, bool) {
	if o == nil || o.AllowPropertiesOnViolations == nil {
		return nil, false
	}
	return o.AllowPropertiesOnViolations, true
}

// HasAllowPropertiesOnViolations returns a boolean if a field has been set.
func (o *TrackSourceSettingsV1) HasAllowPropertiesOnViolations() bool {
	if o != nil && o.AllowPropertiesOnViolations != nil {
		return true
	}

	return false
}

// SetAllowPropertiesOnViolations gets a reference to the given bool and assigns it to the AllowPropertiesOnViolations field.
func (o *TrackSourceSettingsV1) SetAllowPropertiesOnViolations(v bool) {
	o.AllowPropertiesOnViolations = &v
}

// GetCommonEventOnViolations returns the CommonEventOnViolations field value if set, zero value otherwise.
func (o *TrackSourceSettingsV1) GetCommonEventOnViolations() string {
	if o == nil || o.CommonEventOnViolations == nil {
		var ret string
		return ret
	}
	return *o.CommonEventOnViolations
}

// GetCommonEventOnViolationsOk returns a tuple with the CommonEventOnViolations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackSourceSettingsV1) GetCommonEventOnViolationsOk() (*string, bool) {
	if o == nil || o.CommonEventOnViolations == nil {
		return nil, false
	}
	return o.CommonEventOnViolations, true
}

// HasCommonEventOnViolations returns a boolean if a field has been set.
func (o *TrackSourceSettingsV1) HasCommonEventOnViolations() bool {
	if o != nil && o.CommonEventOnViolations != nil {
		return true
	}

	return false
}

// SetCommonEventOnViolations gets a reference to the given string and assigns it to the CommonEventOnViolations field.
func (o *TrackSourceSettingsV1) SetCommonEventOnViolations(v string) {
	o.CommonEventOnViolations = &v
}

func (o TrackSourceSettingsV1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowUnplannedEvents != nil {
		toSerialize["allowUnplannedEvents"] = o.AllowUnplannedEvents
	}
	if o.AllowUnplannedEventProperties != nil {
		toSerialize["allowUnplannedEventProperties"] = o.AllowUnplannedEventProperties
	}
	if o.AllowEventOnViolations != nil {
		toSerialize["allowEventOnViolations"] = o.AllowEventOnViolations
	}
	if o.AllowPropertiesOnViolations != nil {
		toSerialize["allowPropertiesOnViolations"] = o.AllowPropertiesOnViolations
	}
	if o.CommonEventOnViolations != nil {
		toSerialize["commonEventOnViolations"] = o.CommonEventOnViolations
	}
	return json.Marshal(toSerialize)
}

type NullableTrackSourceSettingsV1 struct {
	value *TrackSourceSettingsV1
	isSet bool
}

func (v NullableTrackSourceSettingsV1) Get() *TrackSourceSettingsV1 {
	return v.value
}

func (v *NullableTrackSourceSettingsV1) Set(val *TrackSourceSettingsV1) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackSourceSettingsV1) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackSourceSettingsV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackSourceSettingsV1(val *TrackSourceSettingsV1) *NullableTrackSourceSettingsV1 {
	return &NullableTrackSourceSettingsV1{value: val, isSet: true}
}

func (v NullableTrackSourceSettingsV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrackSourceSettingsV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
