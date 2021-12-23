package runner

import (
	"hc21f/pkg/database"
	"hc21f/pkg/job"
	"hc21f/pkg/twitter"
	"log"
	"os"
)

func Run() {
	if err := twitter.Init(); err != nil {
		log.Fatalf("[ ERROR ]: twitter.Init() \n%s", err)
	}

	if err := database.Init(); err != nil {
		log.Fatalf("[ ERROR ]: database.Init() \n%s", err)
	}

	if len(os.Args) < 2 {
		log.Fatal("[ ERROR ]: must specify job as os.Args")
	}

	arg := os.Args[1]

	t := twitter.Get()
	db := database.Get()

	r, err := job.New(t, db, arg)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	if err := r.Do(); err != nil {
		log.Fatalf("[ ERROR ]: %s", err)
	}

}
