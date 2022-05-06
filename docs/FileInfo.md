# FileInfo

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
**RelativeFiles** | Pointer to [**[]RelativeFiles**](RelativeFiles.md) |  | [optional] 
**FormCollectingStatus** | Pointer to **string** |  | [optional] 
**FormCollectingEndTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewFileInfo

`func NewFileInfo() *FileInfo`

NewFileInfo instantiates a new FileInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileInfoWithDefaults

`func NewFileInfoWithDefaults() *FileInfo`

NewFileInfoWithDefaults instantiates a new FileInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FileInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FileInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *FileInfo) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FileInfo) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FileInfo) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FileInfo) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *FileInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FileInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FileInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FileInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *FileInfo) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FileInfo) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FileInfo) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *FileInfo) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetStatus

`func (o *FileInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FileInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FileInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FileInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIsCreator

`func (o *FileInfo) GetIsCreator() bool`

GetIsCreator returns the IsCreator field if non-nil, zero value otherwise.

### GetIsCreatorOk

`func (o *FileInfo) GetIsCreatorOk() (*bool, bool)`

GetIsCreatorOk returns a tuple with the IsCreator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCreator

`func (o *FileInfo) SetIsCreator(v bool)`

SetIsCreator sets IsCreator field to given value.

### HasIsCreator

`func (o *FileInfo) HasIsCreator() bool`

HasIsCreator returns a boolean if a field has been set.

### GetCreateTime

`func (o *FileInfo) GetCreateTime() int32`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *FileInfo) GetCreateTimeOk() (*int32, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *FileInfo) SetCreateTime(v int32)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *FileInfo) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCreatorName

`func (o *FileInfo) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *FileInfo) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *FileInfo) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.

### HasCreatorName

`func (o *FileInfo) HasCreatorName() bool`

HasCreatorName returns a boolean if a field has been set.

### GetIsOwner

`func (o *FileInfo) GetIsOwner() bool`

GetIsOwner returns the IsOwner field if non-nil, zero value otherwise.

### GetIsOwnerOk

`func (o *FileInfo) GetIsOwnerOk() (*bool, bool)`

GetIsOwnerOk returns a tuple with the IsOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOwner

`func (o *FileInfo) SetIsOwner(v bool)`

SetIsOwner sets IsOwner field to given value.

### HasIsOwner

`func (o *FileInfo) HasIsOwner() bool`

HasIsOwner returns a boolean if a field has been set.

### GetOwnerName

`func (o *FileInfo) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *FileInfo) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *FileInfo) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.

### HasOwnerName

`func (o *FileInfo) HasOwnerName() bool`

HasOwnerName returns a boolean if a field has been set.

### GetLastModifyTime

`func (o *FileInfo) GetLastModifyTime() int32`

GetLastModifyTime returns the LastModifyTime field if non-nil, zero value otherwise.

### GetLastModifyTimeOk

`func (o *FileInfo) GetLastModifyTimeOk() (*int32, bool)`

GetLastModifyTimeOk returns a tuple with the LastModifyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifyTime

`func (o *FileInfo) SetLastModifyTime(v int32)`

SetLastModifyTime sets LastModifyTime field to given value.

### HasLastModifyTime

`func (o *FileInfo) HasLastModifyTime() bool`

HasLastModifyTime returns a boolean if a field has been set.

### GetLastModifyName

`func (o *FileInfo) GetLastModifyName() string`

GetLastModifyName returns the LastModifyName field if non-nil, zero value otherwise.

### GetLastModifyNameOk

`func (o *FileInfo) GetLastModifyNameOk() (*string, bool)`

GetLastModifyNameOk returns a tuple with the LastModifyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifyName

`func (o *FileInfo) SetLastModifyName(v string)`

SetLastModifyName sets LastModifyName field to given value.

### HasLastModifyName

`func (o *FileInfo) HasLastModifyName() bool`

HasLastModifyName returns a boolean if a field has been set.

### GetRelativeFiles

`func (o *FileInfo) GetRelativeFiles() []RelativeFiles`

GetRelativeFiles returns the RelativeFiles field if non-nil, zero value otherwise.

### GetRelativeFilesOk

`func (o *FileInfo) GetRelativeFilesOk() (*[]RelativeFiles, bool)`

GetRelativeFilesOk returns a tuple with the RelativeFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeFiles

`func (o *FileInfo) SetRelativeFiles(v []RelativeFiles)`

SetRelativeFiles sets RelativeFiles field to given value.

### HasRelativeFiles

`func (o *FileInfo) HasRelativeFiles() bool`

HasRelativeFiles returns a boolean if a field has been set.

### GetFormCollectingStatus

`func (o *FileInfo) GetFormCollectingStatus() string`

GetFormCollectingStatus returns the FormCollectingStatus field if non-nil, zero value otherwise.

### GetFormCollectingStatusOk

`func (o *FileInfo) GetFormCollectingStatusOk() (*string, bool)`

GetFormCollectingStatusOk returns a tuple with the FormCollectingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormCollectingStatus

`func (o *FileInfo) SetFormCollectingStatus(v string)`

SetFormCollectingStatus sets FormCollectingStatus field to given value.

### HasFormCollectingStatus

`func (o *FileInfo) HasFormCollectingStatus() bool`

HasFormCollectingStatus returns a boolean if a field has been set.

### GetFormCollectingEndTime

`func (o *FileInfo) GetFormCollectingEndTime() int64`

GetFormCollectingEndTime returns the FormCollectingEndTime field if non-nil, zero value otherwise.

### GetFormCollectingEndTimeOk

`func (o *FileInfo) GetFormCollectingEndTimeOk() (*int64, bool)`

GetFormCollectingEndTimeOk returns a tuple with the FormCollectingEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormCollectingEndTime

`func (o *FileInfo) SetFormCollectingEndTime(v int64)`

SetFormCollectingEndTime sets FormCollectingEndTime field to given value.

### HasFormCollectingEndTime

`func (o *FileInfo) HasFormCollectingEndTime() bool`

HasFormCollectingEndTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


