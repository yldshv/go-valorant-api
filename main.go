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
		"map": p.Map,
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
		"map": p.Map,
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