package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/crypto"
)

// start OMIT
// const { getContractAddress } = require('@ethersproject/address')

func getContractAddress(addr common.Address, nonce uint64) string {
	return crypto.CreateAddress(addr, nonce).Hex()
}

func main() {
	addr := common.HexToAddress("0x36448d41Feb5e08001255dbF72B7Fed3F6d483E2")
	for nonce := uint64(0); nonce < 3; nonce++ {
		fmt.Println(
			nonce,
			getContractAddress(addr, nonce),
		)
	}
}

// end OMIT
/*
b4 OMIT
0 0x7B3c35D7D339A3A0f70D06b8852a35D674bEB1dC
1 0x08456D6f845Af15bFb0502deDbc63c6DC62cdb9e
2 0x1e8150050A7a4715aad42b905C08df76883f396F
a4 OMIT
*/
