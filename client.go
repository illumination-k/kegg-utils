package kegg

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Operation string
type Database string
type Format string

type Client struct {
	baseUrl url.URL
}

func NewClient() Client {
	baseUrl, err := url.Parse("https://rest.kegg.jp/")

	if err != nil {
		panic("Parsing base url must be safe")
	}

	return Client{baseUrl: *baseUrl}
}

func request(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	body := resp.Body

	defer body.Close()

	bytes, err := ioutil.ReadAll(body)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (c Client) List(db Database) (string, error) {
	url := c.baseUrl.JoinPath("list", string(db)).String()
	return request(url)
}

func (c Client) Get(db Database, format Format, ids ...string) (string, error) {
	var idPath strings.Builder
	for _, id := range ids {
		if idPath.Len() != 0 {
			idPath.WriteRune('+')
		}
		idPath.WriteString(id)
	}

	url := c.baseUrl.JoinPath(string(db), idPath.String(), string(format)).String()

	return request(url)
}
