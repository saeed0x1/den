package encodings

import (
	"encoding/base64"
	"fmt"
	"html"
	"net/url"
)

type E struct {
	S string
}

type Convert interface {
	ToURL()
	ToHTML()
	ToB64()
	ToHex()
	ToBin()
	ToDec()
	ToAll()
}

// to base64 encoding
func (d *E) ToBase64() {
	eb64 := base64.StdEncoding.EncodeToString([]byte(d.S))
	fmt.Println(eb64)
}

// to URL Encoding
func (d *E) ToURL() {

	eURL := url.QueryEscape(d.S)

	fmt.Println(eURL)
}

// to hex Encoding
func (d *E) ToHex() {
	// ehex := hex.EncodeToString([]byte(d.S))
	fmt.Printf("%x\n", d.S)
}

// to HTML Encoding
func (d *E) ToHTML() {

	ehtml := html.EscapeString(d.S)

	fmt.Println(ehtml)

}

// to binary Encoding
func (d *E) ToBin() {
	for _, v := range d.S {
		fmt.Printf("%b ", v)
	}
}

// to Decimal Encoding
func (d *E) ToDec() {
	addchars := d.S

	for i := 0; i < len(addchars); i++ {

		fmt.Printf("%d ", addchars[i])

	}

}

func (d *E) ToAll() {
	d.ToURL()
	d.ToHTML()
	d.ToBase64()
	d.ToHex()
	d.ToBin()
	d.ToDec()
}
