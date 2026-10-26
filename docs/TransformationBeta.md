# TransformationBeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the Transformation. | 
**Name** | **string** | The name of the Transformation. | 
**SourceId** | **string** | The Source associated with the Transformation. | 
**DestinationMetadataId** | Pointer to **string** | The optional Destination metadata associated with the Transformation. | [optional] 
**Enabled** | **bool** | If the Transformation is enabled. | 
**If** | **string** | If statement ([FQL](https://segment.com/docs/config-api/fql/)) to match events.  For standard event matchers, use the following:  Track -\\&gt; \&quot;event&#x3D;&#39;\\&lt;eventName\\&gt;&#39;\&quot;  Identify -\\&gt; \&quot;type&#x3D;&#39;identify&#39;\&quot;  Group -\\&gt; \&quot;type&#x3D;&#39;group&#39;\&quot; | 
**NewEventName** | Pointer to **string** | Optional new event name for renaming events. Works only for &#39;track&#39; event type. | [optional] 
**PropertyRenames** | Pointer to [**[]PropertyRenameBeta**](PropertyRenameBeta.md) | Optional array for renaming properties collected by your events. | [optional] 
**PropertyValueTransformations** | Pointer to [**[]PropertyValueTransformationBeta**](PropertyValueTransformationBeta.md) | Optional array for transforming properties and values collected by your events. Limited to 10 properties. | [optional] 

## Methods

### NewTransformationBeta

`func NewTransformationBeta(id string, name string, sourceId string, enabled bool, if_ string, ) *TransformationBeta`

NewTransformationBeta instantiates a new TransformationBeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransformationBetaWithDefaults

`func NewTransformationBetaWithDefaults() *TransformationBeta`

NewTransformationBetaWithDefaults instantiates a new TransformationBeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TransformationBeta) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransformationBeta) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransformationBeta) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TransformationBeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TransformationBeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TransformationBeta) SetName(v string)`

SetName sets Name field to given value.


### GetSourceId

`func (o *TransformationBeta) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *TransformationBeta) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *TransformationBeta) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetDestinationMetadataId

`func (o *TransformationBeta) GetDestinationMetadataId() string`

GetDestinationMetadataId returns the DestinationMetadataId field if non-nil, zero value otherwise.

### GetDestinationMetadataIdOk

`func (o *TransformationBeta) GetDestinationMetadataIdOk() (*string, bool)`

GetDestinationMetadataIdOk returns a tuple with the DestinationMetadataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationMetadataId

`func (o *TransformationBeta) SetDestinationMetadataId(v string)`

SetDestinationMetadataId sets DestinationMetadataId field to given value.

### HasDestinationMetadataId

`func (o *TransformationBeta) HasDestinationMetadataId() bool`

HasDestinationMetadataId returns a boolean if a field has been set.

### GetEnabled

`func (o *TransformationBeta) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TransformationBeta) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TransformationBeta) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIf

`func (o *TransformationBeta) GetIf() string`

GetIf returns the If field if non-nil, zero value otherwise.

### GetIfOk

`func (o *TransformationBeta) GetIfOk() (*string, bool)`

GetIfOk returns a tuple with the If field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIf

`func (o *TransformationBeta) SetIf(v string)`

SetIf sets If field to given value.


### GetNewEventName

`func (o *TransformationBeta) GetNewEventName() string`

GetNewEventName returns the NewEventName field if non-nil, zero value otherwise.

### GetNewEventNameOk

`func (o *TransformationBeta) GetNewEventNameOk() (*string, bool)`

GetNewEventNameOk returns a tuple with the NewEventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewEventName

`func (o *TransformationBeta) SetNewEventName(v string)`

SetNewEventName sets NewEventName field to given value.

### HasNewEventName

`func (o *TransformationBeta) HasNewEventName() bool`

HasNewEventName returns a boolean if a field has been set.

### GetPropertyRenames

`func (o *TransformationBeta) GetPropertyRenames() []PropertyRenameBeta`

GetPropertyRenames returns the PropertyRenames field if non-nil, zero value otherwise.

### GetPropertyRenamesOk

`func (o *TransformationBeta) GetPropertyRenamesOk() (*[]PropertyRenameBeta, bool)`

GetPropertyRenamesOk returns a tuple with the PropertyRenames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyRenames

`func (o *TransformationBeta) SetPropertyRenames(v []PropertyRenameBeta)`

SetPropertyRenames sets PropertyRenames field to given value.

### HasPropertyRenames

`func (o *TransformationBeta) HasPropertyRenames() bool`

HasPropertyRenames returns a boolean if a field has been set.

### GetPropertyValueTransformations

`func (o *TransformationBeta) GetPropertyValueTransformations() []PropertyValueTransformationBeta`

GetPropertyValueTransformations returns the PropertyValueTransformations field if non-nil, zero value otherwise.

### GetPropertyValueTransformationsOk

`func (o *TransformationBeta) GetPropertyValueTransformationsOk() (*[]PropertyValueTransformationBeta, bool)`

GetPropertyValueTransformationsOk returns a tuple with the PropertyValueTransformations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyValueTransformations

`func (o *TransformationBeta) SetPropertyValueTransformations(v []PropertyValueTransformationBeta)`

SetPropertyValueTransformations sets PropertyValueTransformations field to given value.

### HasPropertyValueTransformations

`func (o *TransformationBeta) HasPropertyValueTransformations() bool`

HasPropertyValueTransformations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


