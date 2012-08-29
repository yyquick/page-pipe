package dispatcher

import (
	"github.com/yyquick/page-pipe/parser"
	"net/http"
	"io/ioutil"
)

func dispatch(pagelets parser.Pagelets, results chan string) {
	prev_p := 0;
	c := 0;
	i := 0
	r := make(chan string)
	for _, p := range pagelets {
		if p.Priority == prev_p {
			c++
			go getResult(p, r)
		} else {
			for i = 0; i < c; i++ {
				s := <- r
				results <- s
				c = 0
			}
			prev_p = p.Priority
			c++
			go getResult(p, r)
		}
	}
	for i = 0; i < c; i++ {
		s := <- r
		results <- s
	}
	close(results)
}

func getResult(p *parser.Pagelet, r chan string) {
	resp, err := http.Get(p.Href)
	if err == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			r <- string(body)
		} else {
			r <- "error"
		}
	} else {
		r <- "error"
	}
}

func wrapJsForContent(p *parser.Pagelet, content string) string {
	return content
}

func wrapJsForError(p *parser.Pagelet, error string) string {
	return error
}