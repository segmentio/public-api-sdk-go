# UpdateAudienceScheduleForAudienceInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Strategy** | **string** | Strategy of the audience schedule (periodic or specific days). | 
**Config** | [**NullableConfig**](Config.md) |  | 

## Methods

### NewUpdateAudienceScheduleForAudienceInput

`func NewUpdateAudienceScheduleForAudienceInput(strategy string, config NullableConfig, ) *UpdateAudienceScheduleForAudienceInput`

NewUpdateAudienceScheduleForAudienceInput instantiates a new UpdateAudienceScheduleForAudienceInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAudienceScheduleForAudienceInputWithDefaults

`func NewUpdateAudienceScheduleForAudienceInputWithDefaults() *UpdateAudienceScheduleForAudienceInput`

NewUpdateAudienceScheduleForAudienceInputWithDefaults instantiates a new UpdateAudienceScheduleForAudienceInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStrategy

`func (o *UpdateAudienceScheduleForAudienceInput) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *UpdateAudienceScheduleForAudienceInput) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *UpdateAudienceScheduleForAudienceInput) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.


### GetConfig

`func (o *UpdateAudienceScheduleForAudienceInput) GetConfig() Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateAudienceScheduleForAudienceInput) GetConfigOk() (*Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateAudienceScheduleForAudienceInput) SetConfig(v Config)`

SetConfig sets Config field to given value.


### SetConfigNil

`func (o *UpdateAudienceScheduleForAudienceInput) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *UpdateAudienceScheduleForAudienceInput) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


