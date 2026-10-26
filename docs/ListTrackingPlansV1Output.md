# ListTrackingPlansV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrackingPlans** | [**[]TrackingPlanV1**](TrackingPlanV1.md) | A paginated list of Tracking Plans. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewListTrackingPlansV1Output

`func NewListTrackingPlansV1Output(trackingPlans []TrackingPlanV1, pagination PaginationOutput, ) *ListTrackingPlansV1Output`

NewListTrackingPlansV1Output instantiates a new ListTrackingPlansV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTrackingPlansV1OutputWithDefaults

`func NewListTrackingPlansV1OutputWithDefaults() *ListTrackingPlansV1Output`

NewListTrackingPlansV1OutputWithDefaults instantiates a new ListTrackingPlansV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrackingPlans

`func (o *ListTrackingPlansV1Output) GetTrackingPlans() []TrackingPlanV1`

GetTrackingPlans returns the TrackingPlans field if non-nil, zero value otherwise.

### GetTrackingPlansOk

`func (o *ListTrackingPlansV1Output) GetTrackingPlansOk() (*[]TrackingPlanV1, bool)`

GetTrackingPlansOk returns a tuple with the TrackingPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingPlans

`func (o *ListTrackingPlansV1Output) SetTrackingPlans(v []TrackingPlanV1)`

SetTrackingPlans sets TrackingPlans field to given value.


### GetPagination

`func (o *ListTrackingPlansV1Output) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListTrackingPlansV1Output) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListTrackingPlansV1Output) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


