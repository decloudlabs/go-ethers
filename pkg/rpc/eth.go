// @Title        rpc
// @Description  Abstract implementation of network rpc service for evm architecture
// @Author       happyboy
package rpc

import (
	"github.com/decloudlabs/go-ethers/pkg/structs"
	"github.com/decloudlabs/go-ethers/pkg/utils"
	"io"
	"net/http"
	"strconv"
)

// @title         EthSyncing
// @description   rpc method "eth_syncing" abstract implementation
// @auth          happyboy
func (rpc *RPC) EthSyncing() string {
	resp, _ := http.Post(rpc.rpcUrl, CONTENT_TYPE, rpc.BuildJsonRpcMsg("eth_syncing", nil))
	body, _ := io.ReadAll(resp.Body)
	return string(body)
}

// @title         EthMining
// @description   rpc method "eth_mining" abstract implementation
// @auth          happyboy
func (rpc *RPC) EthMining() string {
	resp, _ := http.Post(rpc.rpcUrl, CONTENT_TYPE, rpc.BuildJsonRpcMsg("eth_mining", nil))
	body, _ := io.ReadAll(resp.Body)
	return string(body)
}

// @title         EthHashrate
// @description   rpc method "eth_hashrate" abstract implementation
// @auth          happyboy
func (rpc *RPC) EthHashrate() string {
	resp, _ := http.Post(rpc.rpcUrl, CONTENT_TYPE, rpc.BuildJsonRpcMsg("eth_hashrate", nil))
	body, _ := io.ReadAll(resp.Body)
	return string(body)
}

// @title         EthGasPrice
// @description   rpc method "eth_gasPrice" abstract implementation
// @auth          happyboy
func (rpc *RPC) EthGasPrice() string {
	resp, _ := http.Post(rpc.rpcUrl, CONTENT_TYPE, rpc.BuildJsonRpcMsg("eth_gasPrice", nil))
	body, _ := io.ReadAll(resp.Body)
	return string(body)
}

// @title         EthAccounts
// @description   rpc method "eth_accounts" abstract implementation
// @auth          happyboy
func (rpc *RPC) EthAccounts() string {
	resp, _ := http.Post(rpc.rpcUrl, CONTENT_TYPE, rpc.BuildJsonRpcMsg("eth_accounts", nil))
	body, _ := io.ReadAll(resp.Body)
	return string(body)
}

// @title         EthGetBalance
// @description   rpc method "eth_getBalance" abstract implementation
// @param         address        string         "address to be queried"
// @param         tag            string         "tag"
// @auth          happyboy
func (rpc *RPC) EthGetBalance(address string, tag string) string {
	resp, _ := http.Post(rpc.rpcUrl, CONTENT_TYPE, rpc.BuildJsonRpcMsg("eth_getBalance", []string{address, tag}))
	body, _ := io.ReadAll(resp.Body)
	return string(body)
}

// @title         EthGetStorageAt
// @description   rpc method "eth_getStorageAt" abstract implementation
// @param         address        string         "address to be queried"
// @param         quantity       string         "address to be queried"
// @param         tag            string         "tag"
// @auth          happyboy
func (rpc *RPC) EthGetStorageAt(address string, quantity string, tag string) string {
	resp, _ := http.Post(rpc.rpcUrl, CONTENT_TYPE, rpc.BuildJsonRpcMsg("eth_getStorageAt", []string{address, quantity, tag}))
	body, _ := io.ReadAll(resp.Body)
	return string(body)
}

// @title         EthGetTransactionCount
// @description   rpc method "eth_getTransactionCount" abstract implementation
// @param         address        string         "address to be queried"
// @param         tag            string         "tag"
// @auth          happyboy
func (rpc *RPC) EthGetTransactionCount(address string, tag string) string {
	resp, _ := http.Post(rpc.rpcUrl, CONTENT_TYPE, rpc.BuildJsonRpcMsg("eth_getTransactionCount", []string{address, tag}))
	body, _ := io.ReadAll(resp.Body)
	return string(body)
}

// @title         EthGetBlockTransactionCountByHash
// @description   rpc method "eth_getBlockTransactionCountByHash" abstract implementation
// @param         hash        string         "block hash"
// @auth          happyboy
func (rpc *RPC) EthGetBlockTransactionCountByHash(hash string) string {
	resp, _ := http.Post(rpc.rpcUrl, CONTENT_TYPE, rpc.BuildJsonRpcMsg("eth_getBlockTransactionCountByHash", []string{hash}))
	body, _ := io.ReadAll(resp.Body)
	return string(body)
}

// @title         EthGetBlockTransactionCountByNumber
// @description   rpc method "eth_getBlockTransactionCountByNumber" abstract implementation
// @param         blockNumber        string         "block number"
// @auth          happyboy
func (rpc *RPC) EthGetBlockTransactionCountByNumber(blockNumber int64) string {
	numberHex := "0x" + string(utils.DecimalToHex(blockNumber))
	resp, _ := http.Post(rpc.rpcUrl, CONTENT_TYPE, rpc.BuildJsonRpcMsg("eth_getBlockTransactionCountByNumber", []string{numberHex}))
	body, _ := io.ReadAll(resp.Body)
	return string(body)
}

// @title         EthGetCode
// @description   rpc method "eth_getCode" abstract implementation
// @param         address        string         "address to be queried"
// @param         tag            string         "tag latest|earliest|pending"
// @auth          happyboy
func (rpc *RPC) EthGetCode(address string, tag string) string {
	resp, _ := http.Post(rpc.rpcUrl, CONTENT_TYPE, rpc.BuildJsonRpcMsg("eth_getCode", []string{address, tag}))
	body, _ := io.ReadAll(resp.Body)
	return string(body)
}

// @title         EthSign
// @description   rpc method "eth_sign" abstract implementation
// @param         address            string         "address to be queried"
// @param         signMsg            string         "sign(keccak256("\x19Ethereum Signed Message:\n" + len(message) + message))) sgin messge"
// @auth          happyboy
func (rpc *RPC) EthSign(address string, signMsg string) string {
	resp, _ := http.Post(rpc.rpcUrl, CONTENT_TYPE, rpc.BuildJsonRpcMsg("eth_sign", []string{address, signMsg}))
	body, _ := io.ReadAll(resp.Body)
	return string(body)
}

// @title         EthSendTransaction
// @description   rpc method "eth_sendTransaction" abstract implementation
// @param         from             string         "20 bytes - the source address of the sending transaction"
// @param         to               string         "20 bytes - the destination address of the transaction"
// @param         gas              string         "Available gas amount for transaction execution, optional integer, default value 90000, unused gas will be returned"
// @param         gasPrice         string         "gas price, optional, default value: To-Be-Determined"
// @param         value            string         "The amount sent by the transaction, an optional integer"
// @param         data             string         "Compilation of the contract or the signature and encoding parameters of the called method"
// @param         nonce            string         "nonce, optional. Rewriting of pending transactions can be achieved using the same nonce"
// @auth          happyboy
func (rpc *RPC) EthSendTransaction(from string, to string, gas int64, gasPrice int64, value int64, data string, nonce string, isContractCreation bool) string {
	tJsosn := structs.BuildTransaction(from, to, gas, gasPrice, value, data, nonce, isContractCreation).TransactionToJsonStr()
	resp, _ := http.Post(rpc.rpcUrl, CONTENT_TYPE, rpc.BuildJsonRpcMsg("eth_sendTransaction", []string{tJsosn}))
	body, _ := io.ReadAll(resp.Body)
	return string(body)
}

// @title         EthGetTransactionReceipt
// @description   rpc method "eth_getTransactionReceipt" abstract implementation
// @auth          happyboy
func (rpc *RPC) EthGetTransactionReceipt(address string) string {
	resp, _ := http.Post(rpc.rpcUrl, CONTENT_TYPE, rpc.BuildJsonRpcMsg("eth_getTransactionReceipt", []string{address}))
	body, _ := io.ReadAll(resp.Body)
	return string(body)
}

// @title         EthBlockNumber
// @description   rpc method "eth_blockNumber" abstract implementation
// @auth          happyboy
func (rpc *RPC) EthBlockNumber() *structs.JsonrpcResponse {
	resp, _ := http.Post(rpc.rpcUrl, CONTENT_TYPE, rpc.BuildJsonRpcMsg("eth_blockNumber", []string{}))
	body, _ := io.ReadAll(resp.Body)
	res := structs.NewJsonrpcResponse(string(body))
	res.Result = strconv.Itoa(int(utils.HexToDecimal(res.Result)))
	return res
}

// @title         EthNewBlockFilter
// @description   rpc method "eth_newBlockFilter" abstract implementation
// @auth          happyboy
func (rpc *RPC) EthNewBlockFilter() string {
	resp, _ := http.Post(rpc.rpcUrl, CONTENT_TYPE, rpc.BuildJsonRpcMsg("eth_newBlockFilter", []string{}))
	body, _ := io.ReadAll(resp.Body)
	//res := structs.NewJsonrpcResponse(string(body))
	//res.Result = strconv.Itoa(int(utils.HexToDecimal(res.Result)))
	return string(body)
}

