package controllers

import (
	"github.com/astaxie/beego"
	//	"log"
	"tuts_wiki/models"
)

type HomeController struct {
	beego.Controller
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
		_, inDB := models.WikiM.GetWikiwordByWord(word)

		//if not in DB, skip this word
		if !inDB {
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

func (c *HomeController) Get() {
	c.Data["Website"] = "d1works.com"
	c.Data["Email"] = "yizuoshe@gmail.com"
	c.Data["Version"] = "0.1"
	c.TplNames = "home.html"

	if !isLogin(c.Ctx) {
		c.Redirect("/", 302)
	}

	// view context, init if not defined
	history := c.GetSession("WordPath")
	current := c.GetSession("WordNow")
	if history == nil {
		default_start := beego.AppConfig.String("StartPoint")
		history = []string{default_start}
		c.SetSession("WordPath", history)
	}
	if current == nil {
		current = history.([]string)[0]
		c.SetSession("WordNow", current)
	}

	new_path := history.([]string)
	new_current := current.(string)

	//If word indicated
	WordIndicate := c.Input().Get("select")
	if WordIndicate != "" {

		new_path, new_current = word_path_roam(new_path, new_current, WordIndicate)
		c.SetSession("WordPath", new_path)
		c.SetSession("WordNow", new_current)
	}

	var wordContent string
	wordStruct, inDB := models.WikiM.GetWikiwordByWord(new_current)
	if inDB {
		wordContent = string(wordStruct.Content)
	} else {
		wordContent = "NULL"
	}

	var wordAttrs []models.Wikiwordattr
	wordAttrs, inDB = models.WikiM.GetAttrsByWord(new_current)
	if !inDB {
		wordAttrs = []models.Wikiwordattr{}
	}

	var involved []string
	involved, inDB = models.WikiM.GetInvolvedByWord(new_current)
	if !inDB {
		involved = []string{}
	}

	var beinvolved []string
	beinvolved, inDB = models.WikiM.GetBeInvolvedByWord(new_current)
	if !inDB {
		beinvolved = []string{}
	}

	c.Data["WordCurrent"] = new_current
	c.Data["WordPath"] = new_path
	c.Data["WordContent"] = wordContent
	c.Data["WordAttrs"] = wordAttrs

	c.Data["BeInvolved"] = beinvolved
	c.Data["Involved"] = involved

}
