## Unit Test Golang

#### Available Method From Unit Test Golang
Library: "testing"
* Log() -> Menampilkan log
* Logf() -> Menampilkan log menggunakan format
* Fail() ->	Menandakan terjadi Fail() dan proses testing fungsi tetap diteruskan
* FailNow() ->	Menandakan terjadi Fail() dan proses testing fungsi dihentikan
* Failed() ->	Menampilkan laporan fail
* Error() ->	Log() diikuti dengan Fail()
* Errorf() ->	Logf() diikuti dengan Fail()
* Fatal() ->	Log() diikuti dengan failNow()
* Fatalf() ->	Logf() diikuti dengan failNow()
* Skip() ->	Log() diikuti dengan SkipNow()
* Skipf() ->	Logf() diikuti dengan SkipNow()
* SkipNow() ->	Menghentikan proses testing fungsi, dilanjutkan ke testing fungsi setelahnya
* Skiped() ->	Menampilkan laporan skip
* Parallel() ->	Menge-set bahwa eksekusi testing adalah parallel

#### Testing Use Testify
Using testify should have to download the library from github.com/stretchr/testify
``` go get github.com/stretchr/testify ```

Available package:
* assert ->	Berisikan tools standar untuk testing
* http ->	Berisikan tools untuk keperluan testing http
* mock ->	Berisikan tools untuk mocking object
* require ->	Sama seperti assert, hanya saja jika terjadi fail pada saat test akan menghentikan eksekusi program
* suite ->	Berisikan tools testing yang berhubungan dengan struct dan method