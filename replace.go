//**********************************************************
//
// Copyright (C) 2018 - 2021 J&J Ideenschmiede UG (haftungsbeschränkt) <info@jj-ideenschmiede.de>
//
// This file is part of htmlrepl.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor
//
//**********************************************************

package htmlrepl

import (
	"encoding/json"
	"net/http"
	"strings"
)

// Replace all HTML Tags
func Replace(text string) (string, error) {

	// Get json file
	tags, err := http.Get("https://raw.githubusercontent.com/jojojojonas/htmlrepl/master/files/json/html-tags.json")
	if err != nil {
		return "", err
	}

	// Close file
	defer tags.Body.Close()

	// Decode json file
	var decode []string

	err = json.NewDecoder(tags.Body).Decode(&decode)
	if err != nil {
		return "", err
	}

	// Filter text string
	for _, value := range decode {

		// Replace string
		text = strings.Replace(text, value, "", -1)

		// split tag
		split := strings.Split(value, "<")

		// Creating closing tag
		tag := "</" + split[1]

		// Replace closing tag
		text = strings.Replace(text, tag, "", -1)

	}

	// Return string
	return text, nil

}
