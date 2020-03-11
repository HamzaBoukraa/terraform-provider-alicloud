package vpc

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// CreateRouterInterface invokes the vpc.CreateRouterInterface API synchronously
// api document: https://help.aliyun.com/api/vpc/createrouterinterface.html
func (client *Client) CreateRouterInterface(request *CreateRouterInterfaceRequest) (response *CreateRouterInterfaceResponse, err error) {
	response = CreateCreateRouterInterfaceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateRouterInterfaceWithChan invokes the vpc.CreateRouterInterface API asynchronously
// api document: https://help.aliyun.com/api/vpc/createrouterinterface.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateRouterInterfaceWithChan(request *CreateRouterInterfaceRequest) (<-chan *CreateRouterInterfaceResponse, <-chan error) {
	responseChan := make(chan *CreateRouterInterfaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateRouterInterface(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// CreateRouterInterfaceWithCallback invokes the vpc.CreateRouterInterface API asynchronously
// api document: https://help.aliyun.com/api/vpc/createrouterinterface.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateRouterInterfaceWithCallback(request *CreateRouterInterfaceRequest, callback func(response *CreateRouterInterfaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateRouterInterfaceResponse
		var err error
		defer close(result)
		response, err = client.CreateRouterInterface(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// CreateRouterInterfaceRequest is the request struct for api CreateRouterInterface
type CreateRouterInterfaceRequest struct {
	*requests.RpcRequest
	AccessPointId            string           `position:"Query" name:"AccessPointId"`
	OppositeRouterId         string           `position:"Query" name:"OppositeRouterId"`
	OppositeAccessPointId    string           `position:"Query" name:"OppositeAccessPointId"`
	ResourceOwnerId          requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Role                     string           `position:"Query" name:"Role"`
	ClientToken              string           `position:"Query" name:"ClientToken"`
	HealthCheckTargetIp      string           `position:"Query" name:"HealthCheckTargetIp"`
	Description              string           `position:"Query" name:"Description"`
	Spec                     string           `position:"Query" name:"Spec"`
	OppositeInterfaceId      string           `position:"Query" name:"OppositeInterfaceId"`
	InstanceChargeType       string           `position:"Query" name:"InstanceChargeType"`
	Period                   requests.Integer `position:"Query" name:"Period"`
	AutoPay                  requests.Boolean `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount     string           `position:"Query" name:"ResourceOwnerAccount"`
	OppositeRegionId         string           `position:"Query" name:"OppositeRegionId"`
	OwnerAccount             string           `position:"Query" name:"OwnerAccount"`
	OwnerId                  requests.Integer `position:"Query" name:"OwnerId"`
	OppositeInterfaceOwnerId string           `position:"Query" name:"OppositeInterfaceOwnerId"`
	RouterType               string           `position:"Query" name:"RouterType"`
	HealthCheckSourceIp      string           `position:"Query" name:"HealthCheckSourceIp"`
	RouterId                 string           `position:"Query" name:"RouterId"`
	OppositeRouterType       string           `position:"Query" name:"OppositeRouterType"`
	Name                     string           `position:"Query" name:"Name"`
	PricingCycle             string           `position:"Query" name:"PricingCycle"`
}

// CreateRouterInterfaceResponse is the response struct for api CreateRouterInterface
type CreateRouterInterfaceResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	RouterInterfaceId string `json:"RouterInterfaceId" xml:"RouterInterfaceId"`
	OrderId           int64  `json:"OrderId" xml:"OrderId"`
}

// CreateCreateRouterInterfaceRequest creates a request to invoke CreateRouterInterface API
func CreateCreateRouterInterfaceRequest() (request *CreateRouterInterfaceRequest) {
	request = &CreateRouterInterfaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CreateRouterInterface", "Vpc", "openAPI")
	return
}

// CreateCreateRouterInterfaceResponse creates a response to parse from CreateRouterInterface response
func CreateCreateRouterInterfaceResponse() (response *CreateRouterInterfaceResponse) {
	response = &CreateRouterInterfaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
