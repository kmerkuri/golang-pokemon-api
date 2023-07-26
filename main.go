package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

)

func main(){
	//call routes
	getpokemon()

}
type Response struct{
	Name string `json:"name"`
	Pokemonentries []Pokemonentries `json:"pokemon_entries"`
}
type Pokemonentries struct {
	Entrienumber int `json:"entry_number"`
	PokemonSpecies PokemonSpecies `json:"pokemon_species"`
}
type PokemonSpecies struct {
	Name string `json:"name"`
	Url string `json:"url"`
}
func getpokemon(){
	request, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err !=nil{
		fmt.Print(err.Error())
		os.Exit(1)
	}
	reponsedata, err := ioutil.ReadAll(request.Body)
	if err !=nil{
		log.Fatal("Cant get data from url")
	}
	var responsebody Response
	json.Unmarshal(reponsedata,&responsebody)
	fmt.Println(responsebody.Name)
for i := 0; i < len(responsebody.Pokemonentries); i++ {
	fmt.Println(responsebody.Pokemonentries[i].Entrienumber)
	fmt.Println(responsebody.Pokemonentries[i].PokemonSpecies.Name)
	fmt.Println(responsebody.Pokemonentries[i].PokemonSpecies.Url)

}
}

