package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/byplayer/download_podcast_go/models"
)

func main() {
	response, err := http.Get("https://www.tfm.co.jp/podcasts/asemamire/podcast.xml")

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	rss := models.Rss{}
	decoder := xml.NewDecoder(response.Body)
	err = decoder.Decode(&rss)

	if err != nil {
		log.Fatal(err)
	}

	for _, item := range rss.Channel.Items {
		fname := filepath.Base(item.Enclosure.Url)
		fmt.Printf("'%s','%s'\n", item.Enclosure.Url, fname)

		r2 := regexp.MustCompile(`(.*_vol)([\d]+)(.*)`)
		result := r2.FindAllStringSubmatch(fname, -1)
		// fmt.Println(result[0])
		// fmt.Println(len(result[0]))
		// fmt.Println(result[0][0])
		i, _ := strconv.Atoi(result[0][2])
		newFName := fmt.Sprintf("%s%03d%s", result[0][1], i, result[0][3])
		fmt.Println(newFName)

		//Store Data Here(Establish Prior Connection With DB)
		//Store.Insert(data)
	}
}
