# DeploymentAssetMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | Pointer to **string** |  | [optional] 
**TotalDeploymentRunDuration** | Pointer to **int64** |  | [optional] 
**AverageDeploymentRunDuration** | Pointer to **int64** |  | [optional] 
**MedianDeploymentRunDuration** | Pointer to **int64** |  | [optional] 
**LongestDeploymentRunDuration** | Pointer to **int64** |  | [optional] 
**ShortestDeploymentRunDuration** | Pointer to **int64** |  | [optional] 
**TotalDeploymentRuns** | Pointer to **int32** |  | [optional] 
**TotalDeploymentRunSuccess** | Pointer to **int32** |  | [optional] 
**TotalDeploymentRunError** | Pointer to **int32** |  | [optional] 
**TotalDeploymentRunCancel** | Pointer to **int32** |  | [optional] 
**TotalDeploymentRunUnknown** | Pointer to **int32** |  | [optional] 
**TotalVirtualMachines** | Pointer to **int32** |  | [optional] 

## Methods

### NewDeploymentAssetMetric

`func NewDeploymentAssetMetric() *DeploymentAssetMetric`

NewDeploymentAssetMetric instantiates a new DeploymentAssetMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentAssetMetricWithDefaults

`func NewDeploymentAssetMetricWithDefaults() *DeploymentAssetMetric`

NewDeploymentAssetMetricWithDefaults instantiates a new DeploymentAssetMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *DeploymentAssetMetric) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *DeploymentAssetMetric) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *DeploymentAssetMetric) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *DeploymentAssetMetric) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### GetTotalDeploymentRunDuration

`func (o *DeploymentAssetMetric) GetTotalDeploymentRunDuration() int64`

GetTotalDeploymentRunDuration returns the TotalDeploymentRunDuration field if non-nil, zero value otherwise.

### GetTotalDeploymentRunDurationOk

`func (o *DeploymentAssetMetric) GetTotalDeploymentRunDurationOk() (*int64, bool)`

GetTotalDeploymentRunDurationOk returns a tuple with the TotalDeploymentRunDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDeploymentRunDuration

`func (o *DeploymentAssetMetric) SetTotalDeploymentRunDuration(v int64)`

SetTotalDeploymentRunDuration sets TotalDeploymentRunDuration field to given value.

### HasTotalDeploymentRunDuration

`func (o *DeploymentAssetMetric) HasTotalDeploymentRunDuration() bool`

HasTotalDeploymentRunDuration returns a boolean if a field has been set.

### GetAverageDeploymentRunDuration

`func (o *DeploymentAssetMetric) GetAverageDeploymentRunDuration() int64`

GetAverageDeploymentRunDuration returns the AverageDeploymentRunDuration field if non-nil, zero value otherwise.

### GetAverageDeploymentRunDurationOk

`func (o *DeploymentAssetMetric) GetAverageDeploymentRunDurationOk() (*int64, bool)`

GetAverageDeploymentRunDurationOk returns a tuple with the AverageDeploymentRunDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageDeploymentRunDuration

`func (o *DeploymentAssetMetric) SetAverageDeploymentRunDuration(v int64)`

SetAverageDeploymentRunDuration sets AverageDeploymentRunDuration field to given value.

### HasAverageDeploymentRunDuration

`func (o *DeploymentAssetMetric) HasAverageDeploymentRunDuration() bool`

HasAverageDeploymentRunDuration returns a boolean if a field has been set.

### GetMedianDeploymentRunDuration

`func (o *DeploymentAssetMetric) GetMedianDeploymentRunDuration() int64`

GetMedianDeploymentRunDuration returns the MedianDeploymentRunDuration field if non-nil, zero value otherwise.

### GetMedianDeploymentRunDurationOk

`func (o *DeploymentAssetMetric) GetMedianDeploymentRunDurationOk() (*int64, bool)`

GetMedianDeploymentRunDurationOk returns a tuple with the MedianDeploymentRunDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedianDeploymentRunDuration

`func (o *DeploymentAssetMetric) SetMedianDeploymentRunDuration(v int64)`

SetMedianDeploymentRunDuration sets MedianDeploymentRunDuration field to given value.

### HasMedianDeploymentRunDuration

`func (o *DeploymentAssetMetric) HasMedianDeploymentRunDuration() bool`

HasMedianDeploymentRunDuration returns a boolean if a field has been set.

### GetLongestDeploymentRunDuration

`func (o *DeploymentAssetMetric) GetLongestDeploymentRunDuration() int64`

GetLongestDeploymentRunDuration returns the LongestDeploymentRunDuration field if non-nil, zero value otherwise.

### GetLongestDeploymentRunDurationOk

`func (o *DeploymentAssetMetric) GetLongestDeploymentRunDurationOk() (*int64, bool)`

GetLongestDeploymentRunDurationOk returns a tuple with the LongestDeploymentRunDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongestDeploymentRunDuration

`func (o *DeploymentAssetMetric) SetLongestDeploymentRunDuration(v int64)`

SetLongestDeploymentRunDuration sets LongestDeploymentRunDuration field to given value.

### HasLongestDeploymentRunDuration

`func (o *DeploymentAssetMetric) HasLongestDeploymentRunDuration() bool`

HasLongestDeploymentRunDuration returns a boolean if a field has been set.

### GetShortestDeploymentRunDuration

`func (o *DeploymentAssetMetric) GetShortestDeploymentRunDuration() int64`

GetShortestDeploymentRunDuration returns the ShortestDeploymentRunDuration field if non-nil, zero value otherwise.

### GetShortestDeploymentRunDurationOk

`func (o *DeploymentAssetMetric) GetShortestDeploymentRunDurationOk() (*int64, bool)`

GetShortestDeploymentRunDurationOk returns a tuple with the ShortestDeploymentRunDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortestDeploymentRunDuration

`func (o *DeploymentAssetMetric) SetShortestDeploymentRunDuration(v int64)`

SetShortestDeploymentRunDuration sets ShortestDeploymentRunDuration field to given value.

### HasShortestDeploymentRunDuration

`func (o *DeploymentAssetMetric) HasShortestDeploymentRunDuration() bool`

HasShortestDeploymentRunDuration returns a boolean if a field has been set.

### GetTotalDeploymentRuns

`func (o *DeploymentAssetMetric) GetTotalDeploymentRuns() int32`

GetTotalDeploymentRuns returns the TotalDeploymentRuns field if non-nil, zero value otherwise.

### GetTotalDeploymentRunsOk

`func (o *DeploymentAssetMetric) GetTotalDeploymentRunsOk() (*int32, bool)`

GetTotalDeploymentRunsOk returns a tuple with the TotalDeploymentRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDeploymentRuns

`func (o *DeploymentAssetMetric) SetTotalDeploymentRuns(v int32)`

SetTotalDeploymentRuns sets TotalDeploymentRuns field to given value.

### HasTotalDeploymentRuns

`func (o *DeploymentAssetMetric) HasTotalDeploymentRuns() bool`

HasTotalDeploymentRuns returns a boolean if a field has been set.

### GetTotalDeploymentRunSuccess

`func (o *DeploymentAssetMetric) GetTotalDeploymentRunSuccess() int32`

GetTotalDeploymentRunSuccess returns the TotalDeploymentRunSuccess field if non-nil, zero value otherwise.

### GetTotalDeploymentRunSuccessOk

`func (o *DeploymentAssetMetric) GetTotalDeploymentRunSuccessOk() (*int32, bool)`

GetTotalDeploymentRunSuccessOk returns a tuple with the TotalDeploymentRunSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDeploymentRunSuccess

`func (o *DeploymentAssetMetric) SetTotalDeploymentRunSuccess(v int32)`

SetTotalDeploymentRunSuccess sets TotalDeploymentRunSuccess field to given value.

### HasTotalDeploymentRunSuccess

`func (o *DeploymentAssetMetric) HasTotalDeploymentRunSuccess() bool`

HasTotalDeploymentRunSuccess returns a boolean if a field has been set.

### GetTotalDeploymentRunError

`func (o *DeploymentAssetMetric) GetTotalDeploymentRunError() int32`

GetTotalDeploymentRunError returns the TotalDeploymentRunError field if non-nil, zero value otherwise.

### GetTotalDeploymentRunErrorOk

`func (o *DeploymentAssetMetric) GetTotalDeploymentRunErrorOk() (*int32, bool)`

GetTotalDeploymentRunErrorOk returns a tuple with the TotalDeploymentRunError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDeploymentRunError

`func (o *DeploymentAssetMetric) SetTotalDeploymentRunError(v int32)`

SetTotalDeploymentRunError sets TotalDeploymentRunError field to given value.

### HasTotalDeploymentRunError

`func (o *DeploymentAssetMetric) HasTotalDeploymentRunError() bool`

HasTotalDeploymentRunError returns a boolean if a field has been set.

### GetTotalDeploymentRunCancel

`func (o *DeploymentAssetMetric) GetTotalDeploymentRunCancel() int32`

GetTotalDeploymentRunCancel returns the TotalDeploymentRunCancel field if non-nil, zero value otherwise.

### GetTotalDeploymentRunCancelOk

`func (o *DeploymentAssetMetric) GetTotalDeploymentRunCancelOk() (*int32, bool)`

GetTotalDeploymentRunCancelOk returns a tuple with the TotalDeploymentRunCancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDeploymentRunCancel

`func (o *DeploymentAssetMetric) SetTotalDeploymentRunCancel(v int32)`

SetTotalDeploymentRunCancel sets TotalDeploymentRunCancel field to given value.

### HasTotalDeploymentRunCancel

`func (o *DeploymentAssetMetric) HasTotalDeploymentRunCancel() bool`

HasTotalDeploymentRunCancel returns a boolean if a field has been set.

### GetTotalDeploymentRunUnknown

`func (o *DeploymentAssetMetric) GetTotalDeploymentRunUnknown() int32`

GetTotalDeploymentRunUnknown returns the TotalDeploymentRunUnknown field if non-nil, zero value otherwise.

### GetTotalDeploymentRunUnknownOk

`func (o *DeploymentAssetMetric) GetTotalDeploymentRunUnknownOk() (*int32, bool)`

GetTotalDeploymentRunUnknownOk returns a tuple with the TotalDeploymentRunUnknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDeploymentRunUnknown

`func (o *DeploymentAssetMetric) SetTotalDeploymentRunUnknown(v int32)`

SetTotalDeploymentRunUnknown sets TotalDeploymentRunUnknown field to given value.

### HasTotalDeploymentRunUnknown

`func (o *DeploymentAssetMetric) HasTotalDeploymentRunUnknown() bool`

HasTotalDeploymentRunUnknown returns a boolean if a field has been set.

### GetTotalVirtualMachines

`func (o *DeploymentAssetMetric) GetTotalVirtualMachines() int32`

GetTotalVirtualMachines returns the TotalVirtualMachines field if non-nil, zero value otherwise.

### GetTotalVirtualMachinesOk

`func (o *DeploymentAssetMetric) GetTotalVirtualMachinesOk() (*int32, bool)`

GetTotalVirtualMachinesOk returns a tuple with the TotalVirtualMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVirtualMachines

`func (o *DeploymentAssetMetric) SetTotalVirtualMachines(v int32)`

SetTotalVirtualMachines sets TotalVirtualMachines field to given value.

### HasTotalVirtualMachines

`func (o *DeploymentAssetMetric) HasTotalVirtualMachines() bool`

HasTotalVirtualMachines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


