package main

import "encoding/json"
import "net/http"
import "fmt"
import "log"
import "io/ioutil"
import "container/list"

func GetKanjiForApiKey(key string) *list.List {
	url := "http://www.wanikani.com/api/user/" + key + "/kanji"
	fmt.Println("Getting kanji for url " + url)

	response, err := http.Get(url)
	content, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		log.Fatalln("Error: http.Get.", err)
	}

	result := list.New()

	var jsonData interface{}
	json.Unmarshal(content, &jsonData)	

	requestedInformation := jsonData.(map[string]interface{})

	for _, v := range requestedInformation {
		switch v.(type) {
			case []interface{}:
				for i := range(v.([]interface{})) {
					data := v.([]interface{})[i].(map[string]interface{})
					character := data["character"].(string) 
					if data["stats"] == nil {
						continue
					}
					stats := data["stats"].(map[string]interface{})
					srs := stats["srs"].(string)
					kanjiStats := KanjiStats { srs }
					kanji := Kanji { character, kanjiStats.Status() }

					result.PushBack(kanji)
				}
			default:
				continue
		}
	}

	return result
}
