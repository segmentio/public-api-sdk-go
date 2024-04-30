/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 50.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the UpdateSubscriptionForDestination200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSubscriptionForDestination200Response{}

// UpdateSubscriptionForDestination200Response struct for UpdateSubscriptionForDestination200Response
type UpdateSubscriptionForDestination200Response struct {
	Data *UpdateSubscriptionForDestinationAlphaOutput `json:"data,omitempty"`
}

// NewUpdateSubscriptionForDestination200Response instantiates a new UpdateSubscriptionForDestination200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSubscriptionForDestination200Response() *UpdateSubscriptionForDestination200Response {
	this := UpdateSubscriptionForDestination200Response{}
	return &this
}

// NewUpdateSubscriptionForDestination200ResponseWithDefaults instantiates a new UpdateSubscriptionForDestination200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSubscriptionForDestination200ResponseWithDefaults() *UpdateSubscriptionForDestination200Response {
	this := UpdateSubscriptionForDestination200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UpdateSubscriptionForDestination200Response) GetData() UpdateSubscriptionForDestinationAlphaOutput {
	if o == nil || IsNil(o.Data) {
		var ret UpdateSubscriptionForDestinationAlphaOutput
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSubscriptionForDestination200Response) GetDataOk() (*UpdateSubscriptionForDestinationAlphaOutput, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UpdateSubscriptionForDestination200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given UpdateSubscriptionForDestinationAlphaOutput and assigns it to the Data field.
func (o *UpdateSubscriptionForDestination200Response) SetData(
	v UpdateSubscriptionForDestinationAlphaOutput,
) {
	o.Data = &v
}

func (o UpdateSubscriptionForDestination200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSubscriptionForDestination200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableUpdateSubscriptionForDestination200Response struct {
	value *UpdateSubscriptionForDestination200Response
	isSet bool
}

func (v NullableUpdateSubscriptionForDestination200Response) Get() *UpdateSubscriptionForDestination200Response {
	return v.value
}

func (v *NullableUpdateSubscriptionForDestination200Response) Set(
	val *UpdateSubscriptionForDestination200Response,
) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSubscriptionForDestination200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSubscriptionForDestination200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSubscriptionForDestination200Response(
	val *UpdateSubscriptionForDestination200Response,
) *NullableUpdateSubscriptionForDestination200Response {
	return &NullableUpdateSubscriptionForDestination200Response{value: val, isSet: true}
}

func (v NullableUpdateSubscriptionForDestination200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSubscriptionForDestination200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
