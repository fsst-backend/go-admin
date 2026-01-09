package test

import (
	//"go-admin/models/tools"
	//"os"

	"fmt"
	"testing"

	"golang.org/x/crypto/bcrypt"
	//"text/template"
)

func TestGoModelTemplate(t *testing.T) {
	hash, err := bcrypt.GenerateFromPassword([]byte("adminqwer123"), bcrypt.DefaultCost)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("\n%s\n", string(hash))
	t.Log(string(hash))
}

func TestGoApiTemplate(t *testing.T) {
	//t1, err := template.ParseFiles("api.go.template")
	//if err != nil {
	//	t.Error(err)
	//}
	//table := tools.SysTables{}
	//table.TBName = "sys_tables"
	//tab, _ := table.Get()
	//file, err := os.Create("apis/" + table.PackageName + ".go")
	//if err != nil {
	//	t.Error(err)
	//}
	//defer file.Close()
	//
	//_ = t1.Execute(file, tab)
	t.Log("")
}
