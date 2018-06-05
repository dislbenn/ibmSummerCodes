package main

import (
	"net/http"
	"strings"
)

func mian(){
	http.ListenAndServe(":8080", nil)
}