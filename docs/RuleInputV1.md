# RuleInputV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type for this Tracking Plan rule. | 
**Key** | Pointer to **string** | Key to this rule (free-form string like &#39;Button clicked&#39;). | [optional] 
**JsonSchema** | **interface{}** | JSON Schema of this rule. | 
**Version** | **float32** | Version of this rule. | 

## Methods

### NewRuleInputV1

`func NewRuleInputV1(type_ string, jsonSchema interface{}, version float32, ) *RuleInputV1`

NewRuleInputV1 instantiates a new RuleInputV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleInputV1WithDefaults

`func NewRuleInputV1WithDefaults() *RuleInputV1`

NewRuleInputV1WithDefaults instantiates a new RuleInputV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RuleInputV1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RuleInputV1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RuleInputV1) SetType(v string)`

SetType sets Type field to given value.


### GetKey

`func (o *RuleInputV1) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *RuleInputV1) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *RuleInputV1) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *RuleInputV1) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetJsonSchema

`func (o *RuleInputV1) GetJsonSchema() interface{}`

GetJsonSchema returns the JsonSchema field if non-nil, zero value otherwise.

### GetJsonSchemaOk

`func (o *RuleInputV1) GetJsonSchemaOk() (*interface{}, bool)`

GetJsonSchemaOk returns a tuple with the JsonSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonSchema

`func (o *RuleInputV1) SetJsonSchema(v interface{})`

SetJsonSchema sets JsonSchema field to given value.


### SetJsonSchemaNil

`func (o *RuleInputV1) SetJsonSchemaNil(b bool)`

 SetJsonSchemaNil sets the value for JsonSchema to be an explicit nil

### UnsetJsonSchema
`func (o *RuleInputV1) UnsetJsonSchema()`

UnsetJsonSchema ensures that no value is present for JsonSchema, not even an explicit nil
### GetVersion

`func (o *RuleInputV1) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RuleInputV1) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RuleInputV1) SetVersion(v float32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


