package gw2w

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Handler struct {
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.serveHttp(w, r); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (h *Handler) serveHttp(w http.ResponseWriter, r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	key := r.URL.Path[1:]
	chatID := r.URL.Query().Get("chatid")
	visibleToUser := r.URL.Query().Get("visible_to_user")

	notifyBody := &WebhookNotifierBody{}
	if err = json.Unmarshal(body, notifyBody); err != nil {
		return err
	}

	if err = notifyBody.Trans(chatID, visibleToUser).Send(r.Context(), key); err != nil {
		return err
	}

	return nil
}
