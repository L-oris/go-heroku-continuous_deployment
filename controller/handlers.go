package controller

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/L-oris/go-heroku-continuous_deployment/reverse"
	"github.com/julienschmidt/httprouter"
)

func (c *Controller) PrintMessage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("controller.PrintMessage")

	io.WriteString(w, c.message)
}

func (c *Controller) AddToMessage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("controller.AddToMessage")

	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error reading body: ", err)
		return
	}

	reversedMessage := reverse.Reverse(string(bs))
	c.message = c.message + " | " + reversedMessage
	io.WriteString(w, c.message)
}

func (c *Controller) ResetMessage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("controller.ResetMessage")
	c.message = "Hello World"
	io.WriteString(w, c.message)
}
