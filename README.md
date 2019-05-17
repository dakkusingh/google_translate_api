# Google Translate API in Golang
Note: This is not normal commercial Translator API provided by Google.
Use for research and POC purposes.

## Install:
```
go get github.com/dakkusingh/google_translate_api
```

## Example usage - Punjabi to English:
```
package main

import (
	"fmt"

	gta "github.com/dakkusingh/google_translate_api"
)

func main() {
	text := "ਤੁਸੀ ਕਿਵੇਂ ਹੋ"
	sourceLang := "pa"
	targetLang := "en"

	translation, transliteration, _ := gta.Translate(text, sourceLang, targetLang)
	
	fmt.Println(translation)
	fmt.Println(transliteration)
	// Output: "How are you"
	// Output: "Tusī kivēṁ hō"
}
```
