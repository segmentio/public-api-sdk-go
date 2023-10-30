/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 37.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UpdateRulesInTrackingPlan200Response struct for UpdateRulesInTrackingPlan200Response
type UpdateRulesInTrackingPlan200Response struct {
	Data *UpdateRulesInTrackingPlanV1Output `json:"data,omitempty"`
}

// NewUpdateRulesInTrackingPlan200Response instantiates a new UpdateRulesInTrackingPlan200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRulesInTrackingPlan200Response() *UpdateRulesInTrackingPlan200Response {
	this := UpdateRulesInTrackingPlan200Response{}
	return &this
}

// NewUpdateRulesInTrackingPlan200ResponseWithDefaults instantiates a new UpdateRulesInTrackingPlan200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRulesInTrackingPlan200ResponseWithDefaults() *UpdateRulesInTrackingPlan200Response {
	this := UpdateRulesInTrackingPlan200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UpdateRulesInTrackingPlan200Response) GetData() UpdateRulesInTrackingPlanV1Output {
	if o == nil || o.Data == nil {
		var ret UpdateRulesInTrackingPlanV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRulesInTrackingPlan200Response) GetDataOk() (*UpdateRulesInTrackingPlanV1Output, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UpdateRulesInTrackingPlan200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given UpdateRulesInTrackingPlanV1Output and assigns it to the Data field.
func (o *UpdateRulesInTrackingPlan200Response) SetData(v UpdateRulesInTrackingPlanV1Output) {
	o.Data = &v
}

func (o UpdateRulesInTrackingPlan200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateRulesInTrackingPlan200Response struct {
	value *UpdateRulesInTrackingPlan200Response
	isSet bool
}

func (v NullableUpdateRulesInTrackingPlan200Response) Get() *UpdateRulesInTrackingPlan200Response {
	return v.value
}

func (v *NullableUpdateRulesInTrackingPlan200Response) Set(
	val *UpdateRulesInTrackingPlan200Response,
) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRulesInTrackingPlan200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRulesInTrackingPlan200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRulesInTrackingPlan200Response(
	val *UpdateRulesInTrackingPlan200Response,
) *NullableUpdateRulesInTrackingPlan200Response {
	return &NullableUpdateRulesInTrackingPlan200Response{value: val, isSet: true}
}

func (v NullableUpdateRulesInTrackingPlan200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRulesInTrackingPlan200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
