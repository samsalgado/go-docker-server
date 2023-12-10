package main
import ("errors"
	"fmt"
	"io"
	"net/http"
	"os"
	
)
type Page struct {
	Title string
	Body []byte
}
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}
func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "Live Site\n")
}
func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /welcome request\n")
	io.WriteString(w, "Welcome, HTTP\n")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/welcome",welcome)
	p1 := &Page{Title: "Test Page", Body: []byte("Sample")}
	p1.save()
	err := http.ListenAndServe(":3000", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}
