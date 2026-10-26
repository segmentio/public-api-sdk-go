# RemoveRuleV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type for this Tracking Plan rule. | 
**Key** | Pointer to **string** | Key to this rule (free-form string like &#39;Button clicked&#39;). | [optional] 
**Version** | **float32** | Version of this rule. | 

## Methods

### NewRemoveRuleV1

`func NewRemoveRuleV1(type_ string, version float32, ) *RemoveRuleV1`

NewRemoveRuleV1 instantiates a new RemoveRuleV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveRuleV1WithDefaults

`func NewRemoveRuleV1WithDefaults() *RemoveRuleV1`

NewRemoveRuleV1WithDefaults instantiates a new RemoveRuleV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RemoveRuleV1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RemoveRuleV1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RemoveRuleV1) SetType(v string)`

SetType sets Type field to given value.


### GetKey

`func (o *RemoveRuleV1) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *RemoveRuleV1) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *RemoveRuleV1) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *RemoveRuleV1) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetVersion

`func (o *RemoveRuleV1) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RemoveRuleV1) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RemoveRuleV1) SetVersion(v float32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


