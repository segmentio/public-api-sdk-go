# TransformationV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the Transformation. | 
**Name** | **string** | The name of the Transformation. | 
**SourceId** | **string** | The Source associated with the Transformation. | 
**DestinationMetadataId** | Pointer to **string** | The optional Destination metadata associated with the Transformation. | [optional] 
**Enabled** | **bool** | If the Transformation is enabled. | 
**If** | **string** | If statement ([FQL](https://segment.com/docs/config-api/fql/)) to match events.  For standard event matchers, use the following:  Track -\\&gt; \&quot;event&#x3D;&#39;\\&lt;eventName\\&gt;&#39;\&quot;  Identify -\\&gt; \&quot;type&#x3D;&#39;identify&#39;\&quot;  Group -\\&gt; \&quot;type&#x3D;&#39;group&#39;\&quot; | 
**Drop** | Pointer to **bool** | Optional boolean value if the Transformation should drop the event entirely when the if statement matches, ignores all other transforms. | [optional] 
**NewEventName** | Pointer to **string** | Optional new event name for renaming events. Works only for &#39;track&#39; event type. | [optional] 
**PropertyRenames** | Pointer to [**[]PropertyRenameV1**](PropertyRenameV1.md) | Optional array for renaming properties collected by your events. | [optional] 
**PropertyValueTransformations** | Pointer to [**[]PropertyValueTransformationV1**](PropertyValueTransformationV1.md) | Optional array for transforming properties and values collected by your events. Limited to 10 properties. | [optional] 
**FqlDefinedProperties** | Pointer to [**[]FQLDefinedPropertyV1**](FQLDefinedPropertyV1.md) | Optional array for defining new properties in FQL. Limited to 1 property right now. | [optional] 
**AllowProperties** | Pointer to **[]string** | Optional array for allowing properties from your events. | [optional] 
**HashPropertiesConfiguration** | Pointer to [**HashPropertiesConfiguration**](HashPropertiesConfiguration.md) |  | [optional] 

## Methods

### NewTransformationV1

`func NewTransformationV1(id string, name string, sourceId string, enabled bool, if_ string, ) *TransformationV1`

NewTransformationV1 instantiates a new TransformationV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransformationV1WithDefaults

`func NewTransformationV1WithDefaults() *TransformationV1`

NewTransformationV1WithDefaults instantiates a new TransformationV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TransformationV1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransformationV1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransformationV1) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TransformationV1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TransformationV1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TransformationV1) SetName(v string)`

SetName sets Name field to given value.


### GetSourceId

`func (o *TransformationV1) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *TransformationV1) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *TransformationV1) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetDestinationMetadataId

`func (o *TransformationV1) GetDestinationMetadataId() string`

GetDestinationMetadataId returns the DestinationMetadataId field if non-nil, zero value otherwise.

### GetDestinationMetadataIdOk

`func (o *TransformationV1) GetDestinationMetadataIdOk() (*string, bool)`

GetDestinationMetadataIdOk returns a tuple with the DestinationMetadataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationMetadataId

`func (o *TransformationV1) SetDestinationMetadataId(v string)`

SetDestinationMetadataId sets DestinationMetadataId field to given value.

### HasDestinationMetadataId

`func (o *TransformationV1) HasDestinationMetadataId() bool`

HasDestinationMetadataId returns a boolean if a field has been set.

### GetEnabled

`func (o *TransformationV1) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TransformationV1) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TransformationV1) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIf

`func (o *TransformationV1) GetIf() string`

GetIf returns the If field if non-nil, zero value otherwise.

### GetIfOk

`func (o *TransformationV1) GetIfOk() (*string, bool)`

GetIfOk returns a tuple with the If field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIf

`func (o *TransformationV1) SetIf(v string)`

SetIf sets If field to given value.


### GetDrop

`func (o *TransformationV1) GetDrop() bool`

GetDrop returns the Drop field if non-nil, zero value otherwise.

### GetDropOk

`func (o *TransformationV1) GetDropOk() (*bool, bool)`

GetDropOk returns a tuple with the Drop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrop

`func (o *TransformationV1) SetDrop(v bool)`

SetDrop sets Drop field to given value.

### HasDrop

`func (o *TransformationV1) HasDrop() bool`

HasDrop returns a boolean if a field has been set.

### GetNewEventName

`func (o *TransformationV1) GetNewEventName() string`

GetNewEventName returns the NewEventName field if non-nil, zero value otherwise.

### GetNewEventNameOk

`func (o *TransformationV1) GetNewEventNameOk() (*string, bool)`

GetNewEventNameOk returns a tuple with the NewEventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewEventName

`func (o *TransformationV1) SetNewEventName(v string)`

SetNewEventName sets NewEventName field to given value.

### HasNewEventName

`func (o *TransformationV1) HasNewEventName() bool`

HasNewEventName returns a boolean if a field has been set.

### GetPropertyRenames

`func (o *TransformationV1) GetPropertyRenames() []PropertyRenameV1`

GetPropertyRenames returns the PropertyRenames field if non-nil, zero value otherwise.

### GetPropertyRenamesOk

`func (o *TransformationV1) GetPropertyRenamesOk() (*[]PropertyRenameV1, bool)`

GetPropertyRenamesOk returns a tuple with the PropertyRenames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyRenames

`func (o *TransformationV1) SetPropertyRenames(v []PropertyRenameV1)`

SetPropertyRenames sets PropertyRenames field to given value.

### HasPropertyRenames

`func (o *TransformationV1) HasPropertyRenames() bool`

HasPropertyRenames returns a boolean if a field has been set.

### GetPropertyValueTransformations

`func (o *TransformationV1) GetPropertyValueTransformations() []PropertyValueTransformationV1`

GetPropertyValueTransformations returns the PropertyValueTransformations field if non-nil, zero value otherwise.

### GetPropertyValueTransformationsOk

`func (o *TransformationV1) GetPropertyValueTransformationsOk() (*[]PropertyValueTransformationV1, bool)`

GetPropertyValueTransformationsOk returns a tuple with the PropertyValueTransformations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyValueTransformations

`func (o *TransformationV1) SetPropertyValueTransformations(v []PropertyValueTransformationV1)`

SetPropertyValueTransformations sets PropertyValueTransformations field to given value.

### HasPropertyValueTransformations

`func (o *TransformationV1) HasPropertyValueTransformations() bool`

HasPropertyValueTransformations returns a boolean if a field has been set.

### GetFqlDefinedProperties

`func (o *TransformationV1) GetFqlDefinedProperties() []FQLDefinedPropertyV1`

GetFqlDefinedProperties returns the FqlDefinedProperties field if non-nil, zero value otherwise.

### GetFqlDefinedPropertiesOk

`func (o *TransformationV1) GetFqlDefinedPropertiesOk() (*[]FQLDefinedPropertyV1, bool)`

GetFqlDefinedPropertiesOk returns a tuple with the FqlDefinedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqlDefinedProperties

`func (o *TransformationV1) SetFqlDefinedProperties(v []FQLDefinedPropertyV1)`

SetFqlDefinedProperties sets FqlDefinedProperties field to given value.

### HasFqlDefinedProperties

`func (o *TransformationV1) HasFqlDefinedProperties() bool`

HasFqlDefinedProperties returns a boolean if a field has been set.

### GetAllowProperties

`func (o *TransformationV1) GetAllowProperties() []string`

GetAllowProperties returns the AllowProperties field if non-nil, zero value otherwise.

### GetAllowPropertiesOk

`func (o *TransformationV1) GetAllowPropertiesOk() (*[]string, bool)`

GetAllowPropertiesOk returns a tuple with the AllowProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowProperties

`func (o *TransformationV1) SetAllowProperties(v []string)`

SetAllowProperties sets AllowProperties field to given value.

### HasAllowProperties

`func (o *TransformationV1) HasAllowProperties() bool`

HasAllowProperties returns a boolean if a field has been set.

### GetHashPropertiesConfiguration

`func (o *TransformationV1) GetHashPropertiesConfiguration() HashPropertiesConfiguration`

GetHashPropertiesConfiguration returns the HashPropertiesConfiguration field if non-nil, zero value otherwise.

### GetHashPropertiesConfigurationOk

`func (o *TransformationV1) GetHashPropertiesConfigurationOk() (*HashPropertiesConfiguration, bool)`

GetHashPropertiesConfigurationOk returns a tuple with the HashPropertiesConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashPropertiesConfiguration

`func (o *TransformationV1) SetHashPropertiesConfiguration(v HashPropertiesConfiguration)`

SetHashPropertiesConfiguration sets HashPropertiesConfiguration field to given value.

### HasHashPropertiesConfiguration

`func (o *TransformationV1) HasHashPropertiesConfiguration() bool`

HasHashPropertiesConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


