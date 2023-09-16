package govapi

import (
	"encoding/json"
	"fmt"
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

