package main

//ctrl+c to exit code run
import (
	"net/http"

	"github.com/thomasfan100/Go_Code_Pluralsight/controllers"
)

func main() {
	//new user example
	/*
		u := models.User{
			ID:        2,
			FirstName: "Tricia",
			LastName:  "McMillan",
		}
		fmt.Println(u)
	*/
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}

//---------------------Examples------------------------
//Constants
/*
const (
	first = 1
	second = "second"
	third = iota + 6
	fourth = 2 << iota
	fifth
)
*/
//Variables
/*
	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	firstName := "Arthur"
	fmt.Println(firstName)

	b:=true
	fmt.Println(b)

	c:=complex(3,4)
	fmt.Println(c)

	r,im := real(c),imag(c)
	fmt.Println(r,im)
*/
//Pointers
/*
	var firstName *string = new(string)
	*firstName = "Arthur"
	fmt.Println(*firstName)

	firstName2 := "Arthur"
	fmt.Println(firstName)
	ptr := &firstName
	fmt.Println(ptr,*ptr)

	firstName = "Tricia"
	fmt.Println(ptr,*ptr)
*/
/*
	const pi = 3.1415
	fmt.Println(pi)

	const c int= 3
	fmt.Println(c + 3)
	fmt.Println(float32(c) + 1.2)
*/
//Arrays
/*
	var arr[3]int //variable called arr that is a 3 element array of integers
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)
	fmt.Println(arr[1])

	arr2 := [3]int{1,2,3}
	fmt.Println(arr2)
*/
//Slice
/*
	arr := [3]int{1,2,3}

	slice := arr[:]

	fmt.Println(arr,slice)

	arr[1] = 42
	slice[2] = 27

	fmt.Println(arr,slice)

	slice2 := []int{1,2,3}
	fmt.Println(slice2)

	slice2 = append(slice2 , 4, 42, 27)
	fmt.Println(slice2)

	s2 := slice2[1:]
	s3 := slice2[:2]
	s4 := slice2[1:2]
	fmt.Println(s2,s3,s4)
*/
//Maps
/*
	m := map[string]int{"foo":42}
	fmt.Println(m)
	fmt.Println(m["foo"])
	m["foo"] = 27
	fmt.Println(m)
	delete(m, "foo")
	fmt.Println(m)
*/
//Structs
/*
	type user struct{
		ID int
		FirstName string
		LastName string
	}
	var u user
	u.ID = 1
	u.FirstName = "Arthur"
	u.LastName = "Dent"
	fmt.Println(u)

	u2 := user{
		    ID:1,
		    FirstName: "Arthur",
		    LastName: "Dent",
		  }
	fmt.Println(u2)
*/
//Functions
/*
	func main() {
		port := 3000
		//underscore is a write only variable that ignores the first result
		_, err := startWebServer(port,2)
		fmt.Println( err)

	}
	//import "errors"
	//go accepts comma-delimited list before a data type as all the same data type
	func startWebServer(port, numberOfRetries int) (int,error) {
		fmt.Println("Starting server...")
		//do important things
		fmt.Println("Server started on port", port)
		fmt.Println("Number of retries", numberOfRetries)
		return port,nil
		//return errors.New("Something went wrong")
	}
*/
//loop til condition
/*

	var i int
	for i < 5 {
		println(i)
		i++
		if i == 3 {
			break
		}
		continue

*/
//loop til condition with a post clause
/*
	for i := 0;i < 5 ; i++{
		println(i)
	}
*/
//infinite loop
/*
var i int
	//or remove the two semicolons
	for ; ;{
		if i == 5 {
		 	break
		}
		println(i)
		i++
	}
*/
//looping over iterables
/*
	slice := []int{1,2,3}
	for i := 0; i < len(slice); i++ {
		println(slice[i])
	}
	for i,v := range slice {
		println(i,v)
	}
	wellKnownPorts:= map[string]int{"http":80,"https":443}
	for k,o := range wellKnownPorts{
		println(k,o)
	}
	wellKnownPorts2:= map[string]int{"http":80,"https":443}
	for _,o := range wellKnownPorts2{
		println(o)
	}
*/
//panic functions
//panic("Something bad just happened")
//if and else
/*
	type User struct{
	ID int
	FirstName string
	LastName string
	}
	u1 := User{
		ID:	1,
		FirstName: "Arthur",
		LastName:  "Dent",
	}
	u2 := User{
		ID:	2,
		FirstName:  "Ford",
		LastName:  "Prefect",
	}
	if u1.ID == u2.ID{
		println("Same user!")
	} else if u1.FirstName == u2.FirstName {
		println("Similar user.")
	} else {
		println("Differnet users!")
	}
*/
//Switch Statements
/*
type HTTPRequest struct{
		Method string
	}
	r := HTTPRequest {Method: "HEAD"}

	//go has implicit breaks, use fallthrough to go under.
	switch r.Method{
	case "GET":
		println("Get request")
	case "DELETE":
		println("Delete request")
	case "POST":
		println("Post request")
	case "PUT":
		println("Put request")
	default:
		println("Unandled method")
	}
*/
