/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 53.2.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the UsersPerSourceSnapshotV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsersPerSourceSnapshotV1{}

// UsersPerSourceSnapshotV1 A snapshot of MTU metrics for a given Source within the given usage period.
type UsersPerSourceSnapshotV1 struct {
	// The Source id.
	SourceId string `json:"sourceId"`
	// The start of the usage period, in unix time (seconds since epoch).
	PeriodStart float32 `json:"periodStart"`
	// The end of the usage period, in unix time (seconds since epoch).
	PeriodEnd float32 `json:"periodEnd"`
	// A bigint of the number of anonymous users in this snapshot.
	Anonymous string `json:"anonymous"`
	// A bigint of the number of anonymous identified users in this snapshot.
	AnonymousIdentified string `json:"anonymousIdentified"`
	// A bigint of the number of identified users in this snapshot.
	Identified string `json:"identified"`
	// A bigint of the number of never identified users in this snapshot.
	NeverIdentified string `json:"neverIdentified"`
	// The timestamp of this snapshot within the billing cycle, in the ISO-8601 format.
	Timestamp string `json:"timestamp"`
}

// NewUsersPerSourceSnapshotV1 instantiates a new UsersPerSourceSnapshotV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersPerSourceSnapshotV1(
	sourceId string,
	periodStart float32,
	periodEnd float32,
	anonymous string,
	anonymousIdentified string,
	identified string,
	neverIdentified string,
	timestamp string,
) *UsersPerSourceSnapshotV1 {
	this := UsersPerSourceSnapshotV1{}
	this.SourceId = sourceId
	this.PeriodStart = periodStart
	this.PeriodEnd = periodEnd
	this.Anonymous = anonymous
	this.AnonymousIdentified = anonymousIdentified
	this.Identified = identified
	this.NeverIdentified = neverIdentified
	this.Timestamp = timestamp
	return &this
}

// NewUsersPerSourceSnapshotV1WithDefaults instantiates a new UsersPerSourceSnapshotV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersPerSourceSnapshotV1WithDefaults() *UsersPerSourceSnapshotV1 {
	this := UsersPerSourceSnapshotV1{}
	return &this
}

// GetSourceId returns the SourceId field value
func (o *UsersPerSourceSnapshotV1) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *UsersPerSourceSnapshotV1) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *UsersPerSourceSnapshotV1) SetSourceId(v string) {
	o.SourceId = v
}

// GetPeriodStart returns the PeriodStart field value
func (o *UsersPerSourceSnapshotV1) GetPeriodStart() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PeriodStart
}

// GetPeriodStartOk returns a tuple with the PeriodStart field value
// and a boolean to check if the value has been set.
func (o *UsersPerSourceSnapshotV1) GetPeriodStartOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PeriodStart, true
}

// SetPeriodStart sets field value
func (o *UsersPerSourceSnapshotV1) SetPeriodStart(v float32) {
	o.PeriodStart = v
}

// GetPeriodEnd returns the PeriodEnd field value
func (o *UsersPerSourceSnapshotV1) GetPeriodEnd() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PeriodEnd
}

// GetPeriodEndOk returns a tuple with the PeriodEnd field value
// and a boolean to check if the value has been set.
func (o *UsersPerSourceSnapshotV1) GetPeriodEndOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PeriodEnd, true
}

// SetPeriodEnd sets field value
func (o *UsersPerSourceSnapshotV1) SetPeriodEnd(v float32) {
	o.PeriodEnd = v
}

// GetAnonymous returns the Anonymous field value
func (o *UsersPerSourceSnapshotV1) GetAnonymous() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Anonymous
}

// GetAnonymousOk returns a tuple with the Anonymous field value
// and a boolean to check if the value has been set.
func (o *UsersPerSourceSnapshotV1) GetAnonymousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Anonymous, true
}

// SetAnonymous sets field value
func (o *UsersPerSourceSnapshotV1) SetAnonymous(v string) {
	o.Anonymous = v
}

// GetAnonymousIdentified returns the AnonymousIdentified field value
func (o *UsersPerSourceSnapshotV1) GetAnonymousIdentified() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AnonymousIdentified
}

// GetAnonymousIdentifiedOk returns a tuple with the AnonymousIdentified field value
// and a boolean to check if the value has been set.
func (o *UsersPerSourceSnapshotV1) GetAnonymousIdentifiedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnonymousIdentified, true
}

// SetAnonymousIdentified sets field value
func (o *UsersPerSourceSnapshotV1) SetAnonymousIdentified(v string) {
	o.AnonymousIdentified = v
}

// GetIdentified returns the Identified field value
func (o *UsersPerSourceSnapshotV1) GetIdentified() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identified
}

// GetIdentifiedOk returns a tuple with the Identified field value
// and a boolean to check if the value has been set.
func (o *UsersPerSourceSnapshotV1) GetIdentifiedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identified, true
}

// SetIdentified sets field value
func (o *UsersPerSourceSnapshotV1) SetIdentified(v string) {
	o.Identified = v
}

// GetNeverIdentified returns the NeverIdentified field value
func (o *UsersPerSourceSnapshotV1) GetNeverIdentified() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NeverIdentified
}

// GetNeverIdentifiedOk returns a tuple with the NeverIdentified field value
// and a boolean to check if the value has been set.
func (o *UsersPerSourceSnapshotV1) GetNeverIdentifiedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NeverIdentified, true
}

// SetNeverIdentified sets field value
func (o *UsersPerSourceSnapshotV1) SetNeverIdentified(v string) {
	o.NeverIdentified = v
}

// GetTimestamp returns the Timestamp field value
func (o *UsersPerSourceSnapshotV1) GetTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *UsersPerSourceSnapshotV1) GetTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *UsersPerSourceSnapshotV1) SetTimestamp(v string) {
	o.Timestamp = v
}

func (o UsersPerSourceSnapshotV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsersPerSourceSnapshotV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sourceId"] = o.SourceId
	toSerialize["periodStart"] = o.PeriodStart
	toSerialize["periodEnd"] = o.PeriodEnd
	toSerialize["anonymous"] = o.Anonymous
	toSerialize["anonymousIdentified"] = o.AnonymousIdentified
	toSerialize["identified"] = o.Identified
	toSerialize["neverIdentified"] = o.NeverIdentified
	toSerialize["timestamp"] = o.Timestamp
	return toSerialize, nil
}

type NullableUsersPerSourceSnapshotV1 struct {
	value *UsersPerSourceSnapshotV1
	isSet bool
}

func (v NullableUsersPerSourceSnapshotV1) Get() *UsersPerSourceSnapshotV1 {
	return v.value
}

func (v *NullableUsersPerSourceSnapshotV1) Set(val *UsersPerSourceSnapshotV1) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersPerSourceSnapshotV1) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersPerSourceSnapshotV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersPerSourceSnapshotV1(
	val *UsersPerSourceSnapshotV1,
) *NullableUsersPerSourceSnapshotV1 {
	return &NullableUsersPerSourceSnapshotV1{value: val, isSet: true}
}

func (v NullableUsersPerSourceSnapshotV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersPerSourceSnapshotV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
