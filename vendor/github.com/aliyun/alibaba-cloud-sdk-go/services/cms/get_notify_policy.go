package cms

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

// GetNotifyPolicy invokes the cms.GetNotifyPolicy API synchronously
// api document: https://help.aliyun.com/api/cms/getnotifypolicy.html
func (client *Client) GetNotifyPolicy(request *GetNotifyPolicyRequest) (response *GetNotifyPolicyResponse, err error) {
	response = CreateGetNotifyPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// GetNotifyPolicyWithChan invokes the cms.GetNotifyPolicy API asynchronously
// api document: https://help.aliyun.com/api/cms/getnotifypolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetNotifyPolicyWithChan(request *GetNotifyPolicyRequest) (<-chan *GetNotifyPolicyResponse, <-chan error) {
	responseChan := make(chan *GetNotifyPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetNotifyPolicy(request)
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

// GetNotifyPolicyWithCallback invokes the cms.GetNotifyPolicy API asynchronously
// api document: https://help.aliyun.com/api/cms/getnotifypolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetNotifyPolicyWithCallback(request *GetNotifyPolicyRequest, callback func(response *GetNotifyPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetNotifyPolicyResponse
		var err error
		defer close(result)
		response, err = client.GetNotifyPolicy(request)
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

// GetNotifyPolicyRequest is the request struct for api GetNotifyPolicy
type GetNotifyPolicyRequest struct {
	*requests.RpcRequest
	AlertName  string `position:"Query" name:"AlertName"`
	PolicyType string `position:"Query" name:"PolicyType"`
	Id         string `position:"Query" name:"Id"`
	Dimensions string `position:"Query" name:"Dimensions"`
}

// GetNotifyPolicyResponse is the response struct for api GetNotifyPolicy
type GetNotifyPolicyResponse struct {
	*responses.BaseResponse
	Code    string `json:"code" xml:"code"`
	Message string `json:"message" xml:"message"`
	Success string `json:"success" xml:"success"`
	TraceId string `json:"traceId" xml:"traceId"`
	Result  Result `json:"Result" xml:"Result"`
}

// CreateGetNotifyPolicyRequest creates a request to invoke GetNotifyPolicy API
func CreateGetNotifyPolicyRequest() (request *GetNotifyPolicyRequest) {
	request = &GetNotifyPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2018-03-08", "GetNotifyPolicy", "cms", "openAPI")
	return
}

// CreateGetNotifyPolicyResponse creates a response to parse from GetNotifyPolicy response
func CreateGetNotifyPolicyResponse() (response *GetNotifyPolicyResponse) {
	response = &GetNotifyPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}