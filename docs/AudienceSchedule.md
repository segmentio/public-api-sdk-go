# AudienceSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Distinct identifier for the schedule. | 
**Strategy** | **string** | Strategy of the audience schedule (periodic or specific days). | 
**Config** | [**NullableConfig**](Config.md) |  | 
**NextExecution** | Pointer to **string** | The next scheduled execution time (RFC3339). | [optional] 

## Methods

### NewAudienceSchedule

`func NewAudienceSchedule(id string, strategy string, config NullableConfig, ) *AudienceSchedule`

NewAudienceSchedule instantiates a new AudienceSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceScheduleWithDefaults

`func NewAudienceScheduleWithDefaults() *AudienceSchedule`

NewAudienceScheduleWithDefaults instantiates a new AudienceSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AudienceSchedule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AudienceSchedule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AudienceSchedule) SetId(v string)`

SetId sets Id field to given value.


### GetStrategy

`func (o *AudienceSchedule) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *AudienceSchedule) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *AudienceSchedule) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.


### GetConfig

`func (o *AudienceSchedule) GetConfig() Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AudienceSchedule) GetConfigOk() (*Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AudienceSchedule) SetConfig(v Config)`

SetConfig sets Config field to given value.


### SetConfigNil

`func (o *AudienceSchedule) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *AudienceSchedule) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetNextExecution

`func (o *AudienceSchedule) GetNextExecution() string`

GetNextExecution returns the NextExecution field if non-nil, zero value otherwise.

### GetNextExecutionOk

`func (o *AudienceSchedule) GetNextExecutionOk() (*string, bool)`

GetNextExecutionOk returns a tuple with the NextExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextExecution

`func (o *AudienceSchedule) SetNextExecution(v string)`

SetNextExecution sets NextExecution field to given value.

### HasNextExecution

`func (o *AudienceSchedule) HasNextExecution() bool`

HasNextExecution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


