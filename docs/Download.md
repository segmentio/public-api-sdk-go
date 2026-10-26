# Download

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Urls** | **[]string** | List of presigned URLs from S3. | 

## Methods

### NewDownload

`func NewDownload(urls []string, ) *Download`

NewDownload instantiates a new Download object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDownloadWithDefaults

`func NewDownloadWithDefaults() *Download`

NewDownloadWithDefaults instantiates a new Download object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrls

`func (o *Download) GetUrls() []string`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *Download) GetUrlsOk() (*[]string, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *Download) SetUrls(v []string)`

SetUrls sets Urls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


