package web

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"ruleta/domain"
	"ruleta/service/port/in"
	"strconv"
)

type GirarRuletaHandler struct {
	in.GirarRuletaPortIN
}

func (h GirarRuletaHandler) GirarRuleta(w http.ResponseWriter, req *http.Request) {

	var requestBodyReadCloser io.ReadCloser

	requestBodyReadCloser = req.Body

	requestBodyBytes, _ := ioutil.ReadAll(requestBodyReadCloser)

	type girarRuletaRequestBody struct {
		NumeroSeleccionado int
	}

	var requestBodyJson girarRuletaRequestBody

	json.Unmarshal(requestBodyBytes, &requestBodyJson)

	var numeroSeleccionado string = strconv.Itoa(requestBodyJson.NumeroSeleccionado)

	numeroSeleccionadoCommand := in.GirarRuletaCommand{
		NumeroSeleccionado: domain.Numero{Value: requestBodyJson.NumeroSeleccionado},
	}

	err := h.GirarRuletaPortIN(numeroSeleccionadoCommand)

	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	
	fmt.Fprintf(w, "numero seleccionado : "+numeroSeleccionado)

	

}
