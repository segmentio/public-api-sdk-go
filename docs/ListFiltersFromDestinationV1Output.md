# ListFiltersFromDestinationV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | [**[]DestinationFilterV1**](DestinationFilterV1.md) | A list of the filters that belong to the specified Destination instance. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewListFiltersFromDestinationV1Output

`func NewListFiltersFromDestinationV1Output(filters []DestinationFilterV1, pagination PaginationOutput, ) *ListFiltersFromDestinationV1Output`

NewListFiltersFromDestinationV1Output instantiates a new ListFiltersFromDestinationV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFiltersFromDestinationV1OutputWithDefaults

`func NewListFiltersFromDestinationV1OutputWithDefaults() *ListFiltersFromDestinationV1Output`

NewListFiltersFromDestinationV1OutputWithDefaults instantiates a new ListFiltersFromDestinationV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *ListFiltersFromDestinationV1Output) GetFilters() []DestinationFilterV1`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ListFiltersFromDestinationV1Output) GetFiltersOk() (*[]DestinationFilterV1, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ListFiltersFromDestinationV1Output) SetFilters(v []DestinationFilterV1)`

SetFilters sets Filters field to given value.


### GetPagination

`func (o *ListFiltersFromDestinationV1Output) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListFiltersFromDestinationV1Output) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListFiltersFromDestinationV1Output) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


