package handler

import (
	"encoding/json"
	"fmt"
	"github.com/jpastorm/redis-cache/cache/model"
	"github.com/jpastorm/redis-cache/cache/usecase"
	"io/ioutil"
	"net/http"
)

type handler struct {
	cacheUsecase usecase.Usecase
}

func NewHandler(cacheUsecase usecase.Usecase) handler {
	return handler{cacheUsecase: cacheUsecase}
}

func (h handler) Set(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if  err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	cache := model.Cache{}
	err = json.Unmarshal(body, &cache)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(cache)
	err = h.cacheUsecase.Set(cache.Name, cache)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}

func (h handler) Get(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("q")
	if param == "" {
		http.Error(w, "No found query param", http.StatusBadRequest)
		return
	}
	result, err := h.cacheUsecase.Get(param)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(result))
}