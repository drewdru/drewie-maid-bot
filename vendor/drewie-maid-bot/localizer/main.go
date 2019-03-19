package localizer

import (
	"fmt"
	"io/ioutil"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	yaml "gopkg.in/yaml.v2"
)

var bundle *i18n.Bundle

func Translate(key, locale string) string {
	localizer := i18n.NewLocalizer(bundle, locale)
	msg, err := localizer.Localize(
		&i18n.LocalizeConfig{
			MessageID: key,
		},
	)
	if err != nil {
		fmt.Errorf("Error initialising localization, %v", err)
		panic(err)
	}
	return msg
}

func main() {
	langFiles := []string{"en-US.yaml", "ru-RU.yaml"}
	var err error

	bundle, err = CreateLocalizerBundle(langFiles)
	if err != nil {
		fmt.Errorf("Error initialising localization, %v", err)
		panic(err)
	}
}

// CreateLocalizerBundle reads language files and registers them in i18n bundle
func CreateLocalizerBundle(langFiles []string) (*i18n.Bundle, error) {
	// Bundle stores a set of messages
	bundle := &i18n.Bundle{DefaultLanguage: language.English}

	// Enable bundle to understand yaml
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	var translations []byte
	var err error
	for _, file := range langFiles {

		// Read our language yaml file
		translations, err = ioutil.ReadFile(file)
		if err != nil {
			fmt.Errorf("Unable to read translation file %s", file)
			return nil, err
		}

		// It parses the bytes in buffer to add translations to the bundle
		bundle.MustParseMessageFileBytes(translations, file)
	}

	return bundle, nil
}
