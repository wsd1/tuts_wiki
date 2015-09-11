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

	if nil != models.WikiM.SaveWikiword(&w) {
		c.Data["json"] = "{\"Word\":\"" + w.Word + "\"}"
	} else {
		c.Data["json"] = "{\"Word\":\"\"}"
	}

	c.ServeJson()
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
	} else {
		c.Data["json"] = "{\"Word\":\"\"}"
	}

	c.ServeJson()
}

func (c *WordsController) Get() {
	var wordStruct *models.Wikiwordstruct
	var wordAttrs []models.Wikiwordattr
	var involved []string
	var beInvolved []string

	// router: /words/?:word
	WordIndicate := c.Ctx.Input.Param(":word")

	//Check login
	if !isLogin(c.Ctx) {
		c.Redirect("/", 302)
	}

	// Word visit history.
	wordPath, wordCurrent := session_word_path(c, WordIndicate)
	//wordCurrent := WordIndicate
	//wordPath := []string{}

	wordStruct = models.WikiM.GetWikiwordByWord(wordCurrent)
	if nil == wordStruct {
		wordStruct = new(models.Wikiwordstruct)
	}

	wordAttrs = models.WikiM.GetAttrsByWord(wordCurrent)
	if nil == wordAttrs {
		wordAttrs = []models.Wikiwordattr{}
	}

	involved = models.WikiM.GetInvolvedByWord(wordCurrent)
	if nil == involved {
		involved = []string{}
	}

	beInvolved = models.WikiM.GetBeInvolvedByWord(wordCurrent)
	if nil == beInvolved {
		beInvolved = []string{}
	}

	c.TplNames = "words.html"

	c.Data["Website"] = "d1works.com"
	c.Data["Email"] = "yizuoshe@gmail.com"
	c.Data["Version"] = "0.1"

	c.Data["WordCurrent"] = wordCurrent
	c.Data["WordPath"] = wordPath

	c.Data["WordContent"] = wordStruct.Content
	c.Data["WordCreate"] = wordStruct.Created
	c.Data["WordModify"] = wordStruct.Modified
	c.Data["WordVisit"] = wordStruct.Visited

	c.Data["WordAttrs"] = wordAttrs
	c.Data["BeInvolved"] = beInvolved
	c.Data["Involved"] = involved

}

func isInHistory(str string, strs []string) (bool, int) {
	for i, v := range strs {
		if v == str {
			return true, i
		}
	}
	return false, len(strs)
}

//input path , current word (be sure in path) and suggest word
//return new path and practical current word
func word_path_roam(path []string, current string, word string) ([]string, string) {

	inPath, _ := isInHistory(word, path)
	//if word in path, just return original path and word as current
	if inPath {
		return path, word
	} else {
		//if word not in path, check if it is in DB
		w := models.WikiM.GetWikiwordByWord(word)

		//if not in DB, skip this word
		if nil == w {
			return path, current

		} else {

			//if in DB, trim path, make sure current word is the last one, and append word as last one, make it current word.

			_, idx := isInHistory(current, path)

			new_path := path[:idx+1]

			new_path = append(new_path, word)

			return new_path, word
		}
	}
}

// Input suggest word, return paht and current.
// This func will update session storage also.
func session_word_path(c *WordsController, new_word string) ([]string, string) {

	// view context, init if not defined
	history := c.GetSession("WordPath")
	if history == nil {
		default_start := beego.AppConfig.String("StartPoint")
		history = []string{default_start}
		c.SetSession("WordPath", history)
	}

	current := c.GetSession("WordNow")
	if current == nil {
		current = history.([]string)[0]
		c.SetSession("WordNow", current)
	}

	new_path := history.([]string)
	new_current := current.(string)

	//If word indicated
	if new_word != "" {
		new_path, new_current = word_path_roam(new_path, new_current, new_word)
		c.SetSession("WordPath", new_path)
		c.SetSession("WordNow", new_current)
	}

	return new_path, new_current
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
