package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
	"time"
	"tuts_wiki/models"
)

type WordsController struct {
	beego.Controller
}

//POST /words：新建一个
func (c *WordsController) Post() {

	// router: /words/?:word
	if c.Ctx.Input.Param(":word") != "" {
		log.Println("Error:Target must be: /words")
		c.Abort("409") //"409" : Conflict
	}

	w := models.Wikiwordstruct{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &w)
	//	log.Println("Parse to:", w)

	//update time stamp
	inheritTimeStamp(&w, nil)
	//	log.Println("Timestamp:", w)

	if nil != models.WikiM.NewWikiword(&w) {
		c.Data["json"] = "{\"Word\":\"" + w.Word + "\"}"
		c.ServeJson()
	} else {
		c.Abort("500")
	}

}

//PUT /words/xxxx：更新某个指定信息（提供该全部信息）
func (c *WordsController) Put() {

	// router: /words/?:word
	WordIndicate := c.Ctx.Input.Param(":word")
	//	log.Println("Update word:", WordIndicate)
	//	log.Println("Get body:", string(c.Ctx.Input.RequestBody))

	if WordIndicate == "" {
		log.Println("Error:Target must be: /words/xxxx")
		c.Abort("409") //"409" : Conflict
	}

	w := models.Wikiwordstruct{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &w)
	//	log.Println("Parse to:", w)

	if WordIndicate != w.Word {
		log.Printf("Err: router[%s] != w.word[%s]\r\n", WordIndicate, w.Word)
		c.Abort("409") //"409" : Conflict
	}

	//update time stamp
	inheritTimeStamp(&w, models.WikiM.GetWikiwordByWord(WordIndicate))
	//	log.Println("Timestamp:", w)

	if nil != models.WikiM.SaveWikiword(&w) {
		c.Data["json"] = "{\"Word\":\"" + w.Word + "\"}"
		c.ServeJson()
	} else {
		c.Abort("500")
	}

}

func (c *WordsController) Get() {
	var wordStruct *models.Wikiwordstruct
	var wordAttrs []models.Wikiwordattr
	var involved []string
	var beInvolved []string
	var isNew bool

	// router: /words/?:word
	WordIndicate := c.Ctx.Input.Param(":word")

	//Check login
	if !isLogin(c.Ctx) {
		c.Redirect("/", 302)
	}

	// /words will redirect to /words/xxx
	if "" == WordIndicate {
		start := beego.AppConfig.String("StartPoint")
		c.Redirect("/words/"+start, 302)
	}

	wordStruct = models.WikiM.GetWikiwordByWord(WordIndicate)
	if nil == wordStruct {
		wordStruct = new(models.Wikiwordstruct)
	}

	// Word visit history.
	//wordPath, wordCurrent := session_word_path(c, WordIndicate)
	//wordCurrent := WordIndicate
	//wordPath := []string{}

	wordStruct = models.WikiM.GetWikiwordByWord(WordIndicate)
	if nil == wordStruct {
		wordStruct = &models.Wikiwordstruct{
			Word:        WordIndicate,
			Content:     "",
			Compression: false,
			Encryption:  false,
			Created:     0.0,
			Modified:    0.0,
			Visited:     0.0,
			Readonly:    false,
		}
		isNew = true
	}

	wordAttrs = models.WikiM.GetAttrsByWord(WordIndicate)
	if nil == wordAttrs {
		wordAttrs = []models.Wikiwordattr{}
	}

	involved = models.WikiM.GetInvolvedByWord(WordIndicate)
	if nil == involved {
		involved = []string{}
	}

	beInvolved = models.WikiM.GetBeInvolvedByWord(WordIndicate)
	if nil == beInvolved {
		beInvolved = []string{}
	}

	c.TplNames = "words.html"

	c.Data["Website"] = "d1works.com"
	c.Data["Email"] = "yizuoshe@gmail.com"
	c.Data["Version"] = "0.1"

	c.Data["isNew"] = isNew

	c.Data["WordPath"] = session_word_path(c, WordIndicate)

	c.Data["WordCurrent"] = wordStruct.Word
	c.Data["WordContent"] = wordStruct.Content
	c.Data["WordCreate"] = wordStruct.Created
	c.Data["WordModify"] = wordStruct.Modified
	c.Data["WordVisit"] = wordStruct.Visited

	c.Data["WordAttrs"] = wordAttrs
	c.Data["BeInvolved"] = beInvolved
	c.Data["Involved"] = involved

}

func isInHistory(strs []string, str string) (bool, int) {
	for i, v := range strs {
		if v == str {
			return true, i
		}
	}
	return false, len(strs)
}

// Input suggest word, return paht and current.
// This func will update session storage also.
func session_word_path(c *WordsController, new_word string) []string {

	var history []string
	var index int

	if new_word == "" {
		return nil
	}

	// Init if not defined
	if nil == c.GetSession("WordPath") {
		history = []string{new_word}
		c.SetSession("WordPath", history)
	} else {
		history = c.GetSession("WordPath").([]string)
	}

	if nil == c.GetSession("WordPathIdx") {
		index = 0
		c.SetSession("WordPathIdx", index)
	} else {
		index = c.GetSession("WordPathIdx").(int)
	}

	ok, idx := isInHistory(history, new_word)

	if ok {
		//if new word in history, update index only
		c.SetSession("WordPathIdx", idx)

	} else {

		// Add new into history, append it after last indexed word
		history = history[:index+1]
		history = append(history, new_word)

		c.SetSession("WordPath", history)
		c.SetSession("WordPathIdx", index+1)
	}

	return history
}

//This func will setup new word's ts property via old one
func inheritTimeStamp(w, old *models.Wikiwordstruct) {
	if nil == old {
		w.Created = float64(time.Now().Unix())
		w.Modified = w.Created
		w.Visited = w.Created
	} else {
		w.Created = old.Created
		w.Modified = float64(time.Now().Unix())
		w.Visited = old.Visited
	}
}
