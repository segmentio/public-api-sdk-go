# GetDestinationsCatalogV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationsCatalog** | [**[]DestinationMetadataV1**](DestinationMetadataV1.md) | All Destination catalog items contained within the requested page. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewGetDestinationsCatalogV1Output

`func NewGetDestinationsCatalogV1Output(destinationsCatalog []DestinationMetadataV1, pagination PaginationOutput, ) *GetDestinationsCatalogV1Output`

NewGetDestinationsCatalogV1Output instantiates a new GetDestinationsCatalogV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDestinationsCatalogV1OutputWithDefaults

`func NewGetDestinationsCatalogV1OutputWithDefaults() *GetDestinationsCatalogV1Output`

NewGetDestinationsCatalogV1OutputWithDefaults instantiates a new GetDestinationsCatalogV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationsCatalog

`func (o *GetDestinationsCatalogV1Output) GetDestinationsCatalog() []DestinationMetadataV1`

GetDestinationsCatalog returns the DestinationsCatalog field if non-nil, zero value otherwise.

### GetDestinationsCatalogOk

`func (o *GetDestinationsCatalogV1Output) GetDestinationsCatalogOk() (*[]DestinationMetadataV1, bool)`

GetDestinationsCatalogOk returns a tuple with the DestinationsCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationsCatalog

`func (o *GetDestinationsCatalogV1Output) SetDestinationsCatalog(v []DestinationMetadataV1)`

SetDestinationsCatalog sets DestinationsCatalog field to given value.


### GetPagination

`func (o *GetDestinationsCatalogV1Output) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetDestinationsCatalogV1Output) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetDestinationsCatalogV1Output) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


