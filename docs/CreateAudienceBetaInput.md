# CreateAudienceBetaInput

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

### NewCreateAudienceBetaInput

`func NewCreateAudienceBetaInput(name string, definition AudienceDefinition, audienceType string, ) *CreateAudienceBetaInput`

NewCreateAudienceBetaInput instantiates a new CreateAudienceBetaInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAudienceBetaInputWithDefaults

`func NewCreateAudienceBetaInputWithDefaults() *CreateAudienceBetaInput`

NewCreateAudienceBetaInputWithDefaults instantiates a new CreateAudienceBetaInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateAudienceBetaInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAudienceBetaInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAudienceBetaInput) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *CreateAudienceBetaInput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateAudienceBetaInput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateAudienceBetaInput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateAudienceBetaInput) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *CreateAudienceBetaInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAudienceBetaInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAudienceBetaInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateAudienceBetaInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefinition

`func (o *CreateAudienceBetaInput) GetDefinition() AudienceDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *CreateAudienceBetaInput) GetDefinitionOk() (*AudienceDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *CreateAudienceBetaInput) SetDefinition(v AudienceDefinition)`

SetDefinition sets Definition field to given value.


### GetAudienceType

`func (o *CreateAudienceBetaInput) GetAudienceType() string`

GetAudienceType returns the AudienceType field if non-nil, zero value otherwise.

### GetAudienceTypeOk

`func (o *CreateAudienceBetaInput) GetAudienceTypeOk() (*string, bool)`

GetAudienceTypeOk returns a tuple with the AudienceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceType

`func (o *CreateAudienceBetaInput) SetAudienceType(v string)`

SetAudienceType sets AudienceType field to given value.


### GetOptions

`func (o *CreateAudienceBetaInput) GetOptions() AudienceOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CreateAudienceBetaInput) GetOptionsOk() (*AudienceOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CreateAudienceBetaInput) SetOptions(v AudienceOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CreateAudienceBetaInput) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


