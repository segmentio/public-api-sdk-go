# CreateDestinationV1Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceId** | **string** | The id of the Source to connect to this Destination instance.  Config API note: analogous to &#x60;parent&#x60;. | 
**MetadataId** | **string** | The id of the metadata to link to the new Destination. | 
**Enabled** | Pointer to **bool** | Whether this Destination should receive data. | [optional] 
**Name** | Pointer to **string** | Defines the display name of the Destination.  Config API note: equal to &#x60;displayName&#x60;. | [optional] 
**Settings** | **map[string]interface{}** | An object that contains settings for the Destination based on the \&quot;required\&quot; and \&quot;advanced\&quot; settings present in the Destination metadata.  Config API note: equal to &#x60;config&#x60;. | 

## Methods

### NewCreateDestinationV1Input

`func NewCreateDestinationV1Input(sourceId string, metadataId string, settings map[string]interface{}, ) *CreateDestinationV1Input`

NewCreateDestinationV1Input instantiates a new CreateDestinationV1Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDestinationV1InputWithDefaults

`func NewCreateDestinationV1InputWithDefaults() *CreateDestinationV1Input`

NewCreateDestinationV1InputWithDefaults instantiates a new CreateDestinationV1Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceId

`func (o *CreateDestinationV1Input) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *CreateDestinationV1Input) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *CreateDestinationV1Input) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetMetadataId

`func (o *CreateDestinationV1Input) GetMetadataId() string`

GetMetadataId returns the MetadataId field if non-nil, zero value otherwise.

### GetMetadataIdOk

`func (o *CreateDestinationV1Input) GetMetadataIdOk() (*string, bool)`

GetMetadataIdOk returns a tuple with the MetadataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataId

`func (o *CreateDestinationV1Input) SetMetadataId(v string)`

SetMetadataId sets MetadataId field to given value.


### GetEnabled

`func (o *CreateDestinationV1Input) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateDestinationV1Input) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateDestinationV1Input) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateDestinationV1Input) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *CreateDestinationV1Input) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDestinationV1Input) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDestinationV1Input) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateDestinationV1Input) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSettings

`func (o *CreateDestinationV1Input) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *CreateDestinationV1Input) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *CreateDestinationV1Input) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


