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
	AGI := A.GenerateIndividualKey()
	// fmt.Println("GenerateIndividualKey:", A.PrivateKey)

	BGI := B.GenerateIndividualKey()
	// fmt.Println("GenerateIndividualKey:", B)

	AKey := A.GetSharedKey(BGI.PublicKey)
	fmt.Println("A key", AKey)

	BKey := B.GetSharedKey(AGI.PublicKey)
	fmt.Println("B key", BKey)
	fmt.Println(AKey.String() == BKey.String())
	elapsed := time.Since(currentTime)
	fmt.Println("time：", elapsed)
}
