/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.3.4
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ListSubscriptionsFromDestinationAlphaOutput Output for ListDestinationSubscriptionsAlpha.
type ListSubscriptionsFromDestinationAlphaOutput struct {
	// A list of Destination subscriptions.
	Subscriptions []DestinationSubscription `json:"subscriptions"`
	Pagination    *Pagination               `json:"pagination,omitempty"`
}

// NewListSubscriptionsFromDestinationAlphaOutput instantiates a new ListSubscriptionsFromDestinationAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSubscriptionsFromDestinationAlphaOutput(
	subscriptions []DestinationSubscription,
) *ListSubscriptionsFromDestinationAlphaOutput {
	this := ListSubscriptionsFromDestinationAlphaOutput{}
	this.Subscriptions = subscriptions
	return &this
}

// NewListSubscriptionsFromDestinationAlphaOutputWithDefaults instantiates a new ListSubscriptionsFromDestinationAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSubscriptionsFromDestinationAlphaOutputWithDefaults() *ListSubscriptionsFromDestinationAlphaOutput {
	this := ListSubscriptionsFromDestinationAlphaOutput{}
	return &this
}

// GetSubscriptions returns the Subscriptions field value
func (o *ListSubscriptionsFromDestinationAlphaOutput) GetSubscriptions() []DestinationSubscription {
	if o == nil {
		var ret []DestinationSubscription
		return ret
	}

	return o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value
// and a boolean to check if the value has been set.
func (o *ListSubscriptionsFromDestinationAlphaOutput) GetSubscriptionsOk() ([]DestinationSubscription, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subscriptions, true
}

// SetSubscriptions sets field value
func (o *ListSubscriptionsFromDestinationAlphaOutput) SetSubscriptions(
	v []DestinationSubscription,
) {
	o.Subscriptions = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListSubscriptionsFromDestinationAlphaOutput) GetPagination() Pagination {
	if o == nil || o.Pagination == nil {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSubscriptionsFromDestinationAlphaOutput) GetPaginationOk() (*Pagination, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ListSubscriptionsFromDestinationAlphaOutput) HasPagination() bool {
	if o != nil && o.Pagination != nil {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *ListSubscriptionsFromDestinationAlphaOutput) SetPagination(v Pagination) {
	o.Pagination = &v
}

func (o ListSubscriptionsFromDestinationAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["subscriptions"] = o.Subscriptions
	}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableListSubscriptionsFromDestinationAlphaOutput struct {
	value *ListSubscriptionsFromDestinationAlphaOutput
	isSet bool
}

func (v NullableListSubscriptionsFromDestinationAlphaOutput) Get() *ListSubscriptionsFromDestinationAlphaOutput {
	return v.value
}

func (v *NullableListSubscriptionsFromDestinationAlphaOutput) Set(
	val *ListSubscriptionsFromDestinationAlphaOutput,
) {
	v.value = val
	v.isSet = true
}

func (v NullableListSubscriptionsFromDestinationAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableListSubscriptionsFromDestinationAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSubscriptionsFromDestinationAlphaOutput(
	val *ListSubscriptionsFromDestinationAlphaOutput,
) *NullableListSubscriptionsFromDestinationAlphaOutput {
	return &NullableListSubscriptionsFromDestinationAlphaOutput{value: val, isSet: true}
}

func (v NullableListSubscriptionsFromDestinationAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSubscriptionsFromDestinationAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
