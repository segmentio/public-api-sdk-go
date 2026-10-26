# BatchQueryActivationsForSpaceAlphaInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudienceIds** | Pointer to **[]string** | Audience IDs to list activations for. Max 100 IDs per call. Duplicates are de-duplicated server-side but still count toward the cap.  Omit (or pass an empty array) to scan every audience in the space instead - see the &#x60;pagination&#x60; doc below for how that changes what a \&quot;page\&quot; means. Omitting this list is materially higher latency per call than a scoped batch, since each page has to enumerate audiences before it can look up their activations. | [optional] 
**Pagination** | Pointer to [**PaginationInput**](PaginationInput.md) |  | [optional] 

## Methods

### NewBatchQueryActivationsForSpaceAlphaInput

`func NewBatchQueryActivationsForSpaceAlphaInput() *BatchQueryActivationsForSpaceAlphaInput`

NewBatchQueryActivationsForSpaceAlphaInput instantiates a new BatchQueryActivationsForSpaceAlphaInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchQueryActivationsForSpaceAlphaInputWithDefaults

`func NewBatchQueryActivationsForSpaceAlphaInputWithDefaults() *BatchQueryActivationsForSpaceAlphaInput`

NewBatchQueryActivationsForSpaceAlphaInputWithDefaults instantiates a new BatchQueryActivationsForSpaceAlphaInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudienceIds

`func (o *BatchQueryActivationsForSpaceAlphaInput) GetAudienceIds() []string`

GetAudienceIds returns the AudienceIds field if non-nil, zero value otherwise.

### GetAudienceIdsOk

`func (o *BatchQueryActivationsForSpaceAlphaInput) GetAudienceIdsOk() (*[]string, bool)`

GetAudienceIdsOk returns a tuple with the AudienceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceIds

`func (o *BatchQueryActivationsForSpaceAlphaInput) SetAudienceIds(v []string)`

SetAudienceIds sets AudienceIds field to given value.

### HasAudienceIds

`func (o *BatchQueryActivationsForSpaceAlphaInput) HasAudienceIds() bool`

HasAudienceIds returns a boolean if a field has been set.

### GetPagination

`func (o *BatchQueryActivationsForSpaceAlphaInput) GetPagination() PaginationInput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *BatchQueryActivationsForSpaceAlphaInput) GetPaginationOk() (*PaginationInput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *BatchQueryActivationsForSpaceAlphaInput) SetPagination(v PaginationInput)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *BatchQueryActivationsForSpaceAlphaInput) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


