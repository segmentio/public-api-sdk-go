# ListDestinationsFromAudienceAlphaOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connections** | [**[]SimpleDestination**](SimpleDestination.md) | A list of connection results. | 
**Pagination** | Pointer to [**PaginationOutput**](PaginationOutput.md) |  | [optional] 

## Methods

### NewListDestinationsFromAudienceAlphaOutput

`func NewListDestinationsFromAudienceAlphaOutput(connections []SimpleDestination, ) *ListDestinationsFromAudienceAlphaOutput`

NewListDestinationsFromAudienceAlphaOutput instantiates a new ListDestinationsFromAudienceAlphaOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDestinationsFromAudienceAlphaOutputWithDefaults

`func NewListDestinationsFromAudienceAlphaOutputWithDefaults() *ListDestinationsFromAudienceAlphaOutput`

NewListDestinationsFromAudienceAlphaOutputWithDefaults instantiates a new ListDestinationsFromAudienceAlphaOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnections

`func (o *ListDestinationsFromAudienceAlphaOutput) GetConnections() []SimpleDestination`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *ListDestinationsFromAudienceAlphaOutput) GetConnectionsOk() (*[]SimpleDestination, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *ListDestinationsFromAudienceAlphaOutput) SetConnections(v []SimpleDestination)`

SetConnections sets Connections field to given value.


### GetPagination

`func (o *ListDestinationsFromAudienceAlphaOutput) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListDestinationsFromAudienceAlphaOutput) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListDestinationsFromAudienceAlphaOutput) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListDestinationsFromAudienceAlphaOutput) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


