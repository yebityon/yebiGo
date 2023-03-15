package main 

import (
	"strconv"
	"bufio"
	"os"
)


func nextInt() int {
	var sc = bufio.NewScanner(os.Stdin)
	sc.Scan()
	i, e := strconv.Atoi((sc.Text()))
	if e != nil {
		panic(e)
	}
	return int(i)
}

func nextInt64() int64 {
	
	var sc = bufio.NewScanner(os.Stdin)
	sc.Scan()
	i, e := strconv.Atoi((sc.Text()))
	if e != nil {
		panic(e)
	}
	return int64(i)
}
