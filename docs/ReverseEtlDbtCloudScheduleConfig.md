# ReverseEtlDbtCloudScheduleConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | **string** | The dbt Cloud job used to trigger a sync for a Reverse ETL Connection. | 
**AccountId** | **string** | The dbt Cloud account where the job belongs to. | 

## Methods

### NewReverseEtlDbtCloudScheduleConfig

`func NewReverseEtlDbtCloudScheduleConfig(jobId string, accountId string, ) *ReverseEtlDbtCloudScheduleConfig`

NewReverseEtlDbtCloudScheduleConfig instantiates a new ReverseEtlDbtCloudScheduleConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReverseEtlDbtCloudScheduleConfigWithDefaults

`func NewReverseEtlDbtCloudScheduleConfigWithDefaults() *ReverseEtlDbtCloudScheduleConfig`

NewReverseEtlDbtCloudScheduleConfigWithDefaults instantiates a new ReverseEtlDbtCloudScheduleConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *ReverseEtlDbtCloudScheduleConfig) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *ReverseEtlDbtCloudScheduleConfig) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *ReverseEtlDbtCloudScheduleConfig) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetAccountId

`func (o *ReverseEtlDbtCloudScheduleConfig) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ReverseEtlDbtCloudScheduleConfig) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ReverseEtlDbtCloudScheduleConfig) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


