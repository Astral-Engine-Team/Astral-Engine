package model

type Sword struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	MaxStack  int    `json:"maxStack"`
	Value     int    `json:"value"`
	Rarity    string `json:"rarity"`
}
