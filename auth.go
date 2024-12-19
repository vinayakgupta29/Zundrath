package main

import "net/http"

func AuthoriseRequest(h *http.Header) bool {
	client := h.Get("X-Client")
	authHeader := h.Get("Authorization")
	var mkp MasterKey
	hmac := GetHMAC256(client, string(mkp.GetMasterKey()))
	return hmac == authHeader

}
