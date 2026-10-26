# UpdateTransformationV1Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the Transformation. | [optional] 
**SourceId** | Pointer to **string** | The optional Source to be associated with the Transformation. | [optional] 
**DestinationMetadataId** | Pointer to **string** | The optional Destination metadata to be associated with the Transformation. | [optional] 
**Enabled** | Pointer to **bool** | If the Transformation should be enabled. | [optional] 
**If** | Pointer to **string** | If statement ([FQL](https://segment.com/docs/config-api/fql/)) to match events.  For standard event matchers, use the following:  Track -\\&gt; \&quot;event&#x3D;&#39;\\&lt;eventName\\&gt;&#39;\&quot;  Identify -\\&gt; \&quot;type&#x3D;&#39;identify&#39;\&quot;  Group -\\&gt; \&quot;type&#x3D;&#39;group&#39;\&quot; | [optional] 
**Drop** | Pointer to **bool** | Optional boolean value if the Transformation should drop the event entirely when the if statement matches, ignores all other transforms. | [optional] 
**NewEventName** | Pointer to **string** | Optional new event name for renaming events. Works only for &#39;track&#39; event type. | [optional] 
**PropertyRenames** | Pointer to [**[]PropertyRenameV1**](PropertyRenameV1.md) | Optional array for renaming properties collected by your events. | [optional] 
**PropertyValueTransformations** | Pointer to [**[]PropertyValueTransformationV1**](PropertyValueTransformationV1.md) | Optional array for transforming properties and values collected by your events. Limited to 10 properties. | [optional] 
**FqlDefinedProperties** | Pointer to [**[]FQLDefinedPropertyV1**](FQLDefinedPropertyV1.md) | Optional array for updating properties defined in [FQL](https://segment.com/docs/config-api/fql/). Currently limited to 1 property. | [optional] 
**AllowProperties** | Pointer to **[]string** | Optional array for allowing properties from your events. | [optional] 
**HashPropertiesConfiguration** | Pointer to [**HashPropertiesConfiguration**](HashPropertiesConfiguration.md) |  | [optional] 

## Methods

### NewUpdateTransformationV1Input

`func NewUpdateTransformationV1Input() *UpdateTransformationV1Input`

NewUpdateTransformationV1Input instantiates a new UpdateTransformationV1Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTransformationV1InputWithDefaults

`func NewUpdateTransformationV1InputWithDefaults() *UpdateTransformationV1Input`

NewUpdateTransformationV1InputWithDefaults instantiates a new UpdateTransformationV1Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateTransformationV1Input) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTransformationV1Input) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTransformationV1Input) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateTransformationV1Input) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSourceId

`func (o *UpdateTransformationV1Input) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *UpdateTransformationV1Input) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *UpdateTransformationV1Input) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *UpdateTransformationV1Input) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetDestinationMetadataId

`func (o *UpdateTransformationV1Input) GetDestinationMetadataId() string`

GetDestinationMetadataId returns the DestinationMetadataId field if non-nil, zero value otherwise.

### GetDestinationMetadataIdOk

`func (o *UpdateTransformationV1Input) GetDestinationMetadataIdOk() (*string, bool)`

GetDestinationMetadataIdOk returns a tuple with the DestinationMetadataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationMetadataId

`func (o *UpdateTransformationV1Input) SetDestinationMetadataId(v string)`

SetDestinationMetadataId sets DestinationMetadataId field to given value.

### HasDestinationMetadataId

`func (o *UpdateTransformationV1Input) HasDestinationMetadataId() bool`

HasDestinationMetadataId returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateTransformationV1Input) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateTransformationV1Input) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateTransformationV1Input) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateTransformationV1Input) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIf

`func (o *UpdateTransformationV1Input) GetIf() string`

GetIf returns the If field if non-nil, zero value otherwise.

### GetIfOk

`func (o *UpdateTransformationV1Input) GetIfOk() (*string, bool)`

GetIfOk returns a tuple with the If field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIf

`func (o *UpdateTransformationV1Input) SetIf(v string)`

SetIf sets If field to given value.

### HasIf

`func (o *UpdateTransformationV1Input) HasIf() bool`

HasIf returns a boolean if a field has been set.

### GetDrop

`func (o *UpdateTransformationV1Input) GetDrop() bool`

GetDrop returns the Drop field if non-nil, zero value otherwise.

### GetDropOk

`func (o *UpdateTransformationV1Input) GetDropOk() (*bool, bool)`

GetDropOk returns a tuple with the Drop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrop

`func (o *UpdateTransformationV1Input) SetDrop(v bool)`

SetDrop sets Drop field to given value.

### HasDrop

`func (o *UpdateTransformationV1Input) HasDrop() bool`

HasDrop returns a boolean if a field has been set.

### GetNewEventName

`func (o *UpdateTransformationV1Input) GetNewEventName() string`

GetNewEventName returns the NewEventName field if non-nil, zero value otherwise.

### GetNewEventNameOk

`func (o *UpdateTransformationV1Input) GetNewEventNameOk() (*string, bool)`

GetNewEventNameOk returns a tuple with the NewEventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewEventName

`func (o *UpdateTransformationV1Input) SetNewEventName(v string)`

SetNewEventName sets NewEventName field to given value.

### HasNewEventName

`func (o *UpdateTransformationV1Input) HasNewEventName() bool`

HasNewEventName returns a boolean if a field has been set.

### GetPropertyRenames

`func (o *UpdateTransformationV1Input) GetPropertyRenames() []PropertyRenameV1`

GetPropertyRenames returns the PropertyRenames field if non-nil, zero value otherwise.

### GetPropertyRenamesOk

`func (o *UpdateTransformationV1Input) GetPropertyRenamesOk() (*[]PropertyRenameV1, bool)`

GetPropertyRenamesOk returns a tuple with the PropertyRenames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyRenames

`func (o *UpdateTransformationV1Input) SetPropertyRenames(v []PropertyRenameV1)`

SetPropertyRenames sets PropertyRenames field to given value.

### HasPropertyRenames

`func (o *UpdateTransformationV1Input) HasPropertyRenames() bool`

HasPropertyRenames returns a boolean if a field has been set.

### GetPropertyValueTransformations

`func (o *UpdateTransformationV1Input) GetPropertyValueTransformations() []PropertyValueTransformationV1`

GetPropertyValueTransformations returns the PropertyValueTransformations field if non-nil, zero value otherwise.

### GetPropertyValueTransformationsOk

`func (o *UpdateTransformationV1Input) GetPropertyValueTransformationsOk() (*[]PropertyValueTransformationV1, bool)`

GetPropertyValueTransformationsOk returns a tuple with the PropertyValueTransformations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyValueTransformations

`func (o *UpdateTransformationV1Input) SetPropertyValueTransformations(v []PropertyValueTransformationV1)`

SetPropertyValueTransformations sets PropertyValueTransformations field to given value.

### HasPropertyValueTransformations

`func (o *UpdateTransformationV1Input) HasPropertyValueTransformations() bool`

HasPropertyValueTransformations returns a boolean if a field has been set.

### GetFqlDefinedProperties

`func (o *UpdateTransformationV1Input) GetFqlDefinedProperties() []FQLDefinedPropertyV1`

GetFqlDefinedProperties returns the FqlDefinedProperties field if non-nil, zero value otherwise.

### GetFqlDefinedPropertiesOk

`func (o *UpdateTransformationV1Input) GetFqlDefinedPropertiesOk() (*[]FQLDefinedPropertyV1, bool)`

GetFqlDefinedPropertiesOk returns a tuple with the FqlDefinedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqlDefinedProperties

`func (o *UpdateTransformationV1Input) SetFqlDefinedProperties(v []FQLDefinedPropertyV1)`

SetFqlDefinedProperties sets FqlDefinedProperties field to given value.

### HasFqlDefinedProperties

`func (o *UpdateTransformationV1Input) HasFqlDefinedProperties() bool`

HasFqlDefinedProperties returns a boolean if a field has been set.

### GetAllowProperties

`func (o *UpdateTransformationV1Input) GetAllowProperties() []string`

GetAllowProperties returns the AllowProperties field if non-nil, zero value otherwise.

### GetAllowPropertiesOk

`func (o *UpdateTransformationV1Input) GetAllowPropertiesOk() (*[]string, bool)`

GetAllowPropertiesOk returns a tuple with the AllowProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowProperties

`func (o *UpdateTransformationV1Input) SetAllowProperties(v []string)`

SetAllowProperties sets AllowProperties field to given value.

### HasAllowProperties

`func (o *UpdateTransformationV1Input) HasAllowProperties() bool`

HasAllowProperties returns a boolean if a field has been set.

### GetHashPropertiesConfiguration

`func (o *UpdateTransformationV1Input) GetHashPropertiesConfiguration() HashPropertiesConfiguration`

GetHashPropertiesConfiguration returns the HashPropertiesConfiguration field if non-nil, zero value otherwise.

### GetHashPropertiesConfigurationOk

`func (o *UpdateTransformationV1Input) GetHashPropertiesConfigurationOk() (*HashPropertiesConfiguration, bool)`

GetHashPropertiesConfigurationOk returns a tuple with the HashPropertiesConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashPropertiesConfiguration

`func (o *UpdateTransformationV1Input) SetHashPropertiesConfiguration(v HashPropertiesConfiguration)`

SetHashPropertiesConfiguration sets HashPropertiesConfiguration field to given value.

### HasHashPropertiesConfiguration

`func (o *UpdateTransformationV1Input) HasHashPropertiesConfiguration() bool`

HasHashPropertiesConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


