# CreateAudienceInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the audience. | 
**Enabled** | Pointer to **bool** | Determines whether a computation is enabled. | [optional] 
**Description** | Pointer to **string** | Description of the audience. | [optional] 
**Definition** | [**AudienceDefinition**](AudienceDefinition.md) |  | 
**AudienceType** | **string** | Denotes the type of audience product.  Possible values: USERS, ACCOUNTS, LINKED. | 
**Options** | Pointer to [**AudienceOptions**](AudienceOptions.md) |  | [optional] 

## Methods

### NewCreateAudienceInput

`func NewCreateAudienceInput(name string, definition AudienceDefinition, audienceType string, ) *CreateAudienceInput`

NewCreateAudienceInput instantiates a new CreateAudienceInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAudienceInputWithDefaults

`func NewCreateAudienceInputWithDefaults() *CreateAudienceInput`

NewCreateAudienceInputWithDefaults instantiates a new CreateAudienceInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateAudienceInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAudienceInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAudienceInput) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *CreateAudienceInput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateAudienceInput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateAudienceInput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateAudienceInput) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *CreateAudienceInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAudienceInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAudienceInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateAudienceInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefinition

`func (o *CreateAudienceInput) GetDefinition() AudienceDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *CreateAudienceInput) GetDefinitionOk() (*AudienceDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *CreateAudienceInput) SetDefinition(v AudienceDefinition)`

SetDefinition sets Definition field to given value.


### GetAudienceType

`func (o *CreateAudienceInput) GetAudienceType() string`

GetAudienceType returns the AudienceType field if non-nil, zero value otherwise.

### GetAudienceTypeOk

`func (o *CreateAudienceInput) GetAudienceTypeOk() (*string, bool)`

GetAudienceTypeOk returns a tuple with the AudienceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceType

`func (o *CreateAudienceInput) SetAudienceType(v string)`

SetAudienceType sets AudienceType field to given value.


### GetOptions

`func (o *CreateAudienceInput) GetOptions() AudienceOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CreateAudienceInput) GetOptionsOk() (*AudienceOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CreateAudienceInput) SetOptions(v AudienceOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CreateAudienceInput) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


