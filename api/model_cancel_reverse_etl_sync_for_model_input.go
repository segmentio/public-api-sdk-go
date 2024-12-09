/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 57.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the CancelReverseETLSyncForModelInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CancelReverseETLSyncForModelInput{}

// CancelReverseETLSyncForModelInput Defines the parameters needed to cancel a sync for a RETL connection.
type CancelReverseETLSyncForModelInput struct {
	// The reason for canceling the sync. - IncorrectModel = 0 - IncorrectDest = 1 - IncorrectKeys = 2 - IncorrectMapping = 3 - Other = 4
	ReasonForCanceling *float32 `json:"reasonForCanceling,omitempty"`
}

// NewCancelReverseETLSyncForModelInput instantiates a new CancelReverseETLSyncForModelInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelReverseETLSyncForModelInput() *CancelReverseETLSyncForModelInput {
	this := CancelReverseETLSyncForModelInput{}
	return &this
}

// NewCancelReverseETLSyncForModelInputWithDefaults instantiates a new CancelReverseETLSyncForModelInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelReverseETLSyncForModelInputWithDefaults() *CancelReverseETLSyncForModelInput {
	this := CancelReverseETLSyncForModelInput{}
	return &this
}

// GetReasonForCanceling returns the ReasonForCanceling field value if set, zero value otherwise.
func (o *CancelReverseETLSyncForModelInput) GetReasonForCanceling() float32 {
	if o == nil || IsNil(o.ReasonForCanceling) {
		var ret float32
		return ret
	}
	return *o.ReasonForCanceling
}

// GetReasonForCancelingOk returns a tuple with the ReasonForCanceling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelReverseETLSyncForModelInput) GetReasonForCancelingOk() (*float32, bool) {
	if o == nil || IsNil(o.ReasonForCanceling) {
		return nil, false
	}
	return o.ReasonForCanceling, true
}

// HasReasonForCanceling returns a boolean if a field has been set.
func (o *CancelReverseETLSyncForModelInput) HasReasonForCanceling() bool {
	if o != nil && !IsNil(o.ReasonForCanceling) {
		return true
	}

	return false
}

// SetReasonForCanceling gets a reference to the given float32 and assigns it to the ReasonForCanceling field.
func (o *CancelReverseETLSyncForModelInput) SetReasonForCanceling(v float32) {
	o.ReasonForCanceling = &v
}

func (o CancelReverseETLSyncForModelInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelReverseETLSyncForModelInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReasonForCanceling) {
		toSerialize["reasonForCanceling"] = o.ReasonForCanceling
	}
	return toSerialize, nil
}

type NullableCancelReverseETLSyncForModelInput struct {
	value *CancelReverseETLSyncForModelInput
	isSet bool
}

func (v NullableCancelReverseETLSyncForModelInput) Get() *CancelReverseETLSyncForModelInput {
	return v.value
}

func (v *NullableCancelReverseETLSyncForModelInput) Set(val *CancelReverseETLSyncForModelInput) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelReverseETLSyncForModelInput) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelReverseETLSyncForModelInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelReverseETLSyncForModelInput(
	val *CancelReverseETLSyncForModelInput,
) *NullableCancelReverseETLSyncForModelInput {
	return &NullableCancelReverseETLSyncForModelInput{value: val, isSet: true}
}

func (v NullableCancelReverseETLSyncForModelInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelReverseETLSyncForModelInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
