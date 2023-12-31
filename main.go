package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var quotes = []string{
	"Should you be here?",
	"I have a bad feeling about this.",
	"ROOOAAR",
	"insert funny text here",
	"twitch funny moments youtube search",
	"wth am i?",
	"bro u lost?",
	"have you played jedi fallen order? great game",
	"i use arch btw",
	"i love you... wrong person",
	"c++ hello world speedrun",
	"fortnite battlepass",
	"WOOOW",
	"not funny didn't laugh",
	"You know, Amy, anytime someone calls attention to the breaking of gender roles, it ultimately undermines the concept of gender equality by implying that this is an exception and not the status quo.",
	"who am i?",
	"Say my name.",
	"Better call Saul!",
	"somebody once told me",
	"shrek 1 is the best movie change my mind",
	"thank you markiplier",
	"thank you pewdiepie",
	"tseries can suck my :pleading_face:",
	"what the hell is an error 69?",
	"have you found all the easter eggs in this site?",
	"bro's lost lmao",
	"Luke, I am your father.",
	"i lost 4 hours just thinking random quotes lol",
	"beta testing available right now over xvideos.com!",
	"wanna see my duck duck DUCK I MEANT DUCK DUCK nvm",
	"DAMN YOU!",
	"why u here?",
	"Hello there.",
	"bleading edge software is kinda funny",
	"i rate wen peopel can not speel thing rights",
	"accurate",
	"pls mommy give me da JUICE",
	"life is overrated",
	"check this out: https://github.com/ValveSoftware/steam-for-linux/issues/1890 it's a goldmine",
	"I mean... Ducks are a great source of food.",
	"ithinkitsprettynormalforsomeonetonotusespaces",
	"bro's finding funny just F5'ing the page",
	"TO THE MOON",
	"DANG IT",
	"i love you arch linux... i love you",
	"imagine going to a 404 page to see which is the quote of the day",
	"I sense good luck on you.",
	"I sense duck on you.",
	"shame on you, i wash my ass every time i poo.",
	"why people say 'you are welcome'?",
	"dumb question: do you still love me veronica??",
	"open-source software > close-source software",
	"github.com/_______",
	"im tired of writing quotes but i cant stop",
	"pls veronica come back to me",
	"nothing here. return later",
	"const life = null;",
	"let veronica = lover::new().thinkofher();",
	"I-I love my girlfriend",
	"Rawr",
	"Execute order 66",
	"wtf man?",
	"I sense great force in you",
	"I'm not locked in here with you, you're locked in here with me.",
	"You have become the very thing you swore to destroy.",
}

var people = []string{
	"Dad", "Mom", "Luke Skywalker", "Anakin Skywalker", "Darth Vader",
	"Yoda", "Chewbacca", "Gerald", "Shrek", "Knuckles",
	"Sonic", "Shadow", "Garcello", "devkcud", "Amy",
	"Han Solo", "Pewdiepie", "Markiplier", "Ed Murphy", "Walter White",
	"Jesse Pinkman", "Cal Kestis", "Second Sister", "Ninth Sister", "Grue",
	"Gru", "Wreck-it, Ralph", "Donkey", "Lord Farquaad", "Fiona",
	"Peter Pan", "Magic Mirror", "Geppetto", "Robin Hood", "Princess Leia",
	"Boba Fett", "Palpatine", "Obi-Wan Kenobi", "Ben Kenobi", "C-3PO",
	"R2-D2", "Mario", "Wario", "Luigi", "Yoshi",
	"Waluigi", "Vanellope von Schweetz", "Zangief", "Chun-Li", "Ryu",
	"Ken", "Barbie", "Sticks", "Saul", "Jango Fett",
	"Jar Jar Binks", "Rorschach",
}

var emojis = []string{
	"💀", "😎", "🤗", "😑", "🥱", "🫠",
	"🤭", "🤤", "🤢", "🤮", "💩", "🤡",
	"👽", "😈", "🤠", "😱", "😭", "🥵",
	"🥶", "🤯", "🥺", "😩", "🤓", "🥰",
	"😘", "☺️", "😁", "😀", "😜", "😉",
	"🥸", "🙁", "😕", "😔", "😒", "🤪",
	"🥳", "😡", "😠", "😤", "😨", "🤫",
	"🤥", "🙄", "😬", "😐", "🤔", "🥴",
	"👊", "👍", "👎", "💪", "🫦", "💋",
	"👀", "🤝", "🖕",
}

type Quote struct {
	Quote  string `json:"quote"`
	Person string `json:"person"`
	Emoji  string `json:"emoji"`
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(CORSMiddleware())

	r.GET("/", getQuotes)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	err := r.Run("0.0.0.0:" + port)

	if err != nil {
		fmt.Println(err)
	}
}

func getQuotes(c *gin.Context) {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	c.JSON(200, Quote{
		Quote:  quotes[random.Intn(len(quotes))],
		Person: people[random.Intn(len(people))],
		Emoji:  emojis[random.Intn(len(emojis))],
	})
}
