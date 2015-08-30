package models

import (
	"log"
)

//sql tutorial can be found at:
//	http://segmentfault.com/a/1190000003036452
//	http://go-database-sql.org/retrieving.html

type Wikiwordstruct struct {
	Word        string
	Content     []byte
	Compression bool
	Encryption  bool

	Created  float64
	Modified float64
	Visited  float64

	Readonly bool
	//	Metadataprocessed     int
	//	Presentationdatablock string
}

type Wikiwordattr struct {
	Key   string
	Value string
}
type WikiwordModel struct {
	wordsCache map[string]*Wikiwordstruct
	attrsCache map[string][]Wikiwordattr
	beinvCache map[string][]string
	involCache map[string][]string
}

func (this *WikiwordModel) Reset() {
	this.wordsCache = make(map[string]*Wikiwordstruct)
	this.attrsCache = make(map[string][]Wikiwordattr)

	this.beinvCache = make(map[string][]string)
	this.involCache = make(map[string][]string)
	//	this.GetAllUser()
}

func NewWikiModel() *WikiwordModel {
	wikiM := new(WikiwordModel)
	wikiM.Reset()
	return wikiM
}

// get word struct by word.
// if no cached, query from db and cache it.
func (this *WikiwordModel) GetWikiwordByWord(w string) (*Wikiwordstruct, bool) {

	//cache check
	c, ok := this.wordsCache[w]
	if ok {
		return c, true
	}

	//sql retrive
	word := new(Wikiwordstruct)

	row := SqlDb.QueryRow("SELECT word,content,compression,encryption,created,modified,visited,readonly FROM wikiwordcontent WHERE word=?", w)
	err := row.Scan(&word.Word, &word.Content, &word.Compression, &word.Encryption, &word.Created, &word.Modified, &word.Visited, &word.Readonly)
	if err != nil {
		log.Println("Get err when query: " + w)
		log.Println(err)
		return nil, false
	}

	// cache and return it
	if word.Word == w {
		this.wordsCache[w] = word
		return word, true
	}

	return nil, false

}

func (this *WikiwordModel) GetAttrsByWord(w string) ([]Wikiwordattr, bool) {

	//cache check
	c, ok := this.attrsCache[w]
	if ok {
		return c, true
	}

	// Execute the query
	rows, err := SqlDb.Query("SELECT key, value FROM wikiwordattrs WHERE word = ?", w)
	if err != nil {
		log.Println(err)
		return nil, false
	}
	defer rows.Close()

	attr := Wikiwordattr{}
	attrs := []Wikiwordattr{}

	for rows.Next() {
		err = rows.Scan(&attr.Key, &attr.Value)
		if err != nil {
			log.Println(err)
			return nil, false
		}
		attrs = append(attrs, attr)
	}

	err = rows.Err()
	if err != nil {
		log.Println(err)
		return nil, false
	}

	// cache and return it
	if len(attrs) > 0 {
		this.attrsCache[w] = attrs
		return attrs, true
	}

	return nil, false

}

func (this *WikiwordModel) GetBeInvolvedByWord(w string) ([]string, bool) {

	//cache check
	c, ok := this.beinvCache[w]
	if ok {
		return c, true
	}

	// Execute the query
	rows, err := SqlDb.Query("SELECT word FROM wikirelations WHERE relation = ?", w)
	if err != nil {
		log.Println(err)
		return nil, false
	}
	defer rows.Close()

	beinvolves := []string{}
	var str string

	for rows.Next() {
		err = rows.Scan(&str)
		if err != nil {
			log.Println(err)
			return nil, false
		}
		beinvolves = append(beinvolves, str)
	}

	err = rows.Err()
	if err != nil {
		log.Println(err)
		return nil, false
	}

	// cache and return it
	if len(beinvolves) > 0 {
		this.beinvCache[w] = beinvolves
		return beinvolves, true
	}

	return nil, false

}

func (this *WikiwordModel) GetInvolvedByWord(w string) ([]string, bool) {

	//cache check
	c, ok := this.involCache[w]
	if ok {
		return c, true
	}

	// Execute the query
	rows, err := SqlDb.Query("SELECT relation FROM wikirelations WHERE word = ?", w)
	if err != nil {
		log.Println(err)
		return nil, false
	}
	defer rows.Close()

	involves := []string{}
	var str string

	for rows.Next() {
		err = rows.Scan(&str)
		if err != nil {
			log.Println(err)
			return nil, false
		}
		involves = append(involves, str)
	}

	err = rows.Err()
	if err != nil {
		log.Println(err)
		return nil, false
	}

	// cache and return it
	if len(involves) > 0 {
		this.involCache[w] = involves
		return involves, true
	}

	return nil, false

}
