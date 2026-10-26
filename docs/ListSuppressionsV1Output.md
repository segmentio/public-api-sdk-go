# ListSuppressionsV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Suppressed** | [**[]SuppressedInner**](SuppressedInner.md) | An array that lists the suppressions from the Workspace.  Config API note: equal to &#x60;attributes&#x60;. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewListSuppressionsV1Output

`func NewListSuppressionsV1Output(suppressed []SuppressedInner, pagination PaginationOutput, ) *ListSuppressionsV1Output`

NewListSuppressionsV1Output instantiates a new ListSuppressionsV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSuppressionsV1OutputWithDefaults

`func NewListSuppressionsV1OutputWithDefaults() *ListSuppressionsV1Output`

NewListSuppressionsV1OutputWithDefaults instantiates a new ListSuppressionsV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuppressed

`func (o *ListSuppressionsV1Output) GetSuppressed() []SuppressedInner`

GetSuppressed returns the Suppressed field if non-nil, zero value otherwise.

### GetSuppressedOk

`func (o *ListSuppressionsV1Output) GetSuppressedOk() (*[]SuppressedInner, bool)`

GetSuppressedOk returns a tuple with the Suppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressed

`func (o *ListSuppressionsV1Output) SetSuppressed(v []SuppressedInner)`

SetSuppressed sets Suppressed field to given value.


### GetPagination

`func (o *ListSuppressionsV1Output) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListSuppressionsV1Output) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListSuppressionsV1Output) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


