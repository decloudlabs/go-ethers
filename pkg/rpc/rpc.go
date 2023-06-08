// @Title        rpc
// @Description  Abstract implementation of network rpc service for evm architecture
// @Author       happyboy
package rpc

import (
	"encoding/json"
	"strings"
)

// rpc connection related constants
const (
	CONTENT_TYPE     = "application/json"
	JSON_RPC_VERSION = "2.0"
)

// rpc service connection handle struct
type RPC struct {
	rpcUrl  string `json:"rpc_url"`  // rpc url
	chainId int16  `json:"chain_id"` // chain id
}

// json rpc request struct
type JsonrpcMsg struct {
	Id      int16    `json:"id"`      // chain id
	Jsonrpc string   `json:"jsonrpc"` // json rpc version
	Params  []string `json:"params"`  // params
	Method  string   `json:"method"`  // method name
}

// @title    NewRPC
// @description   rpc service connection handle instantiation function
// @auth      happyboy
// @param     rpcUrl        string         rpcurl
// @param     chainId       int16          chain id
// @return    rpc           *RPC           rpc service connection handle
func NewRPC(rpcUrl string, chainId int16) *RPC {
	rpc := new(RPC)
	rpc.chainId = chainId
	rpc.rpcUrl = rpcUrl
	return rpc
}

// @title    BuildJsonRpcMsg
// @description   build json rpc message struct
// @auth      happyboy
// @param     params        []byte         params
// @return    msg           string          json rpc message struct strin
func (rpc *RPC) BuildJsonRpcMsg(method string, params []string) *strings.Reader {
	msg := JsonrpcMsg{}
	msg.Id = rpc.chainId
	msg.Jsonrpc = JSON_RPC_VERSION
	msg.Method = method
	if params != nil {
		msg.Params = params
	}
	j, _ := json.Marshal(msg)
	return strings.NewReader(string(j))
}
