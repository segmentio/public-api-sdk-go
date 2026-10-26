# LabelV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The key that represents the name of this label. | 
**Value** | **string** | The value associated with the key of this label. | 
**Description** | Pointer to **string** | An optional description of the purpose of this label. | [optional] 

## Methods

### NewLabelV1

`func NewLabelV1(key string, value string, ) *LabelV1`

NewLabelV1 instantiates a new LabelV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelV1WithDefaults

`func NewLabelV1WithDefaults() *LabelV1`

NewLabelV1WithDefaults instantiates a new LabelV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *LabelV1) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *LabelV1) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *LabelV1) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *LabelV1) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *LabelV1) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *LabelV1) SetValue(v string)`

SetValue sets Value field to given value.


### GetDescription

`func (o *LabelV1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LabelV1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LabelV1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LabelV1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


