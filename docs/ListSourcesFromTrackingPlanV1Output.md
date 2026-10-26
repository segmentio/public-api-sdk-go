# ListSourcesFromTrackingPlanV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sources** | [**[]SourceV1**](SourceV1.md) | A paginated list of Sources associated with the Tracking Plan. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewListSourcesFromTrackingPlanV1Output

`func NewListSourcesFromTrackingPlanV1Output(sources []SourceV1, pagination PaginationOutput, ) *ListSourcesFromTrackingPlanV1Output`

NewListSourcesFromTrackingPlanV1Output instantiates a new ListSourcesFromTrackingPlanV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSourcesFromTrackingPlanV1OutputWithDefaults

`func NewListSourcesFromTrackingPlanV1OutputWithDefaults() *ListSourcesFromTrackingPlanV1Output`

NewListSourcesFromTrackingPlanV1OutputWithDefaults instantiates a new ListSourcesFromTrackingPlanV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSources

`func (o *ListSourcesFromTrackingPlanV1Output) GetSources() []SourceV1`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *ListSourcesFromTrackingPlanV1Output) GetSourcesOk() (*[]SourceV1, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *ListSourcesFromTrackingPlanV1Output) SetSources(v []SourceV1)`

SetSources sets Sources field to given value.


### GetPagination

`func (o *ListSourcesFromTrackingPlanV1Output) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListSourcesFromTrackingPlanV1Output) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListSourcesFromTrackingPlanV1Output) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


