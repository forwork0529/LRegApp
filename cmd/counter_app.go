package main

import (
	"LRegApp/packages/counter"
	"LRegApp/packages/fileNames"
	"LRegApp/packages/pars"
	"bufio"
	"fmt"
	"os"
	// "fmt"
	"io/ioutil"
	"log"
)


func main(){

	var err error
	var data []byte
	var outputFile *os.File
	var logFile *os.File

	logFile, err = os.Create("files/logs.txt")
	logs := log.New(logFile, "INFO\t",log.Ldate|log.Ltime)

	input := fileNames.GetName("Type name of the INPUT file") //  example = input.txt
	logs.Printf("got input file name=%v\n", input)
	data , err = ioutil.ReadFile("files/"+input)
	if err != nil{
		fmt.Println("INPUT file must store in ./files ..")
		log.Fatal(err)
	}
	logs.Println("got data from input file")
	output := fileNames.GetName("Type name of the output file")
	outputFile,  err  = os.Create("files/"+output)
	if err != nil{
		log.Fatal(err)
	}
	logs.Println("got output file")

	wr := bufio.NewWriter(outputFile)

	variables := pars.Act(data) //  Returns Expression, arg1, action, arg2
	count := 0                  // count successful tries for logger
	for _, exp := range variables{
		if len(exp) < 4{    // Expression, arg1, action, arg2 - that s why compare with 4
			continue
		}
		result := counter.CountStrings(exp[1], exp[2], exp[3])
		if result == ""{
			continue
		}
		count += 1
		_, err = wr.Write([]byte(exp[0]+result+"\n"))
		if err != nil{
			log.Fatal(err)
		}
	}
	logs.Printf("processed %v expressions\n", count)

	err = wr.Flush()
	if err != nil{
		log.Fatal(err)
	}

}

