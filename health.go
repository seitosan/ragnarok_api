package main

import (
	"encoding/json"
	"net/http"
)

type healthResponse struct {
	Status  int     `json:"status"`
	Version string  `json:"version"`
	Host    string  `json:"host"`
	Port    string  `json:"port"`
	Routes  []Route `json:"routes"`
}
type Route struct {
	Name   string `json:"name"`
	Status int    `json:"status"`
}

func getHealth(w http.ResponseWriter, r *http.Request) {
	config, err := newConfig("./config.yml")
	ExitIfError(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	var allRoutes []Route
	var routeTotest = []string{"/", "render"}
	allRoutes = testRoute(allRoutes, routeTotest)
	var responseObject = healthResponse{
		Status:  200,
		Version: config.Application.Version,
		Host:    config.Server.Host,
		Port:    config.Server.Port,
		Routes:  allRoutes,
	}
	/* #nosec */
	json.NewEncoder(w).Encode(responseObject)
}

func testRoute(routes []Route, handle []string) []Route {
	var route Route
	config, err := newConfig("./config.yml")
	ExitIfError(err)
	for _, item := range handle {
		route.Name = item
		response, err := http.Get("http://" + config.Server.Host + ":" + config.Server.Port + "/" + item)
		ExitIfError(err)
		route.Status = response.StatusCode
		routes = append(routes, route)
	}
	return routes
}
