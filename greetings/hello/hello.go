package main
import (
	"fmt"
	"log"
	"example.com/greetings"
)

func main (){
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//Request a greeting message.
	message , err := greetings.Hello("tetet")
	if(err !=nil){
		log.Fatal(err)
	}
	fmt.Println(message)
}