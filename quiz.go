/*
QUIZ GAMES 

@author : Arsybai

*/
package main

import (
  "bufio"
  "fmt"
  "os"
  "net/http"
  "net/url"
  "encoding/json"
  "io/ioutil"
  //"strings"
)

var soal = "null"
var kategori = "tehnologi"
var a, b, c, d string = "a", "b", "c", "d"
var right = "null"
var iscategory = false
var isplay = false
var score = 0


func main() {
  WhenRun()
  reader := bufio.NewReader(os.Stdin)
  for {
	fmt.Print("> ")
	masuk, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}
	switch masuk {
	case 'h' :
		HelperGame()
		break
	case 'x':
		KategoriSet("aaaa")
		iscategory = true
		break
	case 'i':
		KategoriSet("i")
		break
	case 't':
		KategoriSet("t")
		break
	case 's':
		if iscategory == true {
			KategoriSet("s")
		} else {
			fmt.Println("Score kamu saat ini adalah : ")
			fmt.Println(score)
		}
		break
	case 'k':
		KategoriSet("k")
		break
	case 'a':
		if isplay == true {
			WhenChoose("a")
		}
		break
	case 'b':
		if isplay == true {
			WhenChoose("b")
		}
		break
	case 'c':
		if isplay == true {
			WhenChoose("c")
		}
		break
	case 'd':
		if isplay == true {
			WhenChoose("d")
		}
		break
	case 'p':
		isplay = true
		NewGame()
		break
	} 
  }
}

func NewGame() {
	isplay = true
	Url, ers := url.Parse("http://arsyquiz.herokuapp.com/req/kat/" + kategori)
	if ers != nil {
		fmt.Println("Something Wrong :(\nCheck your connection or contact owner LINE : arsy22bai")
	}
	response, err := http.Get(Url.String())
	if err != nil {
		fmt.Println("Something Wrong :(\nCheck your connection or contact owner LINE : arsy22bai")
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		var result map[string]interface{}
		json.Unmarshal([]byte(data), &result)
		soal = result["soal"].(string)
		right = result["jawaban"].(map[string]interface{})["t"].(string)
		a = result["jawaban"].(map[string]interface{})["a"].(string)
		b = result["jawaban"].(map[string]interface{})["b"].(string)
		c = result["jawaban"].(map[string]interface{})["c"].(string)
		d = result["jawaban"].(map[string]interface{})["d"].(string)
		this_out := soal + "\n\n"
		this_out += "a. " + a
		this_out += "\nb. " + b
		this_out += "\nc. " + c
		this_out += "\nd. " + d
		this_out += "\n\npilih a, b, c atau d."
		fmt.Println(this_out)
	}

}

func WhenChoose(answer string){
	if isplay == true {
		if answer == "a" {
			if a == right {
				score += 10
				fmt.Println("Kamu benar! score + 10 :D")
			} else {
				fmt.Println("yaah kamu salah :( score - 5")
			}
		} else if answer == "b" {
			if b == right {
				score += 10
				fmt.Println("Kamu benar! score + 10 :D")
			} else {
				fmt.Println("yaah kamu salah :( score - 5")
			}
		} else if answer == "c" {
			if c == right {
				score += 10
				fmt.Println("Kamu benar! score + 10 :D")
			} else {
				fmt.Println("yaah kamu salah :( score - 5")
			}
		} else if answer == "d" {
			if d == right {
				score += 10
				fmt.Println("Kamu benar! score + 10 :D")
			} else {
				fmt.Println("yaah kamu salah :( score - 5")
			}
		}
		NewGame()
	}
}

func WhenRun() {
	var nga = "[ SELAMAT DATANG DI QUIZ ]\n\nGame ini merupakan game multichoice,\njika kalian menjawab benar maka score bertambah 10\njika salah maka score berkurang 5.\nKetik p untuk bermain atau h untuk menu\nselamat bermain :D\n\nmade by : arsybai.xyz"
	fmt.Println(nga)
}

func KategoriSet(kat string) {
	var i = "iman"
	var t = "tehnologi"
	var s = "seputardunia"
	var k = "kesehatan"
	metune := "[ PILIH KATEGORI ]\n"
	metune += "\ni  : iman"
	metune += "\nt  : tehnologi"
	metune += "\ns  : seputardunia"
	metune += "\nk  : kesehatan"
	if iscategory == true {
		if kat == "i" {
			kategori = i
			iscategory = false
			fmt.Println("Kategory Diubah Menjadi : " + i)
		} else if kat == "t" {
			kategori = t
			iscategory = false
			fmt.Println("Kategory Diubah Menjadi : " + t)
		} else if kat == "s" {
			kategori = s
			iscategory = false
			fmt.Println("Kategory Diubah Menjadi : " + s)
		} else if kat == "k" {
			kategori = k
			iscategory = false
			fmt.Println("Kategory Diubah Menjadi : " + k)
		} else {
			fmt.Println("Kategori tidak valid!\nGunakan i, t, s atau k!")
		}
	} else {
		fmt.Println(metune)
	}
}

func HelperGame() {
	mboh := "[ Games Menu ]\n"
	mboh += "\nh  : game menu"
	mboh += "\np  : mulai game"
	mboh += "\ns  : cek score"
	mboh += "\nx  : ganti kategori"
	fmt.Println(mboh)
}