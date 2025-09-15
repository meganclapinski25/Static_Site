package main

import (
	"text/template"
	"os"
	//"fmt"
)

type Page struct{
	Content string
}

func main(){

	//Reads the first post file and stores
	fileContents, err := os.ReadFile("first-post.txt")
	if err != nil {
		panic(err) 
	}

	// Parse the template file into an object = t
	t, err :=template.ParseFiles("template.tmpl")
	if err !=nil{
		panic(err)
	}

	// The page now has to get the title and content filled 
	page := Page{
		Content: string(fileContents),
	}

	
	t.Execute (os.Stdout, page)

	//creates the new file in html 
	newFile, err := os.Create("first-post.html")
	if err!=nil{
		panic(err)
	}

	t.Execute(newFile,page)
	newFile.Close()
}
