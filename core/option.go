package core

import "net/http"

type Option = func(c *SDKClient)

func WithHttpClient(httpClient *http.Client) Option {
	return func(c *SDKClient) {
		c.httpClient = httpClient
	}
}

func WithTracer(namespace string) Option {
	return func(c *SDKClient) {
		c.tracer = NewOtel(namespace, c.AppID())
	}
}

// SetDebug set debug mode
func WithDebug(debug bool) Option {
	return func(c *SDKClient) {
		c.debug = debug
	}
}
