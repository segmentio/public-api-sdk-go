# Config1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interval** | **string** | Duration is specified as a string, EG: 15m, 3h25m30s. | 
**Days** | **[]float32** | Days of the week. | 
**Hours** | **[]float32** | Hours of the day. | 
**Timezone** | **string** | Timezone respected by the cron string. Format must be in IANA Timezone Identifier. Example: &#39;America/Los_Angeles&#39;. | 
**Spec** | **string** | 5 field cron string expression. The cron expression must be larger than 15 minutes. | 
**JobId** | **string** | The dbt Cloud job used to trigger a sync for a Reverse ETL Connection. | 
**AccountId** | **string** | The dbt Cloud account where the job belongs to. | 

## Methods

### NewConfig1

`func NewConfig1(interval string, days []float32, hours []float32, timezone string, spec string, jobId string, accountId string, ) *Config1`

NewConfig1 instantiates a new Config1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfig1WithDefaults

`func NewConfig1WithDefaults() *Config1`

NewConfig1WithDefaults instantiates a new Config1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterval

`func (o *Config1) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *Config1) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *Config1) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetDays

`func (o *Config1) GetDays() []float32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *Config1) GetDaysOk() (*[]float32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *Config1) SetDays(v []float32)`

SetDays sets Days field to given value.


### GetHours

`func (o *Config1) GetHours() []float32`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *Config1) GetHoursOk() (*[]float32, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *Config1) SetHours(v []float32)`

SetHours sets Hours field to given value.


### GetTimezone

`func (o *Config1) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *Config1) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *Config1) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetSpec

`func (o *Config1) GetSpec() string`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *Config1) GetSpecOk() (*string, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *Config1) SetSpec(v string)`

SetSpec sets Spec field to given value.


### GetJobId

`func (o *Config1) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *Config1) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *Config1) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetAccountId

`func (o *Config1) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Config1) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Config1) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


