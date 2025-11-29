package parser

import (
	"encoding/json"
	"os"

	"astral/internal/model"
)

func ParseSword(path string) (model.Sword, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return model.Sword{}, err
	}

	var s model.Sword
	err = json.Unmarshal(data, &s)
	return s, err
}
