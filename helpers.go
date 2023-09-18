package govapi

import (
	"net/http"
	"strconv"
)

func MakeReq(v *VAPI, url string, method string) (*http.Response, error) {
	req, _ := http.NewRequest(method, url, nil)
	if v.Token != "" {
		req.Header.Add("Authorization", v.Token)
	}
	client := &http.Client{}
	return client.Do(req)
}

type QueryParams map[string]string

func MakeQueryStr(qp QueryParams) string {
	query := "?"

	for k, v := range qp {
		if v != "" {
			if len(query) > 1 {
				query = query + "&" + k + "=" + v
			} else {
				query = query + k + "=" + v
			}
		}
	}

	return query
}

func UpdateRatelimits(v *VAPI, h *http.Response) {
	v.Ratelimits.Used = convToInt(h.Header.Get("x-ratelimit-limit"))
	v.Ratelimits.Remaining = convToInt(h.Header.Get("x-ratelimit-remaining"))
	v.Ratelimits.Reset = convToInt(h.Header.Get("x-ratelimit-reset"))
}

func convToInt(str string) int {
	n, _ := strconv.Atoi(str)
	return n
}
