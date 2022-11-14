/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 32.0.5
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// RemoveFilterFromDestination200Response struct for RemoveFilterFromDestination200Response
type RemoveFilterFromDestination200Response struct {
	Data *RemoveFilterFromDestinationV1Output `json:"data,omitempty"`
}

// NewRemoveFilterFromDestination200Response instantiates a new RemoveFilterFromDestination200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveFilterFromDestination200Response() *RemoveFilterFromDestination200Response {
	this := RemoveFilterFromDestination200Response{}
	return &this
}

// NewRemoveFilterFromDestination200ResponseWithDefaults instantiates a new RemoveFilterFromDestination200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveFilterFromDestination200ResponseWithDefaults() *RemoveFilterFromDestination200Response {
	this := RemoveFilterFromDestination200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RemoveFilterFromDestination200Response) GetData() RemoveFilterFromDestinationV1Output {
	if o == nil || o.Data == nil {
		var ret RemoveFilterFromDestinationV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoveFilterFromDestination200Response) GetDataOk() (*RemoveFilterFromDestinationV1Output, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RemoveFilterFromDestination200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given RemoveFilterFromDestinationV1Output and assigns it to the Data field.
func (o *RemoveFilterFromDestination200Response) SetData(v RemoveFilterFromDestinationV1Output) {
	o.Data = &v
}

func (o RemoveFilterFromDestination200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableRemoveFilterFromDestination200Response struct {
	value *RemoveFilterFromDestination200Response
	isSet bool
}

func (v NullableRemoveFilterFromDestination200Response) Get() *RemoveFilterFromDestination200Response {
	return v.value
}

func (v *NullableRemoveFilterFromDestination200Response) Set(val *RemoveFilterFromDestination200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveFilterFromDestination200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveFilterFromDestination200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveFilterFromDestination200Response(val *RemoveFilterFromDestination200Response) *NullableRemoveFilterFromDestination200Response {
	return &NullableRemoveFilterFromDestination200Response{value: val, isSet: true}
}

func (v NullableRemoveFilterFromDestination200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveFilterFromDestination200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


