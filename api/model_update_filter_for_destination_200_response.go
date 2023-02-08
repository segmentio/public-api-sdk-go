/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UpdateFilterForDestination200Response struct for UpdateFilterForDestination200Response
type UpdateFilterForDestination200Response struct {
	Data *UpdateFilterForDestinationV1Output `json:"data,omitempty"`
}

// NewUpdateFilterForDestination200Response instantiates a new UpdateFilterForDestination200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFilterForDestination200Response() *UpdateFilterForDestination200Response {
	this := UpdateFilterForDestination200Response{}
	return &this
}

// NewUpdateFilterForDestination200ResponseWithDefaults instantiates a new UpdateFilterForDestination200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFilterForDestination200ResponseWithDefaults() *UpdateFilterForDestination200Response {
	this := UpdateFilterForDestination200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UpdateFilterForDestination200Response) GetData() UpdateFilterForDestinationV1Output {
	if o == nil || o.Data == nil {
		var ret UpdateFilterForDestinationV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFilterForDestination200Response) GetDataOk() (*UpdateFilterForDestinationV1Output, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UpdateFilterForDestination200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given UpdateFilterForDestinationV1Output and assigns it to the Data field.
func (o *UpdateFilterForDestination200Response) SetData(v UpdateFilterForDestinationV1Output) {
	o.Data = &v
}

func (o UpdateFilterForDestination200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateFilterForDestination200Response struct {
	value *UpdateFilterForDestination200Response
	isSet bool
}

func (v NullableUpdateFilterForDestination200Response) Get() *UpdateFilterForDestination200Response {
	return v.value
}

func (v *NullableUpdateFilterForDestination200Response) Set(
	val *UpdateFilterForDestination200Response,
) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFilterForDestination200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFilterForDestination200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFilterForDestination200Response(
	val *UpdateFilterForDestination200Response,
) *NullableUpdateFilterForDestination200Response {
	return &NullableUpdateFilterForDestination200Response{value: val, isSet: true}
}

func (v NullableUpdateFilterForDestination200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFilterForDestination200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
