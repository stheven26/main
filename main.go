package main

import (
	"fmt"
	"time"

	module "github.com/stheven26/module"
)

type Author struct {
	Nama, Password string
}

func main() {
	authorAccess := Author{
		Nama:     "stheven",
		Password: "123456",
	}

	blokirAkses := func(nama, password string) bool {
		return nama != "stheven"
	}
	module.VerifikasiAkses(authorAccess.Nama, authorAccess.Password, blokirAkses)

	if authorAccess.Nama == "stheven" && authorAccess.Password == "123456" {
		fmt.Print("\nActivity:\n1.Belajar\n2.Melakukan Konversi\n3.Berbelanja\n4.Operasi Matematika\n")
		fmt.Print("Enter Option Activity: ")
		var OptionActivity int
		fmt.Scan(&OptionActivity)
		fmt.Println("")

		if OptionActivity == 1 {
			module.Belajar()
		} else if OptionActivity == 2 {
			fmt.Print("Convert:\n1.Convert Feet to Meter\n2.Convert Mile to KM\n3.Convert Liter to CC\n")
			fmt.Print("Enter Option Convert: ")
			var OptionConvert int
			fmt.Scan(&OptionConvert)
			fmt.Println("")

			if OptionConvert == 1 {
				fmt.Print("Enter Feet: ")
				var feet float64
				fmt.Scan(&feet)
				meter := module.FeetToMeter(feet)
				fmt.Printf("'Hasil Konversi %.2f Feet adalah %.2f Meter'\n", feet, meter)
			} else if OptionConvert == 2 {
				fmt.Print("Enter Mile: ")
				var Mile float64
				fmt.Scan(&Mile)
				km := module.MileToKM(Mile)
				fmt.Printf("'Hasil Konversi %.2f Mile adalah %.2f KM\n'", Mile, km)
			} else if OptionConvert == 3 {
				fmt.Print("Enter CC: ")
				var CC float64
				fmt.Scan(&CC)
				liter := module.CCToLiter(CC)
				fmt.Printf("'Hasil Konversi %.1f CC adalah %.2f Liter'\n", CC, liter)
			} else {
				fmt.Println("Masukkan pilihan konversi anda dengan benar!")
			}
		} else if OptionActivity == 3 {
			go module.Belanja()
			time.Sleep(1 * time.Millisecond)
			go module.Membayar()
			time.Sleep(3 * time.Millisecond)
			go module.SelesaiBelanja()
			time.Sleep(5 * time.Millisecond)
		} else if OptionActivity == 4 {
			module.Matematika()
		} else {
			fmt.Println("Masukkan pilihan aktivitas anda dengan benar!")
		}
	} else if authorAccess.Nama == "stheven" && authorAccess.Password != "123456" {
		fmt.Println("'Masukkan password dengan benar!'")
	} else if authorAccess.Nama != "stheven" && authorAccess.Password == "123456" {
		fmt.Println("Masukkan nama dengan benar!'")
	} else {
		fmt.Print("Anda tidak memiliki akses!'")
	}
}
