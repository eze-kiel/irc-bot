package commands

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
)

type chuck struct {
	Value string `json:"value"`
}

type xkcdComic struct {
	Title string `json:"title"`
	Img   string `json:"img"`
}

type yesOrNo struct {
	Answer string `json:"answer"`
}

type idea struct {
	This string `json:"this"`
	That string `json:"that"`
}

type kanyeQuote struct {
	Quote string `json:"quote"`
}

type motivSpeech struct {
	Affirmation string `json:"affirmation"`
}

// ChuckFact returns a Chuck Norris fact
func ChuckFact() string {
	resp, err := http.Get("https://api.chucknorris.io/jokes/random")
	if err != nil {
		return "I had a problem..."
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	var fact chuck
	json.Unmarshal([]byte(bodyBytes), &fact)
	return fact.Value
}

// Help displays the differents options available
func Help() string {
	return "Commands are detailled here : https://github.com/eze-kiel/irc-bot/blob/master/README.md"
}

// Trivia returns a fun fact about a number
func Trivia() string {
	resp, err := http.Get("http://numbersapi.com/random/trivia")
	if err != nil {
		return "I had a problem..."
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	return string(bodyBytes)
}

// Year returns a fun fact about a specific year
func Year() string {
	resp, err := http.Get("http://numbersapi.com/random/year")
	if err != nil {
		return "I had a problem..."
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	return string(bodyBytes)
}

// Date returns a fun fact about a specific date
func Date() string {
	resp, err := http.Get("http://numbersapi.com/random/date")
	if err != nil {
		return "I had a problem..."
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	return string(bodyBytes)
}

// Math returns a fun fact about math
func Math() string {
	resp, err := http.Get("http://numbersapi.com/random/math")
	if err != nil {
		return "I had a problem..."
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	return string(bodyBytes)
}

// Xkcd returns a XKCD's random comic
func Xkcd() string {
	comicNumber := rand.Intn(2317)
	comicNumberStr := strconv.Itoa(comicNumber)
	resp, err := http.Get("https://xkcd.com/" + comicNumberStr + "/info.0.json")
	if err != nil {
		return "I had a problem..."
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	var comic xkcdComic
	json.Unmarshal([]byte(bodyBytes), &comic)
	return comic.Title + " : " + comic.Img
}

// YesOrNo returns... yes or no
func YesOrNo() string {
	resp, err := http.Get("https://yesno.wtf/api/")
	if err != nil {
		return "I had a problem..."
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	var answer yesOrNo
	json.Unmarshal([]byte(bodyBytes), &answer)
	return answer.Answer
}

// Idea returns a business idea
func Idea() string {
	resp, err := http.Get("http://itsthisforthat.com/api.php?json")
	if err != nil {
		return "I had a problem..."
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	var concept idea
	json.Unmarshal([]byte(bodyBytes), &concept)
	return "you should create " + concept.This + " for " + concept.That
}

// Kanye returns a Kanye West quote
func Kanye() string {
	resp, err := http.Get("https://api.kanye.rest/")
	if err != nil {
		return "I had a problem..."
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	var quote kanyeQuote
	json.Unmarshal([]byte(bodyBytes), &quote)
	return quote.Quote
}

// Motivation returns a motivation quote
func Motivation() string {
	resp, err := http.Get("https://www.affirmations.dev/")
	if err != nil {
		return "I had a problem..."
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	var motiv motivSpeech
	json.Unmarshal([]byte(bodyBytes), &motiv)
	return motiv.Affirmation
}
