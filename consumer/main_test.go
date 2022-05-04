package consumer

import (
	"fmt"
	"net/url"
	"os"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
)

var client *Client

func TestMain(m *testing.M) {
	var exitCode int

	setup()

	exitCode = m.Run()

	if err := pact.WritePact(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pact.Teardown()
	os.Exit(exitCode)
}

var pact dsl.Pact
var term = dsl.Term

type request = dsl.Request

func setup() {
	pact = createPact()

	pact.Setup(true)

	u, _ := url.Parse(fmt.Sprintf("http://localhost:%d", pact.Server.Port))

	client = &Client{
		BaseURL: u,
	}

}

func createPact() dsl.Pact {
	return dsl.Pact{
		Consumer: os.Getenv("CONSUMER_NAME"),
		Provider: os.Getenv("PROVIDER_NAME"),
		LogDir:   os.Getenv("LOG_DIR"),
		PactDir:  os.Getenv("PACT_DIR"),
		LogLevel: "INFO",
	}
}
