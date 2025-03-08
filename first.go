package main
import "fmt"
func main(){
	//Variable dicleration 
	//var a int = 10      //Technique 1
	a:= "hello"           //Technique 2(Shortcut)
	// a= 10              //this is not happen because of previously we take string value in variable a
	a = "Matin"            // take value in a variable for 2nd time,no need colon(:)

	//Constant Value
	const p = 1344
	//const p=12            //COnstant is immutable







	fmt.Println(a)
	fmt.Println(p)
}