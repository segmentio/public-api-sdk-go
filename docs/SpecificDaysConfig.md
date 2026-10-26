# SpecificDaysConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | **[]float32** | Days of week for schedule (0&#x3D;Sunday, 6&#x3D;Saturday). | 
**Hours** | **[]float32** | Hours of day for schedule (0-23). | 
**Timezone** | **string** | TZ database time zone identifier; for example, America/New_York. | 

## Methods

### NewSpecificDaysConfig

`func NewSpecificDaysConfig(days []float32, hours []float32, timezone string, ) *SpecificDaysConfig`

NewSpecificDaysConfig instantiates a new SpecificDaysConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpecificDaysConfigWithDefaults

`func NewSpecificDaysConfigWithDefaults() *SpecificDaysConfig`

NewSpecificDaysConfigWithDefaults instantiates a new SpecificDaysConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *SpecificDaysConfig) GetDays() []float32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *SpecificDaysConfig) GetDaysOk() (*[]float32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *SpecificDaysConfig) SetDays(v []float32)`

SetDays sets Days field to given value.


### GetHours

`func (o *SpecificDaysConfig) GetHours() []float32`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *SpecificDaysConfig) GetHoursOk() (*[]float32, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *SpecificDaysConfig) SetHours(v []float32)`

SetHours sets Hours field to given value.


### GetTimezone

`func (o *SpecificDaysConfig) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *SpecificDaysConfig) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *SpecificDaysConfig) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


