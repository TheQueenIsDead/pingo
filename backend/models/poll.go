package models

import (
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
	"io"
	"net/http"
	"sync"
	"time"
)

type Poll struct {
	gorm.Model
	TargetIds       []int     `json:"target_ids,omitempty" gorm:"-"`                               // Skip persisting to db, but use in JSON for transfer
	Targets         []Target  `json:"-" gorm:"many2many:poll_targets;foreignKey:ID;references:ID"` // Skip representing Target objects in JSON, but use foreign keys
	DurationSeconds int       `json:"duration"`
	StartedAt       time.Time `json:"started"`
}

func (p *Poll) AfterSave(tx *gorm.DB) error {
	log.Info("AfterSave hook triggered for poll.")

	var pl []Poll
	err := tx.Preload("Targets").Find(&pl, p.ID).Error
	if err != nil {
		log.Error(err)
		return err
	}
	go handle(pl)

	return nil
}

// handle is a method that takes a fully inflated Poll object with
// associated Target structs, and orchestrates poll jobs,
// collating the results
func handle(pl []Poll) {

	log.Info("Handling poll...")

	var results []Result

	var wg sync.WaitGroup
	var ch chan Result

	for _, p := range pl {
		for _, t := range p.Targets {
			log.Info(t.ID, t.Source)
			wg.Add(1)
			go poll(t, wg, ch)
		}
	}

	log.Info("Waiting for poll results...")
	wg.Wait()
	log.Info("Wait group complete!")

	for r := range ch {
		log.Info("Code", r)
		results = append(results, r)
	}
	log.Info("Done!")
}

// poll takes a target, wait-group and channel of results. It will poll the target before
// pausing for the specified target interval, and will repeat for the specified count
// of hits on the target.
func poll(t Target, wg sync.WaitGroup, ch chan Result) {

	defer wg.Done()

	for count := 0; count <= t.Frequency; count++ {

		log.Infof("[%d] Polling: %v", count, t)

		res, err := http.Get(t.Source)

		var dur time.Duration
		switch t.Unit {
		case "SECOND":
			dur = time.Second
		case "MINUTE":
			dur = time.Minute
		default:
			dur = time.Second
		}

		time.Sleep(dur)

		var msg string
		if err != nil {
			msg = err.Error()

		} else {
			body, _ := io.ReadAll(res.Body)
			msg = string(body)
		}

		errb := false
		if err != nil {
			errb = true
		}
		ch <- Result{
			Poll:    Poll{},
			Target:  t,
			Code:    res.StatusCode,
			Message: msg,
			Error:   errb,
		}
	}
	log.Info("All polls complete for target.")
	wg.Done()
}
