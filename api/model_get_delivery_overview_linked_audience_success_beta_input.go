/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 50.3.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GetDeliveryOverviewLinkedAudienceSuccessBetaInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeliveryOverviewLinkedAudienceSuccessBetaInput{}

// GetDeliveryOverviewLinkedAudienceSuccessBetaInput Input of the Delivery Overview Successful Delivery step.
type GetDeliveryOverviewLinkedAudienceSuccessBetaInput struct {
	// The sourceId for the Workspace.
	SourceId string `json:"sourceId"`
	// The id tied to a Workspace Destination.
	DestinationConfigId string `json:"destinationConfigId"`
	// The ISO8601 formatted timestamp corresponding to the beginning of the requested timeframe, inclusive.
	StartTime string `json:"startTime"`
	// The ISO8601 formatted timestamp corresponding to the end of the requested timeframe, noninclusive.
	EndTime string `json:"endTime"`
	// A comma-delimited list of strings representing one or more dimensions to group the result by.  Valid options are: `eventName`, `eventType`, `activationId`, `audienceId`, and `spaceId`.
	GroupBy []string `json:"groupBy,omitempty"`
	// The size of each bucket in the requested window.  Based on the granularity chosen, there are restrictions on the time range you can query:  **Minute**: - Max time range: 4 hours - Oldest possible start time: 48 hours in the past  **Hour**: - Max Time range: 14 days - Oldest possible start time: 30 days in the past  **Day**: - Max time range: 30 days - Oldest possible start time: 30 days in the past
	Granularity string                            `json:"granularity"`
	Filter      *DeliveryOverviewAudienceFilterBy `json:"filter,omitempty"`
	Pagination  *PaginationInput                  `json:"pagination,omitempty"`
}

// NewGetDeliveryOverviewLinkedAudienceSuccessBetaInput instantiates a new GetDeliveryOverviewLinkedAudienceSuccessBetaInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeliveryOverviewLinkedAudienceSuccessBetaInput(
	sourceId string,
	destinationConfigId string,
	startTime string,
	endTime string,
	granularity string,
) *GetDeliveryOverviewLinkedAudienceSuccessBetaInput {
	this := GetDeliveryOverviewLinkedAudienceSuccessBetaInput{}
	this.SourceId = sourceId
	this.DestinationConfigId = destinationConfigId
	this.StartTime = startTime
	this.EndTime = endTime
	this.Granularity = granularity
	return &this
}

// NewGetDeliveryOverviewLinkedAudienceSuccessBetaInputWithDefaults instantiates a new GetDeliveryOverviewLinkedAudienceSuccessBetaInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeliveryOverviewLinkedAudienceSuccessBetaInputWithDefaults() *GetDeliveryOverviewLinkedAudienceSuccessBetaInput {
	this := GetDeliveryOverviewLinkedAudienceSuccessBetaInput{}
	return &this
}

// GetSourceId returns the SourceId field value
func (o *GetDeliveryOverviewLinkedAudienceSuccessBetaInput) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *GetDeliveryOverviewLinkedAudienceSuccessBetaInput) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *GetDeliveryOverviewLinkedAudienceSuccessBetaInput) SetSourceId(v string) {
	o.SourceId = v
}

// GetDestinationConfigId returns the DestinationConfigId field value
func (o *GetDeliveryOverviewLinkedAudienceSuccessBetaInput) GetDestinationConfigId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationConfigId
}

// GetDestinationConfigIdOk returns a tuple with the DestinationConfigId field value
// and a boolean to check if the value has been set.
func (o *GetDeliveryOverviewLinkedAudienceSuccessBetaInput) GetDestinationConfigIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationConfigId, true
}

// SetDestinationConfigId sets field value
func (o *GetDeliveryOverviewLinkedAudienceSuccessBetaInput) SetDestinationConfigId(v string) {
	o.DestinationConfigId = v
}

// GetStartTime returns the StartTime field value
func (o *GetDeliveryOverviewLinkedAudienceSuccessBetaInput) GetStartTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *GetDeliveryOverviewLinkedAudienceSuccessBetaInput) GetStartTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *GetDeliveryOverviewLinkedAudienceSuccessBetaInput) SetStartTime(v string) {
	o.StartTime = v
}

// GetEndTime returns the EndTime field value
func (o *GetDeliveryOverviewLinkedAudienceSuccessBetaInput) GetEndTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *GetDeliveryOverviewLinkedAudienceSuccessBetaInput) GetEndTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *GetDeliveryOverviewLinkedAudienceSuccessBetaInput) SetEndTime(v string) {
	o.EndTime = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *GetDeliveryOverviewLinkedAudienceSuccessBetaInput) GetGroupBy() []string {
	if o == nil || IsNil(o.GroupBy) {
		var ret []string
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeliveryOverviewLinkedAudienceSuccessBetaInput) GetGroupByOk() ([]string, bool) {
	if o == nil || IsNil(o.GroupBy) {
		return nil, false
	}
	return o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *GetDeliveryOverviewLinkedAudienceSuccessBetaInput) HasGroupBy() bool {
	if o != nil && !IsNil(o.GroupBy) {
		return true
	}

	return false
}

// SetGroupBy gets a reference to the given []string and assigns it to the GroupBy field.
func (o *GetDeliveryOverviewLinkedAudienceSuccessBetaInput) SetGroupBy(v []string) {
	o.GroupBy = v
}

// GetGranularity returns the Granularity field value
func (o *GetDeliveryOverviewLinkedAudienceSuccessBetaInput) GetGranularity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Granularity
}

// GetGranularityOk returns a tuple with the Granularity field value
// and a boolean to check if the value has been set.
func (o *GetDeliveryOverviewLinkedAudienceSuccessBetaInput) GetGranularityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Granularity, true
}

// SetGranularity sets field value
func (o *GetDeliveryOverviewLinkedAudienceSuccessBetaInput) SetGranularity(v string) {
	o.Granularity = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *GetDeliveryOverviewLinkedAudienceSuccessBetaInput) GetFilter() DeliveryOverviewAudienceFilterBy {
	if o == nil || IsNil(o.Filter) {
		var ret DeliveryOverviewAudienceFilterBy
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeliveryOverviewLinkedAudienceSuccessBetaInput) GetFilterOk() (*DeliveryOverviewAudienceFilterBy, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *GetDeliveryOverviewLinkedAudienceSuccessBetaInput) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given DeliveryOverviewAudienceFilterBy and assigns it to the Filter field.
func (o *GetDeliveryOverviewLinkedAudienceSuccessBetaInput) SetFilter(
	v DeliveryOverviewAudienceFilterBy,
) {
	o.Filter = &v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *GetDeliveryOverviewLinkedAudienceSuccessBetaInput) GetPagination() PaginationInput {
	if o == nil || IsNil(o.Pagination) {
		var ret PaginationInput
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeliveryOverviewLinkedAudienceSuccessBetaInput) GetPaginationOk() (*PaginationInput, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *GetDeliveryOverviewLinkedAudienceSuccessBetaInput) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given PaginationInput and assigns it to the Pagination field.
func (o *GetDeliveryOverviewLinkedAudienceSuccessBetaInput) SetPagination(v PaginationInput) {
	o.Pagination = &v
}

func (o GetDeliveryOverviewLinkedAudienceSuccessBetaInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDeliveryOverviewLinkedAudienceSuccessBetaInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sourceId"] = o.SourceId
	toSerialize["destinationConfigId"] = o.DestinationConfigId
	toSerialize["startTime"] = o.StartTime
	toSerialize["endTime"] = o.EndTime
	if !IsNil(o.GroupBy) {
		toSerialize["groupBy"] = o.GroupBy
	}
	toSerialize["granularity"] = o.Granularity
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableGetDeliveryOverviewLinkedAudienceSuccessBetaInput struct {
	value *GetDeliveryOverviewLinkedAudienceSuccessBetaInput
	isSet bool
}

func (v NullableGetDeliveryOverviewLinkedAudienceSuccessBetaInput) Get() *GetDeliveryOverviewLinkedAudienceSuccessBetaInput {
	return v.value
}

func (v *NullableGetDeliveryOverviewLinkedAudienceSuccessBetaInput) Set(
	val *GetDeliveryOverviewLinkedAudienceSuccessBetaInput,
) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeliveryOverviewLinkedAudienceSuccessBetaInput) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeliveryOverviewLinkedAudienceSuccessBetaInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeliveryOverviewLinkedAudienceSuccessBetaInput(
	val *GetDeliveryOverviewLinkedAudienceSuccessBetaInput,
) *NullableGetDeliveryOverviewLinkedAudienceSuccessBetaInput {
	return &NullableGetDeliveryOverviewLinkedAudienceSuccessBetaInput{value: val, isSet: true}
}

func (v NullableGetDeliveryOverviewLinkedAudienceSuccessBetaInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeliveryOverviewLinkedAudienceSuccessBetaInput) UnmarshalJSON(
	src []byte,
) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
