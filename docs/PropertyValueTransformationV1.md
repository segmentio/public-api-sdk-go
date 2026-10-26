# PropertyValueTransformationV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropertyPaths** | **[]string** | The property paths. The maximum number of paths is 10. | 
**PropertyValue** | **string** | The new value of the property paths. | 

## Methods

### NewPropertyValueTransformationV1

`func NewPropertyValueTransformationV1(propertyPaths []string, propertyValue string, ) *PropertyValueTransformationV1`

NewPropertyValueTransformationV1 instantiates a new PropertyValueTransformationV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyValueTransformationV1WithDefaults

`func NewPropertyValueTransformationV1WithDefaults() *PropertyValueTransformationV1`

NewPropertyValueTransformationV1WithDefaults instantiates a new PropertyValueTransformationV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropertyPaths

`func (o *PropertyValueTransformationV1) GetPropertyPaths() []string`

GetPropertyPaths returns the PropertyPaths field if non-nil, zero value otherwise.

### GetPropertyPathsOk

`func (o *PropertyValueTransformationV1) GetPropertyPathsOk() (*[]string, bool)`

GetPropertyPathsOk returns a tuple with the PropertyPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyPaths

`func (o *PropertyValueTransformationV1) SetPropertyPaths(v []string)`

SetPropertyPaths sets PropertyPaths field to given value.


### GetPropertyValue

`func (o *PropertyValueTransformationV1) GetPropertyValue() string`

GetPropertyValue returns the PropertyValue field if non-nil, zero value otherwise.

### GetPropertyValueOk

`func (o *PropertyValueTransformationV1) GetPropertyValueOk() (*string, bool)`

GetPropertyValueOk returns a tuple with the PropertyValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyValue

`func (o *PropertyValueTransformationV1) SetPropertyValue(v string)`

SetPropertyValue sets PropertyValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


