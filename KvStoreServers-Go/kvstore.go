package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"

)

func main() {
	c := make(chan os.Signal, 2)

	signal.Notify(c, os.Interrupt)

	addrs := []string{":9001", ":9002", ":9003", ":9004", ":9005"}

	for _, addr := range addrs {
		s := NewServer(log.New(os.Stdout, "", 0), addr)

		s.log("Listening on http://localhost%s", addr)

		go http.ListenAndServe(addr, s)
	}

	<-c

}

func NewServer(logger *log.Logger, addr string) *Server {
	return &Server{
		data:   map[int]string{},
		logger: logger,
		addr:   addr,
	}
}

type Server struct {
	sync.RWMutex
	data   map[int]string
	logger *log.Logger
	addr   string
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.log("%s http://%s%s", r.Method, r.Host, r.URL.Path)

	switch r.Method {
	case http.MethodGet:
		s.get(w, r)
	case http.MethodPut:
		s.put(w, r)
	case http.MethodPost:
		s.post(w, r)
	case http.MethodDelete:
		s.delete(w, r)
	}
}

func (s *Server) get(w http.ResponseWriter, r *http.Request) {
	s.RLock()
	defer s.RUnlock()

	if strings.HasPrefix(r.URL.Path, "/api/values/") {
		if id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/values/")); err == nil {
			if value, ok := s.data[id]; ok {
				w.Write([]byte(value))
			} else {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			}
		}

		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(s.data)
}

func (s *Server) put(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/values/"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var value string

	json.NewDecoder(r.Body).Decode(&value)

	s.set(id, value)

	w.Write([]byte("Updated Sucessfully\n"))
}

func (s *Server) post(w http.ResponseWriter, r *http.Request) {
	var lastID int

	s.RLock()
	for id := range s.data {
		if id > lastID {
			lastID = id
		}
	}
	s.RUnlock()

	var value string

	json.NewDecoder(r.Body).Decode(&value)

	s.set(lastID+1, value)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Created Sucessfully\n"))
}

func (s *Server) delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/values/"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	s.Lock()
	defer s.Unlock()

	if _, ok := s.data[id]; ok {
		delete(s.data, id)

		w.Write([]byte("Deleted Sucessfully\n"))
		return
	}

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func (s *Server) set(id int, value string) {
	s.Lock()
	s.data[id] = value
	s.Unlock()

	s.log("Set id=%d value=%q", id, value)
}

func (s *Server) log(format string, v ...interface{}) {
	s.logger.Printf(format+"\n", v...)
}