# DeliveryMetricsSummaryBeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceId** | **string** | The Source id.  Config API note: analogous to &#x60;parent&#x60;. | 
**DestinationMetadataId** | **string** | The Destination metadata id. | 
**Metrics** | [**[]MetricBeta**](MetricBeta.md) | The summary of event delivery metrics for the requested Destination. | 

## Methods

### NewDeliveryMetricsSummaryBeta

`func NewDeliveryMetricsSummaryBeta(sourceId string, destinationMetadataId string, metrics []MetricBeta, ) *DeliveryMetricsSummaryBeta`

NewDeliveryMetricsSummaryBeta instantiates a new DeliveryMetricsSummaryBeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryMetricsSummaryBetaWithDefaults

`func NewDeliveryMetricsSummaryBetaWithDefaults() *DeliveryMetricsSummaryBeta`

NewDeliveryMetricsSummaryBetaWithDefaults instantiates a new DeliveryMetricsSummaryBeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceId

`func (o *DeliveryMetricsSummaryBeta) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *DeliveryMetricsSummaryBeta) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *DeliveryMetricsSummaryBeta) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetDestinationMetadataId

`func (o *DeliveryMetricsSummaryBeta) GetDestinationMetadataId() string`

GetDestinationMetadataId returns the DestinationMetadataId field if non-nil, zero value otherwise.

### GetDestinationMetadataIdOk

`func (o *DeliveryMetricsSummaryBeta) GetDestinationMetadataIdOk() (*string, bool)`

GetDestinationMetadataIdOk returns a tuple with the DestinationMetadataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationMetadataId

`func (o *DeliveryMetricsSummaryBeta) SetDestinationMetadataId(v string)`

SetDestinationMetadataId sets DestinationMetadataId field to given value.


### GetMetrics

`func (o *DeliveryMetricsSummaryBeta) GetMetrics() []MetricBeta`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *DeliveryMetricsSummaryBeta) GetMetricsOk() (*[]MetricBeta, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *DeliveryMetricsSummaryBeta) SetMetrics(v []MetricBeta)`

SetMetrics sets Metrics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


