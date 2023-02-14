package main

import (
	"log"

	"github.com/icepie/oh-my-lit/client/jwc"
)

func main() {

	jwcUser := jwc.NewJwCUser()

	var next string
	for {

		data, err := jwcUser.GetGGTZPostList(next)
		if err != nil {
			panic(err)
		}

		next = data.Next

		if data.Next == "" {
			break
		}

		for _, post := range data.Posts {
			data, err := jwcUser.GetGGTZPost(post.Url)
			if err != nil {
				panic(err)
			}

			log.Println(data)
		}

		log.Println(data.Posts)

	}

	// log.Println(data)
}
