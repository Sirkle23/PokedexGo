package PokedexAPI

type Map_strct struct {
	Count    int    
	Next     *string
	Previous *string
	Results  []struct {
		Name string
		URL  string
	} 
}