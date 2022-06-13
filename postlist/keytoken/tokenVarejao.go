package keytoken

import (
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var keyvarejao = "varejao"

func CreateTokenVarejao() (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Minute * 10).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)

	tokens, err := token.SignedString([]byte(keyvarejao))
	if err != nil {
		return "", nil
	}

	return tokens, nil
}

func TokenvalidoVarejao(r *http.Request) bool {
	tokenstr := verificaTokenVarejao(r)
	token, err := jwt.Parse(tokenstr, compareTokenVarejao)
	if err != nil {
		return false
	}

	return token.Valid
}

func verificaTokenVarejao(r *http.Request) string {
	token := r.Header.Get("authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func compareTokenVarejao(t *jwt.Token) (interface{}, error) {
	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, nil
	}

	return []byte(keyvarejao), nil
}
