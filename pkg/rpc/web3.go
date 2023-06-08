// @Title        rpc
// @Description  Abstract implementation of network rpc service for evm architecture
// @Author       happyboy
package rpc

import (
	"io"
	"net/http"
)

// @title         NewRPC
// @description   rpc method "web3_clientVersion" abstract implementation
// @auth          happyboy
func (rpc *RPC) WebClientVersion() string {
	resp, _ := http.Post(rpc.rpcUrl, CONTENT_TYPE, rpc.BuildJsonRpcMsg("web3_clientVersion", nil))
	body, _ := io.ReadAll(resp.Body)
	return string(body)
}
