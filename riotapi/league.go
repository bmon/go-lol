package riotapi

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type LeagueEntry struct {
	FreshBlood       bool   `json:"freshBlood"`
	HotStreak        bool   `json:"hotStreak"`
	Inactive         bool   `json:"inactive"`
	LeaguePoints     int    `json:"leaguePoints"`
	Losses           int    `json:"losses"`
	PlayerOrTeamID   string `json:"playerOrTeamId"`
	PlayerOrTeamName string `json:"playerOrTeamName"`
	Rank             string `json:"rank"`
	Veteran          bool   `json:"veteran"`
	Wins             int    `json:"wins"`
}

type LeagueList struct {
	Entries  []LeagueEntry `json:"entries"`
	LeagueID string        `json:"leagueId"`
	Name     string        `json:"name"`
	Queue    string        `json:"queue"`
	Tier     string        `json:"tier"`
}

func (r *Region) GetChallengerLeague(queue Queue) *LeagueList {
	endpoint := r.ApiEndpoint("/lol/league/v3/challengerleagues/by-queue/" + queue.LongName)

	list := &LeagueList{}
	err := GlobalClient.GetAndUnmarshal(endpoint, list)
	if err != nil {
		log.Error(errors.Wrap(err, "GetChallengerLeague"))
	}
	return list
}
