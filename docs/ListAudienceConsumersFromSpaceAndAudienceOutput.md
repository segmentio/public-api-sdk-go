# ListAudienceConsumersFromSpaceAndAudienceOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audiences** | [**[]AudienceSummary**](AudienceSummary.md) | The list of audience consumers. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewListAudienceConsumersFromSpaceAndAudienceOutput

`func NewListAudienceConsumersFromSpaceAndAudienceOutput(audiences []AudienceSummary, pagination PaginationOutput, ) *ListAudienceConsumersFromSpaceAndAudienceOutput`

NewListAudienceConsumersFromSpaceAndAudienceOutput instantiates a new ListAudienceConsumersFromSpaceAndAudienceOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAudienceConsumersFromSpaceAndAudienceOutputWithDefaults

`func NewListAudienceConsumersFromSpaceAndAudienceOutputWithDefaults() *ListAudienceConsumersFromSpaceAndAudienceOutput`

NewListAudienceConsumersFromSpaceAndAudienceOutputWithDefaults instantiates a new ListAudienceConsumersFromSpaceAndAudienceOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudiences

`func (o *ListAudienceConsumersFromSpaceAndAudienceOutput) GetAudiences() []AudienceSummary`

GetAudiences returns the Audiences field if non-nil, zero value otherwise.

### GetAudiencesOk

`func (o *ListAudienceConsumersFromSpaceAndAudienceOutput) GetAudiencesOk() (*[]AudienceSummary, bool)`

GetAudiencesOk returns a tuple with the Audiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudiences

`func (o *ListAudienceConsumersFromSpaceAndAudienceOutput) SetAudiences(v []AudienceSummary)`

SetAudiences sets Audiences field to given value.


### GetPagination

`func (o *ListAudienceConsumersFromSpaceAndAudienceOutput) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListAudienceConsumersFromSpaceAndAudienceOutput) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListAudienceConsumersFromSpaceAndAudienceOutput) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


