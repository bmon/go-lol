package main

import (
	"fmt"

	"github.com/bmon/go-lol/riotapi"
)

func main() {
	riotapi.GlobalClient.ApiKey = "RGAPI-e81f1d63-912a-47da-9f66-5f560dd9748e"
	riotapi.GlobalClient.Debug = true

	r := riotapi.OCE.GetChallengerLeague(riotapi.RANKED_SOLO_5x5)
	fmt.Printf("%+v\n", r)
}
