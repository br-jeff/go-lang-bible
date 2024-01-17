import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)

		if err = nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n" err)
			os.Extit()
		}

		b, err := ioutil.ReadAll(resp.body)
		resp.Body.Close()

		if err != nil {
			ftm.Fprintf(os.Stderr, "fetch: reading url: %s: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("%s", b)
	}
}