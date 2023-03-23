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

// CreateDestinationSubscription200Response struct for CreateDestinationSubscription200Response
type CreateDestinationSubscription200Response struct {
	Data *CreateDestinationSubscriptionAlphaOutput `json:"data,omitempty"`
}

// NewCreateDestinationSubscription200Response instantiates a new CreateDestinationSubscription200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDestinationSubscription200Response() *CreateDestinationSubscription200Response {
	this := CreateDestinationSubscription200Response{}
	return &this
}

// NewCreateDestinationSubscription200ResponseWithDefaults instantiates a new CreateDestinationSubscription200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDestinationSubscription200ResponseWithDefaults() *CreateDestinationSubscription200Response {
	this := CreateDestinationSubscription200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateDestinationSubscription200Response) GetData() CreateDestinationSubscriptionAlphaOutput {
	if o == nil || o.Data == nil {
		var ret CreateDestinationSubscriptionAlphaOutput
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDestinationSubscription200Response) GetDataOk() (*CreateDestinationSubscriptionAlphaOutput, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateDestinationSubscription200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given CreateDestinationSubscriptionAlphaOutput and assigns it to the Data field.
func (o *CreateDestinationSubscription200Response) SetData(
	v CreateDestinationSubscriptionAlphaOutput,
) {
	o.Data = &v
}

func (o CreateDestinationSubscription200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCreateDestinationSubscription200Response struct {
	value *CreateDestinationSubscription200Response
	isSet bool
}

func (v NullableCreateDestinationSubscription200Response) Get() *CreateDestinationSubscription200Response {
	return v.value
}

func (v *NullableCreateDestinationSubscription200Response) Set(
	val *CreateDestinationSubscription200Response,
) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDestinationSubscription200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDestinationSubscription200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDestinationSubscription200Response(
	val *CreateDestinationSubscription200Response,
) *NullableCreateDestinationSubscription200Response {
	return &NullableCreateDestinationSubscription200Response{value: val, isSet: true}
}

func (v NullableCreateDestinationSubscription200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDestinationSubscription200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
