# Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interval** | **string** | Go duration format string, only supporting units \&quot;d\&quot; (days), \&quot;h\&quot; (hours) and \&quot;m\&quot; (minutes).  Acceptable values are: - 15m - 30m - 1h - 2h - 4h - 6h - 8h - 12h - 1d | 
**Days** | **[]float32** | Days of week for schedule (0&#x3D;Sunday, 6&#x3D;Saturday). | 
**Hours** | **[]float32** | Hours of day for schedule (0-23). | 
**Timezone** | **string** | TZ database time zone identifier; for example, America/New_York. | 

## Methods

### NewConfig

`func NewConfig(interval string, days []float32, hours []float32, timezone string, ) *Config`

NewConfig instantiates a new Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigWithDefaults

`func NewConfigWithDefaults() *Config`

NewConfigWithDefaults instantiates a new Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterval

`func (o *Config) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *Config) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *Config) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetDays

`func (o *Config) GetDays() []float32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *Config) GetDaysOk() (*[]float32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *Config) SetDays(v []float32)`

SetDays sets Days field to given value.


### GetHours

`func (o *Config) GetHours() []float32`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *Config) GetHoursOk() (*[]float32, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *Config) SetHours(v []float32)`

SetHours sets Hours field to given value.


### GetTimezone

`func (o *Config) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *Config) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *Config) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


