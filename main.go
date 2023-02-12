package main

import (
	"fmt"
	"strconv"
	"time"
)

type Produk struct {
	Nama    string
	Harga   int
	Dp      int
	Ansuran []int
}

type User struct {
	Name  string
	Saldo int
}
type Ysue struct {
}

func (credit *Produk) Cicilan(val int) {
	arr := make([]int, val)
	harga := (credit.Harga - credit.Dp) / val
	for i := range arr {
		arr[i] = harga
	}
	credit.Ansuran = arr
}

func Buy(user int, produk int) int {
	fmt.Println("MELAKUKAN PEMBELIAN")
	time.Sleep(2 * time.Second)
	if user < produk {
		fmt.Println("MAAF SALDO ANDA TIDAK MENCUKUPI")
	} else {
		user -= produk
		fmt.Println("SUKSES")
	}
	return user
}

func (credit *Produk) properti() map[string]interface{} {
	data := map[string]interface{}{
		"Produk": credit.Nama,
		"Harga":  credit.Harga,
		"DP":     credit.Dp,
		"Cicilan " + strconv.Itoa(len(credit.Ansuran)) + " bulan": credit.Ansuran,
	}

	return data
}

func BayarAnsuran(bulan, saldo int, cicilan []int) int {
	fmt.Println("MEMBAYAR ANSURAN")
	time.Sleep(1 * time.Second)
	count := 0
	for i := 0; i < bulan; i++ {
		if cicilan[i] != 0 {
			if saldo > cicilan[i] && bulan <= len(cicilan) {
				saldo -= cicilan[i]
				cicilan[i] = 0
				fmt.Println("Bulan ", i+1, " sudah lunas")
				count++
				time.Sleep(1 * time.Second)
			} else {
				fmt.Println("Maaf saldo anda tidak mencukupi")
				break
			}
		}
	}
	// fmt.Println("SISA ANSURAN ", len(cicilan)-count, " Bulan lagi ", cicilan)
	if len(cicilan)-count == 0 {
		fmt.Println("pembayaran sudah lunas")
	} else {
		fmt.Println("SISA ANSURAN ", len(cicilan)-count, " Bulan lagi ", cicilan)
	}

	return saldo
}

func main() {
	credit := Produk{
		Nama:  "KLX 160cc",
		Harga: 220000000,
		Dp:    3500000,
	}
	// credit.name("KlX 160cc")
	// credit.price(22000000)
	// credit.DP(3500000)
	credit.Cicilan(24)

	oke := credit.properti()
	for k, v := range oke {
		fmt.Println(k, " : ", v)
		time.Sleep(1 * time.Second)
	}

	user := User{
		Name:  "EFRI",
		Saldo: 5000000000,
	}

	time.Sleep(2 * time.Second)
	user.Saldo = Buy(user.Saldo, credit.Dp)
	fmt.Println(user)

	time.Sleep(1 * time.Second)
	user.Saldo = BayarAnsuran(24, user.Saldo, credit.Ansuran)
	fmt.Println(user)
	// fmt.Sprintf()
}
