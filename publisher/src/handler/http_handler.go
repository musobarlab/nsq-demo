package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/wuriyanto48/nsq-demo/publisher/src/pub"
	"github.com/wuriyanto48/nsq-demo/publisher/src/utils"
)

// HttpHandler model
type HttpHandler struct {
	publisher pub.Publisher
}

// NewHttpHandler function for initialise *HttpHandler
func NewHttpHandler(publisher pub.Publisher) *HttpHandler {
	return &HttpHandler{publisher: publisher}
}

func (h *HttpHandler) PublishMessages() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		if req.Method != "POST" {
			utils.JsonResponse(res, "Invalid Method", http.StatusMethodNotAllowed)
			return
		}

		var message pub.Message
		//Get message from client
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&message)

		if err != nil {
			log.Printf("Error %s", err.Error())
			utils.JsonResponse(res, "Error occured", http.StatusInternalServerError)
			return
		}

		//publish message to nsq
		b, err := message.JSON()

		if err != nil {
			log.Printf("Error %s", err.Error())
			utils.JsonResponse(res, "Error occured", http.StatusInternalServerError)
			return
		}

		err = h.publisher.Publish("wurys", b)

		if err != nil {
			log.Printf("Error %s", err.Error())
			utils.JsonResponse(res, "Error occured", http.StatusInternalServerError)
			return
		}

		utils.JsonResponse(res, "Message sent", http.StatusOK)

	})

}
