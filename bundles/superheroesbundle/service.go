package superheroesbundle

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/carantes/superheroes-api/core"
)

// SuperheroAPI handle api methods
type SuperheroAPI struct {
	baseURL string
	Client  *http.Client
}

// SuperheroAPISearchResponse body
type SuperheroAPISearchResponse struct {
	Results []SuperheroAPISearchResponseResult `json:"results"`
}

// ToSuperhero map API response to a Superhero model
func (res *SuperheroAPISearchResponse) ToSuperhero() []Superhero {
	var shs []Superhero
	var alignment SuperheroAlignment

	for _, item := range res.Results {
		sh := NewSuperhero(
			item.Name,
			item.Biography.FullName,
			alignment.FromString(item.Biography.Alignment),
			core.ParseInt(item.Powerstats.Intelligence),
			core.ParseInt(item.Powerstats.Power),
			item.Work.Occupation,
			item.Image.URL,
			len(strings.SplitAfter(item.Connections.Relatives, ",")),
		)

		groups := strings.Split(item.Connections.Groups, ",")
		for _, group := range groups {
			sh.AddGroup(group)
		}

		shs = append(shs, *sh)
	}

	return shs
}

// SuperheroAPISearchResponseResult result object
type SuperheroAPISearchResponseResult struct {
	Name        string                                      `json:"name"`
	Powerstats  SuperheroAPISearchResponseResultPowerstats  `json:"powerstats"`
	Biography   SuperheroAPISearchResponseResultBiography   `json:"biography"`
	Work        SuperheroAPISearchResponseResultWork        `json:"work"`
	Connections SuperheroAPISearchResponseResultConnections `json:"connections"`
	Image       SuperheroAPISearchResponseResultImage       `json:"image"`
}

// SuperheroAPISearchResponseResultPowerstats result powerstats section
type SuperheroAPISearchResponseResultPowerstats struct {
	Intelligence string `json:"intelligence"`
	Strength     string `json:"strength"`
	Speed        string `json:"speed"`
	Durability   string `json:"durability"`
	Power        string `json:"power"`
	Combat       string `json:"combat"`
}

// SuperheroAPISearchResponseResultBiography result biography section
type SuperheroAPISearchResponseResultBiography struct {
	FullName  string   `json:"full-name"`
	AlterEgos string   `json:"alter-egos"`
	Aliases   []string `json:"aliases"`
	Alignment string   `json:"alignment"`
}

// SuperheroAPISearchResponseResultWork result work section
type SuperheroAPISearchResponseResultWork struct {
	Occupation string `json:"occupation"`
	Base       string `json:"base"`
}

// SuperheroAPISearchResponseResultConnections result connections section
type SuperheroAPISearchResponseResultConnections struct {
	Groups    string `json:"group-affiliation"`
	Relatives string `json:"relatives"`
}

// SuperheroAPISearchResponseResultImage result image url
type SuperheroAPISearchResponseResultImage struct {
	URL string `json:"url"`
}

// SuperheroAPISearchError is a custom error for search method
type SuperheroAPISearchError struct {
	error
	Name string
}

func (e *SuperheroAPISearchError) Error() string {
	return fmt.Sprintf("could not search superhero with name %s ", e.Name)
}

// NewSuperheroAPI instance
func NewSuperheroAPI(client *http.Client, baseURL string) *SuperheroAPI {
	return &SuperheroAPI{
		baseURL: baseURL,
		Client:  client,
	}
}

// Search data from superheroes api
func (api *SuperheroAPI) Search(name string) (SuperheroAPISearchResponse, error) {
	resp, err := api.Client.Get(api.baseURL + "/search/" + name)

	if err != nil {
		return SuperheroAPISearchResponse{}, &SuperheroAPISearchError{Name: name, error: err}
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return SuperheroAPISearchResponse{}, &SuperheroAPISearchError{Name: name, error: err}
	}

	response := SuperheroAPISearchResponse{}
	jsonErr := json.Unmarshal(body, &response)

	if jsonErr != nil {
		return SuperheroAPISearchResponse{}, &SuperheroAPISearchError{Name: name, error: jsonErr}
	}

	return response, nil
}
