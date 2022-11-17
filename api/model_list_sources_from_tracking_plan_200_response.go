/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 32.0.8
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ListSourcesFromTrackingPlan200Response struct for ListSourcesFromTrackingPlan200Response
type ListSourcesFromTrackingPlan200Response struct {
	Data *ListSourcesFromTrackingPlanV1Output `json:"data,omitempty"`
}

// NewListSourcesFromTrackingPlan200Response instantiates a new ListSourcesFromTrackingPlan200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSourcesFromTrackingPlan200Response() *ListSourcesFromTrackingPlan200Response {
	this := ListSourcesFromTrackingPlan200Response{}
	return &this
}

// NewListSourcesFromTrackingPlan200ResponseWithDefaults instantiates a new ListSourcesFromTrackingPlan200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSourcesFromTrackingPlan200ResponseWithDefaults() *ListSourcesFromTrackingPlan200Response {
	this := ListSourcesFromTrackingPlan200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListSourcesFromTrackingPlan200Response) GetData() ListSourcesFromTrackingPlanV1Output {
	if o == nil || o.Data == nil {
		var ret ListSourcesFromTrackingPlanV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSourcesFromTrackingPlan200Response) GetDataOk() (*ListSourcesFromTrackingPlanV1Output, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListSourcesFromTrackingPlan200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ListSourcesFromTrackingPlanV1Output and assigns it to the Data field.
func (o *ListSourcesFromTrackingPlan200Response) SetData(v ListSourcesFromTrackingPlanV1Output) {
	o.Data = &v
}

func (o ListSourcesFromTrackingPlan200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListSourcesFromTrackingPlan200Response struct {
	value *ListSourcesFromTrackingPlan200Response
	isSet bool
}

func (v NullableListSourcesFromTrackingPlan200Response) Get() *ListSourcesFromTrackingPlan200Response {
	return v.value
}

func (v *NullableListSourcesFromTrackingPlan200Response) Set(val *ListSourcesFromTrackingPlan200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListSourcesFromTrackingPlan200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListSourcesFromTrackingPlan200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSourcesFromTrackingPlan200Response(val *ListSourcesFromTrackingPlan200Response) *NullableListSourcesFromTrackingPlan200Response {
	return &NullableListSourcesFromTrackingPlan200Response{value: val, isSet: true}
}

func (v NullableListSourcesFromTrackingPlan200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSourcesFromTrackingPlan200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


