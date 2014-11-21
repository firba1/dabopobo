package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type model interface {
	incr(string) error
	getInt(string) int
}

type cmd struct {
	regex   string
	handler commandHandler
}

type commandHandler func(m model, mutations [][]string, username string) (response []byte, err error)

func mutateKarma(m model, mutations [][]string, username string) (b []byte, err error) {
	if username == "slackbot" {
		return
	}
	for _, mutation := range mutations {
		identifier := mutation[1]
		op := mutation[2]
		if identifier != "" && identifier != username { //users may not mutate themselves
			suffix := canonicalizeSuffix(op)
			key := strings.ToLower(identifier) + suffix
			err = m.incr(key)
			if err != nil {
				err = nil
				return
			}
			fmt.Println(key)
		}
	}
	return
}

func getKarma(m model, identifier [][]string, username string) (response []byte, err error) {
	name := identifier[0][1]
	karmaset := getKarmaSet(m, name)
	text := fmt.Sprintf("%v's karma is %v %v", name, karmaset.value(), karmaset)
	res := map[string]string{
		"text":     text,
		"parse":    "full",
		"username": "dabopobo",
	}
	response, err = json.Marshal(res)
	fmt.Println(text)
	return
}

func getKarmaSet(m model, name string) (k karmaSet) {
	name = strings.ToLower(name)
	k.plusplus = m.getInt(name + "++")
	k.minusminus = m.getInt(name + "--")
	k.plusminus = m.getInt(name + "+-")
	return
}