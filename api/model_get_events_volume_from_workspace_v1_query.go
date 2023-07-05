/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 36.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GetEventsVolumeFromWorkspaceV1Query GetEventVolumeOutputQuery represents the input query sent to output.
type GetEventsVolumeFromWorkspaceV1Query struct {
	// Workspace being requested.
	WorkspaceId string `json:"workspaceId"`
	// Granularity corresponds to the requested bucket granularity.
	Granularity string `json:"granularity"`
	// StartTime is the ISO8601 formatted timestamp corresponding to the beginning of the requested time frame, inclusive.
	StartTime string `json:"startTime"`
	// EndTime is the ISO8601 formatted timestamp corresponding to the end of the requested time frame, noninclusive.
	EndTime string `json:"endTime"`
	// GroupBy is a comma-delimited list of strings representing the dimensions to group the result by. The current options are: `eventName` or `eventType`.
	GroupBy []string `json:"groupBy,omitempty"`
	// List of strings which allow you to restrict the result to just the given Sources.
	SourceId []string `json:"sourceId,omitempty"`
	// EventName is a list of strings which allow you to restrict the result to just the given event names.
	EventName []string `json:"eventName,omitempty"`
	// EventType is a list of strings which allow you to restrict the result to just the given event types.
	EventType []string `json:"eventType,omitempty"`
	// AppVersion is a list of strings which allow you to restrict the result to just the given application versions.
	AppVersion []string `json:"appVersion,omitempty"`
	// Limit is the total number of items in the result.
	Limit *float32 `json:"limit,omitempty"`
}

// NewGetEventsVolumeFromWorkspaceV1Query instantiates a new GetEventsVolumeFromWorkspaceV1Query object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEventsVolumeFromWorkspaceV1Query(
	workspaceId string,
	granularity string,
	startTime string,
	endTime string,
) *GetEventsVolumeFromWorkspaceV1Query {
	this := GetEventsVolumeFromWorkspaceV1Query{}
	this.WorkspaceId = workspaceId
	this.Granularity = granularity
	this.StartTime = startTime
	this.EndTime = endTime
	return &this
}

// NewGetEventsVolumeFromWorkspaceV1QueryWithDefaults instantiates a new GetEventsVolumeFromWorkspaceV1Query object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEventsVolumeFromWorkspaceV1QueryWithDefaults() *GetEventsVolumeFromWorkspaceV1Query {
	this := GetEventsVolumeFromWorkspaceV1Query{}
	return &this
}

// GetWorkspaceId returns the WorkspaceId field value
func (o *GetEventsVolumeFromWorkspaceV1Query) GetWorkspaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value
// and a boolean to check if the value has been set.
func (o *GetEventsVolumeFromWorkspaceV1Query) GetWorkspaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkspaceId, true
}

// SetWorkspaceId sets field value
func (o *GetEventsVolumeFromWorkspaceV1Query) SetWorkspaceId(v string) {
	o.WorkspaceId = v
}

// GetGranularity returns the Granularity field value
func (o *GetEventsVolumeFromWorkspaceV1Query) GetGranularity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Granularity
}

// GetGranularityOk returns a tuple with the Granularity field value
// and a boolean to check if the value has been set.
func (o *GetEventsVolumeFromWorkspaceV1Query) GetGranularityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Granularity, true
}

// SetGranularity sets field value
func (o *GetEventsVolumeFromWorkspaceV1Query) SetGranularity(v string) {
	o.Granularity = v
}

// GetStartTime returns the StartTime field value
func (o *GetEventsVolumeFromWorkspaceV1Query) GetStartTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *GetEventsVolumeFromWorkspaceV1Query) GetStartTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *GetEventsVolumeFromWorkspaceV1Query) SetStartTime(v string) {
	o.StartTime = v
}

// GetEndTime returns the EndTime field value
func (o *GetEventsVolumeFromWorkspaceV1Query) GetEndTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *GetEventsVolumeFromWorkspaceV1Query) GetEndTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *GetEventsVolumeFromWorkspaceV1Query) SetEndTime(v string) {
	o.EndTime = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *GetEventsVolumeFromWorkspaceV1Query) GetGroupBy() []string {
	if o == nil || o.GroupBy == nil {
		var ret []string
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEventsVolumeFromWorkspaceV1Query) GetGroupByOk() ([]string, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *GetEventsVolumeFromWorkspaceV1Query) HasGroupBy() bool {
	if o != nil && o.GroupBy != nil {
		return true
	}

	return false
}

// SetGroupBy gets a reference to the given []string and assigns it to the GroupBy field.
func (o *GetEventsVolumeFromWorkspaceV1Query) SetGroupBy(v []string) {
	o.GroupBy = v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *GetEventsVolumeFromWorkspaceV1Query) GetSourceId() []string {
	if o == nil || o.SourceId == nil {
		var ret []string
		return ret
	}
	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEventsVolumeFromWorkspaceV1Query) GetSourceIdOk() ([]string, bool) {
	if o == nil || o.SourceId == nil {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *GetEventsVolumeFromWorkspaceV1Query) HasSourceId() bool {
	if o != nil && o.SourceId != nil {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given []string and assigns it to the SourceId field.
func (o *GetEventsVolumeFromWorkspaceV1Query) SetSourceId(v []string) {
	o.SourceId = v
}

// GetEventName returns the EventName field value if set, zero value otherwise.
func (o *GetEventsVolumeFromWorkspaceV1Query) GetEventName() []string {
	if o == nil || o.EventName == nil {
		var ret []string
		return ret
	}
	return o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEventsVolumeFromWorkspaceV1Query) GetEventNameOk() ([]string, bool) {
	if o == nil || o.EventName == nil {
		return nil, false
	}
	return o.EventName, true
}

// HasEventName returns a boolean if a field has been set.
func (o *GetEventsVolumeFromWorkspaceV1Query) HasEventName() bool {
	if o != nil && o.EventName != nil {
		return true
	}

	return false
}

// SetEventName gets a reference to the given []string and assigns it to the EventName field.
func (o *GetEventsVolumeFromWorkspaceV1Query) SetEventName(v []string) {
	o.EventName = v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *GetEventsVolumeFromWorkspaceV1Query) GetEventType() []string {
	if o == nil || o.EventType == nil {
		var ret []string
		return ret
	}
	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEventsVolumeFromWorkspaceV1Query) GetEventTypeOk() ([]string, bool) {
	if o == nil || o.EventType == nil {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *GetEventsVolumeFromWorkspaceV1Query) HasEventType() bool {
	if o != nil && o.EventType != nil {
		return true
	}

	return false
}

// SetEventType gets a reference to the given []string and assigns it to the EventType field.
func (o *GetEventsVolumeFromWorkspaceV1Query) SetEventType(v []string) {
	o.EventType = v
}

// GetAppVersion returns the AppVersion field value if set, zero value otherwise.
func (o *GetEventsVolumeFromWorkspaceV1Query) GetAppVersion() []string {
	if o == nil || o.AppVersion == nil {
		var ret []string
		return ret
	}
	return o.AppVersion
}

// GetAppVersionOk returns a tuple with the AppVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEventsVolumeFromWorkspaceV1Query) GetAppVersionOk() ([]string, bool) {
	if o == nil || o.AppVersion == nil {
		return nil, false
	}
	return o.AppVersion, true
}

// HasAppVersion returns a boolean if a field has been set.
func (o *GetEventsVolumeFromWorkspaceV1Query) HasAppVersion() bool {
	if o != nil && o.AppVersion != nil {
		return true
	}

	return false
}

// SetAppVersion gets a reference to the given []string and assigns it to the AppVersion field.
func (o *GetEventsVolumeFromWorkspaceV1Query) SetAppVersion(v []string) {
	o.AppVersion = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *GetEventsVolumeFromWorkspaceV1Query) GetLimit() float32 {
	if o == nil || o.Limit == nil {
		var ret float32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEventsVolumeFromWorkspaceV1Query) GetLimitOk() (*float32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *GetEventsVolumeFromWorkspaceV1Query) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given float32 and assigns it to the Limit field.
func (o *GetEventsVolumeFromWorkspaceV1Query) SetLimit(v float32) {
	o.Limit = &v
}

func (o GetEventsVolumeFromWorkspaceV1Query) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["workspaceId"] = o.WorkspaceId
	}
	if true {
		toSerialize["granularity"] = o.Granularity
	}
	if true {
		toSerialize["startTime"] = o.StartTime
	}
	if true {
		toSerialize["endTime"] = o.EndTime
	}
	if o.GroupBy != nil {
		toSerialize["groupBy"] = o.GroupBy
	}
	if o.SourceId != nil {
		toSerialize["sourceId"] = o.SourceId
	}
	if o.EventName != nil {
		toSerialize["eventName"] = o.EventName
	}
	if o.EventType != nil {
		toSerialize["eventType"] = o.EventType
	}
	if o.AppVersion != nil {
		toSerialize["appVersion"] = o.AppVersion
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	return json.Marshal(toSerialize)
}

type NullableGetEventsVolumeFromWorkspaceV1Query struct {
	value *GetEventsVolumeFromWorkspaceV1Query
	isSet bool
}

func (v NullableGetEventsVolumeFromWorkspaceV1Query) Get() *GetEventsVolumeFromWorkspaceV1Query {
	return v.value
}

func (v *NullableGetEventsVolumeFromWorkspaceV1Query) Set(
	val *GetEventsVolumeFromWorkspaceV1Query,
) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEventsVolumeFromWorkspaceV1Query) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEventsVolumeFromWorkspaceV1Query) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEventsVolumeFromWorkspaceV1Query(
	val *GetEventsVolumeFromWorkspaceV1Query,
) *NullableGetEventsVolumeFromWorkspaceV1Query {
	return &NullableGetEventsVolumeFromWorkspaceV1Query{value: val, isSet: true}
}

func (v NullableGetEventsVolumeFromWorkspaceV1Query) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEventsVolumeFromWorkspaceV1Query) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
