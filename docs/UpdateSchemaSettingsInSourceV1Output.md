# UpdateSchemaSettingsInSourceV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceId** | **string** | The id of the updated Source.  Config API note: analogous to &#x60;parent&#x60; and &#x60;name&#x60;. | 
**Settings** | [**SourceSettingsOutputV1**](SourceSettingsOutputV1.md) |  | 

## Methods

### NewUpdateSchemaSettingsInSourceV1Output

`func NewUpdateSchemaSettingsInSourceV1Output(sourceId string, settings SourceSettingsOutputV1, ) *UpdateSchemaSettingsInSourceV1Output`

NewUpdateSchemaSettingsInSourceV1Output instantiates a new UpdateSchemaSettingsInSourceV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSchemaSettingsInSourceV1OutputWithDefaults

`func NewUpdateSchemaSettingsInSourceV1OutputWithDefaults() *UpdateSchemaSettingsInSourceV1Output`

NewUpdateSchemaSettingsInSourceV1OutputWithDefaults instantiates a new UpdateSchemaSettingsInSourceV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceId

`func (o *UpdateSchemaSettingsInSourceV1Output) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *UpdateSchemaSettingsInSourceV1Output) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *UpdateSchemaSettingsInSourceV1Output) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSettings

`func (o *UpdateSchemaSettingsInSourceV1Output) GetSettings() SourceSettingsOutputV1`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *UpdateSchemaSettingsInSourceV1Output) GetSettingsOk() (*SourceSettingsOutputV1, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *UpdateSchemaSettingsInSourceV1Output) SetSettings(v SourceSettingsOutputV1)`

SetSettings sets Settings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


