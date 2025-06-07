package Utils 

import ( "fmt" 
		"strings")


func Strings() {
	a := "abcd efgh ihjkl aaaaaaa aaaa aaa"
	parts := strings.Split(a," ")
	fmt.Println(parts)

	fmt.Println(strings.Count(a,"a"))

	b := "         sdfaf asasadsrgfdsagdfgsdfgsdssfss   sf  sfsf sdfsd  sfs sfsf      asda"
	fmt.Println(b,"\n",strings.TrimSpace(b))

	c := strings.Join([]string{"ads","fdsa","dasf","fasa","fasfa"},".")
	fmt.Println(c)
}