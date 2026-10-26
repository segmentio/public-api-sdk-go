# Profile1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Properties** | **[]string** | The profile traits included in the event sent to the Destination. | 
**Mapping** | Pointer to **map[string]string** | Maps a profile trait to the name it should be sent under. Each key is a trait, and each value is the name used in the event. | [optional] 

## Methods

### NewProfile1

`func NewProfile1(properties []string, ) *Profile1`

NewProfile1 instantiates a new Profile1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfile1WithDefaults

`func NewProfile1WithDefaults() *Profile1`

NewProfile1WithDefaults instantiates a new Profile1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *Profile1) GetProperties() []string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Profile1) GetPropertiesOk() (*[]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Profile1) SetProperties(v []string)`

SetProperties sets Properties field to given value.


### GetMapping

`func (o *Profile1) GetMapping() map[string]string`

GetMapping returns the Mapping field if non-nil, zero value otherwise.

### GetMappingOk

`func (o *Profile1) GetMappingOk() (*map[string]string, bool)`

GetMappingOk returns a tuple with the Mapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapping

`func (o *Profile1) SetMapping(v map[string]string)`

SetMapping sets Mapping field to given value.

### HasMapping

`func (o *Profile1) HasMapping() bool`

HasMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


