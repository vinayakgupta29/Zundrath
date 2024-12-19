package main

import "net/http"

func AuthoriseRequest(h *http.Header) bool {
	client := h.Get("x-c")
	authHeader := h.Get("Authorization")
	hmac := GetHMAC256(&client)
	return hmac == authHeader

}
