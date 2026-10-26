# ListSchemaSettingsInSourceV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceId** | **string** | Source id.  Config API note: analogous to &#x60;parent&#x60; and &#x60;name&#x60;. | 
**Settings** | [**SourceSettingsOutputV1**](SourceSettingsOutputV1.md) |  | 

## Methods

### NewListSchemaSettingsInSourceV1Output

`func NewListSchemaSettingsInSourceV1Output(sourceId string, settings SourceSettingsOutputV1, ) *ListSchemaSettingsInSourceV1Output`

NewListSchemaSettingsInSourceV1Output instantiates a new ListSchemaSettingsInSourceV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSchemaSettingsInSourceV1OutputWithDefaults

`func NewListSchemaSettingsInSourceV1OutputWithDefaults() *ListSchemaSettingsInSourceV1Output`

NewListSchemaSettingsInSourceV1OutputWithDefaults instantiates a new ListSchemaSettingsInSourceV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceId

`func (o *ListSchemaSettingsInSourceV1Output) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ListSchemaSettingsInSourceV1Output) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ListSchemaSettingsInSourceV1Output) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSettings

`func (o *ListSchemaSettingsInSourceV1Output) GetSettings() SourceSettingsOutputV1`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *ListSchemaSettingsInSourceV1Output) GetSettingsOk() (*SourceSettingsOutputV1, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *ListSchemaSettingsInSourceV1Output) SetSettings(v SourceSettingsOutputV1)`

SetSettings sets Settings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


