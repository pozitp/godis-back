package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}
type allEvents []event

var events = allEvents{
	{
		ID:          "1",
		Title:       "Introduction to Golang",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
	},
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	var mewEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "err")
	}
	json.Unmarshal(reqBody, &mewEvent)
	events = append(events, mewEvent)
	w.WriteHeader(http.StatusCreated)
	// fmt.Printf("%+v", animals)
	json.NewEncoder(w).Encode(mewEvent)
}

/* func prettyprint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
} */
