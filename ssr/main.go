package main

import (
	"bytes"
	_ "embed"
	"io"
	"log"
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

type Element struct {
	TagName    string
	Attributes map[string]string
	Content    any // string, Element, []Element, or nil
}

func (e *Element) Render() []byte {
	var b bytes.Buffer
	b.WriteString("<")
	b.WriteString(e.TagName)
	for key, value := range e.Attributes {
		b.WriteString(" ")
		b.WriteString(key)
		b.WriteString("=\"")
		b.WriteString(value)
		b.WriteString("\"")
	}
	if e.TagName == "img" {
		b.WriteString(" />")
		return b.Bytes()
	}
	b.WriteString(">")
	switch content := e.Content.(type) {
	case string:
		b.WriteString(content)
	case Element:
		b.Write(content.Render())
	case []Element:
		for _, element := range content {
			b.Write(element.Render())
		}
	case nil:
	default:
		panic("invalid Element.Content")
	}
	b.WriteString("</")
	b.WriteString(e.TagName)
	b.WriteString(">")
	return b.Bytes()
}

type Component struct {
	TagName string
	Slots   map[string]Element
}

func (c *Component) Render(writer io.Writer) {
	children := make([]Element, len(c.Slots))
	var index int
	for name, slot := range c.Slots {
		if name != "" {
			if slot.Attributes == nil {
				slot.Attributes = map[string]string{}
			}
			slot.Attributes["slot"] = name
		}
		children[index] = slot
		index++
	}
	element := Element{TagName: c.TagName, Content: children}
	writer.Write(element.Render())
}

//go:embed components/base.html
var baseDoument []byte

func ShowCardView(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	card := Component{"jph2-card", map[string]Element{
		"header": {TagName: "h3", Content: "Hello from Go"},
		"":       {TagName: "p", Content: "This component was populated on the server side."},
		"footer": {TagName: "button", Content: "Action button"},
	}}
	base := bytes.Split(baseDoument, []byte("<!-- content -->"))
	w.Write(base[0])
	card.Render(w)
	w.Write(base[1])
}

func Asset(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	path := filepath.Join("dist", ps.ByName("asset"))
	http.ServeFile(w, r, path)
}

func main() {
	router := httprouter.New()
	router.GET("/", ShowCardView)
	router.GET("/assets/:asset", Asset)
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
