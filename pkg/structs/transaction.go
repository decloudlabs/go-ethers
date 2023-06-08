package structs

import (
	"encoding/json"
	"fmt"
	"github.com/decloudlabs/go-ethers"
	"github.com/decloudlabs/go-ethers/pkg/utils"
	"log"
)

type Transaction struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Gas      string `json:"gas"`
	GasPrice string `json:"gas_price"`
	Value    string `json:"value"`
	Data     string `json:"data"`
	Nonce    string `json:"nonce"`
}

func BuildTransaction(from string, to string, gas int64, gasPrice int64, value int64, data string, nonce string, isContractCreation bool) *Transaction {
	if isContractCreation {
		if utils.StringIsNull(from, to, data) {
			log.Fatal("params cant null")
		}
	} else {
		if utils.StringIsNull(from, data) {
			log.Fatal("params cant null")
		}
	}
	t := &Transaction{}
	t.From = from
	t.To = to
	t.Gas = utils.DecimalToHex(gas)
	t.GasPrice = utils.DecimalToHex(gasPrice)
	t.Value = utils.DecimalToHex(value)
	t.Data = data
	if !utils.StringIsNull(nonce) {
		t.Nonce = nonce
	}
	return t
}

func (t *Transaction) TransactionToJsonStr() string {
	jsons, errs := json.Marshal(*t)
	if errs != nil {
		fmt.Println(errs.Error())
	}
	return string(jsons)
}
