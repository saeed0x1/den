# DEN 

den is a command-line tool for encoding and decoding strings. It provides various flags that allow you to specify the type of encoding or decoding you want to perform. 

## Usage

```
Usage of den:

  -all
        Encode in all Available Format
  -base64
        Encode in Base64
  -bin
        Encode in Binary
  -dec
        Encode in Decimal
  -hex
        Encode in Hexadecimal
  -html
        Encode in HTML
  -url
        Encode in URL
  -dall
        Decode in all Available Formats
  -dbase64
        Decode from Base64
  -dbin
        Decode from Binary
  -ddec
        Decode from Decimal
  -dhex
        Decode from Hex
  -dhtml
        Decode from HTML
  -durl
        Decode from URL
```

### Syntax

```go
# for single line strings

echo "string" | den -<encoding/decoding format>

# for multiple lines put the string in a text file

cat file.txt | den -<encoding/decoding format>
```

## Examples

To encode a string in base64, use the `-base64` flag:

```
echo "hii" | den -base64
```

To decode a string from base64, use the `-dbase64` flag:

```
echo "SGVsbG8sIHdvcmxkIQ==" | den -dbase64
```

To encode a string in all available formats, use the `-all` flag:

```
echo "hi" | den -all
```

## License

den is released under the MIT License. See [LICENSE](https://raw.githubusercontent.com/zerodayrat/den/main/LICENSE) for details.
