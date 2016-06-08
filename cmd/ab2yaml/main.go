package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/MendelGusmao/ab2yaml"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(2)
	}

	content, err := ioutil.ReadFile(os.Args[1])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	content, err = ab2yaml.Convert(content)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%s", content)
}
