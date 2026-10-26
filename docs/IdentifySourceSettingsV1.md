# IdentifySourceSettingsV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowUnplannedTraits** | Pointer to **bool** | Enable to allow unplanned identify traits.  Config API note: equal to &#x60;allowUnplannedIdentifyTraits&#x60;. | [optional] 
**AllowTraitsOnViolations** | Pointer to **bool** | Enable to allow identify traits on violations.  Config API note: equal to &#x60;allowIdentifyTraitsOnViolations&#x60;. | [optional] 
**CommonEventOnViolations** | Pointer to **string** | The common identify event on violations.  Config API note: equal to &#x60;commonIdentifyEventOnViolations&#x60;. | [optional] 

## Methods

### NewIdentifySourceSettingsV1

`func NewIdentifySourceSettingsV1() *IdentifySourceSettingsV1`

NewIdentifySourceSettingsV1 instantiates a new IdentifySourceSettingsV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentifySourceSettingsV1WithDefaults

`func NewIdentifySourceSettingsV1WithDefaults() *IdentifySourceSettingsV1`

NewIdentifySourceSettingsV1WithDefaults instantiates a new IdentifySourceSettingsV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowUnplannedTraits

`func (o *IdentifySourceSettingsV1) GetAllowUnplannedTraits() bool`

GetAllowUnplannedTraits returns the AllowUnplannedTraits field if non-nil, zero value otherwise.

### GetAllowUnplannedTraitsOk

`func (o *IdentifySourceSettingsV1) GetAllowUnplannedTraitsOk() (*bool, bool)`

GetAllowUnplannedTraitsOk returns a tuple with the AllowUnplannedTraits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUnplannedTraits

`func (o *IdentifySourceSettingsV1) SetAllowUnplannedTraits(v bool)`

SetAllowUnplannedTraits sets AllowUnplannedTraits field to given value.

### HasAllowUnplannedTraits

`func (o *IdentifySourceSettingsV1) HasAllowUnplannedTraits() bool`

HasAllowUnplannedTraits returns a boolean if a field has been set.

### GetAllowTraitsOnViolations

`func (o *IdentifySourceSettingsV1) GetAllowTraitsOnViolations() bool`

GetAllowTraitsOnViolations returns the AllowTraitsOnViolations field if non-nil, zero value otherwise.

### GetAllowTraitsOnViolationsOk

`func (o *IdentifySourceSettingsV1) GetAllowTraitsOnViolationsOk() (*bool, bool)`

GetAllowTraitsOnViolationsOk returns a tuple with the AllowTraitsOnViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTraitsOnViolations

`func (o *IdentifySourceSettingsV1) SetAllowTraitsOnViolations(v bool)`

SetAllowTraitsOnViolations sets AllowTraitsOnViolations field to given value.

### HasAllowTraitsOnViolations

`func (o *IdentifySourceSettingsV1) HasAllowTraitsOnViolations() bool`

HasAllowTraitsOnViolations returns a boolean if a field has been set.

### GetCommonEventOnViolations

`func (o *IdentifySourceSettingsV1) GetCommonEventOnViolations() string`

GetCommonEventOnViolations returns the CommonEventOnViolations field if non-nil, zero value otherwise.

### GetCommonEventOnViolationsOk

`func (o *IdentifySourceSettingsV1) GetCommonEventOnViolationsOk() (*string, bool)`

GetCommonEventOnViolationsOk returns a tuple with the CommonEventOnViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonEventOnViolations

`func (o *IdentifySourceSettingsV1) SetCommonEventOnViolations(v string)`

SetCommonEventOnViolations sets CommonEventOnViolations field to given value.

### HasCommonEventOnViolations

`func (o *IdentifySourceSettingsV1) HasCommonEventOnViolations() bool`

HasCommonEventOnViolations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


