package main

import (
	"log"
	"net/http"

	b "github.com/rinwaowuogba/cloud-resume-project/backend"
)

func main() {
	client := b.GetFirestoreClient()
	server := b.NewVisitCountServer(b.MakeFirestoreClientAdapter(client))
	log.Fatal(http.ListenAndServe(":5000", server))
}
