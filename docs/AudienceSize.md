# AudienceSize

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **float32** | The total audience membership count. Refer to the type field to determine the unit for this field (profiles, accounts, etc). | [optional] 
**Type** | **string** | The unit type for the count(s) being returned. | 
**UniqueCount** | Pointer to **float32** | The unique audience membership count. | [optional] 

## Methods

### NewAudienceSize

`func NewAudienceSize(type_ string, ) *AudienceSize`

NewAudienceSize instantiates a new AudienceSize object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceSizeWithDefaults

`func NewAudienceSizeWithDefaults() *AudienceSize`

NewAudienceSizeWithDefaults instantiates a new AudienceSize object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *AudienceSize) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AudienceSize) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AudienceSize) SetCount(v float32)`

SetCount sets Count field to given value.

### HasCount

`func (o *AudienceSize) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetType

`func (o *AudienceSize) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AudienceSize) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AudienceSize) SetType(v string)`

SetType sets Type field to given value.


### GetUniqueCount

`func (o *AudienceSize) GetUniqueCount() float32`

GetUniqueCount returns the UniqueCount field if non-nil, zero value otherwise.

### GetUniqueCountOk

`func (o *AudienceSize) GetUniqueCountOk() (*float32, bool)`

GetUniqueCountOk returns a tuple with the UniqueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueCount

`func (o *AudienceSize) SetUniqueCount(v float32)`

SetUniqueCount sets UniqueCount field to given value.

### HasUniqueCount

`func (o *AudienceSize) HasUniqueCount() bool`

HasUniqueCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


