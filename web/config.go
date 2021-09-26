package web

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	kratos "github.com/ory/kratos-client-go"
)

var (
	KratosPublicURL    = "http://127.0.0.1:4000"
	KratosPublicClient = NewKratosClient(KratosPublicURL)
	KratosCookieName   = "ory_kratos_session"
)

func NewKratosClient(endpoint string) *kratos.APIClient {
	conf := kratos.NewConfiguration()
	conf.Servers = kratos.ServerConfigurations{{URL: endpoint}}
	return kratos.NewAPIClient(conf)
}

func SessionInfo(w http.ResponseWriter, r *http.Request) (kratos.Session, string, error) {
	var s kratos.Session
	c, err := r.Cookie(KratosCookieName)
	if err != nil || c.Value == "" {
		return s, "", err
	}

	req, err := http.NewRequest(http.MethodGet, KratosPublicURL+"/sessions/whoami", bytes.NewBuffer([]byte{}))
	if err != nil {
		return s, "", err
	}
	req.Header = map[string][]string{
		"Accept": {"application/json"},
		"Cookie": {KratosCookieName + "=" + c.Value},
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return s, "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return s, "", err
	}
	if err := json.Unmarshal(body, &s); err != nil {
		return s, "", err
	}
	if resp.StatusCode != http.StatusOK {
		return s, "", errors.New(resp.Status)
	}
	if s.Active == nil || !*s.Active {
		return s, "", errors.New("session inactive")
	}
	return s, string(body), nil
}
