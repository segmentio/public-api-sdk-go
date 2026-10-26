# RuleV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type for this Tracking Plan rule. | 
**Key** | Pointer to **string** | Key to this rule (free-form string like &#39;Button clicked&#39;). | [optional] 
**JsonSchema** | **interface{}** | JSON Schema of this rule. | 
**Version** | **float32** | Version of this rule. | 
**CreatedAt** | Pointer to **string** | The timestamp of this rule&#39;s creation. | [optional] 
**UpdatedAt** | Pointer to **string** | The timestamp of this rule&#39;s last change. | [optional] 
**DeprecatedAt** | Pointer to **string** | The timestamp of this rule&#39;s deprecation. | [optional] 

## Methods

### NewRuleV1

`func NewRuleV1(type_ string, jsonSchema interface{}, version float32, ) *RuleV1`

NewRuleV1 instantiates a new RuleV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleV1WithDefaults

`func NewRuleV1WithDefaults() *RuleV1`

NewRuleV1WithDefaults instantiates a new RuleV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RuleV1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RuleV1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RuleV1) SetType(v string)`

SetType sets Type field to given value.


### GetKey

`func (o *RuleV1) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *RuleV1) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *RuleV1) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *RuleV1) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetJsonSchema

`func (o *RuleV1) GetJsonSchema() interface{}`

GetJsonSchema returns the JsonSchema field if non-nil, zero value otherwise.

### GetJsonSchemaOk

`func (o *RuleV1) GetJsonSchemaOk() (*interface{}, bool)`

GetJsonSchemaOk returns a tuple with the JsonSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonSchema

`func (o *RuleV1) SetJsonSchema(v interface{})`

SetJsonSchema sets JsonSchema field to given value.


### SetJsonSchemaNil

`func (o *RuleV1) SetJsonSchemaNil(b bool)`

 SetJsonSchemaNil sets the value for JsonSchema to be an explicit nil

### UnsetJsonSchema
`func (o *RuleV1) UnsetJsonSchema()`

UnsetJsonSchema ensures that no value is present for JsonSchema, not even an explicit nil
### GetVersion

`func (o *RuleV1) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RuleV1) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RuleV1) SetVersion(v float32)`

SetVersion sets Version field to given value.


### GetCreatedAt

`func (o *RuleV1) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RuleV1) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RuleV1) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RuleV1) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RuleV1) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RuleV1) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RuleV1) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RuleV1) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeprecatedAt

`func (o *RuleV1) GetDeprecatedAt() string`

GetDeprecatedAt returns the DeprecatedAt field if non-nil, zero value otherwise.

### GetDeprecatedAtOk

`func (o *RuleV1) GetDeprecatedAtOk() (*string, bool)`

GetDeprecatedAtOk returns a tuple with the DeprecatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecatedAt

`func (o *RuleV1) SetDeprecatedAt(v string)`

SetDeprecatedAt sets DeprecatedAt field to given value.

### HasDeprecatedAt

`func (o *RuleV1) HasDeprecatedAt() bool`

HasDeprecatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


