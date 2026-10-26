# DeliveryOverviewMetricsDataset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventName** | Pointer to **string** | The name of the event if group By[] included &#39;event Name&#39; in the request. | [optional] 
**AppVersion** | Pointer to **string** | The version of the app if group By[] included &#39;app Version&#39; in the request. | [optional] 
**EventType** | Pointer to **string** | The event type if group By[] included &#39;event Type&#39; in the request. | [optional] 
**DiscardReason** | Pointer to **string** | The discard reason for dropped events if group By[] included &#39;discard Reason&#39; in the request. | [optional] 
**Total** | **float32** | Holds the count of all event counts over the time frame of the series. | 
**Series** | [**[]DeliveryOverviewMetricsDatapoint**](DeliveryOverviewMetricsDatapoint.md) | A list of the event counts broken down by the requested granularity, time frame, and group By options. | 
**TotalRetryCount** | Pointer to **float32** | The number of events successfully delivered upon retry. | [optional] 

## Methods

### NewDeliveryOverviewMetricsDataset

`func NewDeliveryOverviewMetricsDataset(total float32, series []DeliveryOverviewMetricsDatapoint, ) *DeliveryOverviewMetricsDataset`

NewDeliveryOverviewMetricsDataset instantiates a new DeliveryOverviewMetricsDataset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryOverviewMetricsDatasetWithDefaults

`func NewDeliveryOverviewMetricsDatasetWithDefaults() *DeliveryOverviewMetricsDataset`

NewDeliveryOverviewMetricsDatasetWithDefaults instantiates a new DeliveryOverviewMetricsDataset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventName

`func (o *DeliveryOverviewMetricsDataset) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *DeliveryOverviewMetricsDataset) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *DeliveryOverviewMetricsDataset) SetEventName(v string)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *DeliveryOverviewMetricsDataset) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### GetAppVersion

`func (o *DeliveryOverviewMetricsDataset) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *DeliveryOverviewMetricsDataset) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *DeliveryOverviewMetricsDataset) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *DeliveryOverviewMetricsDataset) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.

### GetEventType

`func (o *DeliveryOverviewMetricsDataset) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *DeliveryOverviewMetricsDataset) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *DeliveryOverviewMetricsDataset) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *DeliveryOverviewMetricsDataset) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetDiscardReason

`func (o *DeliveryOverviewMetricsDataset) GetDiscardReason() string`

GetDiscardReason returns the DiscardReason field if non-nil, zero value otherwise.

### GetDiscardReasonOk

`func (o *DeliveryOverviewMetricsDataset) GetDiscardReasonOk() (*string, bool)`

GetDiscardReasonOk returns a tuple with the DiscardReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscardReason

`func (o *DeliveryOverviewMetricsDataset) SetDiscardReason(v string)`

SetDiscardReason sets DiscardReason field to given value.

### HasDiscardReason

`func (o *DeliveryOverviewMetricsDataset) HasDiscardReason() bool`

HasDiscardReason returns a boolean if a field has been set.

### GetTotal

`func (o *DeliveryOverviewMetricsDataset) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *DeliveryOverviewMetricsDataset) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *DeliveryOverviewMetricsDataset) SetTotal(v float32)`

SetTotal sets Total field to given value.


### GetSeries

`func (o *DeliveryOverviewMetricsDataset) GetSeries() []DeliveryOverviewMetricsDatapoint`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *DeliveryOverviewMetricsDataset) GetSeriesOk() (*[]DeliveryOverviewMetricsDatapoint, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *DeliveryOverviewMetricsDataset) SetSeries(v []DeliveryOverviewMetricsDatapoint)`

SetSeries sets Series field to given value.


### GetTotalRetryCount

`func (o *DeliveryOverviewMetricsDataset) GetTotalRetryCount() float32`

GetTotalRetryCount returns the TotalRetryCount field if non-nil, zero value otherwise.

### GetTotalRetryCountOk

`func (o *DeliveryOverviewMetricsDataset) GetTotalRetryCountOk() (*float32, bool)`

GetTotalRetryCountOk returns a tuple with the TotalRetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRetryCount

`func (o *DeliveryOverviewMetricsDataset) SetTotalRetryCount(v float32)`

SetTotalRetryCount sets TotalRetryCount field to given value.

### HasTotalRetryCount

`func (o *DeliveryOverviewMetricsDataset) HasTotalRetryCount() bool`

HasTotalRetryCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


