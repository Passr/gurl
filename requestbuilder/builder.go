package requestbuilder

import (
	"flag"
        "github.com/spf13/pflag"
	"net/http"
)

type RequestBuilder struct {
	args   []string
	method string
	url    string
}

func New(args []string) *RequestBuilder {
	return &RequestBuilder{args: args}
}

func (b *RequestBuilder) Build() (*http.Request, error) {
	b.parseArgs()
	req, err := http.NewRequest(b.method, b.url, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (b *RequestBuilder) parseArgs() {
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
        pflag.Parse
        fs := flag.NewFlagSet("gurl", flag.PanicOnError)

	fs.StringVar(&b.method, "X", "GET", "http request method")
	fs.Parse(b.args)
	b.url = fs.Arg(0)
}
