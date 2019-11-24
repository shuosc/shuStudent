package fetchName

import (
	"bytes"
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"shuStudent/config"
	"shuStudent/model"
	"strings"
)

func FetchName(student model.Student, token string) (model.Student, error) {
	form := struct {
		Url string `json:"url"`
	}{
		config.GET_NAME_URL,
	}
	body, _ := json.Marshal(form)
	req, _ := http.NewRequest("POST", config.COURSE_PROXY_ENDPOINT+"get", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Cannot fetch student name from " + config.GET_NAME_URL + " via proxy " + config.COURSE_PROXY_ENDPOINT)
		return model.Student{}, err
	}
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Println("Cannot fetch student name from " + config.GET_NAME_URL + " via proxy " + config.COURSE_PROXY_ENDPOINT)
		return model.Student{}, err
	}
	selection := doc.Find("#leftMenu>#leftmenu_Accordion .ui-accordion-content div:nth-of-type(2)")
	if !strings.Contains(selection.Text(), "姓名：") {
		log.Println("Cannot find nameinfo in", selection.Text())
		html, _ := doc.Html()
		log.Println("Doc is:\n", html)
	}
	student.Name = strings.Trim(selection.Text(), " \n")[len("姓名："):]
	return student, nil
}
