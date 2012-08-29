package parser

import (
	"regexp"
	"strings"
	"strconv"
	"sort"
)

type Pagelet struct {
	Id string
	Priority int
	Href string
}

type Pagelets []*Pagelet

func (p Pagelets) Len() int { return len(p) }
func (p Pagelets) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p Pagelets) Less(i, j int) bool { return p[i].Priority < p[j].Priority }
func (p Pagelets) Sort() {sort.Sort(p)}

func GetPageletData(div string) (id string, priority int, href string) {
	blank_reg, _ := regexp.Compile(" +")
	div = blank_reg.ReplaceAllString(div, " ")
	parts := strings.Split(div, " ")
	for _, attr := range parts {
		vparts := strings.Split(attr, "=")
		if len(vparts) == 2 {
			value := strings.Trim(vparts[1], "\"")
			if vparts[0] == "id" {
				id = value
			}
			if vparts[0] == "priority" {
				p, _ := strconv.ParseInt(value, 10, 0)
				priority = int(p)
			}
			if vparts[0] == "href" {
				href = value
			}
		}
	}
	return
}

func NewPagelet(div string) *Pagelet {
	id, priority, href := GetPageletData(div)
	return &Pagelet{id, priority, href}
}

func FindPagelets(html string) Pagelets {
	pl_reg, _ := regexp.Compile("<div.*pagelet.*>")
	results := pl_reg.FindAllString(html, -1)
	pagelets := make([]*Pagelet, len(results))
	for i, value := range results {
		pagelets[i] = NewPagelet(value)
	}
	result := Pagelets(pagelets)
	result.Sort()
	return result
}
