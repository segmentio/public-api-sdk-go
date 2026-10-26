# GetEventsVolumeFromWorkspaceV1Query

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceId** | **string** | Workspace being requested. | 
**Granularity** | **string** | Granularity corresponds to the requested bucket granularity. | 
**StartTime** | **string** | StartTime is the ISO8601 formatted timestamp corresponding to the beginning of the requested time frame, inclusive. | 
**EndTime** | **string** | EndTime is the ISO8601 formatted timestamp corresponding to the end of the requested time frame, noninclusive. | 
**GroupBy** | Pointer to **[]string** | GroupBy is a comma-delimited list of strings representing the dimensions to group the result by. The current options are: &#x60;eventName&#x60; or &#x60;eventType&#x60;. | [optional] 
**SourceId** | Pointer to **[]string** | List of strings which allow you to restrict the result to just the given Sources. | [optional] 
**EventName** | Pointer to **[]string** | EventName is a list of strings which allow you to restrict the result to just the given event names. | [optional] 
**EventType** | Pointer to **[]string** | EventType is a list of strings which allow you to restrict the result to just the given event types. | [optional] 
**AppVersion** | Pointer to **[]string** | AppVersion is a list of strings which allow you to restrict the result to just the given application versions. | [optional] 
**Limit** | Pointer to **float32** | Limit is the total number of items in the result. | [optional] 

## Methods

### NewGetEventsVolumeFromWorkspaceV1Query

`func NewGetEventsVolumeFromWorkspaceV1Query(workspaceId string, granularity string, startTime string, endTime string, ) *GetEventsVolumeFromWorkspaceV1Query`

NewGetEventsVolumeFromWorkspaceV1Query instantiates a new GetEventsVolumeFromWorkspaceV1Query object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEventsVolumeFromWorkspaceV1QueryWithDefaults

`func NewGetEventsVolumeFromWorkspaceV1QueryWithDefaults() *GetEventsVolumeFromWorkspaceV1Query`

NewGetEventsVolumeFromWorkspaceV1QueryWithDefaults instantiates a new GetEventsVolumeFromWorkspaceV1Query object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceId

`func (o *GetEventsVolumeFromWorkspaceV1Query) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *GetEventsVolumeFromWorkspaceV1Query) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *GetEventsVolumeFromWorkspaceV1Query) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetGranularity

`func (o *GetEventsVolumeFromWorkspaceV1Query) GetGranularity() string`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *GetEventsVolumeFromWorkspaceV1Query) GetGranularityOk() (*string, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *GetEventsVolumeFromWorkspaceV1Query) SetGranularity(v string)`

SetGranularity sets Granularity field to given value.


### GetStartTime

`func (o *GetEventsVolumeFromWorkspaceV1Query) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *GetEventsVolumeFromWorkspaceV1Query) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *GetEventsVolumeFromWorkspaceV1Query) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *GetEventsVolumeFromWorkspaceV1Query) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *GetEventsVolumeFromWorkspaceV1Query) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *GetEventsVolumeFromWorkspaceV1Query) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.


### GetGroupBy

`func (o *GetEventsVolumeFromWorkspaceV1Query) GetGroupBy() []string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *GetEventsVolumeFromWorkspaceV1Query) GetGroupByOk() (*[]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *GetEventsVolumeFromWorkspaceV1Query) SetGroupBy(v []string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *GetEventsVolumeFromWorkspaceV1Query) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetSourceId

`func (o *GetEventsVolumeFromWorkspaceV1Query) GetSourceId() []string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *GetEventsVolumeFromWorkspaceV1Query) GetSourceIdOk() (*[]string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *GetEventsVolumeFromWorkspaceV1Query) SetSourceId(v []string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *GetEventsVolumeFromWorkspaceV1Query) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetEventName

`func (o *GetEventsVolumeFromWorkspaceV1Query) GetEventName() []string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *GetEventsVolumeFromWorkspaceV1Query) GetEventNameOk() (*[]string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *GetEventsVolumeFromWorkspaceV1Query) SetEventName(v []string)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *GetEventsVolumeFromWorkspaceV1Query) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### GetEventType

`func (o *GetEventsVolumeFromWorkspaceV1Query) GetEventType() []string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *GetEventsVolumeFromWorkspaceV1Query) GetEventTypeOk() (*[]string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *GetEventsVolumeFromWorkspaceV1Query) SetEventType(v []string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *GetEventsVolumeFromWorkspaceV1Query) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetAppVersion

`func (o *GetEventsVolumeFromWorkspaceV1Query) GetAppVersion() []string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *GetEventsVolumeFromWorkspaceV1Query) GetAppVersionOk() (*[]string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *GetEventsVolumeFromWorkspaceV1Query) SetAppVersion(v []string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *GetEventsVolumeFromWorkspaceV1Query) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.

### GetLimit

`func (o *GetEventsVolumeFromWorkspaceV1Query) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetEventsVolumeFromWorkspaceV1Query) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetEventsVolumeFromWorkspaceV1Query) SetLimit(v float32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *GetEventsVolumeFromWorkspaceV1Query) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


