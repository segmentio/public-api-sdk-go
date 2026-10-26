# CreateDbtModelSyncTriggerInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceId** | **string** | The Source id to trigger a dbt model sync. | 

## Methods

### NewCreateDbtModelSyncTriggerInput

`func NewCreateDbtModelSyncTriggerInput(sourceId string, ) *CreateDbtModelSyncTriggerInput`

NewCreateDbtModelSyncTriggerInput instantiates a new CreateDbtModelSyncTriggerInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDbtModelSyncTriggerInputWithDefaults

`func NewCreateDbtModelSyncTriggerInputWithDefaults() *CreateDbtModelSyncTriggerInput`

NewCreateDbtModelSyncTriggerInputWithDefaults instantiates a new CreateDbtModelSyncTriggerInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceId

`func (o *CreateDbtModelSyncTriggerInput) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *CreateDbtModelSyncTriggerInput) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *CreateDbtModelSyncTriggerInput) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


