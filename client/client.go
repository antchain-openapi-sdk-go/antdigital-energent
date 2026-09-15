// This file is auto-generated, don't edit it. Thanks.
package client

import (
	rpcutil "github.com/alibabacloud-go/tea-rpc-utils/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	antchainutil "github.com/antchain-openapi-sdk-go/antchain-util/service"
)

// Description:
//
// Model for initing client
type Config struct {
	// accesskey id
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	// accesskey secret
	AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	// security token
	SecurityToken *string `json:"securityToken,omitempty" xml:"securityToken,omitempty"`
	// http protocol
	//
	// example:
	//
	// http
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// read timeout
	//
	// example:
	//
	// 10
	ReadTimeout *int `json:"readTimeout,omitempty" xml:"readTimeout,omitempty"`
	// connect timeout
	//
	// example:
	//
	// 10
	ConnectTimeout *int `json:"connectTimeout,omitempty" xml:"connectTimeout,omitempty"`
	// http proxy
	//
	// example:
	//
	// http://localhost
	HttpProxy *string `json:"httpProxy,omitempty" xml:"httpProxy,omitempty"`
	// https proxy
	//
	// example:
	//
	// https://localhost
	HttpsProxy *string `json:"httpsProxy,omitempty" xml:"httpsProxy,omitempty"`
	// endpoint
	//
	// example:
	//
	// cs.aliyuncs.com
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// proxy white list
	//
	// example:
	//
	// http://localhost
	NoProxy *string `json:"noProxy,omitempty" xml:"noProxy,omitempty"`
	// max idle conns
	//
	// example:
	//
	// 3
	MaxIdleConns *int `json:"maxIdleConns,omitempty" xml:"maxIdleConns,omitempty"`
	// user agent
	//
	// example:
	//
	// Alibabacloud/1
	UserAgent *string `json:"userAgent,omitempty" xml:"userAgent,omitempty"`
	// socks5 proxy
	Socks5Proxy *string `json:"socks5Proxy,omitempty" xml:"socks5Proxy,omitempty"`
	// socks5 network
	//
	// example:
	//
	// TCP
	Socks5NetWork *string `json:"socks5NetWork,omitempty" xml:"socks5NetWork,omitempty"`
	// 长链接最大空闲时长
	MaxIdleTimeMillis *int `json:"maxIdleTimeMillis,omitempty" xml:"maxIdleTimeMillis,omitempty"`
	// 长链接最大连接时长
	KeepAliveDurationMillis *int `json:"keepAliveDurationMillis,omitempty" xml:"keepAliveDurationMillis,omitempty"`
	// 最大连接数（长链接最大总数）
	MaxRequests *int `json:"maxRequests,omitempty" xml:"maxRequests,omitempty"`
	// 每个目标主机的最大连接数（分主机域名的长链接最大总数
	MaxRequestsPerHost *int `json:"maxRequestsPerHost,omitempty" xml:"maxRequestsPerHost,omitempty"`
}

func (s Config) String() string {
	return tea.Prettify(s)
}

func (s Config) GoString() string {
	return s.String()
}

func (s *Config) SetAccessKeyId(v string) *Config {
	s.AccessKeyId = &v
	return s
}

func (s *Config) SetAccessKeySecret(v string) *Config {
	s.AccessKeySecret = &v
	return s
}

func (s *Config) SetSecurityToken(v string) *Config {
	s.SecurityToken = &v
	return s
}

func (s *Config) SetProtocol(v string) *Config {
	s.Protocol = &v
	return s
}

func (s *Config) SetReadTimeout(v int) *Config {
	s.ReadTimeout = &v
	return s
}

func (s *Config) SetConnectTimeout(v int) *Config {
	s.ConnectTimeout = &v
	return s
}

func (s *Config) SetHttpProxy(v string) *Config {
	s.HttpProxy = &v
	return s
}

func (s *Config) SetHttpsProxy(v string) *Config {
	s.HttpsProxy = &v
	return s
}

func (s *Config) SetEndpoint(v string) *Config {
	s.Endpoint = &v
	return s
}

func (s *Config) SetNoProxy(v string) *Config {
	s.NoProxy = &v
	return s
}

func (s *Config) SetMaxIdleConns(v int) *Config {
	s.MaxIdleConns = &v
	return s
}

func (s *Config) SetUserAgent(v string) *Config {
	s.UserAgent = &v
	return s
}

func (s *Config) SetSocks5Proxy(v string) *Config {
	s.Socks5Proxy = &v
	return s
}

func (s *Config) SetSocks5NetWork(v string) *Config {
	s.Socks5NetWork = &v
	return s
}

func (s *Config) SetMaxIdleTimeMillis(v int) *Config {
	s.MaxIdleTimeMillis = &v
	return s
}

func (s *Config) SetKeepAliveDurationMillis(v int) *Config {
	s.KeepAliveDurationMillis = &v
	return s
}

func (s *Config) SetMaxRequests(v int) *Config {
	s.MaxRequests = &v
	return s
}

func (s *Config) SetMaxRequestsPerHost(v int) *Config {
	s.MaxRequestsPerHost = &v
	return s
}

// 召回出参明细
type RecallDataDetail struct {
	// 文件名
	// example:
	//
	// 浙江省电力交易细则
	Docname *string `json:"docname,omitempty" xml:"docname,omitempty"`
	// 片段内容
	// example:
	//
	// 浙江省交易细则
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 重排得分
	// example:
	//
	// 1.2
	Rerankscore *int64 `json:"rerankscore,omitempty" xml:"rerankscore,omitempty"`
	// 重排位次
	// example:
	//
	// 1
	Rerankrank *int64 `json:"rerankrank,omitempty" xml:"rerankrank,omitempty"`
	// 综合得分
	// example:
	//
	// 1
	Score *int64 `json:"score,omitempty" xml:"score,omitempty"`
	// chunk标题
	// example:
	//
	// 浙江省交易细则
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 知识库id
	// example:
	//
	// 1
	Knowledgebaseids *string `json:"knowledgebaseids,omitempty" xml:"knowledgebaseids,omitempty"`
	// official 官方知识库、custom个人知识库
	// example:
	//
	// official
	Knowledgebasetag *string `json:"knowledgebasetag,omitempty" xml:"knowledgebasetag,omitempty"`
}

func (s RecallDataDetail) String() string {
	return tea.Prettify(s)
}

func (s RecallDataDetail) GoString() string {
	return s.String()
}

func (s *RecallDataDetail) SetDocname(v string) *RecallDataDetail {
	s.Docname = &v
	return s
}

func (s *RecallDataDetail) SetContent(v string) *RecallDataDetail {
	s.Content = &v
	return s
}

func (s *RecallDataDetail) SetRerankscore(v int64) *RecallDataDetail {
	s.Rerankscore = &v
	return s
}

func (s *RecallDataDetail) SetRerankrank(v int64) *RecallDataDetail {
	s.Rerankrank = &v
	return s
}

func (s *RecallDataDetail) SetScore(v int64) *RecallDataDetail {
	s.Score = &v
	return s
}

func (s *RecallDataDetail) SetTitle(v string) *RecallDataDetail {
	s.Title = &v
	return s
}

func (s *RecallDataDetail) SetKnowledgebaseids(v string) *RecallDataDetail {
	s.Knowledgebaseids = &v
	return s
}

func (s *RecallDataDetail) SetKnowledgebasetag(v string) *RecallDataDetail {
	s.Knowledgebasetag = &v
	return s
}

type QueryKnowledgeRagRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 用户问题
	Query *string `json:"query,omitempty" xml:"query,omitempty" require:"true"`
	// 知识库id
	KnowledgeBaseIds []*string `json:"knowledge_base_ids,omitempty" xml:"knowledge_base_ids,omitempty" type:"Repeated"`
	// 召回个数
	Topk *int64 `json:"topk,omitempty" xml:"topk,omitempty"`
	// 是否包含引用来源
	IncludeReferences *bool `json:"include_references,omitempty" xml:"include_references,omitempty"`
}

func (s QueryKnowledgeRagRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryKnowledgeRagRequest) GoString() string {
	return s.String()
}

func (s *QueryKnowledgeRagRequest) SetAuthToken(v string) *QueryKnowledgeRagRequest {
	s.AuthToken = &v
	return s
}

func (s *QueryKnowledgeRagRequest) SetProductInstanceId(v string) *QueryKnowledgeRagRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *QueryKnowledgeRagRequest) SetQuery(v string) *QueryKnowledgeRagRequest {
	s.Query = &v
	return s
}

func (s *QueryKnowledgeRagRequest) SetKnowledgeBaseIds(v []*string) *QueryKnowledgeRagRequest {
	s.KnowledgeBaseIds = v
	return s
}

func (s *QueryKnowledgeRagRequest) SetTopk(v int64) *QueryKnowledgeRagRequest {
	s.Topk = &v
	return s
}

func (s *QueryKnowledgeRagRequest) SetIncludeReferences(v bool) *QueryKnowledgeRagRequest {
	s.IncludeReferences = &v
	return s
}

type QueryKnowledgeRagResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
}

func (s QueryKnowledgeRagResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryKnowledgeRagResponse) GoString() string {
	return s.String()
}

func (s *QueryKnowledgeRagResponse) SetReqMsgId(v string) *QueryKnowledgeRagResponse {
	s.ReqMsgId = &v
	return s
}

func (s *QueryKnowledgeRagResponse) SetResultCode(v string) *QueryKnowledgeRagResponse {
	s.ResultCode = &v
	return s
}

func (s *QueryKnowledgeRagResponse) SetResultMsg(v string) *QueryKnowledgeRagResponse {
	s.ResultMsg = &v
	return s
}

type QueryEnergyknowledgeRagRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 用户问题
	Query *string `json:"query,omitempty" xml:"query,omitempty" require:"true"`
	// 指定查询的知识库id
	KnowledgeBaseIds []*string `json:"knowledge_base_ids,omitempty" xml:"knowledge_base_ids,omitempty" type:"Repeated"`
	// 召回数量
	Topk *int64 `json:"topk,omitempty" xml:"topk,omitempty"`
	// 是否返回引用docid
	IncludeReferences *bool `json:"include_references,omitempty" xml:"include_references,omitempty"`
}

func (s QueryEnergyknowledgeRagRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryEnergyknowledgeRagRequest) GoString() string {
	return s.String()
}

func (s *QueryEnergyknowledgeRagRequest) SetAuthToken(v string) *QueryEnergyknowledgeRagRequest {
	s.AuthToken = &v
	return s
}

func (s *QueryEnergyknowledgeRagRequest) SetProductInstanceId(v string) *QueryEnergyknowledgeRagRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *QueryEnergyknowledgeRagRequest) SetQuery(v string) *QueryEnergyknowledgeRagRequest {
	s.Query = &v
	return s
}

func (s *QueryEnergyknowledgeRagRequest) SetKnowledgeBaseIds(v []*string) *QueryEnergyknowledgeRagRequest {
	s.KnowledgeBaseIds = v
	return s
}

func (s *QueryEnergyknowledgeRagRequest) SetTopk(v int64) *QueryEnergyknowledgeRagRequest {
	s.Topk = &v
	return s
}

func (s *QueryEnergyknowledgeRagRequest) SetIncludeReferences(v bool) *QueryEnergyknowledgeRagRequest {
	s.IncludeReferences = &v
	return s
}

type QueryEnergyknowledgeRagResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
}

func (s QueryEnergyknowledgeRagResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryEnergyknowledgeRagResponse) GoString() string {
	return s.String()
}

func (s *QueryEnergyknowledgeRagResponse) SetReqMsgId(v string) *QueryEnergyknowledgeRagResponse {
	s.ReqMsgId = &v
	return s
}

func (s *QueryEnergyknowledgeRagResponse) SetResultCode(v string) *QueryEnergyknowledgeRagResponse {
	s.ResultCode = &v
	return s
}

func (s *QueryEnergyknowledgeRagResponse) SetResultMsg(v string) *QueryEnergyknowledgeRagResponse {
	s.ResultMsg = &v
	return s
}

type QueryEnergyknowledgeRagrecallRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// rag检索问题
	Query *string `json:"query,omitempty" xml:"query,omitempty" require:"true"`
	// 召回的知识库id
	Knowledgebaseids []*string `json:"knowledgebaseids,omitempty" xml:"knowledgebaseids,omitempty" type:"Repeated"`
	// 是否包含召回文件名，默认true
	Includereferences *bool `json:"includereferences,omitempty" xml:"includereferences,omitempty"`
	// 召回条数，默认5
	Topk *int64 `json:"topk,omitempty" xml:"topk,omitempty"`
}

func (s QueryEnergyknowledgeRagrecallRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryEnergyknowledgeRagrecallRequest) GoString() string {
	return s.String()
}

func (s *QueryEnergyknowledgeRagrecallRequest) SetAuthToken(v string) *QueryEnergyknowledgeRagrecallRequest {
	s.AuthToken = &v
	return s
}

func (s *QueryEnergyknowledgeRagrecallRequest) SetProductInstanceId(v string) *QueryEnergyknowledgeRagrecallRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *QueryEnergyknowledgeRagrecallRequest) SetQuery(v string) *QueryEnergyknowledgeRagrecallRequest {
	s.Query = &v
	return s
}

func (s *QueryEnergyknowledgeRagrecallRequest) SetKnowledgebaseids(v []*string) *QueryEnergyknowledgeRagrecallRequest {
	s.Knowledgebaseids = v
	return s
}

func (s *QueryEnergyknowledgeRagrecallRequest) SetIncludereferences(v bool) *QueryEnergyknowledgeRagrecallRequest {
	s.Includereferences = &v
	return s
}

func (s *QueryEnergyknowledgeRagrecallRequest) SetTopk(v int64) *QueryEnergyknowledgeRagrecallRequest {
	s.Topk = &v
	return s
}

type QueryEnergyknowledgeRagrecallResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 出参数据
	Datalist []*RecallDataDetail `json:"datalist,omitempty" xml:"datalist,omitempty" type:"Repeated"`
}

func (s QueryEnergyknowledgeRagrecallResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryEnergyknowledgeRagrecallResponse) GoString() string {
	return s.String()
}

func (s *QueryEnergyknowledgeRagrecallResponse) SetReqMsgId(v string) *QueryEnergyknowledgeRagrecallResponse {
	s.ReqMsgId = &v
	return s
}

func (s *QueryEnergyknowledgeRagrecallResponse) SetResultCode(v string) *QueryEnergyknowledgeRagrecallResponse {
	s.ResultCode = &v
	return s
}

func (s *QueryEnergyknowledgeRagrecallResponse) SetResultMsg(v string) *QueryEnergyknowledgeRagrecallResponse {
	s.ResultMsg = &v
	return s
}

func (s *QueryEnergyknowledgeRagrecallResponse) SetDatalist(v []*RecallDataDetail) *QueryEnergyknowledgeRagrecallResponse {
	s.Datalist = v
	return s
}

type Client struct {
	Endpoint                *string
	RegionId                *string
	AccessKeyId             *string
	AccessKeySecret         *string
	Protocol                *string
	UserAgent               *string
	ReadTimeout             *int
	ConnectTimeout          *int
	HttpProxy               *string
	HttpsProxy              *string
	Socks5Proxy             *string
	Socks5NetWork           *string
	NoProxy                 *string
	MaxIdleConns            *int
	SecurityToken           *string
	MaxIdleTimeMillis       *int
	KeepAliveDurationMillis *int
	MaxRequests             *int
	MaxRequestsPerHost      *int
}

// Description:
//
// # Init client with Config
//
// @param config - config contains the necessary information to create a client
func NewClient(config *Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *Config) (_err error) {
	if tea.BoolValue(util.IsUnset(config)) {
		_err = tea.NewSDKError(map[string]interface{}{
			"code":    "ParameterMissing",
			"message": "'config' can not be unset",
		})
		return _err
	}

	client.AccessKeyId = config.AccessKeyId
	client.AccessKeySecret = config.AccessKeySecret
	client.SecurityToken = config.SecurityToken
	client.Endpoint = config.Endpoint
	client.Protocol = config.Protocol
	client.UserAgent = config.UserAgent
	client.ReadTimeout = util.DefaultNumber(config.ReadTimeout, tea.Int(20000))
	client.ConnectTimeout = util.DefaultNumber(config.ConnectTimeout, tea.Int(20000))
	client.HttpProxy = config.HttpProxy
	client.HttpsProxy = config.HttpsProxy
	client.NoProxy = config.NoProxy
	client.Socks5Proxy = config.Socks5Proxy
	client.Socks5NetWork = config.Socks5NetWork
	client.MaxIdleConns = util.DefaultNumber(config.MaxIdleConns, tea.Int(60000))
	client.MaxIdleTimeMillis = util.DefaultNumber(config.MaxIdleTimeMillis, tea.Int(5))
	client.KeepAliveDurationMillis = util.DefaultNumber(config.KeepAliveDurationMillis, tea.Int(5000))
	client.MaxRequests = util.DefaultNumber(config.MaxRequests, tea.Int(100))
	client.MaxRequestsPerHost = util.DefaultNumber(config.MaxRequestsPerHost, tea.Int(100))
	return nil
}

// Description:
//
// # Encapsulate the request and invoke the network
//
// @param action - api name
//
// @param protocol - http or https
//
// @param method - e.g. GET
//
// @param pathname - pathname of every api
//
// @param request - which contains request params
//
// @param runtime - which controls some details of call api, such as retry times
//
// @return the response
func (client *Client) DoRequest(version *string, action *string, protocol *string, method *string, pathname *string, request map[string]interface{}, headers map[string]*string, runtime *util.RuntimeOptions) (_result map[string]interface{}, _err error) {
	_err = tea.Validate(runtime)
	if _err != nil {
		return _result, _err
	}
	_runtime := map[string]interface{}{
		"timeouted":          "retry",
		"readTimeout":        tea.IntValue(util.DefaultNumber(runtime.ReadTimeout, client.ReadTimeout)),
		"connectTimeout":     tea.IntValue(util.DefaultNumber(runtime.ConnectTimeout, client.ConnectTimeout)),
		"httpProxy":          tea.StringValue(util.DefaultString(runtime.HttpProxy, client.HttpProxy)),
		"httpsProxy":         tea.StringValue(util.DefaultString(runtime.HttpsProxy, client.HttpsProxy)),
		"noProxy":            tea.StringValue(util.DefaultString(runtime.NoProxy, client.NoProxy)),
		"maxIdleConns":       tea.IntValue(util.DefaultNumber(runtime.MaxIdleConns, client.MaxIdleConns)),
		"maxIdleTimeMillis":  tea.IntValue(client.MaxIdleTimeMillis),
		"keepAliveDuration":  tea.IntValue(client.KeepAliveDurationMillis),
		"maxRequests":        tea.IntValue(client.MaxRequests),
		"maxRequestsPerHost": tea.IntValue(client.MaxRequestsPerHost),
		"retry": map[string]interface{}{
			"retryable":   tea.BoolValue(runtime.Autoretry),
			"maxAttempts": tea.IntValue(util.DefaultNumber(runtime.MaxAttempts, tea.Int(3))),
		},
		"backoff": map[string]interface{}{
			"policy": tea.StringValue(util.DefaultString(runtime.BackoffPolicy, tea.String("no"))),
			"period": tea.IntValue(util.DefaultNumber(runtime.BackoffPeriod, tea.Int(1))),
		},
		"ignoreSSL": tea.BoolValue(runtime.IgnoreSSL),
	}

	_resp := make(map[string]interface{})
	for _retryTimes := 0; tea.BoolValue(tea.AllowRetry(_runtime["retry"], tea.Int(_retryTimes))); _retryTimes++ {
		if _retryTimes > 0 {
			_backoffTime := tea.GetBackoffTime(_runtime["backoff"], tea.Int(_retryTimes))
			if tea.IntValue(_backoffTime) > 0 {
				tea.Sleep(_backoffTime)
			}
		}

		_resp, _err = func() (map[string]interface{}, error) {
			request_ := tea.NewRequest()
			request_.Protocol = util.DefaultString(client.Protocol, protocol)
			request_.Method = method
			request_.Pathname = pathname
			request_.Query = map[string]*string{
				"method":           action,
				"version":          version,
				"sign_type":        tea.String("HmacSHA1"),
				"req_time":         antchainutil.GetTimestamp(),
				"req_msg_id":       antchainutil.GetNonce(),
				"access_key":       client.AccessKeyId,
				"base_sdk_version": tea.String("TeaSDK-2.0"),
				"sdk_version":      tea.String("1.0.2"),
				"_prod_code":       tea.String("ENERGENT"),
				"_prod_channel":    tea.String("default"),
			}
			if !tea.BoolValue(util.Empty(client.SecurityToken)) {
				request_.Query["security_token"] = client.SecurityToken
			}

			request_.Headers = tea.Merge(map[string]*string{
				"host":       util.DefaultString(client.Endpoint, tea.String("openapi.antchain.antgroup.com")),
				"user-agent": util.GetUserAgent(client.UserAgent),
			}, headers)
			tmp := util.AnyifyMapValue(rpcutil.Query(request))
			request_.Body = tea.ToReader(util.ToFormString(tmp))
			request_.Headers["content-type"] = tea.String("application/x-www-form-urlencoded")
			signedParam := tea.Merge(request_.Query,
				rpcutil.Query(request))
			request_.Query["sign"] = antchainutil.GetSignature(signedParam, client.AccessKeySecret)
			response_, _err := tea.DoRequest(request_, _runtime)
			if _err != nil {
				return _result, _err
			}
			raw, _err := util.ReadAsString(response_.Body)
			if _err != nil {
				return _result, _err
			}

			obj := util.ParseJSON(raw)
			res, _err := util.AssertAsMap(obj)
			if _err != nil {
				return _result, _err
			}

			resp, _err := util.AssertAsMap(res["response"])
			if _err != nil {
				return _result, _err
			}

			if tea.BoolValue(antchainutil.HasError(raw, client.AccessKeySecret)) {
				_err = tea.NewSDKError(map[string]interface{}{
					"message": resp["result_msg"],
					"data":    resp,
					"code":    resp["result_code"],
				})
				return _result, _err
			}

			_result = resp
			return _result, _err
		}()
		if !tea.BoolValue(tea.Retryable(_err)) {
			break
		}
	}

	return _resp, _err
}

// Description:
//
// Description: 知识库rag 检索
//
// Summary: 知识库rag 检索
func (client *Client) QueryKnowledgeRag(request *QueryKnowledgeRagRequest) (_result *QueryKnowledgeRagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryKnowledgeRagResponse{}
	_body, _err := client.QueryKnowledgeRagEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// Description: 知识库rag 检索
//
// Summary: 知识库rag 检索
func (client *Client) QueryKnowledgeRagEx(request *QueryKnowledgeRagRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryKnowledgeRagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryKnowledgeRagResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("antdigital.energent.knowledge.rag.query"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Description:
//
// Description: rag检索
//
// Summary: rag检索
func (client *Client) QueryEnergyknowledgeRag(request *QueryEnergyknowledgeRagRequest) (_result *QueryEnergyknowledgeRagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryEnergyknowledgeRagResponse{}
	_body, _err := client.QueryEnergyknowledgeRagEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// Description: rag检索
//
// Summary: rag检索
func (client *Client) QueryEnergyknowledgeRagEx(request *QueryEnergyknowledgeRagRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryEnergyknowledgeRagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryEnergyknowledgeRagResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("antdigital.energent.energyknowledge.rag.query"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Description:
//
// Description: 知识图谱rag召回
//
// Summary: 知识图谱rag召回
func (client *Client) QueryEnergyknowledgeRagrecall(request *QueryEnergyknowledgeRagrecallRequest) (_result *QueryEnergyknowledgeRagrecallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryEnergyknowledgeRagrecallResponse{}
	_body, _err := client.QueryEnergyknowledgeRagrecallEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// Description: 知识图谱rag召回
//
// Summary: 知识图谱rag召回
func (client *Client) QueryEnergyknowledgeRagrecallEx(request *QueryEnergyknowledgeRagrecallRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryEnergyknowledgeRagrecallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryEnergyknowledgeRagrecallResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("antdigital.energent.energyknowledge.ragrecall.query"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
