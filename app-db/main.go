package main

import (
	"fmt"
	"garination.com/db/cmd"
)

const AsciiArt = `
                             _ _     
  __ _ _ __  _ __         __| | |__  
 / _ | '_ \| '_ \ _____ / _ | '_ \ 
| (_| | |_) | |_) |_____| (_| | |_) |
 \__,_| .__/| .__/       \__,_|_.__/ 
      |_|   |_|                      
`

func main() {
	fmt.Print(AsciiArt)
	cmd.Execute()
}
