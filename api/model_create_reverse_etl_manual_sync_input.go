/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 39.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the CreateReverseETLManualSyncInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateReverseETLManualSyncInput{}

// CreateReverseETLManualSyncInput Defines the parameters needed to trigger a manual sync for a RETL connection.
type CreateReverseETLManualSyncInput struct {
	// The id of the Source.
	SourceId string `json:"sourceId"`
	// The id of the Model.
	ModelId string `json:"modelId"`
	// The id of the Subscription.
	SubscriptionId string `json:"subscriptionId"`
}

// NewCreateReverseETLManualSyncInput instantiates a new CreateReverseETLManualSyncInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateReverseETLManualSyncInput(
	sourceId string,
	modelId string,
	subscriptionId string,
) *CreateReverseETLManualSyncInput {
	this := CreateReverseETLManualSyncInput{}
	this.SourceId = sourceId
	this.ModelId = modelId
	this.SubscriptionId = subscriptionId
	return &this
}

// NewCreateReverseETLManualSyncInputWithDefaults instantiates a new CreateReverseETLManualSyncInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateReverseETLManualSyncInputWithDefaults() *CreateReverseETLManualSyncInput {
	this := CreateReverseETLManualSyncInput{}
	return &this
}

// GetSourceId returns the SourceId field value
func (o *CreateReverseETLManualSyncInput) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *CreateReverseETLManualSyncInput) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *CreateReverseETLManualSyncInput) SetSourceId(v string) {
	o.SourceId = v
}

// GetModelId returns the ModelId field value
func (o *CreateReverseETLManualSyncInput) GetModelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value
// and a boolean to check if the value has been set.
func (o *CreateReverseETLManualSyncInput) GetModelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelId, true
}

// SetModelId sets field value
func (o *CreateReverseETLManualSyncInput) SetModelId(v string) {
	o.ModelId = v
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *CreateReverseETLManualSyncInput) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *CreateReverseETLManualSyncInput) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *CreateReverseETLManualSyncInput) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

func (o CreateReverseETLManualSyncInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateReverseETLManualSyncInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sourceId"] = o.SourceId
	toSerialize["modelId"] = o.ModelId
	toSerialize["subscriptionId"] = o.SubscriptionId
	return toSerialize, nil
}

type NullableCreateReverseETLManualSyncInput struct {
	value *CreateReverseETLManualSyncInput
	isSet bool
}

func (v NullableCreateReverseETLManualSyncInput) Get() *CreateReverseETLManualSyncInput {
	return v.value
}

func (v *NullableCreateReverseETLManualSyncInput) Set(val *CreateReverseETLManualSyncInput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateReverseETLManualSyncInput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateReverseETLManualSyncInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateReverseETLManualSyncInput(
	val *CreateReverseETLManualSyncInput,
) *NullableCreateReverseETLManualSyncInput {
	return &NullableCreateReverseETLManualSyncInput{value: val, isSet: true}
}

func (v NullableCreateReverseETLManualSyncInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateReverseETLManualSyncInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
