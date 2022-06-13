package keytoken

import (
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var keyMacapa = "macapa"

func CreateTokenMacapa() (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Minute * 10).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)

	tokens, err := token.SignedString([]byte(keyMacapa))
	if err != nil {
		return "", nil
	}

	return tokens, nil
}

func TokenvalidoMacapa(r *http.Request) bool {
	tokenstr := verificaTokenMacapa(r)
	token, err := jwt.Parse(tokenstr, compareTokenMacapa)
	if err != nil {
		return false
	}

	return token.Valid
}

func verificaTokenMacapa(r *http.Request) string {
	token := r.Header.Get("authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func compareTokenMacapa(t *jwt.Token) (interface{}, error) {
	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, nil
	}

	return []byte(keyMacapa), nil
}
