# MetricBeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricName** | **string** | The name of the metric. | 
**Total** | **float32** | Number of occurrences of the metric. | 
**Breakdown** | Pointer to [**[]BreakdownBeta**](BreakdownBeta.md) | Breakdown of metrics within a metric. | [optional] 

## Methods

### NewMetricBeta

`func NewMetricBeta(metricName string, total float32, ) *MetricBeta`

NewMetricBeta instantiates a new MetricBeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricBetaWithDefaults

`func NewMetricBetaWithDefaults() *MetricBeta`

NewMetricBetaWithDefaults instantiates a new MetricBeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricName

`func (o *MetricBeta) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *MetricBeta) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *MetricBeta) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.


### GetTotal

`func (o *MetricBeta) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MetricBeta) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MetricBeta) SetTotal(v float32)`

SetTotal sets Total field to given value.


### GetBreakdown

`func (o *MetricBeta) GetBreakdown() []BreakdownBeta`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *MetricBeta) GetBreakdownOk() (*[]BreakdownBeta, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *MetricBeta) SetBreakdown(v []BreakdownBeta)`

SetBreakdown sets Breakdown field to given value.

### HasBreakdown

`func (o *MetricBeta) HasBreakdown() bool`

HasBreakdown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


