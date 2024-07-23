package controller

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gijsdb/go-mem-db/internal/memdb"
	"github.com/rs/zerolog"
)

type WebUIController struct {
	logger zerolog.Logger
	port   string
	db     *memdb.DB
}

func NewWebUIController(logger zerolog.Logger, port string, db *memdb.DB) WebUIController {
	return WebUIController{
		logger: logger,
		port:   port,
		db:     db,
	}
}

func (c *WebUIController) HandleStart() {

	// Serve the static HTML file
	http.Handle("/", http.FileServer(http.Dir("./internal/web")))

	// Route to provide data for HTMX polling
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		data := c.db.List()
		tmpl := `<ul>{{range $key, $value := .}}<li>{{$key}}: {{$value}}</li>{{end}}</ul>`
		t := template.Must(template.New("data").Parse(tmpl))
		w.Header().Set("Content-Type", "text/html")
		if err := t.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	c.logger.Info().Msg(fmt.Sprintf("Web UI started on localhost:%s", c.port))
	http.ListenAndServe(fmt.Sprintf(":%s", c.port), nil)
}
