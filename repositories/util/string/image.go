package string

import (
	"github.com/spf13/viper"
	"regexp"
)

func CheckAvaTar(avaTar string) string {
	if avaTar == "" {
		return avaTar
	}
	host := viper.GetString("img_url")
	if ok, _ := regexp.Match(`((https|http)://)`, []byte(avaTar)); ok == true {
		return avaTar
	}

	avaTar = host + avaTar
	return avaTar
}

func SpliceImages(img []string) []string {
	for i := range img[:] {
		img[i] = viper.GetString("img_url") + img[i]
	}
	return img
}

func SpliceUrl(img string) string {
	if img == "" {
		return img
	}
	img = viper.GetString("img_url") + img
	return img
}
