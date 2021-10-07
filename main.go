package main

import (
	"net/http"
	"strconv"
)

var uni Universal

func main() {
	uni.Get()
	mux := http.NewServeMux()
	mux.HandleFunc("/", HandleConn)
	http.ListenAndServe(":"+strconv.Itoa(uni.Conf.BotConfig.Post), mux)
}
