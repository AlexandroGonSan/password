package main

//https://golang.org/pkg/unicode/utf8/

/*
https://www.ime.usp.br/~pf/algoritmos/aulas/footnotes/utf-8.html
A tabela abaixo mostra a estrutura do código UTF-8. Na coluna esquerda, temos os intervalos de números Unicode, em notação hexadecimal. Na coluna direita, em notação binária, os correspondentes valores válidos dos bytes do código:

números Unicode       byte 1   byte 2   byte 3   byte 4

00000000 .. 0000007F  0xxxxxxx
00000080 .. 000007FF  110xxxxx 10xxxxxx
00000800 .. 0000FFFF  1110xxxx 10xxxxxx 10xxxxxx
00010000 .. 0010FFFF  11110xxx 10xxxxxx 10xxxxxx 10xxxxxx
Agora, a mesma tabela, com os números Unicode e os intervalos de valores dos bytes de código escritos em notação decimal:

       0 .. 127       000..127    
     128 .. 2047      192..223 128..191 
    2048 .. 65535     224..239 128..191 128..191 
   65536 .. 1114111   240..247 128..191 128..191 128..191 
Finalmente, a mesma tabela, com os números Unicode e os intervalos de valores dos bytes escritos em notação hexadecimal:

       0 .. 7F        00..7F    
      80 .. 7FF       C0..DF   80..BF 
     800 .. FFFF      E0..EF   80..BF   80..BF 
   10000 .. 10FFFF    F0..F7   80..BF   80..BF   80..BF
*/

import (
	"fmt"
	"unicode/utf8"
	"math/rand"
	"time"
	"strconv"
	"os"
)

func main() {
	quant_bytes, _ := strconv.Atoi(os.Args[1])
	if quant_bytes < 8 {
		fmt.Printf("A quantidade de bytes requerida precisa " +
			"ser maior ou igual a 8.")
	}
	i := 0
	count_bytes := 0
	debug := false
	for i < 100000{
		rand.Seed(time.Now().UnixNano())
		aux := rand.Intn(100)
		if aux < 5 {
			rand.Seed(time.Now().UnixNano())
			b := []byte{byte(rand.Intn(128))}
			if utf8.Valid(b){
				r, size := utf8.DecodeLastRune(b)
				_ = size
				if debug {
					fmt.Print(b)
				}
				if strconv.IsGraphic(r) && strconv.IsPrint(r) {
					fmt.Printf("%c", r)
					count_bytes++
					i++
				}
			}
		} else if aux < 20 {
			rand.Seed(time.Now().UnixNano())
			b1 := byte(rand.Intn(224-128)+128)
			rand.Seed(time.Now().UnixNano())
			b2 := byte(rand.Intn(192-128)+128)
			b := []byte{b1,b2}
			if utf8.Valid(b){
				r, size := utf8.DecodeLastRune(b)
				_ = size
				if debug {
					fmt.Print(b)
				}
				if strconv.IsGraphic(r) && strconv.IsPrint(r) {
					count_bytes += 2
					fmt.Printf("%c", r)
					i++
				}
			}
		} else if aux < 50 {
			rand.Seed(time.Now().UnixNano())
			b1 := byte(rand.Intn(240-224)+224)
			rand.Seed(time.Now().UnixNano())
			b2 := byte(rand.Intn(192-128)+128)
			rand.Seed(time.Now().UnixNano())
			b3 := byte(rand.Intn(192-128)+128)
			b := []byte{b1,b2,b3}
			if utf8.Valid(b){
				r, size := utf8.DecodeLastRune(b)
				_ = size
				if debug {
					fmt.Print(b)
				}
				if strconv.IsGraphic(r) && strconv.IsPrint(r) {
					count_bytes += 3
					fmt.Printf("%c", r)
					i++
				}
			}
		} else {
			rand.Seed(time.Now().UnixNano())
			b1 := byte(rand.Intn(248-240)+240)
			rand.Seed(time.Now().UnixNano())
			b2 := byte(rand.Intn(192-128)+128)
			rand.Seed(time.Now().UnixNano())
			b3 := byte(rand.Intn(192-128)+128)
			rand.Seed(time.Now().UnixNano())
			b4 := byte(rand.Intn(192-128)+128)
			b := []byte{b1,b2,b3,b4}
			if utf8.Valid(b){
				r, size := utf8.DecodeLastRune(b)
				_ = size
				if debug {
					fmt.Print(b)
				}
				if strconv.IsGraphic(r) && strconv.IsPrint(r) {
					fmt.Printf("%c", r)
					count_bytes += 4
					i++
				}
			}
		}
		// if i <= 1 {
			// asset count_bytes 
		// }
		if( count_bytes > quant_bytes-4 ){
		// if( i == quant_bytes ){
			break
		}
	}
	fmt.Println("\n\nOs caracteres quadrados sao diferentes .")
	fmt.Println("Para notar isso cole dois no mousepad e veja que os\n"+
		"...números escritos dentro de cada quadrado são diferentes .")
}
