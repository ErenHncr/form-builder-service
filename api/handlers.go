package api

import "net/http"

func (s *Server) handleFoo(w http.ResponseWriter, r *http.Request) {}
func (s *Server) handleBar(w http.ResponseWriter, r *http.Request) {}
func (s *Server) handleBaz(w http.ResponseWriter, r *http.Request) {}
