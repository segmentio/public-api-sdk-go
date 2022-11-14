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

// ListConnectedDestinationsFromSource200Response struct for ListConnectedDestinationsFromSource200Response
type ListConnectedDestinationsFromSource200Response struct {
	Data *ListConnectedDestinationsFromSourceV1Output `json:"data,omitempty"`
}

// NewListConnectedDestinationsFromSource200Response instantiates a new ListConnectedDestinationsFromSource200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConnectedDestinationsFromSource200Response() *ListConnectedDestinationsFromSource200Response {
	this := ListConnectedDestinationsFromSource200Response{}
	return &this
}

// NewListConnectedDestinationsFromSource200ResponseWithDefaults instantiates a new ListConnectedDestinationsFromSource200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConnectedDestinationsFromSource200ResponseWithDefaults() *ListConnectedDestinationsFromSource200Response {
	this := ListConnectedDestinationsFromSource200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListConnectedDestinationsFromSource200Response) GetData() ListConnectedDestinationsFromSourceV1Output {
	if o == nil || o.Data == nil {
		var ret ListConnectedDestinationsFromSourceV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectedDestinationsFromSource200Response) GetDataOk() (*ListConnectedDestinationsFromSourceV1Output, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListConnectedDestinationsFromSource200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ListConnectedDestinationsFromSourceV1Output and assigns it to the Data field.
func (o *ListConnectedDestinationsFromSource200Response) SetData(v ListConnectedDestinationsFromSourceV1Output) {
	o.Data = &v
}

func (o ListConnectedDestinationsFromSource200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListConnectedDestinationsFromSource200Response struct {
	value *ListConnectedDestinationsFromSource200Response
	isSet bool
}

func (v NullableListConnectedDestinationsFromSource200Response) Get() *ListConnectedDestinationsFromSource200Response {
	return v.value
}

func (v *NullableListConnectedDestinationsFromSource200Response) Set(val *ListConnectedDestinationsFromSource200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListConnectedDestinationsFromSource200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListConnectedDestinationsFromSource200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConnectedDestinationsFromSource200Response(val *ListConnectedDestinationsFromSource200Response) *NullableListConnectedDestinationsFromSource200Response {
	return &NullableListConnectedDestinationsFromSource200Response{value: val, isSet: true}
}

func (v NullableListConnectedDestinationsFromSource200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConnectedDestinationsFromSource200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


