package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dvbnrg/JokeGenerator/model"
)

func getName() (model.Name, error) {

	name, err := http.Get("http://uinames.com/api/")

	if err != nil {
		log.Printf("name url error: %v", err)
		return model.Name{}, err
	}

	body, err := ioutil.ReadAll(name.Body)

	if err != nil {
		log.Printf("name body read error: %v", err)
		return model.Name{}, err
	}

	n := model.Name{}

	err = json.Unmarshal(body, &n)

	if err != nil {
		log.Printf("name body read error: %v", err)
		return model.Name{}, err
	}

	return n, nil
}

func GetJoke() (model.Joke, error) {

	n, err := getName()

	if err != nil {
		log.Printf("name retrieval error %v", err)
		return model.Joke{}, err
	}

	jokeURL := fmt.Sprintf("http://api.icndb.com/jokes/random?firstName=%v&lastName=%v", n.Name, n.Surname)

	joke, err := http.Get(jokeURL)

	if err != nil {
		log.Printf("joke url error: %v", err)
		return model.Joke{}, err
	}

	body, err := ioutil.ReadAll(joke.Body)

	if err != nil {
		log.Printf("joke body read error: %v", err)
		return model.Joke{}, err
	}

	j := model.Joke{}

	err = json.Unmarshal(body, &j)

	if err != nil {
		log.Printf("joke body parse error: %v", err)
		return model.Joke{}, err
	}

	return j, nil
}
