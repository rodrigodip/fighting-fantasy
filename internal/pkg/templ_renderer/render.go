package tmplrender

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin/render"
)

// TemplateRenderer is a custom Gin HTML renderer that holds
// isolated template sets per page, avoiding {{define}} name collisions
// across pages that all use the same base layout.
type TemplateRenderer struct {
	sets    map[string]*template.Template
	funcMap template.FuncMap
	base    []string
}

// NewTemplateRenderer builds the renderer and registers all pages.
// basePath should point to the root templates directory, e.g:
//
//	./internal/interface/web/templates
func NewTemplateRenderer(basePath string, funcMap template.FuncMap) *TemplateRenderer {
	r := &TemplateRenderer{
		sets:    make(map[string]*template.Template),
		funcMap: funcMap,
		base: []string{
			filepath.Join(basePath, "layouts", "base.html"),
			filepath.Join(basePath, "partials", "auth-error.html"),
		},
	}

	// Register pages here.
	// Each page gets its own isolated template set:
	// base layout + shared partials + the page file.
	// This prevents {{define "content"}} collisions across pages.
	r.add("auth", filepath.Join(basePath, "pages", "auth.html"))
	r.add("dashboard", filepath.Join(basePath, "pages", "dashboard.html"))
	r.add("gameover", filepath.Join(basePath, "pages", "gameover.html"))

	return r
}

// add creates an isolated template set for a named page.
// The name is what handlers pass to c.HTML(), e.g: c.HTML(200, "auth", data)
func (r *TemplateRenderer) add(name string, pageFiles ...string) {
	files := append(r.base, pageFiles...)
	r.sets[name] = template.Must(
		template.New("").Funcs(r.funcMap).ParseFiles(files...),
	)
}

// Instance implements gin's render.HTMLRender interface.
// Gin calls this when c.HTML() is used in a handler.
func (r *TemplateRenderer) Instance(name string, data any) render.Render {
	tmpl, ok := r.sets[name]
	if !ok {
		// Fail loudly — a missing template name is always a dev mistake
		panic("renderer: template not found: " + name)
	}
	return &htmlRender{
		template: tmpl,
		data:     data,
	}
}

// htmlRender implements gin's render.Render interface.
// It always executes "base.html" as the entry point,
// which then pulls in {{define "content"}} from the page file.
type htmlRender struct {
	template *template.Template
	data     any
}

func (h *htmlRender) Render(w http.ResponseWriter) error {
	return h.template.ExecuteTemplate(w, "base.html", h.data)
}

func (h *htmlRender) WriteContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

