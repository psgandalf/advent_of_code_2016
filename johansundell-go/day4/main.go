package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/johansundell/advent_of_code_2016/johansundell-go/adventofcode2016"
)

type room struct {
	id       int
	checksum string
	name     string
}

type char struct {
	c     string
	count int
}

type charList []char

func (c charList) Len() int { return len(c) }
func (c charList) Less(i, j int) bool {
	if c[i].count == c[j].count {
		return c[i].c < c[j].c
	}
	return c[i].count > c[j].count
}
func (c charList) Swap(i, j int) { c[i], c[j] = c[j], c[i] }

func main() {
	data, err := adventofcode2016.GetInput("day4.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputs := strings.Split(data, "\n")
	sum, id := checkRooms(inputs)
	fmt.Println("Sum: ", sum, " id: ", id)
}

func checkRooms(inputs []string) (sum, id int) {
	for _, input := range inputs {
		r := newRoom(input)
		if r.isValid() {
			sum += r.id
			name := r.decryptName()
			if name == "northpole object storage" {
				id = r.id
			}
		}
	}
	return
}

func newRoom(str string) room {
	r := room{}
	r.checksum = str[strings.Index(str, "[")+1 : len(str)-1]
	r.id, _ = strconv.Atoi(str[strings.LastIndex(str, "-")+1 : strings.Index(str, "[")])
	str = str[:strings.LastIndex(str, "-")]
	r.name = str
	return r
}

func (r *room) isValid() bool {
	chars := make(map[string]char)
	str := strings.Replace(r.name, "-", "", -1)
	for i := 0; i < len(str); i++ {
		s := string(str[i])
		c, _ := chars[s]
		c.c = s
		c.count++
		chars[s] = c
	}
	letters := make(charList, len(chars))
	i := 0
	for _, c := range chars {
		letters[i] = c
		i++
	}
	sort.Sort(charList(letters))

	for i := 0; i < len(r.checksum); i++ {
		a := letters[i]
		c := string(r.checksum[i])
		if c == a.c {
			continue
		}
		return false
	}
	return true
}

func (r *room) decryptName() (name string) {
	ch := r.id % 26
	for _, s := range r.name {
		switch {
		case s == 45:
			name += " "
		case int(s)+ch < 123:
			name += string(int(s) + ch)
		default:
			name += string(int(s) + ch - 26)
		}
	}
	return
}
