package pokeapi

// RespShallowLocations -
// 希望字段能够被 JSON 编码和解码，以及在不同包中被访问，需要将它的首字母改为大写
type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
