# Configuration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdConfiguration** | Pointer to **int32** |  | [optional] 
**ConfigurationScriptType** | Pointer to **string** |  | [optional] 
**Script** | **string** |  | 
**ScriptName** | Pointer to **string** |  | [optional] 

## Methods

### NewConfiguration

`func NewConfiguration(script string, ) *Configuration`

NewConfiguration instantiates a new Configuration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationWithDefaults

`func NewConfigurationWithDefaults() *Configuration`

NewConfigurationWithDefaults instantiates a new Configuration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdConfiguration

`func (o *Configuration) GetIdConfiguration() int32`

GetIdConfiguration returns the IdConfiguration field if non-nil, zero value otherwise.

### GetIdConfigurationOk

`func (o *Configuration) GetIdConfigurationOk() (*int32, bool)`

GetIdConfigurationOk returns a tuple with the IdConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdConfiguration

`func (o *Configuration) SetIdConfiguration(v int32)`

SetIdConfiguration sets IdConfiguration field to given value.

### HasIdConfiguration

`func (o *Configuration) HasIdConfiguration() bool`

HasIdConfiguration returns a boolean if a field has been set.

### GetConfigurationScriptType

`func (o *Configuration) GetConfigurationScriptType() string`

GetConfigurationScriptType returns the ConfigurationScriptType field if non-nil, zero value otherwise.

### GetConfigurationScriptTypeOk

`func (o *Configuration) GetConfigurationScriptTypeOk() (*string, bool)`

GetConfigurationScriptTypeOk returns a tuple with the ConfigurationScriptType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationScriptType

`func (o *Configuration) SetConfigurationScriptType(v string)`

SetConfigurationScriptType sets ConfigurationScriptType field to given value.

### HasConfigurationScriptType

`func (o *Configuration) HasConfigurationScriptType() bool`

HasConfigurationScriptType returns a boolean if a field has been set.

### GetScript

`func (o *Configuration) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *Configuration) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *Configuration) SetScript(v string)`

SetScript sets Script field to given value.


### GetScriptName

`func (o *Configuration) GetScriptName() string`

GetScriptName returns the ScriptName field if non-nil, zero value otherwise.

### GetScriptNameOk

`func (o *Configuration) GetScriptNameOk() (*string, bool)`

GetScriptNameOk returns a tuple with the ScriptName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptName

`func (o *Configuration) SetScriptName(v string)`

SetScriptName sets ScriptName field to given value.

### HasScriptName

`func (o *Configuration) HasScriptName() bool`

HasScriptName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


