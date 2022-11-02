/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 32.0.2
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// RemoveSubscriptionFromDestination200Response struct for RemoveSubscriptionFromDestination200Response
type RemoveSubscriptionFromDestination200Response struct {
	Data *RemoveSubscriptionFromDestinationAlphaOutput `json:"data,omitempty"`
}

// NewRemoveSubscriptionFromDestination200Response instantiates a new RemoveSubscriptionFromDestination200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveSubscriptionFromDestination200Response() *RemoveSubscriptionFromDestination200Response {
	this := RemoveSubscriptionFromDestination200Response{}
	return &this
}

// NewRemoveSubscriptionFromDestination200ResponseWithDefaults instantiates a new RemoveSubscriptionFromDestination200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveSubscriptionFromDestination200ResponseWithDefaults() *RemoveSubscriptionFromDestination200Response {
	this := RemoveSubscriptionFromDestination200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RemoveSubscriptionFromDestination200Response) GetData() RemoveSubscriptionFromDestinationAlphaOutput {
	if o == nil || o.Data == nil {
		var ret RemoveSubscriptionFromDestinationAlphaOutput
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoveSubscriptionFromDestination200Response) GetDataOk() (*RemoveSubscriptionFromDestinationAlphaOutput, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RemoveSubscriptionFromDestination200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given RemoveSubscriptionFromDestinationAlphaOutput and assigns it to the Data field.
func (o *RemoveSubscriptionFromDestination200Response) SetData(v RemoveSubscriptionFromDestinationAlphaOutput) {
	o.Data = &v
}

func (o RemoveSubscriptionFromDestination200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableRemoveSubscriptionFromDestination200Response struct {
	value *RemoveSubscriptionFromDestination200Response
	isSet bool
}

func (v NullableRemoveSubscriptionFromDestination200Response) Get() *RemoveSubscriptionFromDestination200Response {
	return v.value
}

func (v *NullableRemoveSubscriptionFromDestination200Response) Set(val *RemoveSubscriptionFromDestination200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveSubscriptionFromDestination200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveSubscriptionFromDestination200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveSubscriptionFromDestination200Response(val *RemoveSubscriptionFromDestination200Response) *NullableRemoveSubscriptionFromDestination200Response {
	return &NullableRemoveSubscriptionFromDestination200Response{value: val, isSet: true}
}

func (v NullableRemoveSubscriptionFromDestination200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveSubscriptionFromDestination200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


