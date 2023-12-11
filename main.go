package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"go.uber.org/zap"
)

var (
	count            int
	currentChallenge = common.Hex2Bytes("7245544800000000000000000000000000000000000000000000000000000000")
	strs_7777777     = "0x0077777777"
)

func main() {
	flag.IntVar(&count, "c", 0, "count")
	flag.Parse()
	if count <= 0 {
		flag.Usage()
		os.Exit(0)
	}

	cpuNum := runtime.NumCPU() - 1
	var strChan = make(chan string, 1)
	for i := 0; i < cpuNum; i++ {
		go func(sc chan<- string) {
			for {
				data := []byte(fmt.Sprintf(`data:application/json,{"p":"rerc-20","op":"mint","tick":"rETH","id":"%s","amt":"10000"}`, randHash()))
				sc <- hexutil.Bytes(data).String()
			}
		}(strChan)
	}

	for str := range strChan {
		fmt.Println(str)
		fmt.Println()
		count--
		if count <= 0 {
			break
		}
	}
}

func randHash() string {
	data := make([]byte, 32) // common.Hash的长度为32字节
	for {
		_, err := rand.Read(data)
		if err != nil {
			log.Fatal("create rand hash", zap.Error(err))
		}

		if strings.HasPrefix(crypto.Keccak256Hash(data, currentChallenge).Hex(), strs_7777777) {
			return common.BytesToHash(data).Hex()
		}
	}
}
