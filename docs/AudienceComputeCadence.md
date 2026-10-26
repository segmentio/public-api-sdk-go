# AudienceComputeCadence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The cadence type on which the audience&#39;s membership is computed. If &#39;BATCH&#39;, the audience is computed on a periodic basis. If &#39;REALTIME&#39;, the audience is continously computed. | 

## Methods

### NewAudienceComputeCadence

`func NewAudienceComputeCadence(type_ string, ) *AudienceComputeCadence`

NewAudienceComputeCadence instantiates a new AudienceComputeCadence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceComputeCadenceWithDefaults

`func NewAudienceComputeCadenceWithDefaults() *AudienceComputeCadence`

NewAudienceComputeCadenceWithDefaults instantiates a new AudienceComputeCadence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AudienceComputeCadence) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AudienceComputeCadence) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AudienceComputeCadence) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


