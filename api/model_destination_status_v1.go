/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 54.3.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the DestinationStatusV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DestinationStatusV1{}

// DestinationStatusV1 DestinationStatus represents status of each Destination in a stream.
type DestinationStatusV1 struct {
	Name       string `json:"name"`
	Id         string `json:"id"`
	Status     string `json:"status"`
	ErrString  string `json:"errString"`
	FinishedAt string `json:"finishedAt"`
}

// NewDestinationStatusV1 instantiates a new DestinationStatusV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestinationStatusV1(
	name string,
	id string,
	status string,
	errString string,
	finishedAt string,
) *DestinationStatusV1 {
	this := DestinationStatusV1{}
	this.Name = name
	this.Id = id
	this.Status = status
	this.ErrString = errString
	this.FinishedAt = finishedAt
	return &this
}

// NewDestinationStatusV1WithDefaults instantiates a new DestinationStatusV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestinationStatusV1WithDefaults() *DestinationStatusV1 {
	this := DestinationStatusV1{}
	return &this
}

// GetName returns the Name field value
func (o *DestinationStatusV1) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DestinationStatusV1) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DestinationStatusV1) SetName(v string) {
	o.Name = v
}

// GetId returns the Id field value
func (o *DestinationStatusV1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DestinationStatusV1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DestinationStatusV1) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *DestinationStatusV1) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DestinationStatusV1) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DestinationStatusV1) SetStatus(v string) {
	o.Status = v
}

// GetErrString returns the ErrString field value
func (o *DestinationStatusV1) GetErrString() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrString
}

// GetErrStringOk returns a tuple with the ErrString field value
// and a boolean to check if the value has been set.
func (o *DestinationStatusV1) GetErrStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrString, true
}

// SetErrString sets field value
func (o *DestinationStatusV1) SetErrString(v string) {
	o.ErrString = v
}

// GetFinishedAt returns the FinishedAt field value
func (o *DestinationStatusV1) GetFinishedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value
// and a boolean to check if the value has been set.
func (o *DestinationStatusV1) GetFinishedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinishedAt, true
}

// SetFinishedAt sets field value
func (o *DestinationStatusV1) SetFinishedAt(v string) {
	o.FinishedAt = v
}

func (o DestinationStatusV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DestinationStatusV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["id"] = o.Id
	toSerialize["status"] = o.Status
	toSerialize["errString"] = o.ErrString
	toSerialize["finishedAt"] = o.FinishedAt
	return toSerialize, nil
}

type NullableDestinationStatusV1 struct {
	value *DestinationStatusV1
	isSet bool
}

func (v NullableDestinationStatusV1) Get() *DestinationStatusV1 {
	return v.value
}

func (v *NullableDestinationStatusV1) Set(val *DestinationStatusV1) {
	v.value = val
	v.isSet = true
}

func (v NullableDestinationStatusV1) IsSet() bool {
	return v.isSet
}

func (v *NullableDestinationStatusV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestinationStatusV1(val *DestinationStatusV1) *NullableDestinationStatusV1 {
	return &NullableDestinationStatusV1{value: val, isSet: true}
}

func (v NullableDestinationStatusV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDestinationStatusV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
