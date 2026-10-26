# ListAudiencesAlphaOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audiences** | [**[]AudienceSummaryWithAudienceTypeAndLookback**](AudienceSummaryWithAudienceTypeAndLookback.md) | A list of audience summary results. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewListAudiencesAlphaOutput

`func NewListAudiencesAlphaOutput(audiences []AudienceSummaryWithAudienceTypeAndLookback, pagination PaginationOutput, ) *ListAudiencesAlphaOutput`

NewListAudiencesAlphaOutput instantiates a new ListAudiencesAlphaOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAudiencesAlphaOutputWithDefaults

`func NewListAudiencesAlphaOutputWithDefaults() *ListAudiencesAlphaOutput`

NewListAudiencesAlphaOutputWithDefaults instantiates a new ListAudiencesAlphaOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudiences

`func (o *ListAudiencesAlphaOutput) GetAudiences() []AudienceSummaryWithAudienceTypeAndLookback`

GetAudiences returns the Audiences field if non-nil, zero value otherwise.

### GetAudiencesOk

`func (o *ListAudiencesAlphaOutput) GetAudiencesOk() (*[]AudienceSummaryWithAudienceTypeAndLookback, bool)`

GetAudiencesOk returns a tuple with the Audiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudiences

`func (o *ListAudiencesAlphaOutput) SetAudiences(v []AudienceSummaryWithAudienceTypeAndLookback)`

SetAudiences sets Audiences field to given value.


### GetPagination

`func (o *ListAudiencesAlphaOutput) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListAudiencesAlphaOutput) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListAudiencesAlphaOutput) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


