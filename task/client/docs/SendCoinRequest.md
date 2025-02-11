# SendCoinRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToUser** | **string** | Имя пользователя, которому нужно отправить монеты. | 
**Amount** | **int32** | Количество монет, которые необходимо отправить. | 

## Methods

### NewSendCoinRequest

`func NewSendCoinRequest(toUser string, amount int32, ) *SendCoinRequest`

NewSendCoinRequest instantiates a new SendCoinRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendCoinRequestWithDefaults

`func NewSendCoinRequestWithDefaults() *SendCoinRequest`

NewSendCoinRequestWithDefaults instantiates a new SendCoinRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToUser

`func (o *SendCoinRequest) GetToUser() string`

GetToUser returns the ToUser field if non-nil, zero value otherwise.

### GetToUserOk

`func (o *SendCoinRequest) GetToUserOk() (*string, bool)`

GetToUserOk returns a tuple with the ToUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToUser

`func (o *SendCoinRequest) SetToUser(v string)`

SetToUser sets ToUser field to given value.


### GetAmount

`func (o *SendCoinRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SendCoinRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SendCoinRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


