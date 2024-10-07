package examples

import "fmt"

type Endpoints struct {
	CurrentUserUrl    string `json:"current_user_url"`
	AuthorizationsUrl string `json:"authorizations_url"`
	RepositoryUrl     string `json:"repository_url"`
}

func GetEndpoints() (*Endpoints, error) {
	resp, err := httpClient.Get("https://api.github.com", nil)
	if err != nil {
		return nil, err
	}

	fmt.Printf((fmt.Sprintf("Status Code: %d", resp.StatusCode())))
	fmt.Printf((fmt.Sprintf("Status: %s", resp.Status())))
	fmt.Printf((fmt.Sprintf("Body: %s\n", resp.String())))

	var endpoints Endpoints
	if err := resp.UnmarshalJson(&endpoints); err != nil {
		return nil, err
	}

	return &endpoints, nil
}
