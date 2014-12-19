package main

import (
    "fmt"
    "strings"
    "net/http"
    "log"
    "io"
    "reflect"

    "./form_helpers"
)

// MyForm

type MyForm struct {
    UserName string `required:"true" field:"name" name:"Имя пользователя" type:"text"`
    UserPassword string `required:"true" field:"password" name:"Пароль пользователя" type:"password"`
    Resident bool `field:"resident" type:"radio" radio:"1;checked" name:"Резидент РФ"`
    NoResident bool `field:"resident" type:"radio" radio:"2" name:"Не резидент РФ"`
    Gender string `field:"gender" name:"Пол" type:"select" select:"Не известный=3;selected,Мужской=1,Женский=2"`
    Age int64 `field:"age" name:"Возраст" type:"text" default:"true"`
    Token string `field:"token" type:"hidden" default:"true"`
}

func htmlLayout(str string) (html_str string) {
    return fmt.Sprintf("<html><body>\n<form action='/create' method='POST'>\n%s\n<input type='submit' value='Submit'>\n</form>\n</body></html>", str)
}

func extractOptions(field reflect.StructField) (options form_helpers.MyFieldOptions) {
     tag := field.Tag
     options = form_helpers.MyFieldOptions{
         Type:     tag.Get("type"),
         Name:     tag.Get("field"),
         Label:    tag.Get("name"),
         Required: tag.Get("required") == "true",
         Default:  tag.Get("default") == "true",
     }

    return options
}

func FormCreate(formData *MyForm) (form string, err error) {

     fields := make([]string, 0)
     fdValue := reflect.ValueOf(formData).Elem()
     fdType := reflect.TypeOf(*formData)

     for i := 0; i < fdValue.NumField(); i++ {
         field := fdValue.Field(i)
         value := fmt.Sprintf("%s", field.Interface()) // FIXME
         options := extractOptions(fdType.Field(i))
         fields = append(fields, form_helpers.FieldCreate(value, options))
     }

     form = strings.Join(fields, "\n")
     return form, err
}

func FormRead(formData *MyForm, request *http.Request) (err error) {
     return err
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/form", 302)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
    var fd *MyForm = &MyForm{
        Age: 18,
        Token: "345625145123451234123412342345",
    }

    form, err := FormCreate(fd)

    if err != nil {
        log.Println(err)
        io.WriteString(w, "can not create form")
    } else {
        io.WriteString(w, htmlLayout(form))
    }
}

func createHandler(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "hello!!!")
}


func main() {

    http.HandleFunc("/", welcomeHandler)
    http.HandleFunc("/form", formHandler)
    http.HandleFunc("/create", createHandler)

    log.Println("Server listening on http://0.0.0.0:4000")
    err := http.ListenAndServe(":4000", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
