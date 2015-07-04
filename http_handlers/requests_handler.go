package http_handlers

import (
	"encoding/json"
	"fmt"
	"github.com/bsm/openrtb"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type RequestsHandler struct {
	Probability float32
}

func (r *RequestsHandler) buildResponse(rd *rand.Rand, req *openrtb.Request) openrtb.Response {
	cur := "RUB"
	pr := rd.Float32()
	markup := "markup here"
	return openrtb.Response{
		Id:      req.Id,
		Seatbid: []openrtb.Seatbid{openrtb.Seatbid{Bid: []openrtb.Bid{openrtb.Bid{Id: req.Imp[0].Id, Adm: &markup, Price: &pr, Impid: req.Imp[0].Id}}}},
		Cur:     &cur,
	}
}

func (r *RequestsHandler) sendEmptyResponse(w http.ResponseWriter) {
	w.WriteHeader(204)
}

func (h *RequestsHandler) HandleResponse(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	req, err := openrtb.ParseRequest(r.Body)
	if nil != err {
		h.sendEmptyResponse(w)
		log.Println("ERROR on parsing the request", err.Error())
	} else {
		log.Println("INFO  Received bid request ", *req.Id)

		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		if r.Float32() > h.Probability {
			h.sendEmptyResponse(w)
			return
		}

		resp := h.buildResponse(r, req)
		bin, err := json.Marshal(resp)
		if nil == err {
			w.WriteHeader(200)
			fmt.Fprintf(w, "%s", string(bin))
		} else {
			w.WriteHeader(500)
			fmt.Fprintf(w, "%s", string(err.Error()))
		}
	}
}
