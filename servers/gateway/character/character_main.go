package character

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type Character struct {
	Strength       int `json:"Strength"`
	Dexterity      int
	Constitution   int
	Intelligence   int
	Wisdom         int
	Charisma       int
	Acrobatics     int
	AnimalHandling int
	Arcana         int
	Athletics      int
	Deception      int
	History        int
	Insight        int
	Intimidation   int
	Investigation  int
	Medicine       int
	Nature         int
	Perception     int
	Performance    int
	Persuasion     int
	Religion       int
	SleightOfHand  int
	Stealth        int
	Survival       int
	SVSTR          int
	SVDEX          int
	SVCON          int
	SVINT          int
	SVWIS          int
	SVCHA          int
}

func CharacterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	query := r.FormValue("url")
	if len(query) == 0 {
		err := fmt.Sprintf("%d Bad Request", http.StatusBadRequest)
		http.Error(w, err, http.StatusBadRequest)
		return
	}

	body, err := fetchHTML(query)
	if err != nil {
		log.Print("Lol")
	}
	defer body.Close()
}

func fetchHTML(pageURL string) (io.ReadCloser, error) {
	res, err := http.Get(pageURL)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("%d HTTP Error Code", res.StatusCode)
	}

	contType := res.Header.Get("Content-Type")
	if !strings.HasPrefix(contType, "text/html") {
		return nil, fmt.Errorf("content is not a web page")
	}

	return res.Body, nil
}

func extractStats(pageURL string) (*Character, error) {
	returnCharacter := Character{}
	c := colly.NewCollector(
		colly.AllowedDomains("dndbeyond.com"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL)
	})

	c.OnHTML(".ct-skills__item", func(e *colly.HTMLElement) {
		checkText(e.ChildText(".ct-skills__col--skill"), &returnCharacter, e)
	})
	return &returnCharacter, nil

}

func checkText(skillName string, character *Character, e *colly.HTMLElement) {
	switch skillName {
	case "Acrobatics":
		character.Acrobatics = getVal(e)
	case "Animal Handling":
		character.AnimalHandling = getVal(e)
	case "Arcana":
		character.Arcana = getVal(e)
	case "Athletics":
		character.Athletics = getVal(e)
	case "Deception":
		character.Deception = getVal(e)
	case "History":
		character.History = getVal(e)
	case "Insight":
		character.Insight = getVal(e)
	case "Intimidation":
		character.Intimidation = getVal(e)
	case "Investigation":
		character.Investigation = getVal(e)
	case "Medicine":
		character.Medicine = getVal(e)
	case "Nature":
		character.Nature = getVal(e)
	case "Perception":
		character.Perception = getVal(e)
	case "Performance":
		character.Performance = getVal(e)
	case "Persuasion":
		character.Persuasion = getVal(e)
	case "Religion":
		character.Religion = getVal(e)
	case "Sleight Of Hand":
		character.SleightOfHand = getVal(e)
	case "Stealth":
		character.Stealth = getVal(e)
	case "Survival":
		character.Survival = getVal(e)
	default:
		return

	}
}

func getVal(e *colly.HTMLElement) int {
	convertedVal, err := strconv.Atoi(e.ChildText(".ct-signed-number__number"))
	if err != nil {
		return 0
	}
	if e.ChildText("ct-signed-number__sign") == "-" {
		convertedVal = -convertedVal
	}
	return convertedVal
}
