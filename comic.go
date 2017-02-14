package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string
	Actors []string
}

type Comic struct {
	Month      string
	Num        int
	Year       string
	Transcript string
}

var movies1 []Movie
var comic Comic

func main() {
	b := `[{"Title":"Great Wall","Actors":["g","y","a"], "date":"2017"},{"Title":"Hero","Actors":["a","b","c"], "data":"2018"}]`
	b1 := `{"month": "4", "num": 571, "link": "", "year": "2009", "news": "", "safe_title": "Can't Sleep", "transcript": "[[Someone is in bed, presumably trying to sleep. The top of each panel is a thought bubble showing sheep leaping over a fence.]]\n1 ... 2 ...\n<<baaa>>\n[[Two sheep are jumping from left to right.]]\n\n... 1,306 ... 1,307 ...\n<<baaa>>\n[[Two sheep are jumping from left to right. The would-be sleeper is holding his pillow.]]\n\n... 32,767 ... -32,768 ...\n<<baaa>> <<baaa>> <<baaa>> <<baaa>> <<baaa>>\n[[A whole flock of sheep is jumping over the fence from right to left. The would-be sleeper is sitting up.]]\nSleeper: ?\n\n... -32,767 ... -32,766 ...\n<<baaa>>\n[[Two sheep are jumping from left to right. The would-be sleeper is holding his pillow over his head.]]\n\n{{Title text: If androids someday DO dream of electric sheep, don't forget to declare sheepCount as a long int.}}", "alt": "If androids someday DO dream of electric sheep, don't forget to declare sheepCount as a long int.", "img": "https://imgs.xkcd.com/comics/cant_sleep.png", "title": "Can't Sleep", "day": "20"}`
	if err := json.Unmarshal([]byte(b), &movies1); err != nil {
		fmt.Printf("json unmarshal failed : %s ", err)
	}

	fmt.Printf("movies1 %v \n", movies1)

	if err := json.Unmarshal([]byte(b1), &comic); err != nil {
		fmt.Printf("json unmarshal failed : %s ", err)
	}

	fmt.Printf("commic %v \n", comic)
}
