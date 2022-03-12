package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type Adventrue map[string]Story

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Story struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

func ParseJson() Adventrue {
	file, err := os.Open("story.json")
	if err != nil {
		panic(err)
	}
	byteValue,_ := ioutil.ReadAll(file)
	var jsonMap Adventrue
	json.Unmarshal(byteValue, &jsonMap)
	return jsonMap
}

func GetStory(name string, stories Adventrue) (Story,error){
	story,ok:=stories[name]
	if !ok{
		return story,errors.New("the story couldn't be found")
	}
	return story,nil
}