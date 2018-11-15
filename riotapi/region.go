package riotapi

import "fmt"

const apiBase = "api.riotgames.com"

type Region struct {
	Name   string
	Domain string
}

var OCE = &Region{"OCE", "OC1"}

func (r *Region) ApiEndpoint(endpoint string) string {
	return fmt.Sprintf("https://%s.%s%s", r.Domain, apiBase, endpoint)
}
