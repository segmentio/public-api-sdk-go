/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.13.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ListDestinations200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListDestinations200Response{}

// ListDestinations200Response struct for ListDestinations200Response
type ListDestinations200Response struct {
	Data *ListDestinationsV1Output `json:"data,omitempty"`
}

// NewListDestinations200Response instantiates a new ListDestinations200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDestinations200Response() *ListDestinations200Response {
	this := ListDestinations200Response{}
	return &this
}

// NewListDestinations200ResponseWithDefaults instantiates a new ListDestinations200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDestinations200ResponseWithDefaults() *ListDestinations200Response {
	this := ListDestinations200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListDestinations200Response) GetData() ListDestinationsV1Output {
	if o == nil || IsNil(o.Data) {
		var ret ListDestinationsV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDestinations200Response) GetDataOk() (*ListDestinationsV1Output, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListDestinations200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ListDestinationsV1Output and assigns it to the Data field.
func (o *ListDestinations200Response) SetData(v ListDestinationsV1Output) {
	o.Data = &v
}

func (o ListDestinations200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListDestinations200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableListDestinations200Response struct {
	value *ListDestinations200Response
	isSet bool
}

func (v NullableListDestinations200Response) Get() *ListDestinations200Response {
	return v.value
}

func (v *NullableListDestinations200Response) Set(val *ListDestinations200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListDestinations200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListDestinations200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDestinations200Response(
	val *ListDestinations200Response,
) *NullableListDestinations200Response {
	return &NullableListDestinations200Response{value: val, isSet: true}
}

func (v NullableListDestinations200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDestinations200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
