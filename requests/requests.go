package requests

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Requests struct {
	BaseURL string
	Host    string
	Key     string
	Client  *http.Client
}

func (r *Requests) NewGetRequest(route string) ([]byte, error) {

	builtURL, _ := buildURL(r.BaseURL, route)

	fmt.Printf("Build URL : %s \n", builtURL)
	//TO DO : Error check

	request, err := http.NewRequest("GET", builtURL, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-type", "application/json")
	request.Header.Set("x-rapidapi-host", r.Host)
	request.Header.Set("x-rapidapi-key", r.Key)

	if err != nil {
		return nil, err
	}

	resp, err := r.Client.Do(request)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func buildURL(base string, route string) (string, error) {
	return fmt.Sprintf(base + route), nil //need to do build error checking with slashes and stuff
}
