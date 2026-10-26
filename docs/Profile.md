# Profile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Properties** | Pointer to **[]string** | The profile traits to include in the event sent to the Destination. | [optional] 
**Mapping** | Pointer to **map[string]string** | Maps a profile trait to the name it should be sent under. Each key is a trait, and each value is the name used in the event. | [optional] 

## Methods

### NewProfile

`func NewProfile() *Profile`

NewProfile instantiates a new Profile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileWithDefaults

`func NewProfileWithDefaults() *Profile`

NewProfileWithDefaults instantiates a new Profile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *Profile) GetProperties() []string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Profile) GetPropertiesOk() (*[]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Profile) SetProperties(v []string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Profile) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetMapping

`func (o *Profile) GetMapping() map[string]string`

GetMapping returns the Mapping field if non-nil, zero value otherwise.

### GetMappingOk

`func (o *Profile) GetMappingOk() (*map[string]string, bool)`

GetMappingOk returns a tuple with the Mapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapping

`func (o *Profile) SetMapping(v map[string]string)`

SetMapping sets Mapping field to given value.

### HasMapping

`func (o *Profile) HasMapping() bool`

HasMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


