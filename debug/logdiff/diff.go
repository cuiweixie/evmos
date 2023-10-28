package main

import (
	"fmt"
	"github.com/evmos/evmos/v15/debuglog"
)

func main() {
	estimateRecord, err := debuglog.LoadRecordsFromFile("/Users/xiecui/debug/estimate_216147.log")
	if err != nil {
		panic(err)
	}
	txRecord, err := debuglog.LoadRecordsFromFile("/Users/xiecui/debug/tx_200000.log")
	if err != nil {
		panic(err)
	}
	fmt.Println("e len:", len(estimateRecord), "t len:", len(txRecord))
	eSum := uint64(0)
	tSum := uint64(0)
	for i := 0; i < len(estimateRecord); i++ {
		if estimateRecord[i].UsedGas != txRecord[i].UsedGas {
			fmt.Println("i", i, "e:", estimateRecord[i].UsedGas, "t:", txRecord[i].UsedGas)
		}
		//if i == 37 {
		//	fmt.Println(estimateRecord[i].Stack)
		//	fmt.Println("====================================")
		//	fmt.Println(txRecord[i].Stack)
		//	break
		//}
		eSum += estimateRecord[i].UsedGas
	}

	for i := 0; i < len(txRecord); i++ {
		tSum += txRecord[i].UsedGas
	}

	println("eSum:", eSum, "tSum:", tSum)
}
