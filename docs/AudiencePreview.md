# AudiencePreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for tracking and retrieving results of an audience preview. | 
**AudienceType** | **string** | The audience type of the preview.  Possible values: USERS, ACCOUNTS, LINKED. | 
**Definition** | [**AudienceDefinition**](AudienceDefinition.md) |  | 
**Options** | [**ReadAudiencePreviewOptions**](ReadAudiencePreviewOptions.md) |  | 
**Status** | **string** | Status for the audience preview. | 
**Results** | Pointer to [**[]AudiencePreviewResult**](AudiencePreviewResult.md) | Sampled result membership for the audience preview. Only has a value if the status is &#39;COMPLETED&#39;. | [optional] 
**Size** | Pointer to [**AudienceSize**](AudienceSize.md) |  | [optional] 
**FailureReason** | Pointer to **string** | Explanation of why the audience preview failed. Only has a value if status is &#39;FAILED&#39;. | [optional] 

## Methods

### NewAudiencePreview

`func NewAudiencePreview(id string, audienceType string, definition AudienceDefinition, options ReadAudiencePreviewOptions, status string, ) *AudiencePreview`

NewAudiencePreview instantiates a new AudiencePreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudiencePreviewWithDefaults

`func NewAudiencePreviewWithDefaults() *AudiencePreview`

NewAudiencePreviewWithDefaults instantiates a new AudiencePreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AudiencePreview) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AudiencePreview) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AudiencePreview) SetId(v string)`

SetId sets Id field to given value.


### GetAudienceType

`func (o *AudiencePreview) GetAudienceType() string`

GetAudienceType returns the AudienceType field if non-nil, zero value otherwise.

### GetAudienceTypeOk

`func (o *AudiencePreview) GetAudienceTypeOk() (*string, bool)`

GetAudienceTypeOk returns a tuple with the AudienceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceType

`func (o *AudiencePreview) SetAudienceType(v string)`

SetAudienceType sets AudienceType field to given value.


### GetDefinition

`func (o *AudiencePreview) GetDefinition() AudienceDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *AudiencePreview) GetDefinitionOk() (*AudienceDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *AudiencePreview) SetDefinition(v AudienceDefinition)`

SetDefinition sets Definition field to given value.


### GetOptions

`func (o *AudiencePreview) GetOptions() ReadAudiencePreviewOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AudiencePreview) GetOptionsOk() (*ReadAudiencePreviewOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AudiencePreview) SetOptions(v ReadAudiencePreviewOptions)`

SetOptions sets Options field to given value.


### GetStatus

`func (o *AudiencePreview) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AudiencePreview) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AudiencePreview) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetResults

`func (o *AudiencePreview) GetResults() []AudiencePreviewResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *AudiencePreview) GetResultsOk() (*[]AudiencePreviewResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *AudiencePreview) SetResults(v []AudiencePreviewResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *AudiencePreview) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetSize

`func (o *AudiencePreview) GetSize() AudienceSize`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AudiencePreview) GetSizeOk() (*AudienceSize, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AudiencePreview) SetSize(v AudienceSize)`

SetSize sets Size field to given value.

### HasSize

`func (o *AudiencePreview) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetFailureReason

`func (o *AudiencePreview) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *AudiencePreview) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *AudiencePreview) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *AudiencePreview) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


