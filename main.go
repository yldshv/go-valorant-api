package govapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type VAPI struct {
	Token string
}

func WithKey(t string) func(*VAPI) {
	return func(v *VAPI) {
		v.Token = t
	}
}

func New(opts ...func(*VAPI)) *VAPI {
	api := &VAPI{}
	for _, o := range opts {
		o(api)
	}
	return api
}

func (v *VAPI) GetAccountByName(p GetAccountByNameParams) (*GetAccountResponse, error) {
	query := MakeQueryStr(QueryParams{
		"force": p.Force,
	})
	url := fmt.Sprintf("https://api.henrikdev.xyz/valorant/v1/account/%v/%v%v", p.Name, p.Tag, query)
	resp, err := MakeReq(v, url, "GET")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var acc *GetAccountResponse

	if err := json.NewDecoder(resp.Body).Decode(&acc); err != nil {
		return nil, err
	}

	return acc, nil
}

func (v *VAPI) GetAccountByPUUID(p GetAccountByPUUIDParams) (*GetAccountResponse, error) {
	query := MakeQueryStr(QueryParams{
		"force": p.Force,
	})
	url := fmt.Sprintf("https://api.henrikdev.xyz/valorant/v1/by-puuid/account/%v%v", p.Puuid, query)
	resp, err := MakeReq(v, url, "GET")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var acc *GetAccountResponse

	if err := json.NewDecoder(resp.Body).Decode(&acc); err != nil {
		return nil, err
	}

	return acc, nil
}

func (v *VAPI) GetLifetimeMatchesByPUUID(p GetLifetimeMatchesByPUUIDParams) (*GetLifetimeMatchesByPUUIDResponse, error) {
	query := MakeQueryStr(QueryParams{
		"mode": p.Mode,
		"map":  p.Map,
		"page": p.Page,
		"size": p.Size,
	})
	url := fmt.Sprintf("https://api.henrikdev.xyz/valorant/v1/by-puuid/lifetime/matches/%v/%v%v", p.Affinity, p.PUUID, query)
	resp, err := MakeReq(v, url, "GET")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var matches *GetLifetimeMatchesByPUUIDResponse

	if err := json.NewDecoder(resp.Body).Decode(&matches); err != nil {
		return nil, err
	}

	return matches, nil
}

func (v *VAPI) GetLifetimeMMRHistoryByPUUID(p GetLifetimeMMRHistoryByPUUIDParams) (*GetLifetimeMMRHistoryByPUUIDResponse, error) {
	query := MakeQueryStr(QueryParams{
		"page": p.Page,
		"size": p.Size,
	})
	url := fmt.Sprintf("https://api.henrikdev.xyz/valorant/v1/by-puuid/lifetime/mmr-history/%v/%v%v", p.Affinity, p.Puuid, query)
	resp, err := MakeReq(v, url, "GET")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var history *GetLifetimeMMRHistoryByPUUIDResponse

	if err := json.NewDecoder(resp.Body).Decode(&history); err != nil {
		return nil, err
	}

	return history, nil
}

func (v *VAPI) GetMatchesByPUUIDv3(p GetMatchesByPUUIDv3Params) (*GetMatchesByPUUIDv3Response, error) {
	query := MakeQueryStr(QueryParams{
		"mode": p.Mode,
		"map":  p.Map,
		"size": p.Size,
	})
	url := fmt.Sprintf("https://api.henrikdev.xyz/valorant/v3/by-puuid/matches/%v/%v%v", p.Affinity, p.Puuid, query)
	resp, err := MakeReq(v, url, "GET")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var matches *GetMatchesByPUUIDv3Response

	if err := json.NewDecoder(resp.Body).Decode(&matches); err != nil {
		return nil, err
	}

	return matches, nil
}

func (v *VAPI) GetMMRByPUUIDv2(p GetMMRByPUUIDv2Params) (*GetMMRByPUUIDv2Response, error) {
	query := MakeQueryStr(QueryParams{
		"season": p.Season,
	})
	url := fmt.Sprintf("https://api.henrikdev.xyz/valorant/v2/by-puuid/mmr/%v/%v%v", p.Affinity, p.Puuid, query)
	resp, err := MakeReq(v, url, "GET")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var mmr *GetMMRByPUUIDv2Response

	if err := json.NewDecoder(resp.Body).Decode(&mmr); err != nil {
		return nil, err
	}

	return mmr, nil
}

func (v *VAPI) GetMMRHistoryByPUUID(p GetMMRHistoryByPUUIDParams) (*GetMMRHistoryByPUUIDResponse, error) {
	url := fmt.Sprintf("https://api.henrikdev.xyz/valorant/v1/by-puuid/mmr-history/%v/%v", p.Affinity, p.Puuid)
	resp, err := MakeReq(v, url, "GET")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var mmrHistory *GetMMRHistoryByPUUIDResponse

	if err := json.NewDecoder(resp.Body).Decode(&mmrHistory); err != nil {
		return nil, err
	}

	return mmrHistory, nil
}

func (v *VAPI) GetContent(p GetContentParams) (*GetContentResponse, error) {
	query := MakeQueryStr(QueryParams{
		"locale": p.Locale,
	})
	url := fmt.Sprintf("https://api.henrikdev.xyz/valorant/v1/content%v", query)
	resp, err := MakeReq(v, url, "GET")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var content *GetContentResponse

	if err := json.NewDecoder(resp.Body).Decode(&content); err != nil {
		return nil, err
	}

	return content, nil
}

func (v *VAPI) GetEsportsSchedule(p GetEsportsScheduleParams) (*GetEsportsScheduleResponse, error) {
	query := MakeQueryStr(QueryParams{
		"region": p.Affinity,
		"league": p.League,
	})
	url := fmt.Sprintf("https://api.henrikdev.xyz/valorant/v1/esports/schedule%v", query)
	resp, err := MakeReq(v, url, "GET")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var schedule *GetEsportsScheduleResponse

	if err := json.NewDecoder(resp.Body).Decode(&schedule); err != nil {
		return nil, err
	}

	return schedule, nil
}

func (v *VAPI) GetLeaderboardV2(p GetLeaderboardV2Params) (*GetLeaderboardV2Response, error) {
	query := MakeQueryStr(QueryParams{
		"affinity": p.Affinity,
		"puuid":    p.Puuid,
		"name":     p.Name,
		"tag":      p.Tag,
		"season":   p.Season,
	})
	url := fmt.Sprintf("https://api.henrikdev.xyz/valorant/v2/leaderboard%v", query)
	resp, err := MakeReq(v, url, "GET")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var lb *GetLeaderboardV2Response

	if err := json.NewDecoder(resp.Body).Decode(&lb); err != nil {
		return nil, err
	}

	return lb, nil
}

func (v *VAPI) GetLifetimeMatchesByName(p GetLifetimeMatchesByNameParams) (*GetLifetimeMatchesByNameResponse, error) {
	query := MakeQueryStr(QueryParams{
		"mode": p.Mode,
		"map":  p.Map,
		"page": p.Page,
		"size": p.Size,
	})
	url := fmt.Sprintf("https://api.henrikdev.xyz/valorant/v1/lifetime/matches/%v/%v/%v%v", p.Affinity, p.Name, p.Tag, query)
	resp, err := MakeReq(v, url, "GET")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var matches *GetLifetimeMatchesByNameResponse

	if err := json.NewDecoder(resp.Body).Decode(&matches); err != nil {
		return nil, err
	}

	return matches, nil
}

func (v *VAPI) GetLifetimeMMRHistoryByName(p GetLifetimeMMRHistoryByNameParams) (*GetLifetimeMMRHistoryByNameResponse, error) {
	query := MakeQueryStr(QueryParams{
		"page": p.Page,
		"size": p.Size,
	})
	url := fmt.Sprintf("https://api.henrikdev.xyz/valorant/v1/lifetime/mmr-history/%v/%v/%v%v", p.Affinity, p.Name, p.Tag, query)
	resp, err := MakeReq(v, url, "GET")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var history *GetLifetimeMMRHistoryByNameResponse

	if err := json.NewDecoder(resp.Body).Decode(&history); err != nil {
		return nil, err
	}

	return history, nil
}

func (v *VAPI) GetMatchesByNameV3(p GetMatchesByNameV3Params) (*GetMatchesByNameV3Response, error) {
	url := fmt.Sprintf("https://api.henrikdev.xyz/valorant/v3/matches/%v/%v/%v", p.Affinity, p.Name, p.Tag)
	resp, err := MakeReq(v, url, "GET")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var matches *GetMatchesByNameV3Response

	if err := json.NewDecoder(resp.Body).Decode(&matches); err != nil {
		return nil, err
	}

	return matches, nil
}

func (v *VAPI) GetMatch(p GetMatchParams) (*GetMatchResponse, error) {
	url := fmt.Sprintf("https://api.henrikdev.xyz/valorant/v2/match/%v", p.MatchId)
	resp, err := MakeReq(v, url, "GET")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var match *GetMatchResponse

	if err := json.NewDecoder(resp.Body).Decode(&match); err != nil {
		return nil, err
	}

	return match, nil
}

func (v *VAPI) GetMMRHistoryByName(p GetMMRHistoryByNameParams) (*GetMMRHistoryByNameResponse, error) {
	url := fmt.Sprintf("https://api.henrikdev.xyz/valorant/v1/mmr-history/%v/%v/%v", p.Affinity, p.Name, p.Tag)
	resp, err := MakeReq(v, url, "GET")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var mmrHistory *GetMMRHistoryByNameResponse

	if err := json.NewDecoder(resp.Body).Decode(&mmrHistory); err != nil {
		return nil, err
	}

	return mmrHistory, nil
}

func (v *VAPI) GetMMRByNameV2(p GetMMRByNameV2Params) (*GetMMRByNameV2Response, error) {
	query := MakeQueryStr(QueryParams{
		"season": p.Season,
	})
	url := fmt.Sprintf("https://api.henrikdev.xyz/valorant/v2/mmr/%v/%v/%v%v", p.Affinity, p.Name, p.Tag, query)
	resp, err := MakeReq(v, url, "GET")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var mmr *GetMMRByNameV2Response

	if err := json.NewDecoder(resp.Body).Decode(&mmr); err != nil {
		return nil, err
	}

	return mmr, nil
}

func (v *VAPI) GetPremierTeam(p GetPremierTeamParams) (*GetPremierTeamResponse, error) {
	url := fmt.Sprintf("https://api.henrikdev.xyz/valorant/v1/premier/%v/%v", p.TeamName, p.TeamTag)
	resp, err := MakeReq(v, url, "GET")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var team *GetPremierTeamResponse

	if err := json.NewDecoder(resp.Body).Decode(&team); err != nil {
		return nil, err
	}

	return team, nil
}

func (v *VAPI) GetPremierTeamHistory(p GetPremierTeamParams) (*GetPremierTeamHistoryResponse, error) {
	url := fmt.Sprintf("https://api.henrikdev.xyz/valorant/v1/premier/%v/%v/history", p.TeamName, p.TeamTag)
	resp, err := MakeReq(v, url, "GET")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var teamHistory *GetPremierTeamHistoryResponse

	if err := json.NewDecoder(resp.Body).Decode(&teamHistory); err != nil {
		return nil, err
	}

	return teamHistory, nil
}

func (v *VAPI) GetPremierTeams(p GetPremierTeamsParams) (*GetPremierTeamsResponse, error) {
	query := MakeQueryStr(QueryParams{
		"name":       p.TeamName,
		"tag":        p.TeamTag,
		"division":   p.Division,
		"conference": p.Conference,
	})
	url := fmt.Sprintf("https://api.henrikdev.xyz/valorant/v1/premier/search%v", query)
	resp, err := MakeReq(v, url, "GET")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var teams *GetPremierTeamsResponse

	if err := json.NewDecoder(resp.Body).Decode(&teams); err != nil {
		return nil, err
	}

	return teams, nil
}

func (v *VAPI) GetPremierConferences() (*GetPremierConferencesResponse, error) {
	url := "https://api.henrikdev.xyz/valorant/v1/premier/conferences"
	resp, err := MakeReq(v, url, "GET")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var conferences *GetPremierConferencesResponse

	if err := json.NewDecoder(resp.Body).Decode(&conferences); err != nil {
		return nil, err
	}

	return conferences, nil
}

func (v *VAPI) GetPremierSeasons(p GetPremierSeasonsParams) (*GetPremierSeasonsResponse, error) {
	url := fmt.Sprintf("https://api.henrikdev.xyz/valorant/v1/premier/seasons/%v", p.Affinity)
	resp, err := MakeReq(v, url, "GET")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var seasons *GetPremierSeasonsResponse

	if err := json.NewDecoder(resp.Body).Decode(&seasons); err != nil {
		return nil, err
	}

	return seasons, nil
}

func (v *VAPI) GetPremierLeaderboard(p GetPremierLeaderboardParams) (*GetPremierLeaderboardResponse, error) {
	url := fmt.Sprintf("https://api.henrikdev.xyz/valorant/v1/premier/leaderboard/%v", p.Affinity)
	resp, err := MakeReq(v, url, "GET")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var lb *GetPremierLeaderboardResponse

	if err := json.NewDecoder(resp.Body).Decode(&lb); err != nil {
		return nil, err
	}

	return lb, nil
}

func (v *VAPI) GetPremierConfLeaderboard(p GetPremierConfLeaderboardParams) (*GetPremierLeaderboardResponse, error) {
	url := fmt.Sprintf("https://api.henrikdev.xyz/valorant/v1/premier/leaderboard/%v/%v", p.Affinity, p.Conference)
	resp, err := MakeReq(v, url, "GET")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var lb *GetPremierLeaderboardResponse

	if err := json.NewDecoder(resp.Body).Decode(&lb); err != nil {
		return nil, err
	}

	return lb, nil
}

func (v *VAPI) GetPremierConfDivLeaderboard(p GetPremierConfDivLeaderboardParams) (*GetPremierLeaderboardResponse, error) {
	url := fmt.Sprintf("https://api.henrikdev.xyz/valorant/v1/premier/leaderboard/%v/%v/%v", p.Affinity, p.Conference, p.Division)
	resp, err := MakeReq(v, url, "GET")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var lb *GetPremierLeaderboardResponse

	if err := json.NewDecoder(resp.Body).Decode(&lb); err != nil {
		return nil, err
	}

	return lb, nil
}

func (v *VAPI) GetQueueStatus(p GetStatusParams) (*GetQueueStatusResponse, error) {
	url := fmt.Sprintf("https://api.henrikdev.xyz/valorant/v1/queue-status/%v", p.Affinity)
	resp, err := MakeReq(v, url, "GET")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var status *GetQueueStatusResponse

	if err := json.NewDecoder(resp.Body).Decode(&status); err != nil {
		return nil, err
	}

	return status, nil
}

func (v *VAPI) GetStatus(p GetStatusParams) (*GetStatusResponse, error) {
	url := fmt.Sprintf("https://api.henrikdev.xyz/valorant/v1/status/%v", p.Affinity)
	resp, err := MakeReq(v, url, "GET")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var status *GetStatusResponse

	if err := json.NewDecoder(resp.Body).Decode(&status); err != nil {
		return nil, err
	}

	return status, nil
}

func (v *VAPI) GetVersion(p GetStatusParams) (*GetVersionResponse, error) {
	url := fmt.Sprintf("https://api.henrikdev.xyz/valorant/v1/version/%v", p.Affinity)
	resp, err := MakeReq(v, url, "GET")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var version *GetVersionResponse

	if err := json.NewDecoder(resp.Body).Decode(&version); err != nil {
		return nil, err
	}

	return version, nil
}

func (v *VAPI) GetRawMatchHistory(p GetRawMatchHistoryParams) (*GetRawMatchHistoryResponse, error) {
	p.Type = "matchhistory"
	body, _ := json.Marshal(p)
	req, _ := http.NewRequest("POST", "https://api.henrikdev.xyz/valorant/v1/raw", bytes.NewReader(body))
	req.Header.Add("Authorization", v.Token)
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var matchHistory *GetRawMatchHistoryResponse

	if err := json.NewDecoder(resp.Body).Decode(&matchHistory); err != nil {
		return nil, err
	}

	return matchHistory, nil
}

func (v *VAPI) GetRawMatchDetails(p GetRawMatchDetailsParams) ([]GetRawMatchDetailsResponse, error) {
	p.Type = "matchdetails"
	body, _ := json.Marshal(p)
	req, _ := http.NewRequest("POST", "https://api.henrikdev.xyz/valorant/v1/raw", bytes.NewReader(body))
	req.Header.Add("Authorization", v.Token)
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var matchDetails []GetRawMatchDetailsResponse

	if err := json.NewDecoder(resp.Body).Decode(&matchDetails); err != nil {
		return nil, err
	}

	return matchDetails, nil
}
