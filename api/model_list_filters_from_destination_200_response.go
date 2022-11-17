/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 32.0.6
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ListFiltersFromDestination200Response struct for ListFiltersFromDestination200Response
type ListFiltersFromDestination200Response struct {
	Data *ListFiltersFromDestinationV1Output `json:"data,omitempty"`
}

// NewListFiltersFromDestination200Response instantiates a new ListFiltersFromDestination200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListFiltersFromDestination200Response() *ListFiltersFromDestination200Response {
	this := ListFiltersFromDestination200Response{}
	return &this
}

// NewListFiltersFromDestination200ResponseWithDefaults instantiates a new ListFiltersFromDestination200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListFiltersFromDestination200ResponseWithDefaults() *ListFiltersFromDestination200Response {
	this := ListFiltersFromDestination200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListFiltersFromDestination200Response) GetData() ListFiltersFromDestinationV1Output {
	if o == nil || o.Data == nil {
		var ret ListFiltersFromDestinationV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFiltersFromDestination200Response) GetDataOk() (*ListFiltersFromDestinationV1Output, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListFiltersFromDestination200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ListFiltersFromDestinationV1Output and assigns it to the Data field.
func (o *ListFiltersFromDestination200Response) SetData(v ListFiltersFromDestinationV1Output) {
	o.Data = &v
}

func (o ListFiltersFromDestination200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListFiltersFromDestination200Response struct {
	value *ListFiltersFromDestination200Response
	isSet bool
}

func (v NullableListFiltersFromDestination200Response) Get() *ListFiltersFromDestination200Response {
	return v.value
}

func (v *NullableListFiltersFromDestination200Response) Set(val *ListFiltersFromDestination200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListFiltersFromDestination200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListFiltersFromDestination200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListFiltersFromDestination200Response(val *ListFiltersFromDestination200Response) *NullableListFiltersFromDestination200Response {
	return &NullableListFiltersFromDestination200Response{value: val, isSet: true}
}

func (v NullableListFiltersFromDestination200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListFiltersFromDestination200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


