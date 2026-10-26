# CreateAudiencePreviewAlphaInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Definition** | [**AudienceDefinition**](AudienceDefinition.md) |  | 
**AudienceType** | **string** | Denotes the type of audience product.  Possible values: USERS, ACCOUNTS, LINKED. | 
**Options** | Pointer to [**CreateAudiencePreviewOptions**](CreateAudiencePreviewOptions.md) |  | [optional] 

## Methods

### NewCreateAudiencePreviewAlphaInput

`func NewCreateAudiencePreviewAlphaInput(definition AudienceDefinition, audienceType string, ) *CreateAudiencePreviewAlphaInput`

NewCreateAudiencePreviewAlphaInput instantiates a new CreateAudiencePreviewAlphaInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAudiencePreviewAlphaInputWithDefaults

`func NewCreateAudiencePreviewAlphaInputWithDefaults() *CreateAudiencePreviewAlphaInput`

NewCreateAudiencePreviewAlphaInputWithDefaults instantiates a new CreateAudiencePreviewAlphaInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefinition

`func (o *CreateAudiencePreviewAlphaInput) GetDefinition() AudienceDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *CreateAudiencePreviewAlphaInput) GetDefinitionOk() (*AudienceDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *CreateAudiencePreviewAlphaInput) SetDefinition(v AudienceDefinition)`

SetDefinition sets Definition field to given value.


### GetAudienceType

`func (o *CreateAudiencePreviewAlphaInput) GetAudienceType() string`

GetAudienceType returns the AudienceType field if non-nil, zero value otherwise.

### GetAudienceTypeOk

`func (o *CreateAudiencePreviewAlphaInput) GetAudienceTypeOk() (*string, bool)`

GetAudienceTypeOk returns a tuple with the AudienceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceType

`func (o *CreateAudiencePreviewAlphaInput) SetAudienceType(v string)`

SetAudienceType sets AudienceType field to given value.


### GetOptions

`func (o *CreateAudiencePreviewAlphaInput) GetOptions() CreateAudiencePreviewOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CreateAudiencePreviewAlphaInput) GetOptionsOk() (*CreateAudiencePreviewOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CreateAudiencePreviewAlphaInput) SetOptions(v CreateAudiencePreviewOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CreateAudiencePreviewAlphaInput) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


