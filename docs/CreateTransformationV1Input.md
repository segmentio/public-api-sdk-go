# CreateTransformationV1Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the Transformation. | 
**SourceId** | **string** | The Source to be associated with the Transformation. | 
**DestinationMetadataId** | Pointer to **string** | The optional Destination metadata id to be associated with the Transformation. | [optional] 
**Enabled** | **bool** | If the Transformation should be enabled. | 
**If** | **string** | If statement ([FQL](https://segment.com/docs/config-api/fql/)) to match events.  For standard event matchers, use the following:  Track -\\&gt; \&quot;event&#x3D;&#39;\\&lt;eventName\\&gt;&#39;\&quot;  Identify -\\&gt; \&quot;type&#x3D;&#39;identify&#39;\&quot;  Group -\\&gt; \&quot;type&#x3D;&#39;group&#39;\&quot; | 
**Drop** | Pointer to **bool** | Optional boolean value if the Transformation should drop the event entirely when the if statement matches, ignores all other transforms. | [optional] 
**NewEventName** | Pointer to **string** | Optional new event name for renaming events. Works only for &#39;track&#39; event type. | [optional] 
**PropertyRenames** | Pointer to [**[]PropertyRenameV1**](PropertyRenameV1.md) | Optional array for renaming properties collected by your events. | [optional] 
**PropertyValueTransformations** | Pointer to [**[]PropertyValueTransformationV1**](PropertyValueTransformationV1.md) | Optional array for transforming properties and values collected by your events. Limited to 10 properties. | [optional] 
**FqlDefinedProperties** | Pointer to [**[]FQLDefinedPropertyV1**](FQLDefinedPropertyV1.md) | Optional array for defining new properties in [FQL](https://segment.com/docs/config-api/fql/). Currently limited to 1 property. | [optional] 
**AllowProperties** | Pointer to **[]string** | Optional array for allowing properties from your events. | [optional] 
**HashPropertiesConfiguration** | Pointer to [**HashPropertiesConfiguration**](HashPropertiesConfiguration.md) |  | [optional] 

## Methods

### NewCreateTransformationV1Input

`func NewCreateTransformationV1Input(name string, sourceId string, enabled bool, if_ string, ) *CreateTransformationV1Input`

NewCreateTransformationV1Input instantiates a new CreateTransformationV1Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTransformationV1InputWithDefaults

`func NewCreateTransformationV1InputWithDefaults() *CreateTransformationV1Input`

NewCreateTransformationV1InputWithDefaults instantiates a new CreateTransformationV1Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateTransformationV1Input) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTransformationV1Input) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTransformationV1Input) SetName(v string)`

SetName sets Name field to given value.


### GetSourceId

`func (o *CreateTransformationV1Input) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *CreateTransformationV1Input) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *CreateTransformationV1Input) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetDestinationMetadataId

`func (o *CreateTransformationV1Input) GetDestinationMetadataId() string`

GetDestinationMetadataId returns the DestinationMetadataId field if non-nil, zero value otherwise.

### GetDestinationMetadataIdOk

`func (o *CreateTransformationV1Input) GetDestinationMetadataIdOk() (*string, bool)`

GetDestinationMetadataIdOk returns a tuple with the DestinationMetadataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationMetadataId

`func (o *CreateTransformationV1Input) SetDestinationMetadataId(v string)`

SetDestinationMetadataId sets DestinationMetadataId field to given value.

### HasDestinationMetadataId

`func (o *CreateTransformationV1Input) HasDestinationMetadataId() bool`

HasDestinationMetadataId returns a boolean if a field has been set.

### GetEnabled

`func (o *CreateTransformationV1Input) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateTransformationV1Input) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateTransformationV1Input) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIf

`func (o *CreateTransformationV1Input) GetIf() string`

GetIf returns the If field if non-nil, zero value otherwise.

### GetIfOk

`func (o *CreateTransformationV1Input) GetIfOk() (*string, bool)`

GetIfOk returns a tuple with the If field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIf

`func (o *CreateTransformationV1Input) SetIf(v string)`

SetIf sets If field to given value.


### GetDrop

`func (o *CreateTransformationV1Input) GetDrop() bool`

GetDrop returns the Drop field if non-nil, zero value otherwise.

### GetDropOk

`func (o *CreateTransformationV1Input) GetDropOk() (*bool, bool)`

GetDropOk returns a tuple with the Drop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrop

`func (o *CreateTransformationV1Input) SetDrop(v bool)`

SetDrop sets Drop field to given value.

### HasDrop

`func (o *CreateTransformationV1Input) HasDrop() bool`

HasDrop returns a boolean if a field has been set.

### GetNewEventName

`func (o *CreateTransformationV1Input) GetNewEventName() string`

GetNewEventName returns the NewEventName field if non-nil, zero value otherwise.

### GetNewEventNameOk

`func (o *CreateTransformationV1Input) GetNewEventNameOk() (*string, bool)`

GetNewEventNameOk returns a tuple with the NewEventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewEventName

`func (o *CreateTransformationV1Input) SetNewEventName(v string)`

SetNewEventName sets NewEventName field to given value.

### HasNewEventName

`func (o *CreateTransformationV1Input) HasNewEventName() bool`

HasNewEventName returns a boolean if a field has been set.

### GetPropertyRenames

`func (o *CreateTransformationV1Input) GetPropertyRenames() []PropertyRenameV1`

GetPropertyRenames returns the PropertyRenames field if non-nil, zero value otherwise.

### GetPropertyRenamesOk

`func (o *CreateTransformationV1Input) GetPropertyRenamesOk() (*[]PropertyRenameV1, bool)`

GetPropertyRenamesOk returns a tuple with the PropertyRenames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyRenames

`func (o *CreateTransformationV1Input) SetPropertyRenames(v []PropertyRenameV1)`

SetPropertyRenames sets PropertyRenames field to given value.

### HasPropertyRenames

`func (o *CreateTransformationV1Input) HasPropertyRenames() bool`

HasPropertyRenames returns a boolean if a field has been set.

### GetPropertyValueTransformations

`func (o *CreateTransformationV1Input) GetPropertyValueTransformations() []PropertyValueTransformationV1`

GetPropertyValueTransformations returns the PropertyValueTransformations field if non-nil, zero value otherwise.

### GetPropertyValueTransformationsOk

`func (o *CreateTransformationV1Input) GetPropertyValueTransformationsOk() (*[]PropertyValueTransformationV1, bool)`

GetPropertyValueTransformationsOk returns a tuple with the PropertyValueTransformations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyValueTransformations

`func (o *CreateTransformationV1Input) SetPropertyValueTransformations(v []PropertyValueTransformationV1)`

SetPropertyValueTransformations sets PropertyValueTransformations field to given value.

### HasPropertyValueTransformations

`func (o *CreateTransformationV1Input) HasPropertyValueTransformations() bool`

HasPropertyValueTransformations returns a boolean if a field has been set.

### GetFqlDefinedProperties

`func (o *CreateTransformationV1Input) GetFqlDefinedProperties() []FQLDefinedPropertyV1`

GetFqlDefinedProperties returns the FqlDefinedProperties field if non-nil, zero value otherwise.

### GetFqlDefinedPropertiesOk

`func (o *CreateTransformationV1Input) GetFqlDefinedPropertiesOk() (*[]FQLDefinedPropertyV1, bool)`

GetFqlDefinedPropertiesOk returns a tuple with the FqlDefinedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqlDefinedProperties

`func (o *CreateTransformationV1Input) SetFqlDefinedProperties(v []FQLDefinedPropertyV1)`

SetFqlDefinedProperties sets FqlDefinedProperties field to given value.

### HasFqlDefinedProperties

`func (o *CreateTransformationV1Input) HasFqlDefinedProperties() bool`

HasFqlDefinedProperties returns a boolean if a field has been set.

### GetAllowProperties

`func (o *CreateTransformationV1Input) GetAllowProperties() []string`

GetAllowProperties returns the AllowProperties field if non-nil, zero value otherwise.

### GetAllowPropertiesOk

`func (o *CreateTransformationV1Input) GetAllowPropertiesOk() (*[]string, bool)`

GetAllowPropertiesOk returns a tuple with the AllowProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowProperties

`func (o *CreateTransformationV1Input) SetAllowProperties(v []string)`

SetAllowProperties sets AllowProperties field to given value.

### HasAllowProperties

`func (o *CreateTransformationV1Input) HasAllowProperties() bool`

HasAllowProperties returns a boolean if a field has been set.

### GetHashPropertiesConfiguration

`func (o *CreateTransformationV1Input) GetHashPropertiesConfiguration() HashPropertiesConfiguration`

GetHashPropertiesConfiguration returns the HashPropertiesConfiguration field if non-nil, zero value otherwise.

### GetHashPropertiesConfigurationOk

`func (o *CreateTransformationV1Input) GetHashPropertiesConfigurationOk() (*HashPropertiesConfiguration, bool)`

GetHashPropertiesConfigurationOk returns a tuple with the HashPropertiesConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashPropertiesConfiguration

`func (o *CreateTransformationV1Input) SetHashPropertiesConfiguration(v HashPropertiesConfiguration)`

SetHashPropertiesConfiguration sets HashPropertiesConfiguration field to given value.

### HasHashPropertiesConfiguration

`func (o *CreateTransformationV1Input) HasHashPropertiesConfiguration() bool`

HasHashPropertiesConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


