package Utils

import "fmt"

type Employee struct {
	contact_details Contact
	name string 
	address string 
}

type Contact struct {
	phone int
	email string
}

func Structure() {
	var e1 Employee;
	e1.contact_details = Contact{8374826890,"adsa@ajfshs.com"}

	e2 := Employee{name : "dsfs",address :"sfgdrsgdf",contact_details: Contact{9678746890,"adasfd@sfef"}}
	fmt.Println(e1)
	fmt.Println(e2)

	var e3 Employee;
	e3.name = "adfasfa"
	e3.address = "fasfas"
	e3.contact_details = Contact{9784060386,"aeewr@wedfa"}

	fmt.Println(e3);

	e3.contact_details.email = "iohjdjfk@skjfjas"
	e3.contact_details.phone = 879805867980
}