# UpsertRuleV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewKey** | Pointer to **string** | This rule&#39;s new intended key. | [optional] 
**Type** | **string** | The type for this Tracking Plan rule. | 
**Key** | Pointer to **string** | Key to this rule (free-form string like &#39;Button clicked&#39;). | [optional] 
**JsonSchema** | **interface{}** | JSON Schema of this rule. | 
**Version** | **float32** | Version of this rule. | 

## Methods

### NewUpsertRuleV1

`func NewUpsertRuleV1(type_ string, jsonSchema interface{}, version float32, ) *UpsertRuleV1`

NewUpsertRuleV1 instantiates a new UpsertRuleV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpsertRuleV1WithDefaults

`func NewUpsertRuleV1WithDefaults() *UpsertRuleV1`

NewUpsertRuleV1WithDefaults instantiates a new UpsertRuleV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewKey

`func (o *UpsertRuleV1) GetNewKey() string`

GetNewKey returns the NewKey field if non-nil, zero value otherwise.

### GetNewKeyOk

`func (o *UpsertRuleV1) GetNewKeyOk() (*string, bool)`

GetNewKeyOk returns a tuple with the NewKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewKey

`func (o *UpsertRuleV1) SetNewKey(v string)`

SetNewKey sets NewKey field to given value.

### HasNewKey

`func (o *UpsertRuleV1) HasNewKey() bool`

HasNewKey returns a boolean if a field has been set.

### GetType

`func (o *UpsertRuleV1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpsertRuleV1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpsertRuleV1) SetType(v string)`

SetType sets Type field to given value.


### GetKey

`func (o *UpsertRuleV1) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpsertRuleV1) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpsertRuleV1) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpsertRuleV1) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetJsonSchema

`func (o *UpsertRuleV1) GetJsonSchema() interface{}`

GetJsonSchema returns the JsonSchema field if non-nil, zero value otherwise.

### GetJsonSchemaOk

`func (o *UpsertRuleV1) GetJsonSchemaOk() (*interface{}, bool)`

GetJsonSchemaOk returns a tuple with the JsonSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonSchema

`func (o *UpsertRuleV1) SetJsonSchema(v interface{})`

SetJsonSchema sets JsonSchema field to given value.


### SetJsonSchemaNil

`func (o *UpsertRuleV1) SetJsonSchemaNil(b bool)`

 SetJsonSchemaNil sets the value for JsonSchema to be an explicit nil

### UnsetJsonSchema
`func (o *UpsertRuleV1) UnsetJsonSchema()`

UnsetJsonSchema ensures that no value is present for JsonSchema, not even an explicit nil
### GetVersion

`func (o *UpsertRuleV1) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpsertRuleV1) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpsertRuleV1) SetVersion(v float32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


