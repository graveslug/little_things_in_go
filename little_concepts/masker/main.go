package main

import (
	"fmt"
	"os"
)

func main() {
	//Waiting for terminal input
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Bruh, wheres the input?")
		return
	}
	const (
		//future note. Should I bother with https? Hm.
		link  = "http://"
		nlink = len(link)
		mask  = '*'
	)

	var (
		text = args[0]
		size = len(text)
		// buff = make([]byte, 0, size)
		buff = []byte(text)

		in bool
	)

	//Often we use a for range loop but for this
	//we want to loop over the bytes of the text
	//not over its runes.
	for i := 0; i < size; i++ {
		// Checks the length of the following text
		// if theres more room for a
		// link thus ensures it never slices beyond its capacity.
		// && compares the input to see if it matches
		//the link pattern
		if len(text[i:]) >= nlink && text[i:i+nlink] == link {
			in = true
			// fmt.Printf(`text[%d : %[1]d+%d] = `, i, nlink)
			// fmt.Println(text[i : i+nlink])

			// buff = append(buff, link...)
			//so the loop jumps after the link pattern
			//preventing duplication from occuring ex: http://http://
			i += nlink
		}
		c := text[i]

		switch c {
		case ' ', '\t', '\n':
			in = false
		}
		if in {
			buff[i] = mask
			// c = mask
		}
		// buff = append(buff, c)
	}

	fmt.Println(string(buff))

}
