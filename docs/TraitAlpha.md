# TraitAlpha

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Trait** | **string** | Trait key. | 
**Type** | [**TraitType**](TraitType.md) |  | 
**LastSeenAt** | **string** | ISO timestamp of when the trait was last seen. | 
**Samples** | Pointer to **[]string** | Sample values for this trait. Only present when includeSampleValues is true. | [optional] 

## Methods

### NewTraitAlpha

`func NewTraitAlpha(trait string, type_ TraitType, lastSeenAt string, ) *TraitAlpha`

NewTraitAlpha instantiates a new TraitAlpha object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraitAlphaWithDefaults

`func NewTraitAlphaWithDefaults() *TraitAlpha`

NewTraitAlphaWithDefaults instantiates a new TraitAlpha object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrait

`func (o *TraitAlpha) GetTrait() string`

GetTrait returns the Trait field if non-nil, zero value otherwise.

### GetTraitOk

`func (o *TraitAlpha) GetTraitOk() (*string, bool)`

GetTraitOk returns a tuple with the Trait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrait

`func (o *TraitAlpha) SetTrait(v string)`

SetTrait sets Trait field to given value.


### GetType

`func (o *TraitAlpha) GetType() TraitType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TraitAlpha) GetTypeOk() (*TraitType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TraitAlpha) SetType(v TraitType)`

SetType sets Type field to given value.


### GetLastSeenAt

`func (o *TraitAlpha) GetLastSeenAt() string`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *TraitAlpha) GetLastSeenAtOk() (*string, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *TraitAlpha) SetLastSeenAt(v string)`

SetLastSeenAt sets LastSeenAt field to given value.


### GetSamples

`func (o *TraitAlpha) GetSamples() []string`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *TraitAlpha) GetSamplesOk() (*[]string, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *TraitAlpha) SetSamples(v []string)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *TraitAlpha) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


