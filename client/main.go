package main

import (
	"log"

	"github.com/icepie/oh-my-lit/client/jwc"
)

func main() {

	jwcUser := jwc.NewJwCUser()

	var next string
	for {

		data, err := jwcUser.GetGGTZPost(next)
		if err != nil {
			panic(err)
		}

		if data.Next != "" {
			log.Println(data.Next)
		}

		next = data.Next

		if data.Next == "" {
			break
		}

		log.Println(data.Posts)

	}

	// log.Println(data)
}
