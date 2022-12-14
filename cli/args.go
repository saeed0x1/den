package cli

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	decode "github.com/zerodayrat/den/decoding"
	encode "github.com/zerodayrat/den/encodings"
)

func Init() {

	base64 := flag.Bool("base64", false, "Encode in Base64")
	hex := flag.Bool("hex", false, "Encode in Hexadecimal")
	url := flag.Bool("url", false, "Encode in URL")
	html := flag.Bool("html", false, "Encode in HTML")
	bin := flag.Bool("bin", false, "Encode in Binary")
	dec := flag.Bool("dec", false, "Encode in Decimal")
	all := flag.Bool("all", false, "Encode in all Available Format")

	dbase64 := flag.Bool("dbase64", false, "Decode from Base64")
	dhex := flag.Bool("dhex", false, "Decode from Hex")
	durl := flag.Bool("durl", false, "Decode from URL")
	dhtml := flag.Bool("dhtml", false, "Decode from HTML")
	dbin := flag.Bool("dbin", false, "Decode from Binary")
	ddec := flag.Bool("ddec", false, "Decode from Decimal")
	dall := flag.Bool("dall", false, "Decode in all Available Formats")

	flag.Parse()

	it := bufio.NewScanner(os.Stdin)

	// check if any flag has passed or not
	// if no flag is passed it'll print the tool usage
	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(0)
	} else if flag.NFlag() > 1 {
		fmt.Println("Error : Pass only one flag at a time.")
		fmt.Println()
		flag.Usage()
		os.Exit(0)
	} else {

		for it.Scan() {
			e := encode.E{S: it.Text()}
			de := decode.DE{S: it.Text()}

			switch {

			// encoding
			case *url:
				e.ToURL()
			case *html:
				e.ToHTML()
			case *base64:
				e.ToBase64()
			case *hex:
				e.ToHex()
			case *bin:
				e.ToBin()
			case *dec:
				e.ToDec()
			case *all:
				e.ToAll()

			// decoding
			case *durl:
				de.FromURL()
			case *dhtml:
				de.FromHTML()
			case *dbase64:
				de.FromBase64()
			case *dhex:
				de.FromHex()
			case *dbin:
				de.FromBin()
			case *ddec:
				de.FromDec()
			case *dall:
				de.ForAll()
			}
		}
	}

}
