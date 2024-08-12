package mainpage

import (
	"github.com/Vackhan/metrics/internal/server/pkg/functionality/mainpage"
	"github.com/Vackhan/metrics/internal/server/pkg/storage"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// New обработчик для update серверов, совместимых со стандартным сервером go
func New(repo storage.UpdateRepo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		command := mainpage.NewMainPage(repo)
		list := command.GetListOfMetrics()
		dir, err := os.ReadDir(".")
		if err != nil {
			panic(err)
		}
		log.Println(os.Getwd())
		log.Println(dir)
		templatesDir, _ := filepath.Abs("./internal/templates/mainpage")
		log.Println(templatesDir)
		tmpl, err := template.ParseFiles("./internal/templates/mainpage/index.html")
		if err != nil {
			panic(err)
		}
		tmpl.Execute(w, list)
	}
}
