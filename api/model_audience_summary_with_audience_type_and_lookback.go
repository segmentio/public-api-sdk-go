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

// checks if the AudienceSummaryWithAudienceTypeAndLookback type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AudienceSummaryWithAudienceTypeAndLookback{}

// AudienceSummaryWithAudienceTypeAndLookback struct for AudienceSummaryWithAudienceTypeAndLookback
type AudienceSummaryWithAudienceTypeAndLookback struct {
	// Discriminator denoting the audience's product type.
	AudienceType   string                       `json:"audienceType"`
	ComputeCadence AudienceComputeCadence       `json:"computeCadence"`
	Size           *AudienceSize                `json:"size,omitempty"`
	Options        *AudienceOptionsWithLookback `json:"options,omitempty"`
	// List of schedules for the audience.
	Schedules []AudienceSchedule `json:"schedules,omitempty"`
	// Audience id.
	Id string `json:"id"`
	// Space id for the audience.
	SpaceId string `json:"spaceId"`
	// Name of the audience.
	Name string `json:"name"`
	// Description of the audience.
	Description *string `json:"description,omitempty"`
	// Key for the audience.
	Key string `json:"key"`
	// Enabled/disabled status for the audience.
	Enabled    bool                       `json:"enabled"`
	Definition NullableAudienceDefinition `json:"definition"`
	// Status for the audience.  Possible values: Backfilling, Computing, Failed, Live, Awaiting Destinations, Disabled.
	Status *string `json:"status,omitempty"`
	// User id who created the audience.
	CreatedBy string `json:"createdBy"`
	// User id who last updated the audience.
	UpdatedBy string `json:"updatedBy"`
	// Date the audience was created.
	CreatedAt string `json:"createdAt"`
	// Date the audience was last updated.
	UpdatedAt string `json:"updatedAt"`
}

// NewAudienceSummaryWithAudienceTypeAndLookback instantiates a new AudienceSummaryWithAudienceTypeAndLookback object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAudienceSummaryWithAudienceTypeAndLookback(
	audienceType string,
	computeCadence AudienceComputeCadence,
	id string,
	spaceId string,
	name string,
	key string,
	enabled bool,
	definition NullableAudienceDefinition,
	createdBy string,
	updatedBy string,
	createdAt string,
	updatedAt string,
) *AudienceSummaryWithAudienceTypeAndLookback {
	this := AudienceSummaryWithAudienceTypeAndLookback{}
	this.AudienceType = audienceType
	this.ComputeCadence = computeCadence
	this.Id = id
	this.SpaceId = spaceId
	this.Name = name
	this.Key = key
	this.Enabled = enabled
	this.Definition = definition
	this.CreatedBy = createdBy
	this.UpdatedBy = updatedBy
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewAudienceSummaryWithAudienceTypeAndLookbackWithDefaults instantiates a new AudienceSummaryWithAudienceTypeAndLookback object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAudienceSummaryWithAudienceTypeAndLookbackWithDefaults() *AudienceSummaryWithAudienceTypeAndLookback {
	this := AudienceSummaryWithAudienceTypeAndLookback{}
	return &this
}

// GetAudienceType returns the AudienceType field value
func (o *AudienceSummaryWithAudienceTypeAndLookback) GetAudienceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AudienceType
}

// GetAudienceTypeOk returns a tuple with the AudienceType field value
// and a boolean to check if the value has been set.
func (o *AudienceSummaryWithAudienceTypeAndLookback) GetAudienceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AudienceType, true
}

// SetAudienceType sets field value
func (o *AudienceSummaryWithAudienceTypeAndLookback) SetAudienceType(v string) {
	o.AudienceType = v
}

// GetComputeCadence returns the ComputeCadence field value
func (o *AudienceSummaryWithAudienceTypeAndLookback) GetComputeCadence() AudienceComputeCadence {
	if o == nil {
		var ret AudienceComputeCadence
		return ret
	}

	return o.ComputeCadence
}

// GetComputeCadenceOk returns a tuple with the ComputeCadence field value
// and a boolean to check if the value has been set.
func (o *AudienceSummaryWithAudienceTypeAndLookback) GetComputeCadenceOk() (*AudienceComputeCadence, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComputeCadence, true
}

// SetComputeCadence sets field value
func (o *AudienceSummaryWithAudienceTypeAndLookback) SetComputeCadence(v AudienceComputeCadence) {
	o.ComputeCadence = v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *AudienceSummaryWithAudienceTypeAndLookback) GetSize() AudienceSize {
	if o == nil || IsNil(o.Size) {
		var ret AudienceSize
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudienceSummaryWithAudienceTypeAndLookback) GetSizeOk() (*AudienceSize, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *AudienceSummaryWithAudienceTypeAndLookback) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given AudienceSize and assigns it to the Size field.
func (o *AudienceSummaryWithAudienceTypeAndLookback) SetSize(v AudienceSize) {
	o.Size = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *AudienceSummaryWithAudienceTypeAndLookback) GetOptions() AudienceOptionsWithLookback {
	if o == nil || IsNil(o.Options) {
		var ret AudienceOptionsWithLookback
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudienceSummaryWithAudienceTypeAndLookback) GetOptionsOk() (*AudienceOptionsWithLookback, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *AudienceSummaryWithAudienceTypeAndLookback) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given AudienceOptionsWithLookback and assigns it to the Options field.
func (o *AudienceSummaryWithAudienceTypeAndLookback) SetOptions(v AudienceOptionsWithLookback) {
	o.Options = &v
}

// GetSchedules returns the Schedules field value if set, zero value otherwise.
func (o *AudienceSummaryWithAudienceTypeAndLookback) GetSchedules() []AudienceSchedule {
	if o == nil || IsNil(o.Schedules) {
		var ret []AudienceSchedule
		return ret
	}
	return o.Schedules
}

// GetSchedulesOk returns a tuple with the Schedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudienceSummaryWithAudienceTypeAndLookback) GetSchedulesOk() ([]AudienceSchedule, bool) {
	if o == nil || IsNil(o.Schedules) {
		return nil, false
	}
	return o.Schedules, true
}

// HasSchedules returns a boolean if a field has been set.
func (o *AudienceSummaryWithAudienceTypeAndLookback) HasSchedules() bool {
	if o != nil && !IsNil(o.Schedules) {
		return true
	}

	return false
}

// SetSchedules gets a reference to the given []AudienceSchedule and assigns it to the Schedules field.
func (o *AudienceSummaryWithAudienceTypeAndLookback) SetSchedules(v []AudienceSchedule) {
	o.Schedules = v
}

// GetId returns the Id field value
func (o *AudienceSummaryWithAudienceTypeAndLookback) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AudienceSummaryWithAudienceTypeAndLookback) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AudienceSummaryWithAudienceTypeAndLookback) SetId(v string) {
	o.Id = v
}

// GetSpaceId returns the SpaceId field value
func (o *AudienceSummaryWithAudienceTypeAndLookback) GetSpaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpaceId
}

// GetSpaceIdOk returns a tuple with the SpaceId field value
// and a boolean to check if the value has been set.
func (o *AudienceSummaryWithAudienceTypeAndLookback) GetSpaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpaceId, true
}

// SetSpaceId sets field value
func (o *AudienceSummaryWithAudienceTypeAndLookback) SetSpaceId(v string) {
	o.SpaceId = v
}

// GetName returns the Name field value
func (o *AudienceSummaryWithAudienceTypeAndLookback) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AudienceSummaryWithAudienceTypeAndLookback) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AudienceSummaryWithAudienceTypeAndLookback) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AudienceSummaryWithAudienceTypeAndLookback) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudienceSummaryWithAudienceTypeAndLookback) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AudienceSummaryWithAudienceTypeAndLookback) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AudienceSummaryWithAudienceTypeAndLookback) SetDescription(v string) {
	o.Description = &v
}

// GetKey returns the Key field value
func (o *AudienceSummaryWithAudienceTypeAndLookback) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *AudienceSummaryWithAudienceTypeAndLookback) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *AudienceSummaryWithAudienceTypeAndLookback) SetKey(v string) {
	o.Key = v
}

// GetEnabled returns the Enabled field value
func (o *AudienceSummaryWithAudienceTypeAndLookback) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AudienceSummaryWithAudienceTypeAndLookback) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AudienceSummaryWithAudienceTypeAndLookback) SetEnabled(v bool) {
	o.Enabled = v
}

// GetDefinition returns the Definition field value
// If the value is explicit nil, the zero value for AudienceDefinition will be returned
func (o *AudienceSummaryWithAudienceTypeAndLookback) GetDefinition() AudienceDefinition {
	if o == nil || o.Definition.Get() == nil {
		var ret AudienceDefinition
		return ret
	}

	return *o.Definition.Get()
}

// GetDefinitionOk returns a tuple with the Definition field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AudienceSummaryWithAudienceTypeAndLookback) GetDefinitionOk() (*AudienceDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Definition.Get(), o.Definition.IsSet()
}

// SetDefinition sets field value
func (o *AudienceSummaryWithAudienceTypeAndLookback) SetDefinition(v AudienceDefinition) {
	o.Definition.Set(&v)
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AudienceSummaryWithAudienceTypeAndLookback) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudienceSummaryWithAudienceTypeAndLookback) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AudienceSummaryWithAudienceTypeAndLookback) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AudienceSummaryWithAudienceTypeAndLookback) SetStatus(v string) {
	o.Status = &v
}

// GetCreatedBy returns the CreatedBy field value
func (o *AudienceSummaryWithAudienceTypeAndLookback) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *AudienceSummaryWithAudienceTypeAndLookback) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *AudienceSummaryWithAudienceTypeAndLookback) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetUpdatedBy returns the UpdatedBy field value
func (o *AudienceSummaryWithAudienceTypeAndLookback) GetUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value
// and a boolean to check if the value has been set.
func (o *AudienceSummaryWithAudienceTypeAndLookback) GetUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedBy, true
}

// SetUpdatedBy sets field value
func (o *AudienceSummaryWithAudienceTypeAndLookback) SetUpdatedBy(v string) {
	o.UpdatedBy = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *AudienceSummaryWithAudienceTypeAndLookback) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AudienceSummaryWithAudienceTypeAndLookback) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AudienceSummaryWithAudienceTypeAndLookback) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *AudienceSummaryWithAudienceTypeAndLookback) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *AudienceSummaryWithAudienceTypeAndLookback) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *AudienceSummaryWithAudienceTypeAndLookback) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o AudienceSummaryWithAudienceTypeAndLookback) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AudienceSummaryWithAudienceTypeAndLookback) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["audienceType"] = o.AudienceType
	toSerialize["computeCadence"] = o.ComputeCadence
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.Schedules) {
		toSerialize["schedules"] = o.Schedules
	}
	toSerialize["id"] = o.Id
	toSerialize["spaceId"] = o.SpaceId
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["key"] = o.Key
	toSerialize["enabled"] = o.Enabled
	toSerialize["definition"] = o.Definition.Get()
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["updatedBy"] = o.UpdatedBy
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	return toSerialize, nil
}

type NullableAudienceSummaryWithAudienceTypeAndLookback struct {
	value *AudienceSummaryWithAudienceTypeAndLookback
	isSet bool
}

func (v NullableAudienceSummaryWithAudienceTypeAndLookback) Get() *AudienceSummaryWithAudienceTypeAndLookback {
	return v.value
}

func (v *NullableAudienceSummaryWithAudienceTypeAndLookback) Set(
	val *AudienceSummaryWithAudienceTypeAndLookback,
) {
	v.value = val
	v.isSet = true
}

func (v NullableAudienceSummaryWithAudienceTypeAndLookback) IsSet() bool {
	return v.isSet
}

func (v *NullableAudienceSummaryWithAudienceTypeAndLookback) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAudienceSummaryWithAudienceTypeAndLookback(
	val *AudienceSummaryWithAudienceTypeAndLookback,
) *NullableAudienceSummaryWithAudienceTypeAndLookback {
	return &NullableAudienceSummaryWithAudienceTypeAndLookback{value: val, isSet: true}
}

func (v NullableAudienceSummaryWithAudienceTypeAndLookback) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAudienceSummaryWithAudienceTypeAndLookback) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
