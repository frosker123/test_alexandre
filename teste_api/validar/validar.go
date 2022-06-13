package validar

import (
	"errors"
	"main/modelos"
	"strconv"
	"strings"
)

func ValidarCampos(contatos []modelos.Contatos) error {

	for i := range contatos {
		if contatos[i].Nome == "" {
			return errors.New("campo nome tem que ser preenchido")

		}
		contatos[i].Nome = strings.ToUpper(contatos[i].Nome)
	}

	for i := range contatos {
		if len(contatos[i].Celular) != 13 {
			return errors.New("quantidade de numero invalido")
		}

		if contatos[i].Celular == "" {
			return errors.New("campo celular tem que ser preenchido")
		}
		valida := isNumeric(contatos[i].Celular)
		if !valida {
			return errors.New("celular pode ter apenas numeros")
		}

	}

	return nil
}

func isNumeric(numero string) bool {
	_, err := strconv.ParseFloat(numero, 64)
	return err == nil
}
