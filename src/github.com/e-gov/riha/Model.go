package riha

import (
	"encoding/json"
	"io/ioutil"

	"os"

	. "github.com/e-gov/riha/util"

	log "github.com/Sirupsen/logrus"
)

var systems []System

var RIHA_BASE = ""

type System struct {
	Name          string  `json:"name"`
	Shortname     string  `json:"shortname"`
	Owner         Company `json:"owner"`
	Documentation string  `json:"documentation"`
	Meta          Meta    `json:"meta"`
	payload       map[string]interface{}
}

type Company struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Meta struct {
	System   Status `json:"system_status"`
	Approval Status `json:"approval_status"`
}

type Status struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

type SystemDetails struct {
	Shortname            string      `json:"shortname"`
	DescriptionTimestamp string      `json:"description_timestamp"`
	Payload              interface{} `json:"payload"`
}

func GetList() (*[]System, *AppError) {
	return &systems, nil
}

func GetDetails(shortname string) (*SystemDetails, *AppError) {
	for i := range systems {
		if systems[i].Shortname == shortname {
			return &SystemDetails{
				DescriptionTimestamp: systems[i].Meta.System.Timestamp,
				Shortname:            systems[i].Shortname,
				Payload:              systems[i].payload,
			}, nil
		}
	}

	return nil, &AppError{Error: nil, Message: "System not found", Code: 404}
}

func InitStorage(fname string) {
	var riha interface{}

	pwd, _ := os.Getwd()
	log.WithFields(log.Fields{
		"pwd": pwd,
	}).Debugf("Loading systems from %s", fname)

	data, err := ioutil.ReadFile(fname)

	if err != nil {
		log.Error("Could not read file", err)
		return
	}

	if json.Unmarshal(data, &riha) != nil {
		log.Error("Could not parse file", err)
		return
	}

	rMap := riha.([]interface{})

	for i := range rMap {
		c := rMap[i].(map[string]interface{})

		s := System{
			Name:      c["nimi"].(string),
			Shortname: c["lyhinimi"].(string),
			Owner: Company{
				Code: c["omanik"].(string),
				Name: c["omanik_nimi"].(string),
			},
			Documentation: RIHA_BASE + c["objekti_url"].(string),
			Meta: Meta{
				System: Status{
					Status:    c["staatus_kood"].(string),
					Timestamp: c["last_modified"].(string),
				},
				Approval: Status{
					Status:    c["staatus_kood"].(string),
					Timestamp: c["last_modified"].(string),
				},
			},
			payload: c,
		}
		systems = append(systems, s)
	}
	log.Debugf("Loaded %d systems from %s", len(systems), fname)
}
