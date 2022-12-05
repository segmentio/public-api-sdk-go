/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 33.0.3
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ListSelectiveSyncsFromWarehouseAndSource200Response struct for ListSelectiveSyncsFromWarehouseAndSource200Response
type ListSelectiveSyncsFromWarehouseAndSource200Response struct {
	Data *ListSelectiveSyncsFromWarehouseAndSourceV1Output `json:"data,omitempty"`
}

// NewListSelectiveSyncsFromWarehouseAndSource200Response instantiates a new ListSelectiveSyncsFromWarehouseAndSource200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSelectiveSyncsFromWarehouseAndSource200Response() *ListSelectiveSyncsFromWarehouseAndSource200Response {
	this := ListSelectiveSyncsFromWarehouseAndSource200Response{}
	return &this
}

// NewListSelectiveSyncsFromWarehouseAndSource200ResponseWithDefaults instantiates a new ListSelectiveSyncsFromWarehouseAndSource200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSelectiveSyncsFromWarehouseAndSource200ResponseWithDefaults() *ListSelectiveSyncsFromWarehouseAndSource200Response {
	this := ListSelectiveSyncsFromWarehouseAndSource200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListSelectiveSyncsFromWarehouseAndSource200Response) GetData() ListSelectiveSyncsFromWarehouseAndSourceV1Output {
	if o == nil || o.Data == nil {
		var ret ListSelectiveSyncsFromWarehouseAndSourceV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSelectiveSyncsFromWarehouseAndSource200Response) GetDataOk() (*ListSelectiveSyncsFromWarehouseAndSourceV1Output, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListSelectiveSyncsFromWarehouseAndSource200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ListSelectiveSyncsFromWarehouseAndSourceV1Output and assigns it to the Data field.
func (o *ListSelectiveSyncsFromWarehouseAndSource200Response) SetData(
	v ListSelectiveSyncsFromWarehouseAndSourceV1Output,
) {
	o.Data = &v
}

func (o ListSelectiveSyncsFromWarehouseAndSource200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListSelectiveSyncsFromWarehouseAndSource200Response struct {
	value *ListSelectiveSyncsFromWarehouseAndSource200Response
	isSet bool
}

func (v NullableListSelectiveSyncsFromWarehouseAndSource200Response) Get() *ListSelectiveSyncsFromWarehouseAndSource200Response {
	return v.value
}

func (v *NullableListSelectiveSyncsFromWarehouseAndSource200Response) Set(
	val *ListSelectiveSyncsFromWarehouseAndSource200Response,
) {
	v.value = val
	v.isSet = true
}

func (v NullableListSelectiveSyncsFromWarehouseAndSource200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListSelectiveSyncsFromWarehouseAndSource200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSelectiveSyncsFromWarehouseAndSource200Response(
	val *ListSelectiveSyncsFromWarehouseAndSource200Response,
) *NullableListSelectiveSyncsFromWarehouseAndSource200Response {
	return &NullableListSelectiveSyncsFromWarehouseAndSource200Response{value: val, isSet: true}
}

func (v NullableListSelectiveSyncsFromWarehouseAndSource200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSelectiveSyncsFromWarehouseAndSource200Response) UnmarshalJSON(
	src []byte,
) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
