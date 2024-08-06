/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 52.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the DeliveryOverviewDestinationFilterBy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryOverviewDestinationFilterBy{}

// DeliveryOverviewDestinationFilterBy The `DeliveryOverviewDestinationFilterBy` object is a map of the filterable fields and their values.
type DeliveryOverviewDestinationFilterBy struct {
	// A list of strings of discard reasons.  See [Discard Record Documentation](https://segment.com/docs/connections/delivery-overview/#troubleshooting) for valid error codes.
	DiscardReason []string `json:"discardReason,omitempty"`
	// A list of strings of event names.
	EventName []string `json:"eventName,omitempty"`
	// A list of strings of event types. Valid options are: `alias`, `group`, `identify`, `page`, `screen`, and `track`.
	EventType []string `json:"eventType,omitempty"`
	// A list of strings of app versions.
	AppVersion []string `json:"appVersion,omitempty"`
	// A list of strings of subscriptionIDs for Actions Destinations.
	SubscriptionId []string `json:"subscriptionId,omitempty"`
	// A list of strings of event context IDs from a Linked Audience mapping/activation.
	ActivationId []string `json:"activationId,omitempty"`
	// A list of strings of audienceIDs for a Linked Audience.
	AudienceId []string `json:"audienceId,omitempty"`
	// A list of strings of spaceIDs for a Linked Audience.
	SpaceId []string `json:"spaceId,omitempty"`
}

// NewDeliveryOverviewDestinationFilterBy instantiates a new DeliveryOverviewDestinationFilterBy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryOverviewDestinationFilterBy() *DeliveryOverviewDestinationFilterBy {
	this := DeliveryOverviewDestinationFilterBy{}
	return &this
}

// NewDeliveryOverviewDestinationFilterByWithDefaults instantiates a new DeliveryOverviewDestinationFilterBy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryOverviewDestinationFilterByWithDefaults() *DeliveryOverviewDestinationFilterBy {
	this := DeliveryOverviewDestinationFilterBy{}
	return &this
}

// GetDiscardReason returns the DiscardReason field value if set, zero value otherwise.
func (o *DeliveryOverviewDestinationFilterBy) GetDiscardReason() []string {
	if o == nil || IsNil(o.DiscardReason) {
		var ret []string
		return ret
	}
	return o.DiscardReason
}

// GetDiscardReasonOk returns a tuple with the DiscardReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryOverviewDestinationFilterBy) GetDiscardReasonOk() ([]string, bool) {
	if o == nil || IsNil(o.DiscardReason) {
		return nil, false
	}
	return o.DiscardReason, true
}

// HasDiscardReason returns a boolean if a field has been set.
func (o *DeliveryOverviewDestinationFilterBy) HasDiscardReason() bool {
	if o != nil && !IsNil(o.DiscardReason) {
		return true
	}

	return false
}

// SetDiscardReason gets a reference to the given []string and assigns it to the DiscardReason field.
func (o *DeliveryOverviewDestinationFilterBy) SetDiscardReason(v []string) {
	o.DiscardReason = v
}

// GetEventName returns the EventName field value if set, zero value otherwise.
func (o *DeliveryOverviewDestinationFilterBy) GetEventName() []string {
	if o == nil || IsNil(o.EventName) {
		var ret []string
		return ret
	}
	return o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryOverviewDestinationFilterBy) GetEventNameOk() ([]string, bool) {
	if o == nil || IsNil(o.EventName) {
		return nil, false
	}
	return o.EventName, true
}

// HasEventName returns a boolean if a field has been set.
func (o *DeliveryOverviewDestinationFilterBy) HasEventName() bool {
	if o != nil && !IsNil(o.EventName) {
		return true
	}

	return false
}

// SetEventName gets a reference to the given []string and assigns it to the EventName field.
func (o *DeliveryOverviewDestinationFilterBy) SetEventName(v []string) {
	o.EventName = v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *DeliveryOverviewDestinationFilterBy) GetEventType() []string {
	if o == nil || IsNil(o.EventType) {
		var ret []string
		return ret
	}
	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryOverviewDestinationFilterBy) GetEventTypeOk() ([]string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *DeliveryOverviewDestinationFilterBy) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given []string and assigns it to the EventType field.
func (o *DeliveryOverviewDestinationFilterBy) SetEventType(v []string) {
	o.EventType = v
}

// GetAppVersion returns the AppVersion field value if set, zero value otherwise.
func (o *DeliveryOverviewDestinationFilterBy) GetAppVersion() []string {
	if o == nil || IsNil(o.AppVersion) {
		var ret []string
		return ret
	}
	return o.AppVersion
}

// GetAppVersionOk returns a tuple with the AppVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryOverviewDestinationFilterBy) GetAppVersionOk() ([]string, bool) {
	if o == nil || IsNil(o.AppVersion) {
		return nil, false
	}
	return o.AppVersion, true
}

// HasAppVersion returns a boolean if a field has been set.
func (o *DeliveryOverviewDestinationFilterBy) HasAppVersion() bool {
	if o != nil && !IsNil(o.AppVersion) {
		return true
	}

	return false
}

// SetAppVersion gets a reference to the given []string and assigns it to the AppVersion field.
func (o *DeliveryOverviewDestinationFilterBy) SetAppVersion(v []string) {
	o.AppVersion = v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *DeliveryOverviewDestinationFilterBy) GetSubscriptionId() []string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret []string
		return ret
	}
	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryOverviewDestinationFilterBy) GetSubscriptionIdOk() ([]string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *DeliveryOverviewDestinationFilterBy) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given []string and assigns it to the SubscriptionId field.
func (o *DeliveryOverviewDestinationFilterBy) SetSubscriptionId(v []string) {
	o.SubscriptionId = v
}

// GetActivationId returns the ActivationId field value if set, zero value otherwise.
func (o *DeliveryOverviewDestinationFilterBy) GetActivationId() []string {
	if o == nil || IsNil(o.ActivationId) {
		var ret []string
		return ret
	}
	return o.ActivationId
}

// GetActivationIdOk returns a tuple with the ActivationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryOverviewDestinationFilterBy) GetActivationIdOk() ([]string, bool) {
	if o == nil || IsNil(o.ActivationId) {
		return nil, false
	}
	return o.ActivationId, true
}

// HasActivationId returns a boolean if a field has been set.
func (o *DeliveryOverviewDestinationFilterBy) HasActivationId() bool {
	if o != nil && !IsNil(o.ActivationId) {
		return true
	}

	return false
}

// SetActivationId gets a reference to the given []string and assigns it to the ActivationId field.
func (o *DeliveryOverviewDestinationFilterBy) SetActivationId(v []string) {
	o.ActivationId = v
}

// GetAudienceId returns the AudienceId field value if set, zero value otherwise.
func (o *DeliveryOverviewDestinationFilterBy) GetAudienceId() []string {
	if o == nil || IsNil(o.AudienceId) {
		var ret []string
		return ret
	}
	return o.AudienceId
}

// GetAudienceIdOk returns a tuple with the AudienceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryOverviewDestinationFilterBy) GetAudienceIdOk() ([]string, bool) {
	if o == nil || IsNil(o.AudienceId) {
		return nil, false
	}
	return o.AudienceId, true
}

// HasAudienceId returns a boolean if a field has been set.
func (o *DeliveryOverviewDestinationFilterBy) HasAudienceId() bool {
	if o != nil && !IsNil(o.AudienceId) {
		return true
	}

	return false
}

// SetAudienceId gets a reference to the given []string and assigns it to the AudienceId field.
func (o *DeliveryOverviewDestinationFilterBy) SetAudienceId(v []string) {
	o.AudienceId = v
}

// GetSpaceId returns the SpaceId field value if set, zero value otherwise.
func (o *DeliveryOverviewDestinationFilterBy) GetSpaceId() []string {
	if o == nil || IsNil(o.SpaceId) {
		var ret []string
		return ret
	}
	return o.SpaceId
}

// GetSpaceIdOk returns a tuple with the SpaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryOverviewDestinationFilterBy) GetSpaceIdOk() ([]string, bool) {
	if o == nil || IsNil(o.SpaceId) {
		return nil, false
	}
	return o.SpaceId, true
}

// HasSpaceId returns a boolean if a field has been set.
func (o *DeliveryOverviewDestinationFilterBy) HasSpaceId() bool {
	if o != nil && !IsNil(o.SpaceId) {
		return true
	}

	return false
}

// SetSpaceId gets a reference to the given []string and assigns it to the SpaceId field.
func (o *DeliveryOverviewDestinationFilterBy) SetSpaceId(v []string) {
	o.SpaceId = v
}

func (o DeliveryOverviewDestinationFilterBy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliveryOverviewDestinationFilterBy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DiscardReason) {
		toSerialize["discardReason"] = o.DiscardReason
	}
	if !IsNil(o.EventName) {
		toSerialize["eventName"] = o.EventName
	}
	if !IsNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	if !IsNil(o.AppVersion) {
		toSerialize["appVersion"] = o.AppVersion
	}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if !IsNil(o.ActivationId) {
		toSerialize["activationId"] = o.ActivationId
	}
	if !IsNil(o.AudienceId) {
		toSerialize["audienceId"] = o.AudienceId
	}
	if !IsNil(o.SpaceId) {
		toSerialize["spaceId"] = o.SpaceId
	}
	return toSerialize, nil
}

type NullableDeliveryOverviewDestinationFilterBy struct {
	value *DeliveryOverviewDestinationFilterBy
	isSet bool
}

func (v NullableDeliveryOverviewDestinationFilterBy) Get() *DeliveryOverviewDestinationFilterBy {
	return v.value
}

func (v *NullableDeliveryOverviewDestinationFilterBy) Set(
	val *DeliveryOverviewDestinationFilterBy,
) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryOverviewDestinationFilterBy) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryOverviewDestinationFilterBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryOverviewDestinationFilterBy(
	val *DeliveryOverviewDestinationFilterBy,
) *NullableDeliveryOverviewDestinationFilterBy {
	return &NullableDeliveryOverviewDestinationFilterBy{value: val, isSet: true}
}

func (v NullableDeliveryOverviewDestinationFilterBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryOverviewDestinationFilterBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
