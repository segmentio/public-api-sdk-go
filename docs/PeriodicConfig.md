# PeriodicConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interval** | **string** | Go duration format string, only supporting units \&quot;d\&quot; (days), \&quot;h\&quot; (hours) and \&quot;m\&quot; (minutes).  Acceptable values are: - 15m - 30m - 1h - 2h - 4h - 6h - 8h - 12h - 1d | 

## Methods

### NewPeriodicConfig

`func NewPeriodicConfig(interval string, ) *PeriodicConfig`

NewPeriodicConfig instantiates a new PeriodicConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeriodicConfigWithDefaults

`func NewPeriodicConfigWithDefaults() *PeriodicConfig`

NewPeriodicConfigWithDefaults instantiates a new PeriodicConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterval

`func (o *PeriodicConfig) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *PeriodicConfig) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *PeriodicConfig) SetInterval(v string)`

SetInterval sets Interval field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


