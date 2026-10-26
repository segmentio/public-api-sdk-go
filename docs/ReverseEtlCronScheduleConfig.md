# ReverseEtlCronScheduleConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | **string** | 5 field cron string expression. The cron expression must be larger than 15 minutes. | 
**Timezone** | **string** | Timezone respected by the cron string. Format must be in IANA Timezone Identifier. Example: &#39;America/Los_Angeles&#39;. | 

## Methods

### NewReverseEtlCronScheduleConfig

`func NewReverseEtlCronScheduleConfig(spec string, timezone string, ) *ReverseEtlCronScheduleConfig`

NewReverseEtlCronScheduleConfig instantiates a new ReverseEtlCronScheduleConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReverseEtlCronScheduleConfigWithDefaults

`func NewReverseEtlCronScheduleConfigWithDefaults() *ReverseEtlCronScheduleConfig`

NewReverseEtlCronScheduleConfigWithDefaults instantiates a new ReverseEtlCronScheduleConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *ReverseEtlCronScheduleConfig) GetSpec() string`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ReverseEtlCronScheduleConfig) GetSpecOk() (*string, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ReverseEtlCronScheduleConfig) SetSpec(v string)`

SetSpec sets Spec field to given value.


### GetTimezone

`func (o *ReverseEtlCronScheduleConfig) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ReverseEtlCronScheduleConfig) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ReverseEtlCronScheduleConfig) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


