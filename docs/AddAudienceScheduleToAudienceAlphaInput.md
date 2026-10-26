# AddAudienceScheduleToAudienceAlphaInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Strategy** | **string** | Strategy of the audience schedule (periodic or specific days). | 
**Config** | [**NullableConfig**](Config.md) |  | 

## Methods

### NewAddAudienceScheduleToAudienceAlphaInput

`func NewAddAudienceScheduleToAudienceAlphaInput(strategy string, config NullableConfig, ) *AddAudienceScheduleToAudienceAlphaInput`

NewAddAudienceScheduleToAudienceAlphaInput instantiates a new AddAudienceScheduleToAudienceAlphaInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAudienceScheduleToAudienceAlphaInputWithDefaults

`func NewAddAudienceScheduleToAudienceAlphaInputWithDefaults() *AddAudienceScheduleToAudienceAlphaInput`

NewAddAudienceScheduleToAudienceAlphaInputWithDefaults instantiates a new AddAudienceScheduleToAudienceAlphaInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStrategy

`func (o *AddAudienceScheduleToAudienceAlphaInput) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *AddAudienceScheduleToAudienceAlphaInput) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *AddAudienceScheduleToAudienceAlphaInput) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.


### GetConfig

`func (o *AddAudienceScheduleToAudienceAlphaInput) GetConfig() Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddAudienceScheduleToAudienceAlphaInput) GetConfigOk() (*Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddAudienceScheduleToAudienceAlphaInput) SetConfig(v Config)`

SetConfig sets Config field to given value.


### SetConfigNil

`func (o *AddAudienceScheduleToAudienceAlphaInput) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *AddAudienceScheduleToAudienceAlphaInput) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


