/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 57.0.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the AuditEventV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditEventV1{}

// AuditEventV1 Represents an Audit Trail event.
type AuditEventV1 struct {
	// Unique identifier for this audit trail event.
	Id string `json:"id"`
	// The timestamp of this event in ISO-8601 format.
	Timestamp string `json:"timestamp"`
	// The type of this event.
	Type string `json:"type"`
	// The user or API token that triggered this event.
	Actor string `json:"actor"`
	// The email of the user that triggered this event.
	ActorEmail *string `json:"actorEmail,omitempty"`
	// The identifier of the resource affected by this event.
	ResourceId string `json:"resourceId"`
	// The kind of resource affected by this event.
	ResourceType string `json:"resourceType"`
	// The name of the resource affected by this event.
	ResourceName string `json:"resourceName"`
}

// NewAuditEventV1 instantiates a new AuditEventV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditEventV1(
	id string,
	timestamp string,
	type_ string,
	actor string,
	resourceId string,
	resourceType string,
	resourceName string,
) *AuditEventV1 {
	this := AuditEventV1{}
	this.Id = id
	this.Timestamp = timestamp
	this.Type = type_
	this.Actor = actor
	this.ResourceId = resourceId
	this.ResourceType = resourceType
	this.ResourceName = resourceName
	return &this
}

// NewAuditEventV1WithDefaults instantiates a new AuditEventV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditEventV1WithDefaults() *AuditEventV1 {
	this := AuditEventV1{}
	return &this
}

// GetId returns the Id field value
func (o *AuditEventV1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AuditEventV1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AuditEventV1) SetId(v string) {
	o.Id = v
}

// GetTimestamp returns the Timestamp field value
func (o *AuditEventV1) GetTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *AuditEventV1) GetTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *AuditEventV1) SetTimestamp(v string) {
	o.Timestamp = v
}

// GetType returns the Type field value
func (o *AuditEventV1) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AuditEventV1) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AuditEventV1) SetType(v string) {
	o.Type = v
}

// GetActor returns the Actor field value
func (o *AuditEventV1) GetActor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Actor
}

// GetActorOk returns a tuple with the Actor field value
// and a boolean to check if the value has been set.
func (o *AuditEventV1) GetActorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Actor, true
}

// SetActor sets field value
func (o *AuditEventV1) SetActor(v string) {
	o.Actor = v
}

// GetActorEmail returns the ActorEmail field value if set, zero value otherwise.
func (o *AuditEventV1) GetActorEmail() string {
	if o == nil || IsNil(o.ActorEmail) {
		var ret string
		return ret
	}
	return *o.ActorEmail
}

// GetActorEmailOk returns a tuple with the ActorEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditEventV1) GetActorEmailOk() (*string, bool) {
	if o == nil || IsNil(o.ActorEmail) {
		return nil, false
	}
	return o.ActorEmail, true
}

// HasActorEmail returns a boolean if a field has been set.
func (o *AuditEventV1) HasActorEmail() bool {
	if o != nil && !IsNil(o.ActorEmail) {
		return true
	}

	return false
}

// SetActorEmail gets a reference to the given string and assigns it to the ActorEmail field.
func (o *AuditEventV1) SetActorEmail(v string) {
	o.ActorEmail = &v
}

// GetResourceId returns the ResourceId field value
func (o *AuditEventV1) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *AuditEventV1) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *AuditEventV1) SetResourceId(v string) {
	o.ResourceId = v
}

// GetResourceType returns the ResourceType field value
func (o *AuditEventV1) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *AuditEventV1) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *AuditEventV1) SetResourceType(v string) {
	o.ResourceType = v
}

// GetResourceName returns the ResourceName field value
func (o *AuditEventV1) GetResourceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value
// and a boolean to check if the value has been set.
func (o *AuditEventV1) GetResourceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceName, true
}

// SetResourceName sets field value
func (o *AuditEventV1) SetResourceName(v string) {
	o.ResourceName = v
}

func (o AuditEventV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditEventV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["type"] = o.Type
	toSerialize["actor"] = o.Actor
	if !IsNil(o.ActorEmail) {
		toSerialize["actorEmail"] = o.ActorEmail
	}
	toSerialize["resourceId"] = o.ResourceId
	toSerialize["resourceType"] = o.ResourceType
	toSerialize["resourceName"] = o.ResourceName
	return toSerialize, nil
}

type NullableAuditEventV1 struct {
	value *AuditEventV1
	isSet bool
}

func (v NullableAuditEventV1) Get() *AuditEventV1 {
	return v.value
}

func (v *NullableAuditEventV1) Set(val *AuditEventV1) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditEventV1) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditEventV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditEventV1(val *AuditEventV1) *NullableAuditEventV1 {
	return &NullableAuditEventV1{value: val, isSet: true}
}

func (v NullableAuditEventV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditEventV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
