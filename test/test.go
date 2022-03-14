package main

import (
	"fmt"
	"time"

	"github.com/MetahorizonLab/dhkea-go"
)

func main() {
	currentTime := time.Now()
	digits := 1024

	A := dhkea.New(digits)
	B := dhkea.New(digits)

	fmt.Println("A Publickey：", A.PublicKey)

	// // Input B PublicKey
	// var InputPublicKey string
	// fmt.Print("B PublicKey：")
	// fmt.Scanln(&InputPublicKey)
	// BPublicKeyBigint := big.NewInt(0)
	// BPublicKeyBigint.SetString(InputPublicKey, 10)
	// AKey := A.GetSharedKey(BPublicKeyBigint)

	AKey := A.GetSharedKey(B.PublicKey)
	fmt.Println("A key", AKey)
	// fmt.Println("A key", BPublicKey)

	// BKey := B.GetSharedKey(A.PublicKey)
	// fmt.Println("B key", BKey)
	// fmt.Println(AKey.String() == BKey.String())
	elapsed := time.Since(currentTime)
	fmt.Println("time：", elapsed)
}
