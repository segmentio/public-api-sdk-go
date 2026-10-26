# CreateCloudSourceRegulationV1Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegulationType** | **string** | The regulation type to create. | 
**SubjectType** | **string** | The subject type. Must be &#x60;objectId&#x60; for Cloud Sources. | 
**SubjectIds** | **[]string** | The list of &#x60;userId&#x60; or &#x60;objectId&#x60; values of the subjects to regulate.  Config API note: equal to &#x60;parent&#x60; but allows an array. | 
**Collection** | **string** | The Cloud Source collection to regulate. | 

## Methods

### NewCreateCloudSourceRegulationV1Input

`func NewCreateCloudSourceRegulationV1Input(regulationType string, subjectType string, subjectIds []string, collection string, ) *CreateCloudSourceRegulationV1Input`

NewCreateCloudSourceRegulationV1Input instantiates a new CreateCloudSourceRegulationV1Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCloudSourceRegulationV1InputWithDefaults

`func NewCreateCloudSourceRegulationV1InputWithDefaults() *CreateCloudSourceRegulationV1Input`

NewCreateCloudSourceRegulationV1InputWithDefaults instantiates a new CreateCloudSourceRegulationV1Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegulationType

`func (o *CreateCloudSourceRegulationV1Input) GetRegulationType() string`

GetRegulationType returns the RegulationType field if non-nil, zero value otherwise.

### GetRegulationTypeOk

`func (o *CreateCloudSourceRegulationV1Input) GetRegulationTypeOk() (*string, bool)`

GetRegulationTypeOk returns a tuple with the RegulationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulationType

`func (o *CreateCloudSourceRegulationV1Input) SetRegulationType(v string)`

SetRegulationType sets RegulationType field to given value.


### GetSubjectType

`func (o *CreateCloudSourceRegulationV1Input) GetSubjectType() string`

GetSubjectType returns the SubjectType field if non-nil, zero value otherwise.

### GetSubjectTypeOk

`func (o *CreateCloudSourceRegulationV1Input) GetSubjectTypeOk() (*string, bool)`

GetSubjectTypeOk returns a tuple with the SubjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectType

`func (o *CreateCloudSourceRegulationV1Input) SetSubjectType(v string)`

SetSubjectType sets SubjectType field to given value.


### GetSubjectIds

`func (o *CreateCloudSourceRegulationV1Input) GetSubjectIds() []string`

GetSubjectIds returns the SubjectIds field if non-nil, zero value otherwise.

### GetSubjectIdsOk

`func (o *CreateCloudSourceRegulationV1Input) GetSubjectIdsOk() (*[]string, bool)`

GetSubjectIdsOk returns a tuple with the SubjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectIds

`func (o *CreateCloudSourceRegulationV1Input) SetSubjectIds(v []string)`

SetSubjectIds sets SubjectIds field to given value.


### GetCollection

`func (o *CreateCloudSourceRegulationV1Input) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *CreateCloudSourceRegulationV1Input) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *CreateCloudSourceRegulationV1Input) SetCollection(v string)`

SetCollection sets Collection field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


