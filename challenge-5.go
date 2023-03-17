package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// var mu sync.Mutex // akan diaktifkan jika data yang diinginkan selang-seling

	bisa := []interface{}{"bisa1", "bisa2", "bisa3"}
	coba := []interface{}{"coba1", "coba2", "coba3"}

	// versi arnanda asshafa (PoiKyun#7350)
	// for i := 0; i < 4; i++ {
	// 	wg.Add(1)
		// go printData1(bisa, coba, i, &wg, &mu) 
		// go printData2(bisa, coba, i, &wg) 
	// }

	// wg.Wait()

	// versi mas-mas stranger di discord
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go func(id int) {
			// nonaktifkan mutex jika ingin secara acak
			// mu.Lock()
			// defer mu.Unlock()
			fmt.Printf("%v %d\n", coba, id)
			fmt.Printf("%v %d\n", bisa, id)
			wg.Done()
		}(i)
	}

	wg.Wait()

	// versi mas rizqy
	//Print Interface Acak
	// printInterface(coba, bisa, &wg, &mu, "acak")
	// fmt.Println()
	//Print Interface Rapih
	// printInterface(coba, bisa, &wg, &mu, "rapih")

}

// versi arnanda asshafa (PoiKyun#7350)
func printData1(data1, data2 interface{}, idx int, wg *sync.WaitGroup, mu *sync.Mutex) { 
    mu.Lock()
    fmt.Printf("%v %d\n", data1, idx+1)
    fmt.Printf("%v %d\n", data2, idx+1)
    mu.Unlock()
    wg.Done()
}

func printData2(data1, data2 interface{}, idx int, wg *sync.WaitGroup) { 
    fmt.Printf("%v %d\n", data1, idx+1)
    fmt.Printf("%v %d\n", data2, idx+1)
    wg.Done()
}

// mas rizqy
func printInterface(coba []interface{}, bisa []interface{}, wg *sync.WaitGroup, mu *sync.Mutex, tipe string) {
	state := true
	wg.Add(8)
loop:
	for i := 1; i <= 4; i++ {
		if tipe == "rapih" {
			mu.Lock()
		}
		go func(i int) {
			defer wg.Done()
			if i%2 == 0 {
				fmt.Println(bisa, i)
			} else {
				fmt.Println(coba, i)
			}
			if tipe == "rapih" {
				mu.Unlock()
			}
		}(i)
		if i == 4 && state {
			state = false
			goto loop
		}
	}
	wg.Wait()
}

