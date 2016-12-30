package models

import (
    "encoding/json"
    "io/ioutil"
)

type Settings struct  {
    Name string
    Url string
    IP string
    HTTPPort int
}

func (s *Settings) ReadFromJsonFile(filename string) {
    settings_json_bytes, err := ioutil.ReadFile(filename)
    if err != nil {
        panic(err)
    }
    unmarshallErr := json.Unmarshal(settings_json_bytes, &s)
    if unmarshallErr != nil {
        panic(unmarshallErr)
    }
}
