package tmpl

import (
	"html/template"
	"io/fs"
	"sync"

	"github.com/gin-gonic/gin"
)

type Renderer struct {
	templates *template.Template
	mutex     sync.RWMutex
	fs        fs.FS
}

func NewRenderer(fs fs.FS, pattern string) *Renderer {
	return &Renderer{
		fs: fs,
	}
}

// LoadAllTemplates parses all templates and sets them in Gin
func (r *Renderer) LoadAllTemplates(engine *gin.Engine) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	tmpl, err := template.ParseFS(r.fs, "./templates/layouts/*.html", "./templates/pages/**/*.html")
	if err != nil {
		return err
	}

	r.templates = tmpl
	engine.SetHTMLTemplate(tmpl)
	return nil
}

// Render is now just a wrapper around Gin's HTML
func (r *Renderer) Render(c *gin.Context, code int, name string, data any) {
	// Gin already has the templates loaded via SetHTMLTemplate
	c.HTML(code, name, data)
}
