package riha

import (
	"encoding/json"
	"io/ioutil"

	"os"

	log "github.com/Sirupsen/logrus"
)

var systems []System

type System struct {
	Name          string  `json:"name"`
	Shortname     string  `json:"shortname"`
	Owner         Company `json:"owner"`
	Documentation string  `json:"documentation"`
}

type Company struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Meta Meta   `json:"meta"`
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

func GetList() (*[]System, error) {
	return &systems, nil
}

func GetDetails(shortname string) (*SystemDetails, error) {
	return nil, nil
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
			Name: c["nimi"].(string),
		}
		systems = append(systems, s)
	}
	log.Debugf("Loaded %d systems from %s", len(systems), fname)
}
