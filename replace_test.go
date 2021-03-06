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
	"fmt"
	"testing"
)

// Testing htmlrepl
func TestReplace(t *testing.T) {

	// Replace test html string
	text, err := Replace("<html><head><title></title></head><body><h1>Hello, World!</h1><p>You can modify the text in the box to the left any way you like, and then click the Show Page button below the box to display the result here. Go ahead and do this as often and as long as you like.</p> <p>You can also use this page to test your Javascript functions and local style declarations. Everything you do will be handled entirely by your own browser; nothing you type into the text box will be sent back to the server.</p> <p>When you are satisfied with your page, you can select all text in the textarea and copy it to a new text file, with an extension of either <b>.htm</b> or <b>.html</b>, depending on your Operating System.This file can then be moved to your Web server.</p></body></html>")
	if err != nil {
		fmt.Println(err)
	}

	// Print out to terminal
	fmt.Println(text)

}
