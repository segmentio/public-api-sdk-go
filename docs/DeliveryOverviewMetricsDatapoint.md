# DeliveryOverviewMetricsDatapoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | **string** | The timestamp corresponding to the beginning of the window given by the requested granularity. | 
**Count** | **float32** | Holds the number of events within the specified granularity and group By options. | 
**RetryCount** | Pointer to **float32** | The number of retried events that were successfully delivered. | [optional] 

## Methods

### NewDeliveryOverviewMetricsDatapoint

`func NewDeliveryOverviewMetricsDatapoint(time string, count float32, ) *DeliveryOverviewMetricsDatapoint`

NewDeliveryOverviewMetricsDatapoint instantiates a new DeliveryOverviewMetricsDatapoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryOverviewMetricsDatapointWithDefaults

`func NewDeliveryOverviewMetricsDatapointWithDefaults() *DeliveryOverviewMetricsDatapoint`

NewDeliveryOverviewMetricsDatapointWithDefaults instantiates a new DeliveryOverviewMetricsDatapoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *DeliveryOverviewMetricsDatapoint) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *DeliveryOverviewMetricsDatapoint) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *DeliveryOverviewMetricsDatapoint) SetTime(v string)`

SetTime sets Time field to given value.


### GetCount

`func (o *DeliveryOverviewMetricsDatapoint) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DeliveryOverviewMetricsDatapoint) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DeliveryOverviewMetricsDatapoint) SetCount(v float32)`

SetCount sets Count field to given value.


### GetRetryCount

`func (o *DeliveryOverviewMetricsDatapoint) GetRetryCount() float32`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *DeliveryOverviewMetricsDatapoint) GetRetryCountOk() (*float32, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *DeliveryOverviewMetricsDatapoint) SetRetryCount(v float32)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *DeliveryOverviewMetricsDatapoint) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


