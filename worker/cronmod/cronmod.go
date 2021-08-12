package cronmod

import (
	"github.com/robfig/cron/v3"
	"log"
)

var CronInstance *cron.Cron

func Init() {
	CronInstance = cron.New()
	_, _ = CronInstance.AddFunc("0 * * * *",
		func() {
			log.Printf("[cronmod, check] info: cronmod at working")
		},
	)
	CronInstance.Start()
	log.Printf("[cronmod, init] info: cronmod module is start")
}

func AddFunc(spec string, cmd func()) (cron.EntryID, error) {
	entryID, err := CronInstance.AddFunc(spec, cmd)
	if err != nil {
		log.Printf("[cronmod, addfunc] error: %v", err)
	}
	return entryID, err
}

func Remove(id cron.EntryID) {
	CronInstance.Remove(id)
}
