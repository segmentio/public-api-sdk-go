# DestinationFilterActionV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The kind of Transformation to apply to any matched properties. | 
**Fields** | Pointer to **map[string]interface{}** | A dictionary of paths to object keys that this filter applies to.  The literal string &#39;&#39; represents the top level of the object. | [optional] 
**Percent** | Pointer to **float32** | A decimal between 0 and 1 used for &#39;sample&#39; type events and influences the likelihood of sampling to occur. | [optional] 
**Path** | Pointer to **string** | The JSON path to a property within a payload object from which Segment generates a deterministic sampling rate. | [optional] 

## Methods

### NewDestinationFilterActionV1

`func NewDestinationFilterActionV1(type_ string, ) *DestinationFilterActionV1`

NewDestinationFilterActionV1 instantiates a new DestinationFilterActionV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationFilterActionV1WithDefaults

`func NewDestinationFilterActionV1WithDefaults() *DestinationFilterActionV1`

NewDestinationFilterActionV1WithDefaults instantiates a new DestinationFilterActionV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DestinationFilterActionV1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DestinationFilterActionV1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DestinationFilterActionV1) SetType(v string)`

SetType sets Type field to given value.


### GetFields

`func (o *DestinationFilterActionV1) GetFields() map[string]interface{}`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *DestinationFilterActionV1) GetFieldsOk() (*map[string]interface{}, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *DestinationFilterActionV1) SetFields(v map[string]interface{})`

SetFields sets Fields field to given value.

### HasFields

`func (o *DestinationFilterActionV1) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetPercent

`func (o *DestinationFilterActionV1) GetPercent() float32`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *DestinationFilterActionV1) GetPercentOk() (*float32, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *DestinationFilterActionV1) SetPercent(v float32)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *DestinationFilterActionV1) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetPath

`func (o *DestinationFilterActionV1) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DestinationFilterActionV1) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DestinationFilterActionV1) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DestinationFilterActionV1) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


