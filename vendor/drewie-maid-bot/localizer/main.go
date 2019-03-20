package localizer

import (
	"fmt"
	"log"
	"sync"

	"io/ioutil"
	"path/filepath"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	yaml "gopkg.in/yaml.v2"
)

type localizer struct {
	bundle *i18n.Bundle
}

var instance *localizer
var once sync.Once

// GetInstance return localizer instance
func GetInstance() *localizer {
	once.Do(func() {
		instance = newLocalizer()
	})
	return instance
}

// newLocalizer create localizer instance
func newLocalizer() *localizer {
	var langFiles []string
	var err error
	log.Println("langFiles: %v", langFiles)

	langFiles, err = filepath.Glob("../localizer/localizations/*.yaml")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("langFiles: %v", langFiles)

	instance := new(localizer)
	instance.bundle, err = createLocalizerBundle(langFiles)
	if err != nil {
		log.Println("Error initialising localization, %v", err)
		panic(err)
	}
	return instance
}

// createLocalizerBundle reads language files and registers them in i18n bundle
func createLocalizerBundle(langFiles []string) (*i18n.Bundle, error) {
	bundle := &i18n.Bundle{DefaultLanguage: language.English}
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	var translations []byte
	var err error
	log.Println("A000000000: %v", bundle)

	for _, file := range langFiles {
		translations, err = ioutil.ReadFile(file)
		if err != nil {
			log.Println("Unable to read translation file %s", file)
			return nil, err
		}
		log.Println("A111111111: %v", translations)
		log.Println("A222222222: %v", file)
		bundle.MustParseMessageFileBytes(translations, file)
	}
	log.Println("A333333333: %v", bundle)

	return bundle, nil
}

// Translate reads string from i18n bundle
func (localizerObj *localizer) Translate(key, locale string) string {
	i18nLocalizer := i18n.NewLocalizer(localizerObj.bundle, locale)
	message, err := i18nLocalizer.Localize(
		&i18n.LocalizeConfig{
			MessageID: key,
		},
	)
	if err != nil {
		log.Println("Error initialising localization, %v", err)
		return fmt.Sprintf("Error: %v. Message NOT Found", err)
	}
	return message
}
