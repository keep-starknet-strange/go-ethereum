package main

import (
	"bytes"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"strings"
)

func big_to_hex_le(num *big.Int) (string) {
	raw := num.Bytes()
	end := len(raw) - 1
	for i := 0; i < len(raw)/2; i++ {
		raw[i], raw[end-i] = raw[end-i], raw[i]
	}
	hex_repr := hex.EncodeToString(raw)
	return fmt.Sprintf("%s%s", hex_repr, strings.Repeat("0", 64 - len(hex_repr)))
}

func main() {
	const n_args int = 2 
	flag.Parse()
	if flag.NArg() != n_args {
		fmt.Printf("Invalid arguments length: takes %d values, received %d\n", n_args, flag.NArg())
		os.Exit(1)
	}
	modulus := new(big.Int)
	modulus, _ = modulus.SetString("3618502788666131213697322783095070105623107215331596699973092056135872020481", 0)
	val := make([]string, n_args)
	num := new(big.Int)
	for i := 0; i < n_args; i++ {
		num, ok := num.SetString(flag.Arg(i), 0)
		if !ok {
			fmt.Printf("Invalid input %d: should be an integer\n", i)
			os.Exit(2)
		}
		num = num.Mod(num, modulus)
		val[i] = big_to_hex_le(num)
	}

	data := "0x" + val[0] + val[1]
	from := "0x2fcf618717eDa0Eb623Fe3AE0bfBF115759f1c9d"
	to := "0x000000000000000000000000000000000000000a"

	jsonStr := fmt.Sprintf(`{"method":"eth_call","params":[{"from":"%s","to":"%s","data":"%s"}, "latest"],"id":1,"jsonrpc":"2.0"}`, from, to, data)
	jsonData := []byte(jsonStr)

	url := "http://localhost:8545"
	request, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	request.Header.Set("content-type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("Error with the HTTP request %v\n", request)
		os.Exit(3)
	}
	defer response.Body.Close()
	
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Printf("Output: %s\n", string(body))
}
