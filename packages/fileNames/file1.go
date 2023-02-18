package fileNames

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var re = regexp.MustCompile(`\w+\d*\.txt`)
var sc = bufio.NewScanner(os.Stdin)

func getName()string{
	sc.Scan()
	return sc.Text()
}

func checkName(fName string)string{

	str := re.FindString(fName)
	if len(str) > 4{
		return str
	}
	return ""
}

func GetName(need string)string{
	name := ""
	for {
		fmt.Println(need+" (file_name[number].txt):")
		name = getName()
		name = checkName(name)
		if len(name) > 5 {     // 5 for char bec .txt
			break
		}
		fmt.Println("Typed name not valid")
	}
	return name
}

