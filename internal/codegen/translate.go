package codegen

import (
	"encoding/json"
	"fmt"
	"os"
)

// TranslationInfo contains the translation information for a field
type TranslationInfo struct {
	ElasticName    string `json:"elasticName"`
	VspecName      string `json:"vspecName"`
	ClickHouseType string `json:"clickHouseType"`
	IsArray        *bool  `json:"isArray"`
	GoType         string `json:"goType"`
}

// Translations is a map of translations from ES to VSPEC and vice versa
type Translations struct {
	FromESName    map[string]*TranslationInfo
	FromVspecName map[string]*TranslationInfo
}

func ReadTranslationJSON(filePath string) (*Translations, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var transInfos []*TranslationInfo
	err = decoder.Decode(&transInfos)
	if err != nil {
		return nil, fmt.Errorf("failed to decode json: %w", err)
	}
	translations := &Translations{
		FromESName:    map[string]*TranslationInfo{},
		FromVspecName: map[string]*TranslationInfo{},
	}
	for _, info := range transInfos {
		translations.FromESName[info.ElasticName] = info
		translations.FromVspecName[info.VspecName] = info
	}

	return translations, nil
}
