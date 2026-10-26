# ListConnectedDestinationsFromSourceAlphaOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destinations** | [**[]DestinationV1**](DestinationV1.md) | A list that contains the Destinations connected to the Source. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewListConnectedDestinationsFromSourceAlphaOutput

`func NewListConnectedDestinationsFromSourceAlphaOutput(destinations []DestinationV1, pagination PaginationOutput, ) *ListConnectedDestinationsFromSourceAlphaOutput`

NewListConnectedDestinationsFromSourceAlphaOutput instantiates a new ListConnectedDestinationsFromSourceAlphaOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConnectedDestinationsFromSourceAlphaOutputWithDefaults

`func NewListConnectedDestinationsFromSourceAlphaOutputWithDefaults() *ListConnectedDestinationsFromSourceAlphaOutput`

NewListConnectedDestinationsFromSourceAlphaOutputWithDefaults instantiates a new ListConnectedDestinationsFromSourceAlphaOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinations

`func (o *ListConnectedDestinationsFromSourceAlphaOutput) GetDestinations() []DestinationV1`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *ListConnectedDestinationsFromSourceAlphaOutput) GetDestinationsOk() (*[]DestinationV1, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *ListConnectedDestinationsFromSourceAlphaOutput) SetDestinations(v []DestinationV1)`

SetDestinations sets Destinations field to given value.


### GetPagination

`func (o *ListConnectedDestinationsFromSourceAlphaOutput) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListConnectedDestinationsFromSourceAlphaOutput) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListConnectedDestinationsFromSourceAlphaOutput) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


