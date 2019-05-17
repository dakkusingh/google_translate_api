package google_translate_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func Translate(source, sourceLang, targetLang string) (string, string, error) {
	var text []string
	var result []interface{}

	url, err := url.Parse("https://translate.googleapis.com/translate_a/single")

	q := url.Query()
	q.Add("client", "gtx")
	q.Add("source", "bh")
	q.Add("hl", "en")
	q.Add("ie", "UTF-8")   // Input encoding
	q.Add("oe", "UTF-8")   // Output encoding
	q.Add("multires", "1") // I don't know
	q.Add("kc", "1")       // I don't know
	q.Add("otf", "1")      // I don't know
	q.Add("pc", "1")       // I don't know
	q.Add("trs", "1")      // I don't know
	q.Add("ssel", "0")     // I don't know
	q.Add("tsel", "0")     // I don't know
	q.Add("sc", "1")       // I don't know
	q.Add("dt", "t")       // Translate
	// q.Add("dt", "bd")  // Full translate with synonym ($bodyArray[1])
	// q.Add("dt", "at")  // Other translate ($bodyArray[5] - in google translate page this shows when click on translated word)
	// q.Add("dt", "ex")  // Example part ($bodyArray[13])
	// q.Add("dt", "ld")  // I don't know ($bodyArray[8])
	// q.Add("dt", "md")  // Definition part with example ($bodyArray[12])
	// q.Add("dt", "qca") // I don't know ($bodyArray[8])
	// q.Add("dt", "rw")  // Read also part ($bodyArray[14])
	q.Add("dt", "rm") // Transliteration?
	// q.Add("dt", "ss")   // Full synonym ($bodyArray[11])
	q.Add("sl", sourceLang) // Source language
	q.Add("tl", targetLang) // Target language
	q.Add("q", source)      // String to translate

	url.RawQuery = q.Encode()

	r, err := http.Get(url.String())
	if err != nil {
		return "err", "", errors.New("Error getting translate.googleapis.com")
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return "err", "", errors.New("Error reading response body")
	}

	bReq := strings.Contains(string(body), `<title>Error 400 (Bad Request)`)
	if bReq {
		return "err", "", errors.New("Error 400 (Bad Request)")
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return "err", "", errors.New("Error unmarshaling data")
	}

	// result2 := string(body)
	// fmt.Println(result2)

	// return "", nil

	if len(result) > 0 {
		inner := result[0]
		for _, slice := range inner.([]interface{}) {
			for _, translatedText := range slice.([]interface{}) {
				text = append(text, fmt.Sprintf("%v", translatedText))
				// 	break
			}
		}

		transliteration := ""
		if len(text) >= 8 {
			transliteration = text[8]
		}
		// fmt.Println(text[0])
		// fmt.Println(text[8])
		// fmt.Println(text)

		return text[0], transliteration, nil
	} else {
		return "err", "err", errors.New("Error getting results")
	}
}
