# ListCredentialsV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | [**[]CredentialV1**](CredentialV1.md) | A list of Credentials that belong to the Workspace. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewListCredentialsV1Output

`func NewListCredentialsV1Output(credentials []CredentialV1, pagination PaginationOutput, ) *ListCredentialsV1Output`

NewListCredentialsV1Output instantiates a new ListCredentialsV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCredentialsV1OutputWithDefaults

`func NewListCredentialsV1OutputWithDefaults() *ListCredentialsV1Output`

NewListCredentialsV1OutputWithDefaults instantiates a new ListCredentialsV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *ListCredentialsV1Output) GetCredentials() []CredentialV1`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ListCredentialsV1Output) GetCredentialsOk() (*[]CredentialV1, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ListCredentialsV1Output) SetCredentials(v []CredentialV1)`

SetCredentials sets Credentials field to given value.


### GetPagination

`func (o *ListCredentialsV1Output) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListCredentialsV1Output) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListCredentialsV1Output) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


