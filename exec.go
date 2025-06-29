package main

import (
	"fmt"
	"os"
	"time"
)

type Rajesh struct {
	rajesh *os.File
}

func (r *Rajesh) Info(comm string) {
	fmt.Fprintf(r.rajesh, "%v-info :%s", time.Now(), comm)
}

func creatorrk(filename string) *Rajesh {
	rajesh, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}
	return &Rajesh{rajesh}
}

func main() {
	var intro *Rajesh
	intro = creatorrk("rk.txt")
	intro.Info("we are on fire buddy")

}
