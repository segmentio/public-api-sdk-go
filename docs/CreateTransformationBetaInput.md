# CreateTransformationBetaInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the Transformation. | 
**SourceId** | **string** | The Source to be associated with the Transformation. | 
**DestinationMetadataId** | Pointer to **string** | The optional Destination metadata id to be associated with the Transformation. | [optional] 
**Enabled** | **bool** | If the Transformation should be enabled. | 
**If** | **string** | If statement ([FQL](https://segment.com/docs/config-api/fql/)) to match events.  For standard event matchers, use the following:  Track -\\&gt; \&quot;event&#x3D;&#39;\\&lt;eventName\\&gt;&#39;\&quot;  Identify -\\&gt; \&quot;type&#x3D;&#39;identify&#39;\&quot;  Group -\\&gt; \&quot;type&#x3D;&#39;group&#39;\&quot; | 
**NewEventName** | Pointer to **string** | Optional new event name for renaming events. Works only for &#39;track&#39; event type. | [optional] 
**PropertyRenames** | Pointer to [**[]PropertyRenameBeta**](PropertyRenameBeta.md) | Optional array for renaming properties collected by your events. | [optional] 
**PropertyValueTransformations** | Pointer to [**[]PropertyValueTransformationBeta**](PropertyValueTransformationBeta.md) | Optional array for transforming properties and values collected by your events. Limited to 10 properties. | [optional] 

## Methods

### NewCreateTransformationBetaInput

`func NewCreateTransformationBetaInput(name string, sourceId string, enabled bool, if_ string, ) *CreateTransformationBetaInput`

NewCreateTransformationBetaInput instantiates a new CreateTransformationBetaInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTransformationBetaInputWithDefaults

`func NewCreateTransformationBetaInputWithDefaults() *CreateTransformationBetaInput`

NewCreateTransformationBetaInputWithDefaults instantiates a new CreateTransformationBetaInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateTransformationBetaInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTransformationBetaInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTransformationBetaInput) SetName(v string)`

SetName sets Name field to given value.


### GetSourceId

`func (o *CreateTransformationBetaInput) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *CreateTransformationBetaInput) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *CreateTransformationBetaInput) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetDestinationMetadataId

`func (o *CreateTransformationBetaInput) GetDestinationMetadataId() string`

GetDestinationMetadataId returns the DestinationMetadataId field if non-nil, zero value otherwise.

### GetDestinationMetadataIdOk

`func (o *CreateTransformationBetaInput) GetDestinationMetadataIdOk() (*string, bool)`

GetDestinationMetadataIdOk returns a tuple with the DestinationMetadataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationMetadataId

`func (o *CreateTransformationBetaInput) SetDestinationMetadataId(v string)`

SetDestinationMetadataId sets DestinationMetadataId field to given value.

### HasDestinationMetadataId

`func (o *CreateTransformationBetaInput) HasDestinationMetadataId() bool`

HasDestinationMetadataId returns a boolean if a field has been set.

### GetEnabled

`func (o *CreateTransformationBetaInput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateTransformationBetaInput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateTransformationBetaInput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIf

`func (o *CreateTransformationBetaInput) GetIf() string`

GetIf returns the If field if non-nil, zero value otherwise.

### GetIfOk

`func (o *CreateTransformationBetaInput) GetIfOk() (*string, bool)`

GetIfOk returns a tuple with the If field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIf

`func (o *CreateTransformationBetaInput) SetIf(v string)`

SetIf sets If field to given value.


### GetNewEventName

`func (o *CreateTransformationBetaInput) GetNewEventName() string`

GetNewEventName returns the NewEventName field if non-nil, zero value otherwise.

### GetNewEventNameOk

`func (o *CreateTransformationBetaInput) GetNewEventNameOk() (*string, bool)`

GetNewEventNameOk returns a tuple with the NewEventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewEventName

`func (o *CreateTransformationBetaInput) SetNewEventName(v string)`

SetNewEventName sets NewEventName field to given value.

### HasNewEventName

`func (o *CreateTransformationBetaInput) HasNewEventName() bool`

HasNewEventName returns a boolean if a field has been set.

### GetPropertyRenames

`func (o *CreateTransformationBetaInput) GetPropertyRenames() []PropertyRenameBeta`

GetPropertyRenames returns the PropertyRenames field if non-nil, zero value otherwise.

### GetPropertyRenamesOk

`func (o *CreateTransformationBetaInput) GetPropertyRenamesOk() (*[]PropertyRenameBeta, bool)`

GetPropertyRenamesOk returns a tuple with the PropertyRenames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyRenames

`func (o *CreateTransformationBetaInput) SetPropertyRenames(v []PropertyRenameBeta)`

SetPropertyRenames sets PropertyRenames field to given value.

### HasPropertyRenames

`func (o *CreateTransformationBetaInput) HasPropertyRenames() bool`

HasPropertyRenames returns a boolean if a field has been set.

### GetPropertyValueTransformations

`func (o *CreateTransformationBetaInput) GetPropertyValueTransformations() []PropertyValueTransformationBeta`

GetPropertyValueTransformations returns the PropertyValueTransformations field if non-nil, zero value otherwise.

### GetPropertyValueTransformationsOk

`func (o *CreateTransformationBetaInput) GetPropertyValueTransformationsOk() (*[]PropertyValueTransformationBeta, bool)`

GetPropertyValueTransformationsOk returns a tuple with the PropertyValueTransformations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyValueTransformations

`func (o *CreateTransformationBetaInput) SetPropertyValueTransformations(v []PropertyValueTransformationBeta)`

SetPropertyValueTransformations sets PropertyValueTransformations field to given value.

### HasPropertyValueTransformations

`func (o *CreateTransformationBetaInput) HasPropertyValueTransformations() bool`

HasPropertyValueTransformations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


