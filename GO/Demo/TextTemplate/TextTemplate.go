package main

import (
	"fmt"
	"os"
	"text/template"
)

type secret struct{
	Name string
	Secret string
}

func main(){

	secret1 := secret{"Soushi", "I'm a traveller from beyond this world"}

	secretPath := "D:/C_Softwares/GO/Demo/TextTemplate/secret.txt"

	sec, err := template.New("secret.txt").ParseFiles(secretPath)

	if err != nil {
		fmt.Println(err)
	}
	err = sec.Execute(os.Stdout, secret1)
	if err != nil {
		fmt.Println(err)
	}





}
