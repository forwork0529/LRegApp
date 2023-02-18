package pars

import "regexp"

var re = regexp.MustCompile(`(\d+)([+-])(\d+)=`)

func Act(data []byte)[][]string{
	return re.FindAllStringSubmatch(string(data), -1)
}


