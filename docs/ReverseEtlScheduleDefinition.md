# ReverseEtlScheduleDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Strategy** | **string** | Strategy supports: Periodic, Specific Days, Manual, CRON and DBT_CLOUD. | 
**Config** | Pointer to [**NullableConfig1**](Config1.md) |  | [optional] 

## Methods

### NewReverseEtlScheduleDefinition

`func NewReverseEtlScheduleDefinition(strategy string, ) *ReverseEtlScheduleDefinition`

NewReverseEtlScheduleDefinition instantiates a new ReverseEtlScheduleDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReverseEtlScheduleDefinitionWithDefaults

`func NewReverseEtlScheduleDefinitionWithDefaults() *ReverseEtlScheduleDefinition`

NewReverseEtlScheduleDefinitionWithDefaults instantiates a new ReverseEtlScheduleDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStrategy

`func (o *ReverseEtlScheduleDefinition) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *ReverseEtlScheduleDefinition) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *ReverseEtlScheduleDefinition) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.


### GetConfig

`func (o *ReverseEtlScheduleDefinition) GetConfig() Config1`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ReverseEtlScheduleDefinition) GetConfigOk() (*Config1, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ReverseEtlScheduleDefinition) SetConfig(v Config1)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ReverseEtlScheduleDefinition) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *ReverseEtlScheduleDefinition) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *ReverseEtlScheduleDefinition) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


