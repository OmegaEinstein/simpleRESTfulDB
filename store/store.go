package store

import (
	"encoding/json"
	log "github.com/auxten/logrus"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
)

func Dump(data map[string]interface{}) {
	j, err := json.Marshal(data)
	if err != nil {
		log.Errorf("Marshal json error: %s", err)
		return
	}
	ioutil.WriteFile("dump.db", j, 0644)

}

func Load() map[string]interface{} {
	j, err := ioutil.ReadFile("dump.db")
	if err != nil {
		log.Errorf("load db error %s", err)
	}
	json_j, err := simplejson.NewJson(j)
	if err != nil {
		log.Errorf("unmarshal error")
	}
	m, _ := json_j.Map()
	return m

}
