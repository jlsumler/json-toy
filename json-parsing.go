package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jeffail/gabs"
)

type Site struct {
	//href:/v2/networks/baerum-bysykkel id:baerum-bysykkel name:Bysykkel]
	Href string
	Id   string
	Name string
}
type BikeStream struct {
	// Networks string
	Sites []Site
	//Networks map[string][]Site
}

//http://api.citybik.es/v2/networks?fields=id,name,href
func getJSON(url string) error {
	var f interface{}
	//bs := BikeStream{}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	//
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	//fmt.Printf("%s\n", string(body))
	json.Unmarshal(body, &f)
	mp := f.(map[string]interface{})
	//fmt.Printf("%v\n", mp)
	maps := mp["networks"].([]interface{})
	//fmt.Printf("%v\n", maps)
	for _, m := range maps {
		m := m
		fmt.Printf("%v\n", m)
	}
	return nil
}

//http://api.citybik.es/v2/networks?fields=id,name,href
func getJSONGabs(url string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	//
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	jsonParsed, err := gabs.ParseJSON(body)
	if err != nil {
		return err
	}
	//fmt.Printf("%v\n", jsonParsed)
	// S is shorthand for Search
	children, err := jsonParsed.S("networks").Children()
	if err != nil {
		return err
	}
	//fmt.Printf("%v\n", children)
	// children should be an array of maps
	for i, child := range children {
		fmt.Printf("idx: %d, value: %v\n", i, child.Data().(map[string]interface{}))
		for k := range child.Data().(map[string]interface{}) {
			fmt.Printf("    key: %s\n", k)
		}
	}

	return nil
}
