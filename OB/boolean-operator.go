package ob

import "fmt"

func Ob() {
	var (
		ujian   = 80
		absensi = 80

		lulusUjian   = ujian >= 80
		lulusAbsensi = absensi >= 80

		lulus = lulusUjian && lulusAbsensi
	)

	fmt.Println(lulus)
}