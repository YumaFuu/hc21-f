package runner

import (
	"hc21f/pkg/database"
	"hc21f/pkg/job"
	"hc21f/pkg/twitter"
	"log"
)

func Run(args []string) {
	if err := twitter.Init(); err != nil {
		log.Fatalf("[ ERROR ]: twitter.Init() \n%s", err)
	}

	if err := database.Init(); err != nil {
		log.Fatalf("[ ERROR ]: database.Init() \n%s", err)
	}

	if len(args) < 2 {
		log.Fatal("[ ERROR ]: must specify job as os.Args")
	}

	t := twitter.Get()
	db := database.Get()

	j, err := job.New(t, db, args[1])
	if err != nil {
		log.Fatalf("[ ERROR ]: job.New() %s", err)
	}

	if err := j.Do(); err != nil {
		log.Fatalf("[ ERROR ]: job.Do() %s", err)
	}
}
