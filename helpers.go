package main

import (
    "fmt"
    "strings"
    "strconv"
    "net/http"
    "reflect"

    "./form_helpers"
)

func htmlLayout(str string) (html_str string) {
    return fmt.Sprintf("<html><body>\n<form action='/create' method='post' enctype='multipart/form-data'>\n%s\n<input type='submit' value='Submit'>\n</form>\n</body></html>", str)
}

func extractOptions(field reflect.StructField) (options form_helpers.MyFieldOptions) {
    tag := field.Tag
    options = form_helpers.MyFieldOptions{
        Type:     tag.Get("type"),
        Name:     tag.Get("field"),
        Label:    tag.Get("name"),
        Ext:      tag.Get(tag.Get("type")),
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
        value := fmt.Sprintf("%v", field.Interface())
        options := extractOptions(fdType.Field(i))
        fields = append(fields, form_helpers.FieldCreate(value, options))
    }

    form = strings.Join(fields, "\n")
    return form, err
}

func setField(field *reflect.Value, fieldType *reflect.StructField, value string) (err error) {

    switch fieldType.Type.Kind() {

    case reflect.Bool:
        switch fieldType.Tag.Get("type") {
        case "radio":
            tag_value := strings.Split(fieldType.Tag.Get("radio"), ";")[0]
            field.SetBool(tag_value == value)

        case "checkbox":
            val, err := strconv.ParseBool(value)
            if err != nil {
                field.SetBool(false)
            } else {
                field.SetBool(val)
            }
        }

    //case reflect.time.Duration:

    case reflect.Float64:
        val, err := strconv.ParseFloat(value, 64)
        if err != nil {
            return err
        } else {
            field.SetFloat(val)
        }

    case reflect.Int:
        val, err := strconv.ParseInt(value, 10, 32)
        if err != nil {
            return err
        } else {
            field.SetInt(val)
        }

    case reflect.Int64:
        val, err := strconv.ParseInt(value, 10, 64)
        if err != nil {
            return err
        } else {
            field.SetInt(val)
        }

    case reflect.Uint:
        val, err := strconv.ParseInt(value, 10, 32)
        if err != nil {
            return err
        } else {
            field.SetInt(val)
        }

    case reflect.Uint64:
        val, err := strconv.ParseUint(value, 10, 64)
        if err != nil {
            return err
        } else {
            field.SetUint(val)
        }

    case reflect.String:
        field.SetString(value)
    }

    return nil
}

func FormRead(formData *MyForm, r *http.Request) (err error) {
    _ = r.ParseMultipartForm(10240)

    fdValue := reflect.ValueOf(formData).Elem()

    for i := 0; i < fdValue.NumField(); i++ {
        fieldType := fdValue.Type().Field(i)
        name  := fieldType.Tag.Get("field")
        value := r.Form.Get(name)
        field := fdValue.Field(i)

        err = setField(&field, &fieldType, value)
        if err != nil {
            fmt.Println(err)
            return err
        }
    }

    fmt.Println(formData)
    return err
}
