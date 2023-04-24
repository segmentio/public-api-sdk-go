/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.3.5
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// RegulationListEntryV1 struct for RegulationListEntryV1
type RegulationListEntryV1 struct {
	Id          string   `json:"id"`
	SubjectType string   `json:"subjectType"`
	Subjects    []string `json:"subjects"`
	Status      string   `json:"status"`
	CreatedAt   string   `json:"createdAt"`
	FinishedAt  *string  `json:"finishedAt,omitempty"`
}

// NewRegulationListEntryV1 instantiates a new RegulationListEntryV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegulationListEntryV1(
	id string,
	subjectType string,
	subjects []string,
	status string,
	createdAt string,
) *RegulationListEntryV1 {
	this := RegulationListEntryV1{}
	this.Id = id
	this.SubjectType = subjectType
	this.Subjects = subjects
	this.Status = status
	this.CreatedAt = createdAt
	return &this
}

// NewRegulationListEntryV1WithDefaults instantiates a new RegulationListEntryV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegulationListEntryV1WithDefaults() *RegulationListEntryV1 {
	this := RegulationListEntryV1{}
	return &this
}

// GetId returns the Id field value
func (o *RegulationListEntryV1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RegulationListEntryV1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RegulationListEntryV1) SetId(v string) {
	o.Id = v
}

// GetSubjectType returns the SubjectType field value
func (o *RegulationListEntryV1) GetSubjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubjectType
}

// GetSubjectTypeOk returns a tuple with the SubjectType field value
// and a boolean to check if the value has been set.
func (o *RegulationListEntryV1) GetSubjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubjectType, true
}

// SetSubjectType sets field value
func (o *RegulationListEntryV1) SetSubjectType(v string) {
	o.SubjectType = v
}

// GetSubjects returns the Subjects field value
func (o *RegulationListEntryV1) GetSubjects() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Subjects
}

// GetSubjectsOk returns a tuple with the Subjects field value
// and a boolean to check if the value has been set.
func (o *RegulationListEntryV1) GetSubjectsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subjects, true
}

// SetSubjects sets field value
func (o *RegulationListEntryV1) SetSubjects(v []string) {
	o.Subjects = v
}

// GetStatus returns the Status field value
func (o *RegulationListEntryV1) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RegulationListEntryV1) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RegulationListEntryV1) SetStatus(v string) {
	o.Status = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *RegulationListEntryV1) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *RegulationListEntryV1) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *RegulationListEntryV1) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise.
func (o *RegulationListEntryV1) GetFinishedAt() string {
	if o == nil || o.FinishedAt == nil {
		var ret string
		return ret
	}
	return *o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegulationListEntryV1) GetFinishedAtOk() (*string, bool) {
	if o == nil || o.FinishedAt == nil {
		return nil, false
	}
	return o.FinishedAt, true
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *RegulationListEntryV1) HasFinishedAt() bool {
	if o != nil && o.FinishedAt != nil {
		return true
	}

	return false
}

// SetFinishedAt gets a reference to the given string and assigns it to the FinishedAt field.
func (o *RegulationListEntryV1) SetFinishedAt(v string) {
	o.FinishedAt = &v
}

func (o RegulationListEntryV1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["subjectType"] = o.SubjectType
	}
	if true {
		toSerialize["subjects"] = o.Subjects
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.FinishedAt != nil {
		toSerialize["finishedAt"] = o.FinishedAt
	}
	return json.Marshal(toSerialize)
}

type NullableRegulationListEntryV1 struct {
	value *RegulationListEntryV1
	isSet bool
}

func (v NullableRegulationListEntryV1) Get() *RegulationListEntryV1 {
	return v.value
}

func (v *NullableRegulationListEntryV1) Set(val *RegulationListEntryV1) {
	v.value = val
	v.isSet = true
}

func (v NullableRegulationListEntryV1) IsSet() bool {
	return v.isSet
}

func (v *NullableRegulationListEntryV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegulationListEntryV1(val *RegulationListEntryV1) *NullableRegulationListEntryV1 {
	return &NullableRegulationListEntryV1{value: val, isSet: true}
}

func (v NullableRegulationListEntryV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegulationListEntryV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
