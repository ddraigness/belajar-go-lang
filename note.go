/**
	Error Package dan Importt
Di Go-Lang versi terbaru, kita sudah diwajibkan menggunakan Go-Modules, yang akan dijelaskan pada PART 2 kelas ini

Jadi pada video selanjutnya, pasti akan mendapat error

Namun jika ingin mencoba fitur lama Go-Lang, sebelum lanjutkan video selanjutnya, silahkan matikan default fitur Go-Modules dengan menggunakan perintah

go env -w  GO111MODULE=off
*/

/**
	sebelum ada go modules :
1. saat kita membuat aplikasi, biasanya kita akan menggunakan library atau dependency dari project lain.
2. sebelum ada go modules, management untuk dependency sangat sulit dilakukan di go-lang.
3. biasanya kita akan meng-copy semua source code library atau dependency lain ke project kita, hal ini
	membuat project kita menjadi bengkak karena penuh dengan library orang lain.
4. atau biasanya library orang lain akan kita save di GOPATH directory, namun hal ini akan sulit jika 
	ternyata beberapa aplikasi butuh library yang sama dengan versi yang berbeda
*/