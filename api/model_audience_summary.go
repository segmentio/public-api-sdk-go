/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 50.4.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the AudienceSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AudienceSummary{}

// AudienceSummary Defines an Audience.
type AudienceSummary struct {
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
	Enabled    bool               `json:"enabled"`
	Definition NullableDefinition `json:"definition"`
	// Status for the audience.  Possible values: Backfilling, Computing, Failed, Live, Awaiting Destinations, Disabled.
	Status *string `json:"status,omitempty"`
	// User id who created the audience.
	CreatedBy string `json:"createdBy"`
	// User id who last updated the audience.
	UpdatedBy string `json:"updatedBy"`
	// Date the audience was created.
	CreatedAt string `json:"createdAt"`
	// Date the audience was last updated.
	UpdatedAt string           `json:"updatedAt"`
	Options   *AudienceOptions `json:"options,omitempty"`
}

// NewAudienceSummary instantiates a new AudienceSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAudienceSummary(
	id string,
	spaceId string,
	name string,
	key string,
	enabled bool,
	definition NullableDefinition,
	createdBy string,
	updatedBy string,
	createdAt string,
	updatedAt string,
) *AudienceSummary {
	this := AudienceSummary{}
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

// NewAudienceSummaryWithDefaults instantiates a new AudienceSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAudienceSummaryWithDefaults() *AudienceSummary {
	this := AudienceSummary{}
	return &this
}

// GetId returns the Id field value
func (o *AudienceSummary) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AudienceSummary) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AudienceSummary) SetId(v string) {
	o.Id = v
}

// GetSpaceId returns the SpaceId field value
func (o *AudienceSummary) GetSpaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpaceId
}

// GetSpaceIdOk returns a tuple with the SpaceId field value
// and a boolean to check if the value has been set.
func (o *AudienceSummary) GetSpaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpaceId, true
}

// SetSpaceId sets field value
func (o *AudienceSummary) SetSpaceId(v string) {
	o.SpaceId = v
}

// GetName returns the Name field value
func (o *AudienceSummary) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AudienceSummary) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AudienceSummary) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AudienceSummary) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudienceSummary) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AudienceSummary) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AudienceSummary) SetDescription(v string) {
	o.Description = &v
}

// GetKey returns the Key field value
func (o *AudienceSummary) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *AudienceSummary) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *AudienceSummary) SetKey(v string) {
	o.Key = v
}

// GetEnabled returns the Enabled field value
func (o *AudienceSummary) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AudienceSummary) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AudienceSummary) SetEnabled(v bool) {
	o.Enabled = v
}

// GetDefinition returns the Definition field value
// If the value is explicit nil, the zero value for Definition will be returned
func (o *AudienceSummary) GetDefinition() Definition {
	if o == nil || o.Definition.Get() == nil {
		var ret Definition
		return ret
	}

	return *o.Definition.Get()
}

// GetDefinitionOk returns a tuple with the Definition field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AudienceSummary) GetDefinitionOk() (*Definition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Definition.Get(), o.Definition.IsSet()
}

// SetDefinition sets field value
func (o *AudienceSummary) SetDefinition(v Definition) {
	o.Definition.Set(&v)
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AudienceSummary) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudienceSummary) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AudienceSummary) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AudienceSummary) SetStatus(v string) {
	o.Status = &v
}

// GetCreatedBy returns the CreatedBy field value
func (o *AudienceSummary) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *AudienceSummary) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *AudienceSummary) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetUpdatedBy returns the UpdatedBy field value
func (o *AudienceSummary) GetUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value
// and a boolean to check if the value has been set.
func (o *AudienceSummary) GetUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedBy, true
}

// SetUpdatedBy sets field value
func (o *AudienceSummary) SetUpdatedBy(v string) {
	o.UpdatedBy = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *AudienceSummary) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AudienceSummary) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AudienceSummary) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *AudienceSummary) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *AudienceSummary) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *AudienceSummary) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *AudienceSummary) GetOptions() AudienceOptions {
	if o == nil || IsNil(o.Options) {
		var ret AudienceOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudienceSummary) GetOptionsOk() (*AudienceOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *AudienceSummary) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given AudienceOptions and assigns it to the Options field.
func (o *AudienceSummary) SetOptions(v AudienceOptions) {
	o.Options = &v
}

func (o AudienceSummary) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AudienceSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

type NullableAudienceSummary struct {
	value *AudienceSummary
	isSet bool
}

func (v NullableAudienceSummary) Get() *AudienceSummary {
	return v.value
}

func (v *NullableAudienceSummary) Set(val *AudienceSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableAudienceSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableAudienceSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAudienceSummary(val *AudienceSummary) *NullableAudienceSummary {
	return &NullableAudienceSummary{value: val, isSet: true}
}

func (v NullableAudienceSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAudienceSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
