package main

import (
	"fmt"
	"sort"
)

type Cauthu struct {
	HT string
	SBT int
	SP int 
	T int
}

func danhsachcauthu() *[]Cauthu {
	var cts []Cauthu
	var n, SBT, SP, T int
	var HT string

	fmt.Print("Nhập số lượng cầu thủ: ")
	fmt.Scanln(&n)
	
	for i := 0; i < n; i++ {
		fmt.Printf("Họ tên cầu thủ %v: ", i+1)
		fmt.Scanln(&HT)

		fmt.Printf("Số bàn thắng cầu thủ %v: ", i+1)
		fmt.Scanln(&SBT)

		fmt.Printf("Số phút cầu thủ %v thi đấu: ", i+1)
		fmt.Scanln(&SP)

		if SP >= 500 && SBT >= 3 {
			T = 5000000
		}else if SP >= 500 || SBT >= 3 {
			T = 2000000
		}else {
			T = 0
		}

		ct := Cauthu {
			HT: HT,
			SBT: SBT,
			SP: SP,
			T: T,
		}

		cts = append(cts, ct)
	}
	return &cts
}

func main() {
	cts := *danhsachcauthu()

	for i, value := range(cts) {
		fmt.Printf("Họ tên cầu thủ %v: %v - Số bàn thắng: %v - Số phút: %v",i, value.HT, value.SBT, value.SP)
		fmt.Println("")
	}

	//danh sach cau thu co tien thuong cao nhat
	
	sort.SliceStable(cts, func(i, j int) bool {
		return cts[i].T > cts[j].T
	})
	fmt.Println("Những cầu thủ có số tiền thưởng cao nhất đội là:")
	for _, value := range(cts) {
		if value.T > 0 {
			fmt.Printf("Họ tên: %v - Số tiền thưởng: %v", value.HT, value.T)
			fmt.Println("")
		}
	}

}


