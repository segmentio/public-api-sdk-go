# CreateWorkspaceRegulationV1Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegulationType** | **string** | The regulation type to create. | 
**SubjectType** | **string** | The subject type. Use &#x60;objectId&#x60; for Cloud Source regulations. | 
**SubjectIds** | **[]string** | The list of &#x60;userId&#x60; or &#x60;objectId&#x60; values of the subjects to regulate.  Config API note: equal to &#x60;parent&#x60; but allows an array. | 

## Methods

### NewCreateWorkspaceRegulationV1Input

`func NewCreateWorkspaceRegulationV1Input(regulationType string, subjectType string, subjectIds []string, ) *CreateWorkspaceRegulationV1Input`

NewCreateWorkspaceRegulationV1Input instantiates a new CreateWorkspaceRegulationV1Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWorkspaceRegulationV1InputWithDefaults

`func NewCreateWorkspaceRegulationV1InputWithDefaults() *CreateWorkspaceRegulationV1Input`

NewCreateWorkspaceRegulationV1InputWithDefaults instantiates a new CreateWorkspaceRegulationV1Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegulationType

`func (o *CreateWorkspaceRegulationV1Input) GetRegulationType() string`

GetRegulationType returns the RegulationType field if non-nil, zero value otherwise.

### GetRegulationTypeOk

`func (o *CreateWorkspaceRegulationV1Input) GetRegulationTypeOk() (*string, bool)`

GetRegulationTypeOk returns a tuple with the RegulationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulationType

`func (o *CreateWorkspaceRegulationV1Input) SetRegulationType(v string)`

SetRegulationType sets RegulationType field to given value.


### GetSubjectType

`func (o *CreateWorkspaceRegulationV1Input) GetSubjectType() string`

GetSubjectType returns the SubjectType field if non-nil, zero value otherwise.

### GetSubjectTypeOk

`func (o *CreateWorkspaceRegulationV1Input) GetSubjectTypeOk() (*string, bool)`

GetSubjectTypeOk returns a tuple with the SubjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectType

`func (o *CreateWorkspaceRegulationV1Input) SetSubjectType(v string)`

SetSubjectType sets SubjectType field to given value.


### GetSubjectIds

`func (o *CreateWorkspaceRegulationV1Input) GetSubjectIds() []string`

GetSubjectIds returns the SubjectIds field if non-nil, zero value otherwise.

### GetSubjectIdsOk

`func (o *CreateWorkspaceRegulationV1Input) GetSubjectIdsOk() (*[]string, bool)`

GetSubjectIdsOk returns a tuple with the SubjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectIds

`func (o *CreateWorkspaceRegulationV1Input) SetSubjectIds(v []string)`

SetSubjectIds sets SubjectIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


