# DriveapiFileInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**IsCreator** | Pointer to **bool** |  | [optional] 
**CreateTime** | Pointer to **int32** |  | [optional] 
**CreatorName** | Pointer to **string** |  | [optional] 
**IsOwner** | Pointer to **bool** |  | [optional] 
**OwnerName** | Pointer to **string** |  | [optional] 
**LastModifyTime** | Pointer to **int32** |  | [optional] 
**LastModifyName** | Pointer to **string** |  | [optional] 
**RelativeFiles** | Pointer to [**[]DriveapiRelativeFiles**](DriveapiRelativeFiles.md) |  | [optional] 
**FormCollectingStatus** | Pointer to **string** |  | [optional] 
**FormCollectingEndTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewDriveapiFileInfo

`func NewDriveapiFileInfo() *DriveapiFileInfo`

NewDriveapiFileInfo instantiates a new DriveapiFileInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriveapiFileInfoWithDefaults

`func NewDriveapiFileInfoWithDefaults() *DriveapiFileInfo`

NewDriveapiFileInfoWithDefaults instantiates a new DriveapiFileInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DriveapiFileInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DriveapiFileInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DriveapiFileInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DriveapiFileInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *DriveapiFileInfo) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DriveapiFileInfo) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DriveapiFileInfo) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DriveapiFileInfo) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *DriveapiFileInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DriveapiFileInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DriveapiFileInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DriveapiFileInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *DriveapiFileInfo) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DriveapiFileInfo) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DriveapiFileInfo) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DriveapiFileInfo) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetStatus

`func (o *DriveapiFileInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DriveapiFileInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DriveapiFileInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DriveapiFileInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIsCreator

`func (o *DriveapiFileInfo) GetIsCreator() bool`

GetIsCreator returns the IsCreator field if non-nil, zero value otherwise.

### GetIsCreatorOk

`func (o *DriveapiFileInfo) GetIsCreatorOk() (*bool, bool)`

GetIsCreatorOk returns a tuple with the IsCreator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCreator

`func (o *DriveapiFileInfo) SetIsCreator(v bool)`

SetIsCreator sets IsCreator field to given value.

### HasIsCreator

`func (o *DriveapiFileInfo) HasIsCreator() bool`

HasIsCreator returns a boolean if a field has been set.

### GetCreateTime

`func (o *DriveapiFileInfo) GetCreateTime() int32`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *DriveapiFileInfo) GetCreateTimeOk() (*int32, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *DriveapiFileInfo) SetCreateTime(v int32)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *DriveapiFileInfo) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCreatorName

`func (o *DriveapiFileInfo) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *DriveapiFileInfo) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *DriveapiFileInfo) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.

### HasCreatorName

`func (o *DriveapiFileInfo) HasCreatorName() bool`

HasCreatorName returns a boolean if a field has been set.

### GetIsOwner

`func (o *DriveapiFileInfo) GetIsOwner() bool`

GetIsOwner returns the IsOwner field if non-nil, zero value otherwise.

### GetIsOwnerOk

`func (o *DriveapiFileInfo) GetIsOwnerOk() (*bool, bool)`

GetIsOwnerOk returns a tuple with the IsOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOwner

`func (o *DriveapiFileInfo) SetIsOwner(v bool)`

SetIsOwner sets IsOwner field to given value.

### HasIsOwner

`func (o *DriveapiFileInfo) HasIsOwner() bool`

HasIsOwner returns a boolean if a field has been set.

### GetOwnerName

`func (o *DriveapiFileInfo) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *DriveapiFileInfo) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *DriveapiFileInfo) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.

### HasOwnerName

`func (o *DriveapiFileInfo) HasOwnerName() bool`

HasOwnerName returns a boolean if a field has been set.

### GetLastModifyTime

`func (o *DriveapiFileInfo) GetLastModifyTime() int32`

GetLastModifyTime returns the LastModifyTime field if non-nil, zero value otherwise.

### GetLastModifyTimeOk

`func (o *DriveapiFileInfo) GetLastModifyTimeOk() (*int32, bool)`

GetLastModifyTimeOk returns a tuple with the LastModifyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifyTime

`func (o *DriveapiFileInfo) SetLastModifyTime(v int32)`

SetLastModifyTime sets LastModifyTime field to given value.

### HasLastModifyTime

`func (o *DriveapiFileInfo) HasLastModifyTime() bool`

HasLastModifyTime returns a boolean if a field has been set.

### GetLastModifyName

`func (o *DriveapiFileInfo) GetLastModifyName() string`

GetLastModifyName returns the LastModifyName field if non-nil, zero value otherwise.

### GetLastModifyNameOk

`func (o *DriveapiFileInfo) GetLastModifyNameOk() (*string, bool)`

GetLastModifyNameOk returns a tuple with the LastModifyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifyName

`func (o *DriveapiFileInfo) SetLastModifyName(v string)`

SetLastModifyName sets LastModifyName field to given value.

### HasLastModifyName

`func (o *DriveapiFileInfo) HasLastModifyName() bool`

HasLastModifyName returns a boolean if a field has been set.

### GetRelativeFiles

`func (o *DriveapiFileInfo) GetRelativeFiles() []DriveapiRelativeFiles`

GetRelativeFiles returns the RelativeFiles field if non-nil, zero value otherwise.

### GetRelativeFilesOk

`func (o *DriveapiFileInfo) GetRelativeFilesOk() (*[]DriveapiRelativeFiles, bool)`

GetRelativeFilesOk returns a tuple with the RelativeFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeFiles

`func (o *DriveapiFileInfo) SetRelativeFiles(v []DriveapiRelativeFiles)`

SetRelativeFiles sets RelativeFiles field to given value.

### HasRelativeFiles

`func (o *DriveapiFileInfo) HasRelativeFiles() bool`

HasRelativeFiles returns a boolean if a field has been set.

### GetFormCollectingStatus

`func (o *DriveapiFileInfo) GetFormCollectingStatus() string`

GetFormCollectingStatus returns the FormCollectingStatus field if non-nil, zero value otherwise.

### GetFormCollectingStatusOk

`func (o *DriveapiFileInfo) GetFormCollectingStatusOk() (*string, bool)`

GetFormCollectingStatusOk returns a tuple with the FormCollectingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormCollectingStatus

`func (o *DriveapiFileInfo) SetFormCollectingStatus(v string)`

SetFormCollectingStatus sets FormCollectingStatus field to given value.

### HasFormCollectingStatus

`func (o *DriveapiFileInfo) HasFormCollectingStatus() bool`

HasFormCollectingStatus returns a boolean if a field has been set.

### GetFormCollectingEndTime

`func (o *DriveapiFileInfo) GetFormCollectingEndTime() int64`

GetFormCollectingEndTime returns the FormCollectingEndTime field if non-nil, zero value otherwise.

### GetFormCollectingEndTimeOk

`func (o *DriveapiFileInfo) GetFormCollectingEndTimeOk() (*int64, bool)`

GetFormCollectingEndTimeOk returns a tuple with the FormCollectingEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormCollectingEndTime

`func (o *DriveapiFileInfo) SetFormCollectingEndTime(v int64)`

SetFormCollectingEndTime sets FormCollectingEndTime field to given value.

### HasFormCollectingEndTime

`func (o *DriveapiFileInfo) HasFormCollectingEndTime() bool`

HasFormCollectingEndTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


