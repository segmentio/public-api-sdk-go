# CreateSourceRegulationV1Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegulationType** | **string** | The regulation type to create. Please note that &#x60;DELETE_ARCHIVE_ONLY&#x60; is only supported for limited Workspaces for Source-scoped regulations. | 
**SubjectType** | **string** | The subject type. Note: &#x60;ANONYMOUS_ID&#x60; is only supported for limited Workspaces for Source-scoped regulations. &#x60;ANONYMOUS_ID&#x60; is only supported when regulationType is &#x60;DELETE_ARCHIVE_ONLY&#x60;. | 
**SubjectIds** | **[]string** | The list of &#x60;userId&#x60; or &#x60;objectId&#x60; or &#x60;anonymousId&#x60; values of the subjects to regulate.  Config API note: equal to &#x60;parent&#x60; but allows an array. | 

## Methods

### NewCreateSourceRegulationV1Input

`func NewCreateSourceRegulationV1Input(regulationType string, subjectType string, subjectIds []string, ) *CreateSourceRegulationV1Input`

NewCreateSourceRegulationV1Input instantiates a new CreateSourceRegulationV1Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSourceRegulationV1InputWithDefaults

`func NewCreateSourceRegulationV1InputWithDefaults() *CreateSourceRegulationV1Input`

NewCreateSourceRegulationV1InputWithDefaults instantiates a new CreateSourceRegulationV1Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegulationType

`func (o *CreateSourceRegulationV1Input) GetRegulationType() string`

GetRegulationType returns the RegulationType field if non-nil, zero value otherwise.

### GetRegulationTypeOk

`func (o *CreateSourceRegulationV1Input) GetRegulationTypeOk() (*string, bool)`

GetRegulationTypeOk returns a tuple with the RegulationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulationType

`func (o *CreateSourceRegulationV1Input) SetRegulationType(v string)`

SetRegulationType sets RegulationType field to given value.


### GetSubjectType

`func (o *CreateSourceRegulationV1Input) GetSubjectType() string`

GetSubjectType returns the SubjectType field if non-nil, zero value otherwise.

### GetSubjectTypeOk

`func (o *CreateSourceRegulationV1Input) GetSubjectTypeOk() (*string, bool)`

GetSubjectTypeOk returns a tuple with the SubjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectType

`func (o *CreateSourceRegulationV1Input) SetSubjectType(v string)`

SetSubjectType sets SubjectType field to given value.


### GetSubjectIds

`func (o *CreateSourceRegulationV1Input) GetSubjectIds() []string`

GetSubjectIds returns the SubjectIds field if non-nil, zero value otherwise.

### GetSubjectIdsOk

`func (o *CreateSourceRegulationV1Input) GetSubjectIdsOk() (*[]string, bool)`

GetSubjectIdsOk returns a tuple with the SubjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectIds

`func (o *CreateSourceRegulationV1Input) SetSubjectIds(v []string)`

SetSubjectIds sets SubjectIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


