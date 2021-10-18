package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Enter the anime's url (if you enter a specific episode url only that episode will be downloaded): ")

	var gogourl string

	fmt.Scanln(&gogourl)

	if strings.Contains(gogourl, "-episode-") {
		dl_one := parse_download_one(get_url(gogourl))
		dl_two := parse_download_two(get_url(dl_one))
		download(".", dl_two)
	} else {
		anime_id := parse_anime_url(gogourl)
		i := 0
		for {
			i++
			episode_page := get_url("https://gogoanime.pe/" + anime_id + "-episode-" + fmt.Sprint(i))
			if strings.Contains(episode_page, ">404</h1>") {
				break
			}
			dl_one := parse_download_one(episode_page)
			dl_two := parse_download_two(get_url(dl_one))
			download("./"+anime_id, dl_two)
		}
	}
}
