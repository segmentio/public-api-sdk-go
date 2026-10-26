# ListFunctionsV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Functions** | [**[]ListFunctionItemV1**](ListFunctionItemV1.md) | An array of Functions. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewListFunctionsV1Output

`func NewListFunctionsV1Output(functions []ListFunctionItemV1, pagination PaginationOutput, ) *ListFunctionsV1Output`

NewListFunctionsV1Output instantiates a new ListFunctionsV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFunctionsV1OutputWithDefaults

`func NewListFunctionsV1OutputWithDefaults() *ListFunctionsV1Output`

NewListFunctionsV1OutputWithDefaults instantiates a new ListFunctionsV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctions

`func (o *ListFunctionsV1Output) GetFunctions() []ListFunctionItemV1`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *ListFunctionsV1Output) GetFunctionsOk() (*[]ListFunctionItemV1, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *ListFunctionsV1Output) SetFunctions(v []ListFunctionItemV1)`

SetFunctions sets Functions field to given value.


### GetPagination

`func (o *ListFunctionsV1Output) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListFunctionsV1Output) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListFunctionsV1Output) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


