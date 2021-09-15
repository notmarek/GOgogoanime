package main

import (
	"regexp"
	"strings"
)

func minify_text(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(s, " ", ""), "\n", ""), "\t", "")
}

func parse_download_one(page string) string {
	re := regexp.MustCompile("dowloads\"><ahref=\"(.*?)\"target")
	return re.FindStringSubmatch(page)[1]
}

func parse_download_two(page string) string {
	re := regexp.MustCompile("\"dowload\"><ahref=\"(.*?)\"download")
	return re.FindStringSubmatch(page)[1]
}

func parse_anime_url(url string) string {
	re := regexp.MustCompile("https://gogoanime.pe/category/(.*?)$")
	return re.FindStringSubmatch(url)[1]
}
