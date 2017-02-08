package main

import (
	"fmt"
//	"os"
)

/*
const (
	i = 100
	pi =3.1415
	prefix  = "GO_"
)

var (
	i int
	pi float32
	prefix string
)
*/

const (
	x = iota //x==0
	y = iota //y==1
	z = iota //z==2
	w        //w==3
)

const v  =  iota //v==0
const (
	h,i,j = iota,iota,iota //h=0,i=o,j=0
)
const(
	a	= iota //a=0
	b	= "B"
	c	= iota
	d,e,f =iota,iota,iota //d=3,e=3,f=3
	g	= iota  //g =4
)

func main() {
	h := "hello\t"
	w := "go"
	hw := h + w
	fmt.Println("hello ,go!2017")
	fmt.Println("%s\n",hw)
	a := `hello
	go
	eagle!`
	//fmt.Printf(a)
	fmt.Println(a,b,c,d,e,f,g,h,i,j,x,y,z,w,v)
}

// test go
