/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 47.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the RemoveRulesFromTrackingPlan200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoveRulesFromTrackingPlan200Response{}

// RemoveRulesFromTrackingPlan200Response struct for RemoveRulesFromTrackingPlan200Response
type RemoveRulesFromTrackingPlan200Response struct {
	Data *RemoveRulesFromTrackingPlanV1Output `json:"data,omitempty"`
}

// NewRemoveRulesFromTrackingPlan200Response instantiates a new RemoveRulesFromTrackingPlan200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveRulesFromTrackingPlan200Response() *RemoveRulesFromTrackingPlan200Response {
	this := RemoveRulesFromTrackingPlan200Response{}
	return &this
}

// NewRemoveRulesFromTrackingPlan200ResponseWithDefaults instantiates a new RemoveRulesFromTrackingPlan200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveRulesFromTrackingPlan200ResponseWithDefaults() *RemoveRulesFromTrackingPlan200Response {
	this := RemoveRulesFromTrackingPlan200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RemoveRulesFromTrackingPlan200Response) GetData() RemoveRulesFromTrackingPlanV1Output {
	if o == nil || IsNil(o.Data) {
		var ret RemoveRulesFromTrackingPlanV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoveRulesFromTrackingPlan200Response) GetDataOk() (*RemoveRulesFromTrackingPlanV1Output, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RemoveRulesFromTrackingPlan200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given RemoveRulesFromTrackingPlanV1Output and assigns it to the Data field.
func (o *RemoveRulesFromTrackingPlan200Response) SetData(v RemoveRulesFromTrackingPlanV1Output) {
	o.Data = &v
}

func (o RemoveRulesFromTrackingPlan200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoveRulesFromTrackingPlan200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableRemoveRulesFromTrackingPlan200Response struct {
	value *RemoveRulesFromTrackingPlan200Response
	isSet bool
}

func (v NullableRemoveRulesFromTrackingPlan200Response) Get() *RemoveRulesFromTrackingPlan200Response {
	return v.value
}

func (v *NullableRemoveRulesFromTrackingPlan200Response) Set(
	val *RemoveRulesFromTrackingPlan200Response,
) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveRulesFromTrackingPlan200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveRulesFromTrackingPlan200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveRulesFromTrackingPlan200Response(
	val *RemoveRulesFromTrackingPlan200Response,
) *NullableRemoveRulesFromTrackingPlan200Response {
	return &NullableRemoveRulesFromTrackingPlan200Response{value: val, isSet: true}
}

func (v NullableRemoveRulesFromTrackingPlan200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveRulesFromTrackingPlan200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
