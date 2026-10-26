# PermissionResourceV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of this resource. | 
**Type** | **string** | The type for this resource. | 
**Labels** | Pointer to [**[]AllowedLabelBeta**](AllowedLabelBeta.md) | The labels that further refine access to this resource. Labels are exclusive to Workspace-level permissions. | [optional] 

## Methods

### NewPermissionResourceV1

`func NewPermissionResourceV1(id string, type_ string, ) *PermissionResourceV1`

NewPermissionResourceV1 instantiates a new PermissionResourceV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionResourceV1WithDefaults

`func NewPermissionResourceV1WithDefaults() *PermissionResourceV1`

NewPermissionResourceV1WithDefaults instantiates a new PermissionResourceV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PermissionResourceV1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PermissionResourceV1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PermissionResourceV1) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *PermissionResourceV1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PermissionResourceV1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PermissionResourceV1) SetType(v string)`

SetType sets Type field to given value.


### GetLabels

`func (o *PermissionResourceV1) GetLabels() []AllowedLabelBeta`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PermissionResourceV1) GetLabelsOk() (*[]AllowedLabelBeta, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PermissionResourceV1) SetLabels(v []AllowedLabelBeta)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *PermissionResourceV1) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


