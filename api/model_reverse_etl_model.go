/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 37.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ReverseEtlModel The created Model.
type ReverseEtlModel struct {
	// The id of the Model.
	Id string `json:"id"`
	// The id for the attached Source.
	SourceId string `json:"sourceId"`
	// A short, human-readable description of the Model.
	Name string `json:"name"`
	// A longer, more descriptive explanation of the Model.
	Description string `json:"description"`
	// Indicates whether the Model should have syncs enabled. When disabled, no syncs will be triggered, regardless of the enabled status of the attached destinations/subscriptions.
	Enabled bool `json:"enabled"`
	// Determines the strategy used for triggering syncs, which will be used in conjunction with scheduleConfig.  Possible values: \"manual\", \"periodic\", \"specific_days\".
	ScheduleStrategy string `json:"scheduleStrategy"`
	// Depending on the chosen strategy, configures the schedule for this model.
	ScheduleConfig NullableModelMap `json:"scheduleConfig"`
	// The SQL query that will be executed to extract data from the connected Source.
	Query string `json:"query"`
	// Indicates the column named in `query` that should be used to uniquely identify the extracted records.
	QueryIdentifierColumn string `json:"queryIdentifierColumn"`
}

// NewReverseEtlModel instantiates a new ReverseEtlModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReverseEtlModel(
	id string,
	sourceId string,
	name string,
	description string,
	enabled bool,
	scheduleStrategy string,
	scheduleConfig NullableModelMap,
	query string,
	queryIdentifierColumn string,
) *ReverseEtlModel {
	this := ReverseEtlModel{}
	this.Id = id
	this.SourceId = sourceId
	this.Name = name
	this.Description = description
	this.Enabled = enabled
	this.ScheduleStrategy = scheduleStrategy
	this.ScheduleConfig = scheduleConfig
	this.Query = query
	this.QueryIdentifierColumn = queryIdentifierColumn
	return &this
}

// NewReverseEtlModelWithDefaults instantiates a new ReverseEtlModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReverseEtlModelWithDefaults() *ReverseEtlModel {
	this := ReverseEtlModel{}
	return &this
}

// GetId returns the Id field value
func (o *ReverseEtlModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReverseEtlModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ReverseEtlModel) SetId(v string) {
	o.Id = v
}

// GetSourceId returns the SourceId field value
func (o *ReverseEtlModel) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *ReverseEtlModel) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *ReverseEtlModel) SetSourceId(v string) {
	o.SourceId = v
}

// GetName returns the Name field value
func (o *ReverseEtlModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ReverseEtlModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ReverseEtlModel) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *ReverseEtlModel) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ReverseEtlModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ReverseEtlModel) SetDescription(v string) {
	o.Description = v
}

// GetEnabled returns the Enabled field value
func (o *ReverseEtlModel) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ReverseEtlModel) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ReverseEtlModel) SetEnabled(v bool) {
	o.Enabled = v
}

// GetScheduleStrategy returns the ScheduleStrategy field value
func (o *ReverseEtlModel) GetScheduleStrategy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScheduleStrategy
}

// GetScheduleStrategyOk returns a tuple with the ScheduleStrategy field value
// and a boolean to check if the value has been set.
func (o *ReverseEtlModel) GetScheduleStrategyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduleStrategy, true
}

// SetScheduleStrategy sets field value
func (o *ReverseEtlModel) SetScheduleStrategy(v string) {
	o.ScheduleStrategy = v
}

// GetScheduleConfig returns the ScheduleConfig field value
// If the value is explicit nil, the zero value for ModelMap will be returned
func (o *ReverseEtlModel) GetScheduleConfig() ModelMap {
	if o == nil || o.ScheduleConfig.Get() == nil {
		var ret ModelMap
		return ret
	}

	return *o.ScheduleConfig.Get()
}

// GetScheduleConfigOk returns a tuple with the ScheduleConfig field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReverseEtlModel) GetScheduleConfigOk() (*ModelMap, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScheduleConfig.Get(), o.ScheduleConfig.IsSet()
}

// SetScheduleConfig sets field value
func (o *ReverseEtlModel) SetScheduleConfig(v ModelMap) {
	o.ScheduleConfig.Set(&v)
}

// GetQuery returns the Query field value
func (o *ReverseEtlModel) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *ReverseEtlModel) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *ReverseEtlModel) SetQuery(v string) {
	o.Query = v
}

// GetQueryIdentifierColumn returns the QueryIdentifierColumn field value
func (o *ReverseEtlModel) GetQueryIdentifierColumn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryIdentifierColumn
}

// GetQueryIdentifierColumnOk returns a tuple with the QueryIdentifierColumn field value
// and a boolean to check if the value has been set.
func (o *ReverseEtlModel) GetQueryIdentifierColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryIdentifierColumn, true
}

// SetQueryIdentifierColumn sets field value
func (o *ReverseEtlModel) SetQueryIdentifierColumn(v string) {
	o.QueryIdentifierColumn = v
}

func (o ReverseEtlModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["sourceId"] = o.SourceId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["scheduleStrategy"] = o.ScheduleStrategy
	}
	if true {
		toSerialize["scheduleConfig"] = o.ScheduleConfig.Get()
	}
	if true {
		toSerialize["query"] = o.Query
	}
	if true {
		toSerialize["queryIdentifierColumn"] = o.QueryIdentifierColumn
	}
	return json.Marshal(toSerialize)
}

type NullableReverseEtlModel struct {
	value *ReverseEtlModel
	isSet bool
}

func (v NullableReverseEtlModel) Get() *ReverseEtlModel {
	return v.value
}

func (v *NullableReverseEtlModel) Set(val *ReverseEtlModel) {
	v.value = val
	v.isSet = true
}

func (v NullableReverseEtlModel) IsSet() bool {
	return v.isSet
}

func (v *NullableReverseEtlModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReverseEtlModel(val *ReverseEtlModel) *NullableReverseEtlModel {
	return &NullableReverseEtlModel{value: val, isSet: true}
}

func (v NullableReverseEtlModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReverseEtlModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
