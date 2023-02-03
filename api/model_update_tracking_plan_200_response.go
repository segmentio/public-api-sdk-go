/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.0.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UpdateTrackingPlan200Response struct for UpdateTrackingPlan200Response
type UpdateTrackingPlan200Response struct {
	Data *UpdateTrackingPlanV1Output `json:"data,omitempty"`
}

// NewUpdateTrackingPlan200Response instantiates a new UpdateTrackingPlan200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTrackingPlan200Response() *UpdateTrackingPlan200Response {
	this := UpdateTrackingPlan200Response{}
	return &this
}

// NewUpdateTrackingPlan200ResponseWithDefaults instantiates a new UpdateTrackingPlan200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTrackingPlan200ResponseWithDefaults() *UpdateTrackingPlan200Response {
	this := UpdateTrackingPlan200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UpdateTrackingPlan200Response) GetData() UpdateTrackingPlanV1Output {
	if o == nil || o.Data == nil {
		var ret UpdateTrackingPlanV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTrackingPlan200Response) GetDataOk() (*UpdateTrackingPlanV1Output, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UpdateTrackingPlan200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given UpdateTrackingPlanV1Output and assigns it to the Data field.
func (o *UpdateTrackingPlan200Response) SetData(v UpdateTrackingPlanV1Output) {
	o.Data = &v
}

func (o UpdateTrackingPlan200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateTrackingPlan200Response struct {
	value *UpdateTrackingPlan200Response
	isSet bool
}

func (v NullableUpdateTrackingPlan200Response) Get() *UpdateTrackingPlan200Response {
	return v.value
}

func (v *NullableUpdateTrackingPlan200Response) Set(val *UpdateTrackingPlan200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTrackingPlan200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTrackingPlan200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTrackingPlan200Response(
	val *UpdateTrackingPlan200Response,
) *NullableUpdateTrackingPlan200Response {
	return &NullableUpdateTrackingPlan200Response{value: val, isSet: true}
}

func (v NullableUpdateTrackingPlan200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTrackingPlan200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
