package base62

import (
	"errors"
	"strings"
)

var (
	alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	length = uint64(len(alphabet))
)

func Encode(number uint64) string {
	var encodedUrl strings.Builder
	for number > 0 {
		encodedUrl.WriteByte(alphabet[number % length])
		number /= length
	}
	return encodedUrl.String()
}

func Decode(encodedUrl string) (uint64, error) {
	var decodedUrl uint64
	for _, charAscii := range encodedUrl {
		index := strings.Index(string(charAscii), alphabet)
		if index == -1 {
			return 0, errors.New("Invalid character")
		}
		decodedUrl += uint64(index) * length
	}
	return decodedUrl, nil
}
