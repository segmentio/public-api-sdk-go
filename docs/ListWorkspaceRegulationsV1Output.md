# ListWorkspaceRegulationsV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Regulations** | [**[]RegulationListEntryV1**](RegulationListEntryV1.md) | List of Workspace-scoped regulations with statuses. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewListWorkspaceRegulationsV1Output

`func NewListWorkspaceRegulationsV1Output(regulations []RegulationListEntryV1, pagination PaginationOutput, ) *ListWorkspaceRegulationsV1Output`

NewListWorkspaceRegulationsV1Output instantiates a new ListWorkspaceRegulationsV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWorkspaceRegulationsV1OutputWithDefaults

`func NewListWorkspaceRegulationsV1OutputWithDefaults() *ListWorkspaceRegulationsV1Output`

NewListWorkspaceRegulationsV1OutputWithDefaults instantiates a new ListWorkspaceRegulationsV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegulations

`func (o *ListWorkspaceRegulationsV1Output) GetRegulations() []RegulationListEntryV1`

GetRegulations returns the Regulations field if non-nil, zero value otherwise.

### GetRegulationsOk

`func (o *ListWorkspaceRegulationsV1Output) GetRegulationsOk() (*[]RegulationListEntryV1, bool)`

GetRegulationsOk returns a tuple with the Regulations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulations

`func (o *ListWorkspaceRegulationsV1Output) SetRegulations(v []RegulationListEntryV1)`

SetRegulations sets Regulations field to given value.


### GetPagination

`func (o *ListWorkspaceRegulationsV1Output) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListWorkspaceRegulationsV1Output) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListWorkspaceRegulationsV1Output) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


