/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 38.4.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ListSubscriptionsFromDestination200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSubscriptionsFromDestination200Response{}

// ListSubscriptionsFromDestination200Response struct for ListSubscriptionsFromDestination200Response
type ListSubscriptionsFromDestination200Response struct {
	Data *ListSubscriptionsFromDestinationAlphaOutput `json:"data,omitempty"`
}

// NewListSubscriptionsFromDestination200Response instantiates a new ListSubscriptionsFromDestination200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSubscriptionsFromDestination200Response() *ListSubscriptionsFromDestination200Response {
	this := ListSubscriptionsFromDestination200Response{}
	return &this
}

// NewListSubscriptionsFromDestination200ResponseWithDefaults instantiates a new ListSubscriptionsFromDestination200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSubscriptionsFromDestination200ResponseWithDefaults() *ListSubscriptionsFromDestination200Response {
	this := ListSubscriptionsFromDestination200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListSubscriptionsFromDestination200Response) GetData() ListSubscriptionsFromDestinationAlphaOutput {
	if o == nil || IsNil(o.Data) {
		var ret ListSubscriptionsFromDestinationAlphaOutput
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSubscriptionsFromDestination200Response) GetDataOk() (*ListSubscriptionsFromDestinationAlphaOutput, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListSubscriptionsFromDestination200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ListSubscriptionsFromDestinationAlphaOutput and assigns it to the Data field.
func (o *ListSubscriptionsFromDestination200Response) SetData(
	v ListSubscriptionsFromDestinationAlphaOutput,
) {
	o.Data = &v
}

func (o ListSubscriptionsFromDestination200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListSubscriptionsFromDestination200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableListSubscriptionsFromDestination200Response struct {
	value *ListSubscriptionsFromDestination200Response
	isSet bool
}

func (v NullableListSubscriptionsFromDestination200Response) Get() *ListSubscriptionsFromDestination200Response {
	return v.value
}

func (v *NullableListSubscriptionsFromDestination200Response) Set(
	val *ListSubscriptionsFromDestination200Response,
) {
	v.value = val
	v.isSet = true
}

func (v NullableListSubscriptionsFromDestination200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListSubscriptionsFromDestination200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSubscriptionsFromDestination200Response(
	val *ListSubscriptionsFromDestination200Response,
) *NullableListSubscriptionsFromDestination200Response {
	return &NullableListSubscriptionsFromDestination200Response{value: val, isSet: true}
}

func (v NullableListSubscriptionsFromDestination200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSubscriptionsFromDestination200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
