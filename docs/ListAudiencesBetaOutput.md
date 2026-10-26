# ListAudiencesBetaOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audiences** | [**[]AudienceSummaryWithAudienceTypeAndLookback**](AudienceSummaryWithAudienceTypeAndLookback.md) | A list of audience summary results. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewListAudiencesBetaOutput

`func NewListAudiencesBetaOutput(audiences []AudienceSummaryWithAudienceTypeAndLookback, pagination PaginationOutput, ) *ListAudiencesBetaOutput`

NewListAudiencesBetaOutput instantiates a new ListAudiencesBetaOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAudiencesBetaOutputWithDefaults

`func NewListAudiencesBetaOutputWithDefaults() *ListAudiencesBetaOutput`

NewListAudiencesBetaOutputWithDefaults instantiates a new ListAudiencesBetaOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudiences

`func (o *ListAudiencesBetaOutput) GetAudiences() []AudienceSummaryWithAudienceTypeAndLookback`

GetAudiences returns the Audiences field if non-nil, zero value otherwise.

### GetAudiencesOk

`func (o *ListAudiencesBetaOutput) GetAudiencesOk() (*[]AudienceSummaryWithAudienceTypeAndLookback, bool)`

GetAudiencesOk returns a tuple with the Audiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudiences

`func (o *ListAudiencesBetaOutput) SetAudiences(v []AudienceSummaryWithAudienceTypeAndLookback)`

SetAudiences sets Audiences field to given value.


### GetPagination

`func (o *ListAudiencesBetaOutput) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListAudiencesBetaOutput) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListAudiencesBetaOutput) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


