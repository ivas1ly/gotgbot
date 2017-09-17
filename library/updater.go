package library

import (
	"encoding/json"
	"log"
	"strconv"
	"time"
	"bot/library/Ext"
)

type Updater struct {
	bot        Ext.Bot
	updates    chan Update
	Dispatcher Dispatcher
}

type InlineQuery struct {

}

type ChosenInlineResult struct {

}

type ShippingQuery struct {

}

type PreCheckoutQuery struct {

}

func NewUpdater(token string) Updater {
	u := Updater{}
	u.bot = Ext.Bot{Token: token}
	u.updates = make(chan Update)
	u.Dispatcher = NewDispatcher(u.bot, u.updates)
	return u
}

func (u Updater) Start_polling() {
	go u.Dispatcher.Start()
	go u.start_polling()
}


func (u Updater) start_polling() {
	m := make(map[string]string)
	m["offset"] = strconv.Itoa(0)
	m["timeout"] = strconv.Itoa(0)
	for {
		r := Ext.Get(u.bot, "getUpdates", m)
		if !r.Ok {
			log.Fatal("You done goofed, API Res for getUpdates was not OK")

		}
		offset := 0
		if r.Result != nil {
			//fmt.Println(r)
			var res []Update
			json.Unmarshal(r.Result, &res)
			for _, upd := range res {
				u.updates <- upd
			}
			if len(res) > 0 {
				offset = res[len(res)-1].Update_id + 1
			}
		}

		m["offset"] = strconv.Itoa(offset)

	}
}

func (u Updater) Idle() {
	for {
		time.Sleep(1 * time.Second)
	}

}