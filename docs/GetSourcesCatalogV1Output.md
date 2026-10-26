# GetSourcesCatalogV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourcesCatalog** | [**[]SourceMetadataV1**](SourceMetadataV1.md) | All Source catalog items contained within the requested page. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewGetSourcesCatalogV1Output

`func NewGetSourcesCatalogV1Output(sourcesCatalog []SourceMetadataV1, pagination PaginationOutput, ) *GetSourcesCatalogV1Output`

NewGetSourcesCatalogV1Output instantiates a new GetSourcesCatalogV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSourcesCatalogV1OutputWithDefaults

`func NewGetSourcesCatalogV1OutputWithDefaults() *GetSourcesCatalogV1Output`

NewGetSourcesCatalogV1OutputWithDefaults instantiates a new GetSourcesCatalogV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourcesCatalog

`func (o *GetSourcesCatalogV1Output) GetSourcesCatalog() []SourceMetadataV1`

GetSourcesCatalog returns the SourcesCatalog field if non-nil, zero value otherwise.

### GetSourcesCatalogOk

`func (o *GetSourcesCatalogV1Output) GetSourcesCatalogOk() (*[]SourceMetadataV1, bool)`

GetSourcesCatalogOk returns a tuple with the SourcesCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcesCatalog

`func (o *GetSourcesCatalogV1Output) SetSourcesCatalog(v []SourceMetadataV1)`

SetSourcesCatalog sets SourcesCatalog field to given value.


### GetPagination

`func (o *GetSourcesCatalogV1Output) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetSourcesCatalogV1Output) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetSourcesCatalogV1Output) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


