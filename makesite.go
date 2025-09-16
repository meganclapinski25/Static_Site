package main

import (
	"text/template"
	"os"
	//"fmt"
	"flag"
	"path/filepath"
	"strings"
)

type Page struct{
	Content string
}

func main(){
	//Flag 
	dir := flag.String("dir", "", "")
	inFile := flag.String("file", "first-post.txt", "")
	flag.Parse()

	//Reads the first post file and stores
	fileContents, err := os.ReadFile(*inFile)
	if err != nil {
		panic(err) 
	}

	// Parse the template file into an object = t
	t, err :=template.ParseFiles("template.tmpl")
	if err !=nil{
		panic(err)
	}

	temp, err:=os.ReadDir(*dir)
	if err != nil{
		panic(err)
	}

	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if strings.HasSuffix(strings.ToLower(name), ".txt") {
			// print each .txt to stdout (full path)
			fmt.Println(filepath.Join(listDir, name))
		}
	}



	// The page now has to get the title and content filled 
	page := Page{
		Content: string(fileContents),
	}
	// Remove the extension 
	outPath := strings.TrimSuffix(*inFile, filepath.Ext(*inFile)) + ".html"
	
	t.Execute (os.Stdout, page)

	//creates the new file in html 
	newFile, err := os.Create(outPath)
	if err!=nil{
		panic(err)
	}

	t.Execute(newFile,page)
	newFile.Close()
}
