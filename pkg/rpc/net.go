// @Title        rpc
// @Description  Abstract implementation of network rpc service for evm architecture
// @Author       happyboy
package rpc

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// @title         NetListening
// @description   rpc method "net_listening" abstract implementation
// @auth          happyboy
func (rpc *RPC) NetListening() string {
	data := make(map[string]interface{})
	data["jsonrpc"] = JSON_RPC_VERSION
	data["method"] = "net_listening"
	data["id"] = rpc.chainId
	bytesData, _ := json.Marshal(data)
	resp, _ := http.Post(rpc.rpcUrl, CONTENT_TYPE, bytes.NewReader(bytesData))
	body, _ := io.ReadAll(resp.Body)
	return string(body)
}

// @title         NetPeerCount
// @description   rpc method "net_listening" abstract implementation
// @auth          happyboy
func (rpc *RPC) NetPeerCount() string {
	data := make(map[string]interface{})
	data["jsonrpc"] = JSON_RPC_VERSION
	data["method"] = "net_peerCount"
	data["id"] = rpc.chainId
	bytesData, _ := json.Marshal(data)
	resp, _ := http.Post(rpc.rpcUrl, CONTENT_TYPE, bytes.NewReader(bytesData))
	body, _ := io.ReadAll(resp.Body)
	return string(body)
}
