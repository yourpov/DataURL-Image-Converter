package main

import (
	"fmt"
	"os"
	"strings"

	"img64/decode"
	"img64/encode"
)

func usage() {
	fmt.Println("Usage:")
	fmt.Println(" encode <image> <file>")
	fmt.Println(" decode <file> <image>")
	os.Exit(1)
}

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		usage()
	}
	mode := args[0]
	input := args[1]
	output := args[2]

	switch mode {
	case "encode":
		fmt.Println("- encoding:", input, "->", output)
		dataURL, err := encode.ToDataURL(input)
		if err != nil {
			fmt.Println("encode err:", err)
			os.Exit(1)
		}
		err = os.WriteFile(output, []byte(dataURL), 0644)
		if err != nil {
			fmt.Println("err writing:", err)
			os.Exit(1)
		}
		fmt.Println("output:", output)
	case "decode":
		fmt.Println("- decoding:", input, "->", output)
		b, err := os.ReadFile(input)
		if err != nil {
			fmt.Println("cannot read file input:", err)
			os.Exit(1)
		}
		data := strings.TrimSpace(string(b))
		f, err := os.Create(output)
		if err != nil {
			fmt.Println("cannot create file:", err)
			os.Exit(1)
		}
		f.Close()
		err = decode.FromDataURL(data, output)
		if err != nil {
			fmt.Println("decode err:", err)
			os.Exit(1)
		}
		fmt.Println("output:", output)
	default:
		usage()
	}
}
