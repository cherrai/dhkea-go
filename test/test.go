package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/MetahorizonLab/dhkea-go"
)

func main() {
	var wg sync.WaitGroup
	currentTime := time.Now()
	digitsArr := []int{768, 1024, 1536, 2048, 3072, 4096, 6144, 8192}

	dhkea.EnableCache = true

	// // Input B PublicKey
	// var InputPublicKey string
	// fmt.Print("B PublicKey：")
	// fmt.Scanln(&InputPublicKey)
	// BPublicKeyBigint := big.NewInt(0)
	// BPublicKeyBigint.SetString(InputPublicKey, 10)
	// AKey := A.GetSharedKey(BPublicKeyBigint)

	successCount := 0
	ch := make(chan struct{}, runtime.NumCPU())

	for i := 0; i < 100; i++ {
		ch <- struct{}{}
		wg.Add(1)
		go func(count int, digits int) {
			currentTime := time.Now()
			A := dhkea.New(digits)
			B := dhkea.New(digits)
			AKey := A.GetSharedKey(B.PublicKey)
			// fmt.Println("A key", AKey)

			BKey := B.GetSharedKey(A.PublicKey)
			// fmt.Println("A.PublicKey", A.PublicKey)
			// fmt.Println("A.PriviteKey", A.PrivateKey)
			// fmt.Println("A key", AKey)
			// fmt.Println("B.PublicKey", B.PublicKey)
			// fmt.Println("B.PriviteKey", B.PrivateKey)
			// fmt.Println("B key", BKey)

			// fmt.Println("==============================")
			if AKey.String() == BKey.String() {
				successCount++
			}
			elapsed := time.Since(currentTime)
			fmt.Println("第"+strconv.Itoa(count)+"次计算", "位数："+strconv.Itoa(digits), "key值对比：", AKey.String() == BKey.String(), "耗时：", elapsed)
			defer wg.Done()
			<-ch
		}(i+1, digitsArr[i%8])
	}
	wg.Wait()
	elapsed := time.Since(currentTime)
	fmt.Println("总耗时：", elapsed)
	fmt.Println("成功次数：", successCount)
}
