/*
 * Copyright 2017 Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package uaa

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"bytes"
	"io/ioutil"

	"crypto/tls"

	"code.cloudfoundry.org/clock"
	"code.cloudfoundry.org/lager"
	client "code.cloudfoundry.org/uaa-go-client"
	uaa "code.cloudfoundry.org/uaa-go-client/config"
)

type Sdk struct {
	logger     *log.Logger
	uaaClient  client.Client
	httpClient *http.Client
	endpoint   string
}

type laggerWrapper struct {
	logger *log.Logger
}

func (lw *laggerWrapper) Log(msg lager.LogFormat) {
	lw.logger.Printf("%#v", msg)
}

func New(cfg *uaa.Config, logger *log.Logger) (*Sdk, error) {
	lag := lager.NewLogger("uaa-sdk")
	lag.RegisterSink(&laggerWrapper{logger})

	uaaClient, err := uaa_go_client.NewClient(lag, cfg, clock.NewClock())
	if err != nil {
		return nil, err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: cfg.SkipVerification},
	}

	return &Sdk{
		logger:     logger,
		uaaClient:  uaaClient,
		httpClient: &http.Client{Transport: tr},
		endpoint:   cfg.UaaEndpoint}, nil
}

func (s *Sdk) doRequest(reqType, path string, data *bytes.Buffer) (map[string]interface{}, error) {
	req, err := http.NewRequest(reqType, fmt.Sprintf("%s/%s", s.endpoint, path), data)
	if err != nil {
		return nil, err
	}

	token, err := s.uaaClient.FetchToken(false)
	if err != nil {
		return nil, fmt.Errorf("fetch UAA token: %v", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Content-Type", "application/json")
	resp, err := s.httpClient.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var bodymap map[string]interface{}
	if err := json.Unmarshal(body, &bodymap); err != nil {
		fmt.Printf("%+v", string(body))
		println(err.Error())
		return nil, err
	}
	return bodymap, nil
}

func (s *Sdk) CreateUser(user *User) error {
	userBytes, err := json.Marshal(&user)
	if err != nil {
		return fmt.Errorf("marshalling user: %v", err)
	}

	resp, err := s.doRequest(http.MethodPost, "Users", bytes.NewBuffer(userBytes))
	if err != nil {
		return fmt.Errorf("creating user: %v", err)
	}

	if id, idOk := resp["id"]; idOk {
		user.Id = id.(string)
	} else {
		if resp["error"] == "scim_resource_already_exists" {
			user.Id = resp["user_id"].(string)
		} else {
			return fmt.Errorf("unknown response: %#v", resp)
		}
	}

	return nil
}
