package structs

import (
	"encoding/json"
	"log"
)

// json rpc response struct
type JsonrpcResponse struct {
	Id      string `json:"id"`      // chain id
	Jsonrpc string `json:"jsonrpc"` // json rpc version
	Result  string `json:"result"`  // method name
}

func NewJsonrpcResponse(reponse string) *JsonrpcResponse {
	jsonrpcResponse := &JsonrpcResponse{}
	err := json.Unmarshal([]byte(reponse), jsonrpcResponse)
	if err != nil {
		log.Fatal("new jsonrpc response function error")
	}
	return jsonrpcResponse
}
