/*
 * Copyright (C) 2018 Onchain <onchain@onchain.com>
 *
 * This file is part of The ontology_Zero.
 *
 * The ontology_Zero is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The ontology_Zero is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The ontology_Zero.  If not, see <http://www.gnu.org/licenses/>.
 */

package config

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

const (
	DefaultConfigFilename = "./config.json"
	MINGENBLOCKTIME = 2
	DEFAULTGENBLOCKTIME = 6
	DBFTMINNODENUM        = 4 //min node number of dbft consensus
	SOLOMINNODENUM        = 1 //min node number of solo consensus
)

var Version string

type Configuration struct {
	Magic           int64    `json:"Magic"`
	Version         int      `json:"Version"`
	SeedList        []string `json:"SeedList"`
	BookKeepers     []string `json:"BookKeepers"` // The default book keepers' publickey
	HttpRestPort    int      `json:"HttpRestPort"`
	RestCertPath    string   `json:"RestCertPath"`
	RestKeyPath     string   `json:"RestKeyPath"`
	HttpInfoPort    uint16   `json:"HttpInfoPort"`
	HttpInfoStart   bool     `json:"HttpInfoStart"`
	HttpWsPort      int      `json:"HttpWsPort"`
	HttpJsonPort    int      `json:"HttpJsonPort"`
	HttpLocalPort   int      `json:"HttpLocalPort"`
	OauthServerUrl  string   `json:"OauthServerUrl"`
	NoticeServerUrl string   `json:"NoticeServerUrl"`
	NodePort        int      `json:"NodePort"`
	NodeType        string   `json:"NodeType"`
	WebSocketPort   int      `json:"WebSocketPort"`
	PrintLevel      int      `json:"PrintLevel"`
	IsTLS           bool     `json:"IsTLS"`
	CertPath        string   `json:"CertPath"`
	KeyPath         string   `json:"KeyPath"`
	CAPath          string   `json:"CAPath"`
	GenBlockTime    uint     `json:"GenBlockTime"`
	MultiCoreNum    uint     `json:"MultiCoreNum"`
	EncryptAlg      string   `json:"EncryptAlg"`
	MaxLogSize      int64    `json:"MaxLogSize"`
	MaxTxInBlock    int      `json:"MaxTransactionInBlock"`
	MaxHdrSyncReqs  int      `json:"MaxConcurrentSyncHeaderReqs"`
	ConsensusType   string           `json:"ConsensusType"`
}

type ConfigFile struct {
	ConfigFile Configuration `json:"Configuration"`
}

var Parameters *Configuration

func init() {
	file, e := ioutil.ReadFile(DefaultConfigFilename)
	if e != nil {
		log.Fatalf("File error: %v\n", e)
		os.Exit(1)
	}
	// Remove the UTF-8 Byte Order Mark
	file = bytes.TrimPrefix(file, []byte("\xef\xbb\xbf"))

	config := ConfigFile{}
	e = json.Unmarshal(file, &config)
	if e != nil {
		log.Fatalf("Unmarshal json file erro %v", e)
		os.Exit(1)
	}
	Parameters = &(config.ConfigFile)
}
