package main
import "fmt"

type Header map[string]string

//相当于php中有键值的数组
func main() {
	// var countrymap map[string]string
	// countrymap = make(map[string]string)
	// countrymap := map[string]string{"france":"paris","china":"beijing"}
	// countrymap["france"] = "paris"
	// countrymap["china"] = "beijing"
	// for country := range countrymap{
	// 	fmt.Println(countrymap[country],country)
	// }
	// delete(countrymap,"france")
	// cap,ok :=countrymap["france"]
	// fmt.Println(ok,cap)

	var test Header
	// var test map[string]string
	test = make(Header)
	test["france"] = "paris"
	fmt.Println(test)
}