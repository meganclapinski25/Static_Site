package main

import (
	"text/template"
	"os"
	"fmt"
	"flag"
	"path/filepath"
	"strings"
	"bytes"
    "github.com/yuin/goldmark"
    "github.com/yuin/goldmark/parser"
    "github.com/yuin/goldmark/renderer/html"
	
	
)

type Page struct{
	Content string
}

func main(){

	inFile := flag.String("file", "first-post.txt", "")
	dir := flag.String("dir", "", " ")
	flag.Parse()

	t, err :=template.ParseFiles("template.tmpl")
	if err !=nil{
		panic(err)
	}

	// if the --dir is asked, and doesn't equal nothing 
	
		added, err := os.ReadDir(*dir)
		if err != nil {
			panic(err)
		}
		for _, e:= range added{
			if e.IsDir(){
				continue
			}
			baseFileName := e.Name()

			if !strings.HasSuffix(strings.ToLower(baseFileName), ".txt"){
				continue
			}

			fullPath :=filepath.Join(*dir, baseFileName)
			fmt.Println(fullPath)

			raw, err := os.ReadFile(fullPath)
			if err != nil {
				panic(err)
			}

			page := Page{
				Content: string(raw),
			}
			
			fmt.Println(page)

			

			outfile := strings.TrimSuffix(fullPath, ".txt") + ".html"
			f, err := os.Create(outfile)
			if err != nil{
				panic(err)
			}
			if err := t.Execute(f,page); err != nil{
				panic(err)
			}
			f.Close()
		}
		
	
	//Reads the first post file and stores
	fileContents, err := os.ReadFile(*inFile)
	if err != nil {
		panic(err) 
	}

	//markdown logic 
	ext:= strings.ToLower(filepath.Ext(*inFile))
	if ext == ".md" || ext == ".markdown"{
		var buf bytes.Buffer

		md:= goldmark.New(
			goldmark.WithParserOptions(
				parser.WithAutoHeadingID(),
			),
			goldmark.WithRendererOptions(
				html.WithHardWraps(),
				html.WithUnsafe(),
			),
		)
		temp := parser.NewContext()
		if err := md.Convert(fileContents, &buf, parser.WithContext(temp)); err !=nil{
			panic(err)
		}
		fileContents = buf.Bytes()
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
