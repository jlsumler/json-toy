package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"
)

type Msg struct {
	Name string
	Body string
	Time int64
}

type ByTime []Msg

func main() {
	yearDays := make([]string, 365)
	msg := Msg{}
	//var f interface{}
	b := []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)
	json.Unmarshal(b, &msg)

	fmt.Printf("%#v\n", msg)
	t := time.Unix(0, msg.Time)

	fmt.Println(t.String())
	fmt.Println(t.YearDay())

	yearDays[t.YearDay()] = t.String()
	//fmt.Printf("%v\n", yearDays)
	for i, d := range yearDays {
		if len(yearDays[i]) > 0 {
			fmt.Printf("%d : %s\n", i, d)
		}
	}
	sortToy()
	fmt.Println("getJSON")
	err := getJSON("http://api.citybik.es/v2/networks?fields=id,name,href")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("getJSONGabs")
	err = getJSONGabs("http://api.citybik.es/v2/networks?fields=id,name,href")
	if err != nil {
		fmt.Println(err)
	}
}

func (a ByTime) Len() int           { return len(a) }
func (a ByTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTime) Less(i, j int) bool { return a[i].Time < a[j].Time }

func sortToy() {
	m := []Msg{
		{Name: "Alice", Body: "Alice Body", Time: 1294706395881547000},
		{Name: "Bob", Body: "Bob Body", Time: 1294706395881548000},
		{Name: "Charles", Body: "Charles Body", Time: 1294706395881545000},
		{Name: "Devin", Body: "Devin Body", Time: 1294706395881540000},
	}

	fmt.Printf("%v\n", m)
	sort.Sort(ByTime(m))
	fmt.Printf("%v\n", m)

}
