package books

import (
	"net/http"
)

type BookCore struct{}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HttpClient
)

func Get(url string) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return Client.Do(request)
}

func init() {
	Client = &http.Client{}
}

// mock section

type ClientMock struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

var (
	mockDoFunc func(req *http.Request) (*http.Response, error)
)

func (m *ClientMock) Do(req *http.Request) (*http.Response, error) {
	return mockDoFunc(req)
}
