# ListRulesFromTrackingPlanV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | [**[]RuleV1**](RuleV1.md) | Rules associated with the given Tracking Plan. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewListRulesFromTrackingPlanV1Output

`func NewListRulesFromTrackingPlanV1Output(rules []RuleV1, pagination PaginationOutput, ) *ListRulesFromTrackingPlanV1Output`

NewListRulesFromTrackingPlanV1Output instantiates a new ListRulesFromTrackingPlanV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRulesFromTrackingPlanV1OutputWithDefaults

`func NewListRulesFromTrackingPlanV1OutputWithDefaults() *ListRulesFromTrackingPlanV1Output`

NewListRulesFromTrackingPlanV1OutputWithDefaults instantiates a new ListRulesFromTrackingPlanV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *ListRulesFromTrackingPlanV1Output) GetRules() []RuleV1`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *ListRulesFromTrackingPlanV1Output) GetRulesOk() (*[]RuleV1, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *ListRulesFromTrackingPlanV1Output) SetRules(v []RuleV1)`

SetRules sets Rules field to given value.


### GetPagination

`func (o *ListRulesFromTrackingPlanV1Output) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListRulesFromTrackingPlanV1Output) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListRulesFromTrackingPlanV1Output) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


