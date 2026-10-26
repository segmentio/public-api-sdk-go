# UpdateTransformationBetaInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransformationId** | **string** | ID of the Transformation to update. | 
**Name** | Pointer to **string** | The name of the Transformation. | [optional] 
**SourceId** | Pointer to **string** | The optional Source to be associated with the Transformation. | [optional] 
**DestinationMetadataId** | Pointer to **string** | The optional Destination metadata to be associated with the Transformation. | [optional] 
**Enabled** | Pointer to **bool** | If the Transformation should be enabled. | [optional] 
**If** | Pointer to **string** | If statement ([FQL](https://segment.com/docs/config-api/fql/)) to match events.  For standard event matchers, use the following:  Track -\\&gt; \&quot;event&#x3D;&#39;\\&lt;eventName\\&gt;&#39;\&quot;  Identify -\\&gt; \&quot;type&#x3D;&#39;identify&#39;\&quot;  Group -\\&gt; \&quot;type&#x3D;&#39;group&#39;\&quot; | [optional] 
**NewEventName** | Pointer to **string** | Optional new event name for renaming events. Works only for &#39;track&#39; event type. | [optional] 
**PropertyRenames** | Pointer to [**[]PropertyRenameBeta**](PropertyRenameBeta.md) | Optional array for renaming properties collected by your events. | [optional] 
**PropertyValueTransformations** | Pointer to [**[]PropertyValueTransformationBeta**](PropertyValueTransformationBeta.md) | Optional array for transforming properties and values collected by your events. Limited to 10 properties. | [optional] 

## Methods

### NewUpdateTransformationBetaInput

`func NewUpdateTransformationBetaInput(transformationId string, ) *UpdateTransformationBetaInput`

NewUpdateTransformationBetaInput instantiates a new UpdateTransformationBetaInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTransformationBetaInputWithDefaults

`func NewUpdateTransformationBetaInputWithDefaults() *UpdateTransformationBetaInput`

NewUpdateTransformationBetaInputWithDefaults instantiates a new UpdateTransformationBetaInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransformationId

`func (o *UpdateTransformationBetaInput) GetTransformationId() string`

GetTransformationId returns the TransformationId field if non-nil, zero value otherwise.

### GetTransformationIdOk

`func (o *UpdateTransformationBetaInput) GetTransformationIdOk() (*string, bool)`

GetTransformationIdOk returns a tuple with the TransformationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformationId

`func (o *UpdateTransformationBetaInput) SetTransformationId(v string)`

SetTransformationId sets TransformationId field to given value.


### GetName

`func (o *UpdateTransformationBetaInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTransformationBetaInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTransformationBetaInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateTransformationBetaInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSourceId

`func (o *UpdateTransformationBetaInput) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *UpdateTransformationBetaInput) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *UpdateTransformationBetaInput) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *UpdateTransformationBetaInput) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetDestinationMetadataId

`func (o *UpdateTransformationBetaInput) GetDestinationMetadataId() string`

GetDestinationMetadataId returns the DestinationMetadataId field if non-nil, zero value otherwise.

### GetDestinationMetadataIdOk

`func (o *UpdateTransformationBetaInput) GetDestinationMetadataIdOk() (*string, bool)`

GetDestinationMetadataIdOk returns a tuple with the DestinationMetadataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationMetadataId

`func (o *UpdateTransformationBetaInput) SetDestinationMetadataId(v string)`

SetDestinationMetadataId sets DestinationMetadataId field to given value.

### HasDestinationMetadataId

`func (o *UpdateTransformationBetaInput) HasDestinationMetadataId() bool`

HasDestinationMetadataId returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateTransformationBetaInput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateTransformationBetaInput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateTransformationBetaInput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateTransformationBetaInput) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIf

`func (o *UpdateTransformationBetaInput) GetIf() string`

GetIf returns the If field if non-nil, zero value otherwise.

### GetIfOk

`func (o *UpdateTransformationBetaInput) GetIfOk() (*string, bool)`

GetIfOk returns a tuple with the If field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIf

`func (o *UpdateTransformationBetaInput) SetIf(v string)`

SetIf sets If field to given value.

### HasIf

`func (o *UpdateTransformationBetaInput) HasIf() bool`

HasIf returns a boolean if a field has been set.

### GetNewEventName

`func (o *UpdateTransformationBetaInput) GetNewEventName() string`

GetNewEventName returns the NewEventName field if non-nil, zero value otherwise.

### GetNewEventNameOk

`func (o *UpdateTransformationBetaInput) GetNewEventNameOk() (*string, bool)`

GetNewEventNameOk returns a tuple with the NewEventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewEventName

`func (o *UpdateTransformationBetaInput) SetNewEventName(v string)`

SetNewEventName sets NewEventName field to given value.

### HasNewEventName

`func (o *UpdateTransformationBetaInput) HasNewEventName() bool`

HasNewEventName returns a boolean if a field has been set.

### GetPropertyRenames

`func (o *UpdateTransformationBetaInput) GetPropertyRenames() []PropertyRenameBeta`

GetPropertyRenames returns the PropertyRenames field if non-nil, zero value otherwise.

### GetPropertyRenamesOk

`func (o *UpdateTransformationBetaInput) GetPropertyRenamesOk() (*[]PropertyRenameBeta, bool)`

GetPropertyRenamesOk returns a tuple with the PropertyRenames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyRenames

`func (o *UpdateTransformationBetaInput) SetPropertyRenames(v []PropertyRenameBeta)`

SetPropertyRenames sets PropertyRenames field to given value.

### HasPropertyRenames

`func (o *UpdateTransformationBetaInput) HasPropertyRenames() bool`

HasPropertyRenames returns a boolean if a field has been set.

### GetPropertyValueTransformations

`func (o *UpdateTransformationBetaInput) GetPropertyValueTransformations() []PropertyValueTransformationBeta`

GetPropertyValueTransformations returns the PropertyValueTransformations field if non-nil, zero value otherwise.

### GetPropertyValueTransformationsOk

`func (o *UpdateTransformationBetaInput) GetPropertyValueTransformationsOk() (*[]PropertyValueTransformationBeta, bool)`

GetPropertyValueTransformationsOk returns a tuple with the PropertyValueTransformations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyValueTransformations

`func (o *UpdateTransformationBetaInput) SetPropertyValueTransformations(v []PropertyValueTransformationBeta)`

SetPropertyValueTransformations sets PropertyValueTransformations field to given value.

### HasPropertyValueTransformations

`func (o *UpdateTransformationBetaInput) HasPropertyValueTransformations() bool`

HasPropertyValueTransformations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


