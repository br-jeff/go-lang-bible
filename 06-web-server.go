package (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandlerFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("local-host:8000", nil)) // handler echoes the Path component of the request URL
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}