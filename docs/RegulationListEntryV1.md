# RegulationListEntryV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the regulate request. | 
**SubjectType** | **string** | The subject type. | 
**Subjects** | **[]string** | The list of &#x60;userId&#x60; or &#x60;objectId&#x60; values of the subjects to regulate. | 
**Status** | **string** | The current status of the regulate request. | 
**CreatedAt** | **string** | The timestamp of the creation of the request. | 
**RegulationType** | **string** | The regulation type. | 
**FinishedAt** | Pointer to **string** | The timestamp of when the request finished. | [optional] 

## Methods

### NewRegulationListEntryV1

`func NewRegulationListEntryV1(id string, subjectType string, subjects []string, status string, createdAt string, regulationType string, ) *RegulationListEntryV1`

NewRegulationListEntryV1 instantiates a new RegulationListEntryV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegulationListEntryV1WithDefaults

`func NewRegulationListEntryV1WithDefaults() *RegulationListEntryV1`

NewRegulationListEntryV1WithDefaults instantiates a new RegulationListEntryV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RegulationListEntryV1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegulationListEntryV1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegulationListEntryV1) SetId(v string)`

SetId sets Id field to given value.


### GetSubjectType

`func (o *RegulationListEntryV1) GetSubjectType() string`

GetSubjectType returns the SubjectType field if non-nil, zero value otherwise.

### GetSubjectTypeOk

`func (o *RegulationListEntryV1) GetSubjectTypeOk() (*string, bool)`

GetSubjectTypeOk returns a tuple with the SubjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectType

`func (o *RegulationListEntryV1) SetSubjectType(v string)`

SetSubjectType sets SubjectType field to given value.


### GetSubjects

`func (o *RegulationListEntryV1) GetSubjects() []string`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *RegulationListEntryV1) GetSubjectsOk() (*[]string, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *RegulationListEntryV1) SetSubjects(v []string)`

SetSubjects sets Subjects field to given value.


### GetStatus

`func (o *RegulationListEntryV1) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RegulationListEntryV1) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RegulationListEntryV1) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *RegulationListEntryV1) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RegulationListEntryV1) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RegulationListEntryV1) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetRegulationType

`func (o *RegulationListEntryV1) GetRegulationType() string`

GetRegulationType returns the RegulationType field if non-nil, zero value otherwise.

### GetRegulationTypeOk

`func (o *RegulationListEntryV1) GetRegulationTypeOk() (*string, bool)`

GetRegulationTypeOk returns a tuple with the RegulationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulationType

`func (o *RegulationListEntryV1) SetRegulationType(v string)`

SetRegulationType sets RegulationType field to given value.


### GetFinishedAt

`func (o *RegulationListEntryV1) GetFinishedAt() string`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *RegulationListEntryV1) GetFinishedAtOk() (*string, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *RegulationListEntryV1) SetFinishedAt(v string)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *RegulationListEntryV1) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


