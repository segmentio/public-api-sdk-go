# GetDeliveryOverviewMetricsBetaOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **float32** | The total number of events for the returned dataset. | 
**Dataset** | [**[]DeliveryOverviewMetricsDataset**](DeliveryOverviewMetricsDataset.md) | Represents the list of series broken down by the dimensions and time frame requested. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewGetDeliveryOverviewMetricsBetaOutput

`func NewGetDeliveryOverviewMetricsBetaOutput(total float32, dataset []DeliveryOverviewMetricsDataset, pagination PaginationOutput, ) *GetDeliveryOverviewMetricsBetaOutput`

NewGetDeliveryOverviewMetricsBetaOutput instantiates a new GetDeliveryOverviewMetricsBetaOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeliveryOverviewMetricsBetaOutputWithDefaults

`func NewGetDeliveryOverviewMetricsBetaOutputWithDefaults() *GetDeliveryOverviewMetricsBetaOutput`

NewGetDeliveryOverviewMetricsBetaOutputWithDefaults instantiates a new GetDeliveryOverviewMetricsBetaOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *GetDeliveryOverviewMetricsBetaOutput) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetDeliveryOverviewMetricsBetaOutput) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetDeliveryOverviewMetricsBetaOutput) SetTotal(v float32)`

SetTotal sets Total field to given value.


### GetDataset

`func (o *GetDeliveryOverviewMetricsBetaOutput) GetDataset() []DeliveryOverviewMetricsDataset`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *GetDeliveryOverviewMetricsBetaOutput) GetDatasetOk() (*[]DeliveryOverviewMetricsDataset, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *GetDeliveryOverviewMetricsBetaOutput) SetDataset(v []DeliveryOverviewMetricsDataset)`

SetDataset sets Dataset field to given value.


### GetPagination

`func (o *GetDeliveryOverviewMetricsBetaOutput) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetDeliveryOverviewMetricsBetaOutput) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetDeliveryOverviewMetricsBetaOutput) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


