package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func processJson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	interest := vars["activity"]

	if strings.Compare(interest, "Animals") == 0 {
		todo := Resp{
			Title: "Evanston Animal Shelter",
			Body:  "Evanston Animal Shelter needs your help! Come help out these poor animals whenever you can. 2x a week, an hour each",
			Href:  "img/animalshelter.img",
		}
		if err := json.NewEncoder(w).Encode(todo); err != nil {
			panic(err)
		}
	} else if strings.Compare(interest, "Art") == 0 || strings.Compare(interest, "Elderly") == 0 {
		todo := Resp{
			Title: "The Art Institute of Chicago",
			Body:  "The Art Institute of Chicago is looking for volunteers to help hosting an event for elderly this weekend.",
			Href:  "img/aic.jpg",
		}
		if err := json.NewEncoder(w).Encode(todo); err != nil {
			panic(err)
		}
	} else if strings.Compare(interest, "Handicapped") == 0 {
		todo := Resp{
			Title: "Lambs Farm",
			Body:  "A rehabitational shelther for the developmentally disabled is looking for volunteers to talk to our residents once a week.",
			Href:  "img/lambsfarm.jpg",
		}
		if err := json.NewEncoder(w).Encode(todo); err != nil {
			panic(err)
		}

	} else if strings.Compare(interest, "Environment") == 0 {
		todo := Resp{
			Title: "City of Evanston",
			Body:  "Looking for volunteers to pick up trash and clean up the Evanston environment.",
			Href:  "img/cityevanston.jpg",
		}
		if err := json.NewEncoder(w).Encode(todo); err != nil {
			panic(err)
		}
	} else if strings.Compare(interest, "Faith") == 0 {
		todo := Resp{
			Title: "Sheil Catholic Center",
			Body:  "Looking for volunteers at the Sheil Catholic Server for weekly serving",
			Href:  "img/sheil.jpg",
		}
		if err := json.NewEncoder(w).Encode(todo); err != nil {
			panic(err)
		}
	} else if strings.Compare(interest, "Healthcare") == 0 {
		todo := Resp{
			Title: "Rehabilitation Institute of Chicago",
			Body:  "RIC is looking for volunteers to talk to our patients.",
			Href:  "img/ric.jpg",
		}
		if err := json.NewEncoder(w).Encode(todo); err != nil {
			panic(err)
		}
	} else if strings.Compare(interest, "Human Rights") == 0 || strings.Compare(interest, "LGBT") == 0 {
		todo := Resp{
			Title: "Human Rights Campaign Illinois",
			Body:  "The HRC fights for LGBT equality in Illinois alongside state and local groups and lawmakers. We're looking for Spanish translators",
			Href:  "img/hrci.jpg",
		}
		if err := json.NewEncoder(w).Encode(todo); err != nil {
			panic(err)
		}
	} else if strings.Compare(interest, "Hunger") == 0 {
		todo := Resp{
			Title: "Lake Street Church",
			Body:  "Looking for dedicated volunteers to help Lake Street Church's weekend event for the homeless, 'Food for Thought'",
			Href:  "img/lsc.jpg",
		}
		if err := json.NewEncoder(w).Encode(todo); err != nil {
			panic(err)
		}
	} else if strings.Compare(interest, "Justice") == 0 || strings.Compare(interest, "Kids") == 0 {
		todo := Resp{
			Title: "Restorative Justice Evanston",
			Body:  "We're looking for volunteers to talk to youth offenders who have comitted minor criminal offenses",
			Href:  "img/rje.jpg",
		}
		if err := json.NewEncoder(w).Encode(todo); err != nil {
			panic(err)
		}
	} else if strings.Compare(interest, "Media") == 0 {
		todo := Resp{
			Title: "Northwestern University Department of RTVF",
			Body:  "Looking for volunteers to help out with this weekend's film set work.",
			Href:  "img/rtvf.jpg",
		}
		if err := json.NewEncoder(w).Encode(todo); err != nil {
			panic(err)
		}
	} else if strings.Compare(interest, "Outdoor") == 0 || strings.Compare(interest, "Sports") == 0 {
		todo := Resp{
			Title: "City of Evanston",
			Body:  "Looking for volunteers at Evanston Ice Link",
			Href:  "img/icelink.jpg",
		}
		if err := json.NewEncoder(w).Encode(todo); err != nil {
			panic(err)
		}
	} else if strings.Compare(interest, "Technology") == 0 {
		todo := Resp{
			Title: "Rehabilitation Institute of Chicago",
			Body:  "RIC is looking for volunteers to assist with our website development",
			Href:  "img/ric.jpg",
		}
		if err := json.NewEncoder(w).Encode(todo); err != nil {
			panic(err)
		}
	} else if strings.Compare(interest, "Tutoring") == 0 {
		todo := Resp{
			Title: "Northwestern University Athletics",
			Body:  "Go CATS! NU Athletics is looking for volunteers to tutor our athletes.",
			Href:  "img/nua.jpg",
		}
		if err := json.NewEncoder(w).Encode(todo); err != nil {
			panic(err)
		}
	}
}
