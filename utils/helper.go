package utils

import (
	"net/url"
	"os"
)

func Getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func UrlValidation(urladdress string) bool {
	_, err := url.ParseRequestURI(urladdress)
	res := true
	if err != nil {
		res = false
	}
	return res
}
