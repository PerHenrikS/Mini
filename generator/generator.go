package generator

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"

	"github.com/tkanos/gonfig"
)

//Configuration : contains paths from json
type Configuration struct {
	PagesPath         string
	TemplateIndexPath string
	TemplatePostPath  string
	WebpagePath       string
	PostsPath         string
}

//Generator : struct to contain information needed to page generation
type Generator struct {
	config Configuration
}

//New : creates a page generator
func New() Generator {
	generator := Generator{}
	err := gonfig.GetConf("./conf/config.json", &generator.config)
	if err != nil {
		panic(err)
	}
	return generator
}

//GeneratePage : Creates a folder with webpage
func (g *Generator) GeneratePage() {
	posts := []Post{}

	files, err := ioutil.ReadDir(g.config.PagesPath)
	check(err)

	createDir(g.config.WebpagePath)

	os.RemoveAll(g.config.PostsPath) //To not get duplicates, inefficient but works for now
	createDir(g.config.PostsPath)

	for _, f := range files {
		/*
			This could be made more efficient by using a buffer
			but as long as the post count is low there is no need
		*/
		var path = g.config.PagesPath + "/" + f.Name()
		newPost, err := NewPost(path, f.Name())
		check(err)
		posts = append(posts, newPost)

		//TODO: Create the post file from newPost
		templateContent, err := ioutil.ReadFile(g.config.TemplatePostPath)
		check(err)

		t, err := template.New("Post").Parse(string(templateContent))
		check(err)

		file, err := os.Create("./webpage/posts/" + newPost.Number + "-" + newPost.Title + ".html")
		check(err)

		err = t.Execute(file, newPost)
		check(err)
	}

	//TODO: Take site title and author from config file somewhere?
	webpage := NewPage("My Site", "Henrik Söderberg", posts)

	fileContent, err := ioutil.ReadFile(g.config.TemplateIndexPath)
	check(err)

	t, err := template.New("Homepage").Parse(string(fileContent))
	check(err)

	file, err := os.Create("./webpage/index.html")
	check(err)
	defer file.Close()

	err = t.Execute(file, webpage)
	check(err)
}

//What is the golang way of handling these?
func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func createDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0700) //Read Write Execute permissions TODO: Maybe change this to something more suitable?
	}
}
