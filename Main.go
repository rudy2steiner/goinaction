package main

import (
	"log"
	"os"

	_ "goaction/mathers"
	"goaction/search"
	)

func init(){
	log.SetOutput(os.Stdout)
}
func main(){
	search.Run("president")
}
