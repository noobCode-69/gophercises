package story

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
)

var tmpl = template.Must(template.New("story").Parse(`
<!DOCTYPE html>
<html>
<head>
	<title>{{.Title}}</title>
</head>
<body>

<h1>{{.Title}}</h1>

{{range .Paragraphs}}
<p>{{.}}</p>
{{end}}

<ul>
{{range .Options}}
<li><a href="/{{.Arc}}">{{.Text}}</a></li>
{{end}}
</ul>

</body>
</html>
`))



type Story map[string]Chapter

type Chapter struct {
	Title string `json:"title"`
	Paragraphs []string `json:"story"`
	Options []Option `json:"options"`
}


type Option struct {
	Text string `json:"text"`
	Arc string `json:"arc"`
}


type Handler struct {
	story Story
}


func (h Handler )ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:]
	var myChapter Chapter
	if chapter , ok := h.story[path] ; ok {
		myChapter = chapter
	} else {
		myChapter = h.story["intro"]
	}
	err := tmpl.Execute(w, myChapter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}


func GetHandler(story Story) (Handler, error) {
	return Handler{story: story}, nil
}	


func GetStory(path string) (Story, error){
	data , err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var story Story;
	if err := json.Unmarshal(data, &story); err != nil {
		return nil, err
	}	
	return story, nil
}