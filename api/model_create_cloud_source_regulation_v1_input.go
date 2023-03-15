/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.4.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CreateCloudSourceRegulationV1Input The input to create a Cloud Source-scoped regulation.
type CreateCloudSourceRegulationV1Input struct {
	// The regulation type to create.
	RegulationType string `json:"regulationType"`
	// The subject type. Must be `objectId` for Cloud Sources.
	SubjectType string `json:"subjectType"`
	// The user or object ids of the subjects to regulate.  Config API note: equal to `parent` but allows an array.
	SubjectIds []string `json:"subjectIds"`
	// The Cloud Source collection to regulate.
	Collection string `json:"collection"`
}

// NewCreateCloudSourceRegulationV1Input instantiates a new CreateCloudSourceRegulationV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCloudSourceRegulationV1Input(
	regulationType string,
	subjectType string,
	subjectIds []string,
	collection string,
) *CreateCloudSourceRegulationV1Input {
	this := CreateCloudSourceRegulationV1Input{}
	this.RegulationType = regulationType
	this.SubjectType = subjectType
	this.SubjectIds = subjectIds
	this.Collection = collection
	return &this
}

// NewCreateCloudSourceRegulationV1InputWithDefaults instantiates a new CreateCloudSourceRegulationV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCloudSourceRegulationV1InputWithDefaults() *CreateCloudSourceRegulationV1Input {
	this := CreateCloudSourceRegulationV1Input{}
	return &this
}

// GetRegulationType returns the RegulationType field value
func (o *CreateCloudSourceRegulationV1Input) GetRegulationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegulationType
}

// GetRegulationTypeOk returns a tuple with the RegulationType field value
// and a boolean to check if the value has been set.
func (o *CreateCloudSourceRegulationV1Input) GetRegulationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegulationType, true
}

// SetRegulationType sets field value
func (o *CreateCloudSourceRegulationV1Input) SetRegulationType(v string) {
	o.RegulationType = v
}

// GetSubjectType returns the SubjectType field value
func (o *CreateCloudSourceRegulationV1Input) GetSubjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubjectType
}

// GetSubjectTypeOk returns a tuple with the SubjectType field value
// and a boolean to check if the value has been set.
func (o *CreateCloudSourceRegulationV1Input) GetSubjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubjectType, true
}

// SetSubjectType sets field value
func (o *CreateCloudSourceRegulationV1Input) SetSubjectType(v string) {
	o.SubjectType = v
}

// GetSubjectIds returns the SubjectIds field value
func (o *CreateCloudSourceRegulationV1Input) GetSubjectIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SubjectIds
}

// GetSubjectIdsOk returns a tuple with the SubjectIds field value
// and a boolean to check if the value has been set.
func (o *CreateCloudSourceRegulationV1Input) GetSubjectIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubjectIds, true
}

// SetSubjectIds sets field value
func (o *CreateCloudSourceRegulationV1Input) SetSubjectIds(v []string) {
	o.SubjectIds = v
}

// GetCollection returns the Collection field value
func (o *CreateCloudSourceRegulationV1Input) GetCollection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value
// and a boolean to check if the value has been set.
func (o *CreateCloudSourceRegulationV1Input) GetCollectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Collection, true
}

// SetCollection sets field value
func (o *CreateCloudSourceRegulationV1Input) SetCollection(v string) {
	o.Collection = v
}

func (o CreateCloudSourceRegulationV1Input) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["regulationType"] = o.RegulationType
	}
	if true {
		toSerialize["subjectType"] = o.SubjectType
	}
	if true {
		toSerialize["subjectIds"] = o.SubjectIds
	}
	if true {
		toSerialize["collection"] = o.Collection
	}
	return json.Marshal(toSerialize)
}

type NullableCreateCloudSourceRegulationV1Input struct {
	value *CreateCloudSourceRegulationV1Input
	isSet bool
}

func (v NullableCreateCloudSourceRegulationV1Input) Get() *CreateCloudSourceRegulationV1Input {
	return v.value
}

func (v *NullableCreateCloudSourceRegulationV1Input) Set(val *CreateCloudSourceRegulationV1Input) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCloudSourceRegulationV1Input) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCloudSourceRegulationV1Input) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCloudSourceRegulationV1Input(
	val *CreateCloudSourceRegulationV1Input,
) *NullableCreateCloudSourceRegulationV1Input {
	return &NullableCreateCloudSourceRegulationV1Input{value: val, isSet: true}
}

func (v NullableCreateCloudSourceRegulationV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCloudSourceRegulationV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
