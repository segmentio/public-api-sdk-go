# ListDestinationsV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destinations** | [**[]DestinationV1**](DestinationV1.md) | The list that contains the Destinations connected to the Workspace. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewListDestinationsV1Output

`func NewListDestinationsV1Output(destinations []DestinationV1, pagination PaginationOutput, ) *ListDestinationsV1Output`

NewListDestinationsV1Output instantiates a new ListDestinationsV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDestinationsV1OutputWithDefaults

`func NewListDestinationsV1OutputWithDefaults() *ListDestinationsV1Output`

NewListDestinationsV1OutputWithDefaults instantiates a new ListDestinationsV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinations

`func (o *ListDestinationsV1Output) GetDestinations() []DestinationV1`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *ListDestinationsV1Output) GetDestinationsOk() (*[]DestinationV1, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *ListDestinationsV1Output) SetDestinations(v []DestinationV1)`

SetDestinations sets Destinations field to given value.


### GetPagination

`func (o *ListDestinationsV1Output) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListDestinationsV1Output) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListDestinationsV1Output) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


