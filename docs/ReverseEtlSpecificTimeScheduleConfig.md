# ReverseEtlSpecificTimeScheduleConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | **[]float32** | Days of the week. | 
**Hours** | **[]float32** | Hours of the day. | 
**Timezone** | **string** | Timezone for the specified times. | 

## Methods

### NewReverseEtlSpecificTimeScheduleConfig

`func NewReverseEtlSpecificTimeScheduleConfig(days []float32, hours []float32, timezone string, ) *ReverseEtlSpecificTimeScheduleConfig`

NewReverseEtlSpecificTimeScheduleConfig instantiates a new ReverseEtlSpecificTimeScheduleConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReverseEtlSpecificTimeScheduleConfigWithDefaults

`func NewReverseEtlSpecificTimeScheduleConfigWithDefaults() *ReverseEtlSpecificTimeScheduleConfig`

NewReverseEtlSpecificTimeScheduleConfigWithDefaults instantiates a new ReverseEtlSpecificTimeScheduleConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *ReverseEtlSpecificTimeScheduleConfig) GetDays() []float32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *ReverseEtlSpecificTimeScheduleConfig) GetDaysOk() (*[]float32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *ReverseEtlSpecificTimeScheduleConfig) SetDays(v []float32)`

SetDays sets Days field to given value.


### GetHours

`func (o *ReverseEtlSpecificTimeScheduleConfig) GetHours() []float32`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *ReverseEtlSpecificTimeScheduleConfig) GetHoursOk() (*[]float32, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *ReverseEtlSpecificTimeScheduleConfig) SetHours(v []float32)`

SetHours sets Hours field to given value.


### GetTimezone

`func (o *ReverseEtlSpecificTimeScheduleConfig) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ReverseEtlSpecificTimeScheduleConfig) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ReverseEtlSpecificTimeScheduleConfig) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


