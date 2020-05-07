package main
import (
	"fmt"
	"io/ioutil"
	"strconv"
)

type student struct {
	name string
	age int
	rollno int
}

func newStudent(name string, age int, rollno int) *student {

	st := &student{name:name,age:age,rollno: rollno}
	return st;
}
func updatedetails(s *student, age int){
	s.age = age;
}
func test_s_age(s *student) string{
	if (s.age<17){
		return "Not eligible"
	}else{
		return "Eligible"
	}

}

func (s student) String() string{
	return "Name : "+s.name+" Age : "+strconv.Itoa(s.age)+" Roll no:"+strconv.Itoa(s.rollno)
}
func main() {
	student1 := newStudent("rahul",18,2)
	fmt.Println(student1)
	updatedetails(student1,20)
    fmt.Println(student1)
	test_result := test_s_age(student1)
	fmt.Println(test_result)
	fmt.Println(student1.String())
	err := ioutil.WriteFile("sample",[]byte(student1.String()),0644)
	if(err !=nil){
		fmt.Println("error occured")
	}

}

