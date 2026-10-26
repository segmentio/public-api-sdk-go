# GroupSourceSettingsV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowUnplannedTraits** | Pointer to **bool** | Enable to allow unplanned group traits.  Config API note: equal to &#x60;allowUnplannedGroupTraits&#x60;. | [optional] 
**AllowTraitsOnViolations** | Pointer to **bool** | Enable to allow group traits on violations.  Config API note: equal to &#x60;allowGroupTraitsOnViolations&#x60;. | [optional] 
**CommonEventOnViolations** | Pointer to **string** | The common group event on violations.  Config API note: equal to &#x60;commonGroupEventOnViolations&#x60;. | [optional] 

## Methods

### NewGroupSourceSettingsV1

`func NewGroupSourceSettingsV1() *GroupSourceSettingsV1`

NewGroupSourceSettingsV1 instantiates a new GroupSourceSettingsV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupSourceSettingsV1WithDefaults

`func NewGroupSourceSettingsV1WithDefaults() *GroupSourceSettingsV1`

NewGroupSourceSettingsV1WithDefaults instantiates a new GroupSourceSettingsV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowUnplannedTraits

`func (o *GroupSourceSettingsV1) GetAllowUnplannedTraits() bool`

GetAllowUnplannedTraits returns the AllowUnplannedTraits field if non-nil, zero value otherwise.

### GetAllowUnplannedTraitsOk

`func (o *GroupSourceSettingsV1) GetAllowUnplannedTraitsOk() (*bool, bool)`

GetAllowUnplannedTraitsOk returns a tuple with the AllowUnplannedTraits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUnplannedTraits

`func (o *GroupSourceSettingsV1) SetAllowUnplannedTraits(v bool)`

SetAllowUnplannedTraits sets AllowUnplannedTraits field to given value.

### HasAllowUnplannedTraits

`func (o *GroupSourceSettingsV1) HasAllowUnplannedTraits() bool`

HasAllowUnplannedTraits returns a boolean if a field has been set.

### GetAllowTraitsOnViolations

`func (o *GroupSourceSettingsV1) GetAllowTraitsOnViolations() bool`

GetAllowTraitsOnViolations returns the AllowTraitsOnViolations field if non-nil, zero value otherwise.

### GetAllowTraitsOnViolationsOk

`func (o *GroupSourceSettingsV1) GetAllowTraitsOnViolationsOk() (*bool, bool)`

GetAllowTraitsOnViolationsOk returns a tuple with the AllowTraitsOnViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTraitsOnViolations

`func (o *GroupSourceSettingsV1) SetAllowTraitsOnViolations(v bool)`

SetAllowTraitsOnViolations sets AllowTraitsOnViolations field to given value.

### HasAllowTraitsOnViolations

`func (o *GroupSourceSettingsV1) HasAllowTraitsOnViolations() bool`

HasAllowTraitsOnViolations returns a boolean if a field has been set.

### GetCommonEventOnViolations

`func (o *GroupSourceSettingsV1) GetCommonEventOnViolations() string`

GetCommonEventOnViolations returns the CommonEventOnViolations field if non-nil, zero value otherwise.

### GetCommonEventOnViolationsOk

`func (o *GroupSourceSettingsV1) GetCommonEventOnViolationsOk() (*string, bool)`

GetCommonEventOnViolationsOk returns a tuple with the CommonEventOnViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonEventOnViolations

`func (o *GroupSourceSettingsV1) SetCommonEventOnViolations(v string)`

SetCommonEventOnViolations sets CommonEventOnViolations field to given value.

### HasCommonEventOnViolations

`func (o *GroupSourceSettingsV1) HasCommonEventOnViolations() bool`

HasCommonEventOnViolations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


