package web

import (
	"embed"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/user"
	"text/template"

	"github.com/webview/webview"
)

var title string = "Title"

// Embed the entire directory.
//go:embed templates
var indexHTML embed.FS

//go:embed app
var appFiles embed.FS

func handleError(e error) {
	if e != nil {
		panic(e)
	}
}

func Start() {
	User, _ := user.Current()

	// Note the call to ParseFS instead of Parse
	t, err := template.ParseFS(indexHTML, "templates/index.html.tmpl")
	handleError(err)

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	handleError(err)
	defer ln.Close()

	// http.FS can be used to create a http Filesystem
	appFS := http.FS(appFiles)
	fs := http.FileServer(appFS)

	cwd, err := os.Getwd()
	handleError(err)
	source, err := ioutil.ReadFile(cwd + "/README.md")

	go func() {
		// Serve static files
		http.Handle("/app/", fs) // with trailing slash

		http.HandleFunc("/exit", func(write http.ResponseWriter, r *http.Request) {
			os.Exit(0)
		})

		// Handle all other requests
		http.HandleFunc("/", func(write http.ResponseWriter, req *http.Request) {
			var Path = req.URL.Path
			log.Println("Serving request for path", Path)
			write.Header().Add("Content-Type", "text/html")

			// respond with the output of template execution
			t.Execute(write, struct {
				Title string
				Path  string
				User  *user.User
			}{Title: title, Path: Path, User: User})

			err = t.ExecuteTemplate(write, "Content", string(source))
			handleError(err)
		})

		http.HandleFunc("/page", func(write http.ResponseWriter, req *http.Request) {
			var Path = req.URL.Path
			write.Header().Add("Content-Type", "text/html")
			log.Println("Serving request for path", Path)
			// respond with the output of template execution
			t.Execute(write, struct {
				Title string
				Path  string
				User  *user.User
			}{Title: "Page", Path: Path, User: User})

		})

		log.Fatal(http.Serve(ln, nil))
	}()

	wv := webview.New(true)

	wv.Dispatch(func() {
		wv.SetTitle(title)
	})

	wv.SetSize(1280, 1080, webview.HintNone)
	wv.Navigate("http://" + ln.Addr().String())
	wv.Bind("quit", func() {
		wv.Terminate()
	})
	defer wv.Destroy()
	wv.Run()
}
