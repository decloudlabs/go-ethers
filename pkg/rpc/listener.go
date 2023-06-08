// @Title        rpc
// @Description  Abstract implementation of network rpc service for evm architecture
// @Author       happyboy
package rpc

import (
	"github.com/decloudlabs/go-ethers/pkg/structs"
	"strconv"
)

// @title         NewBlockNumberListener
// @description   new block listener
// @param         c            chan         "new block listener channel"
// @auth          happyboy
func (rpc *RPC) NewBlockNumberListener(c chan structs.JsonrpcResponse) {
	i, _ := strconv.Atoi(rpc.EthBlockNumber().Result)
	c <- *rpc.EthBlockNumber()
	for {
		i2, _ := strconv.Atoi(rpc.EthBlockNumber().Result)
		if i2 > i {
			c <- *rpc.EthBlockNumber()
			i = i2
			continue
		}
	}
}
