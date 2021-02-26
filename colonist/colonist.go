package colonist

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/rs/xid"
)

type DesireType string

const (
	Fulfillment DesireType = "Fulfillment"
	Belonging   DesireType = "Belonging"
	Esteem      DesireType = "Esteem"
)

type Desires map[DesireType]float64

func (d Desires) Increase(t DesireType, amount float64) {
	d[t] = math.Min(100, d[t]+amount)
}

func (d Desires) Decrease(t DesireType, amount float64) {
	d[t] = math.Max(0, d[t]-amount)
}

type NeedType string

const (
	Hunger     NeedType = "Hunger"
	Thirst     NeedType = "Thirst"
	Stress     NeedType = "Stress"
	Exhaustion NeedType = "Exhaustion"
	Fear       NeedType = "Fear"
	Pain       NeedType = "Pain"
)

type Needs map[NeedType]float64

func (n Needs) Increase(t NeedType, amount float64) {
	n[t] = math.Min(100, n[t]+amount)
}

func (n Needs) Decrease(t NeedType, amount float64) {
	n[t] = math.Max(0, n[t]-amount)
}

type SkillType string

const (
	Hunting     SkillType = "Hunting"
	Crafting    SkillType = "Crafting"
	Cooking     SkillType = "Cooking"
	Building    SkillType = "Building"
	Gathering   SkillType = "Gathering"
	Mining      SkillType = "Mining"
	Woodcutting SkillType = "Woodcutting"
	Science     SkillType = "Science"
	Combat      SkillType = "Combat"
	Charisma    SkillType = "Charisma"
	Medicine    SkillType = "Medicine"
	Tinkering   SkillType = "Tinkering"
)

type Skills map[SkillType]float64

func (n Skills) Increase(t SkillType, amount float64) {
	n[t] += amount
}

func (n Skills) Decrease(t SkillType, amount float64) {
	n[t] = math.Max(0, n[t]-amount)
}

// Colonist struct
type Colonist struct {
	Key    string `json:"key"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Age    uint   `json:"age"`

	Bag       *Bag       `json:"bag"`
	Equipment *Equipment `json:"equipment"`

	Desires Desires `json:"desires"`
	Needs   Needs   `json:"needs"`

	Skills Skills `json:"skills"`
}

// GenerateColonist returns a new colonist with the given name
// and a random set of skills and stats.
func NewColonist() *Colonist {
	return &Colonist{
		Key:  xid.New().String(),
		Name: generateName(),
		Age:  generateAge(),
		Bag: &Bag{
			Size: 30,
		},
		Equipment: &Equipment{},
		Desires:   createDesires(),
		Needs:     createNeeds(),
		Skills:    generateSkills(),
	}
}

func generateName() string {
	firstNames := []string{
		"Merek",
		"Carac",
		"Ulric",
		"Tybalt",
		"Borin",
		"Sadon",
		"Terrowin",
		"Rowan",
		"Forthwind",
		"Althalos",
		"Fendrel",
		"Brom",
		"Hadrian",
		"Benedict",
		"Gregory",
		"Peter",
		"Henry",
		"Frederick",
		"Walter",
		"Thomas",
		"Arthur",
		"Bryce",
		"Donald",
		"Leofrick",
		"Letholdus",
		"Lief",
		"Barda",
		"Rulf",
		"Robin",
		"Gavin",
		"Terrin",
		"Terryn",
		"Ronald",
		"Jarin",
		"Cassius",
		"Leo",
		"Cedric",
		"Gavin",
		"Peyton",
		"Josef",
		"Janshai",
		"Doran",
		"Asher",
		"Quinn",
		"Zane  ",
		"Xalvador",
		"Favian",
		"Destrian",
		"Dain",
		"Berinon",
		"Tristan",
		"Gorvenal",
		"Alfie",
		"Andy",
		"Basil",
		"Buddy",
		"Carter",
		"Charlie",
		"Danny",
		"Eddie",
		"Finn",
		"Freddie",
		"George",
		"Harrison",
		"Hank",
		"Jack",
		"Jonny",
		"Karl",
		"Leo",
		"Leonard",
		"Manny",
		"Mason",
		"Noah",
		"Oscar",
		"Pete",
		"Robin",
		"Sammy",
		"Tim",
		"Toby",
		"Tyler",
		"Victor",
		"Will",
		"Zack",
	}
	lastNames := []string{
		"Ashdown",
		"Baker",
		"Bennett",
		"Bigge",
		"Brickenden",
		"Brooker",
		"Browne",
		"Carpenter",
		"Cheeseman",
		"Clarke",
		"Cooper",
		"Fletcher",
		"Foreman",
		"Godfrey",
		"Hughes",
		"Mannering",
		"Payne",
		"Rolfe",
		"Taylor",
		"Walter",
		"Ward",
		"Webb",
		"Wood",
		"Abbey",
		"Arkwright",
		"Bauer",
		"Bell",
		"Brewster",
		"Chamberlain",
		"Chandler",
		"Chapman",
		"Clarke",
		"Collier",
		"Cooper",
		"Dempster",
		"Harper",
		"Inman",
		"Jenner",
		"Kemp",
		"Kitchener",
		"Knight",
		"Lister",
		"Miller",
		"Packard",
		"Page",
		"Payne",
		"Palmer",
		"Parker",
		"Porter",
		"Rolfe",
		"Ryder",
		"Saylor",
		"Scrivens",
		"Sommer",
		"Smith",
		"Spinner",
	}

	nickNames := []string{
		"Stormbringer",
		"Swiftflight",
		"Sorrowsweet",
		"Shadowmere",
		"Shadowfax",
		"Keirstrider",
		"Bruno",
		"Fireflanel",
		"Curonious",
		"Suntaria",
		"Gypsy",
		"Titanium",
		"Curse",
		"Lightning",
		"Thunder",
		"Storm",
		"Storm Chaser",
		"Goldentrot",
		"Stareyes",
		"Rock",
		"Morningstar",
		"Firefreeze",
		"Veillantif",
		"Winddodger",
		"Hafaleil",
		"Starflare",
		"Elton",
		"Cobalt",
		"Darktonian",
		"Death Bringer",
		"Death Storm",
		"Grand Prancer",
		"Big Heart",
		"Canterwell",
		"Fury",
		"Wildfire",
		"Tempest",
		"Silvermane",
		"Bowrider",
		"Bucephalus",
		"Rocinante",
		"Ares",
		"Blade",
		"Maverick",
		"Tank",
		"Crazy Eyes",
		"Fish",
		"Fox",
		"Grant",
		"Hardy",
		"Hawk",
		"Hendman",
		"Keen",
		"Lightfoot",
		"Mannering",
		"Moody",
		"Mundy",
		"Peacock",
		"Power",
		"Pratt",
		"Proude",
		"Pruitt",
		"Puttock",
		"Quick",
		"Rey",
		"Rose",
		"Russ",
		"Selly",
		"Sharp",
		"Short",
		"Sommer",
		"Sparks",
		"Spear",
		"Stern",
		"Sweet",
		"Tait",
		"Terrell",
		"Truman",
	}

	firstName := firstNames[rand.Intn(len(firstNames))]
	lastName := lastNames[rand.Intn(len(lastNames))]
	nickName := nickNames[rand.Intn(len(nickNames))]

	return fmt.Sprintf("%s \"%s\" %s", firstName, nickName, lastName)
}

func generateAge() uint {
	return 20 + uint(rand.Intn(400))
}

func (c *Colonist) SetStatus(s string) {
	c.Status = s
}

func createDesires() map[DesireType]float64 {
	return map[DesireType]float64{
		Fulfillment: float64(rand.Intn(60)),
		Belonging:   float64(rand.Intn(60)),
		Esteem:      float64(rand.Intn(60)),
	}
}

func createNeeds() map[NeedType]float64 {
	return map[NeedType]float64{
		Hunger:     float64(rand.Intn(80)),
		Thirst:     float64(rand.Intn(80)),
		Stress:     float64(rand.Intn(80)),
		Exhaustion: float64(rand.Intn(80)),
		Fear:       float64(rand.Intn(80)),
		Pain:       float64(rand.Intn(80)),
	}
}

func generateSkills() map[SkillType]float64 {
	var values [11]float64

	for i := 0; i < 2; i++ {
		values[i] = 4 + (rand.Float64() * 10)
	}

	for i := 2; i < 8; i++ {
		values[i] = 3 + (rand.Float64() * 5)
	}

	for i := 8; i < 11; i++ {
		values[i] = rand.Float64() * 2
	}

	rand.Shuffle(len(values), func(i, j int) {
		values[i], values[j] = values[j], values[i]
	})

	return map[SkillType]float64{
		Hunting:     values[0],
		Crafting:    values[1],
		Cooking:     values[2],
		Building:    values[3],
		Gathering:   values[4],
		Mining:      values[5],
		Woodcutting: values[6],
		Science:     values[7],
		Combat:      values[8],
		Charisma:    values[9],
		Medicine:    values[10],
		Tinkering:   values[10],
	}
}
