# LivePluginsAlpha

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Live Plugin id. | 
**SourceId** | **string** | The Source id. | 
**CreatedAt** | **string** | Creation date. | 
**CreatedBy** | **string** | Creating user&#39;s id. | 
**DownloadURL** | **string** | The CDN URL that can be used to fetch your latest Live Plugin bundle. | 
**Version** | **float32** | Revision number associated with the latest Live Plugin. | 
**Code** | Pointer to **NullableString** | The code of the Live Plugin. | [optional] 

## Methods

### NewLivePluginsAlpha

`func NewLivePluginsAlpha(id string, sourceId string, createdAt string, createdBy string, downloadURL string, version float32, ) *LivePluginsAlpha`

NewLivePluginsAlpha instantiates a new LivePluginsAlpha object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLivePluginsAlphaWithDefaults

`func NewLivePluginsAlphaWithDefaults() *LivePluginsAlpha`

NewLivePluginsAlphaWithDefaults instantiates a new LivePluginsAlpha object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LivePluginsAlpha) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LivePluginsAlpha) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LivePluginsAlpha) SetId(v string)`

SetId sets Id field to given value.


### GetSourceId

`func (o *LivePluginsAlpha) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *LivePluginsAlpha) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *LivePluginsAlpha) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetCreatedAt

`func (o *LivePluginsAlpha) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LivePluginsAlpha) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LivePluginsAlpha) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *LivePluginsAlpha) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *LivePluginsAlpha) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *LivePluginsAlpha) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDownloadURL

`func (o *LivePluginsAlpha) GetDownloadURL() string`

GetDownloadURL returns the DownloadURL field if non-nil, zero value otherwise.

### GetDownloadURLOk

`func (o *LivePluginsAlpha) GetDownloadURLOk() (*string, bool)`

GetDownloadURLOk returns a tuple with the DownloadURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadURL

`func (o *LivePluginsAlpha) SetDownloadURL(v string)`

SetDownloadURL sets DownloadURL field to given value.


### GetVersion

`func (o *LivePluginsAlpha) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *LivePluginsAlpha) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *LivePluginsAlpha) SetVersion(v float32)`

SetVersion sets Version field to given value.


### GetCode

`func (o *LivePluginsAlpha) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *LivePluginsAlpha) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *LivePluginsAlpha) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *LivePluginsAlpha) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *LivePluginsAlpha) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *LivePluginsAlpha) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


