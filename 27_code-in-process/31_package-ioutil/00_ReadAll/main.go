package main

import (
	"strings"
	"io/ioutil"
	"log"
	"fmt"
)

func main() {
	myPhrase := "mmm, locorice"
	rdr := strings.NewReader(myPhrase)

	bs, err := ioutil.ReadAll(rdr)
	if err != nil {
		log.Fatalln("my program broke again")
	}
	str := string(bs)
	fmt.Println(str)
}
