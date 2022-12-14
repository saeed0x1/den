package decode

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"html"
	"net/url"
	"strconv"
	"strings"
)

type DE struct {
	S string
}

type DecodeMethods interface {
	FromURL()
	FromHTML()
	FromBase64()
	FromHex()
	FromBin()
	FromDec()
	ForAll()
}

func (d *DE) FromURL() {
	decodeurl, _ := url.QueryUnescape(d.S)
	fmt.Println(decodeurl)
}

func (d *DE) FromHTML() {
	decodehtml := html.UnescapeString(d.S)

	fmt.Println(decodehtml)
}

func (d *DE) FromBase64() {
	decodeb64, _ := base64.StdEncoding.DecodeString(d.S)

	fmt.Println(string(decodeb64))
}

func (d *DE) FromHex() {
	b, _ := hex.DecodeString(d.S)

	fmt.Println(string(b))
}

func (d *DE) FromBin() {
	splitbin := strings.Split(d.S, " ")

	for _, v := range splitbin {

		// binary to decimal
		eachchar, _ := strconv.ParseInt(v, 2, 0)
		fmt.Printf("%s", string(eachchar))
	}

}

func (d *DE) FromDec() {

	splited := strings.Split(d.S, " ")

	for _, v := range splited {
		conv, _ := strconv.Atoi(v)

		fmt.Printf("%s", string(conv))
	}
}

func (d *DE) ForAll() {
	d.FromURL()
	d.FromHTML()
	d.FromBase64()
	d.FromHex()
	d.FromBin()
	d.FromDec()
}
