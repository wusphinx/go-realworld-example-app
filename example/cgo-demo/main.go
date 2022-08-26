package main

/*
#include <stdio.h>

void printint(int v) {
    printf("printint: %d\n", v);
}
*/
import "C"

func inline() {
	v := 1024
	//nolint
	C.printint(C.int(v))
}

func main() {
	inline()
}
