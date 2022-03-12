package main

import (
	"html/template"
	"net/http"
)

const port=":8080"
const templatePath="./templates/storytemplate.html"

type storyHandler struct{
	adventure Adventrue
}

func(sh storyHandler) ServeHTTP(w http.ResponseWriter,r *http.Request){
	storyName:=r.URL.Path
	if storyName==""{
		storyName="/intro"
	}
	storyName=storyName[1:]
	story,err:=GetStory(storyName,sh.adventure)
	if err!=nil {
		http.NotFound(w,r)
		return
	}
	t,_:=template.ParseFiles(templatePath)
	t.Execute(w,story)
}

func handlerInit() http.Handler{
	story:=ParseJson()
	return storyHandler{adventure: story}
}

func main() {
	handler:=handlerInit()
	http.ListenAndServe(port,handler)
}