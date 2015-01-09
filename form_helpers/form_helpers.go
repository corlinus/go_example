package form_helpers

import (
    "fmt"
    "strings"
)

type MyFieldOptions struct {
    Type string
    Name string
    Label string
    Ext string
    Required bool
    Default bool
}

type tag_attrs map[string]string

func tag(tag_name string, attributes tag_attrs) (tag_html string) {
    var attributes_html string

    for k, v := range attributes {
        attributes_html += " "
        attributes_html += k
        if v != "" {
            attributes_html += "=\""
            attributes_html += v
            attributes_html += "\""
        }
    }

    return fmt.Sprintf("<%s%s>", tag_name, attributes_html)
}

func contentTag(tag_name, content string, attributes tag_attrs) (tag_html string) {
    return fmt.Sprintf("%s%s</%s>", tag(tag_name, attributes), content, tag_name)
}

func simpleInputTag(input_type, name, value string) (field_html string) {
    return tag("input", tag_attrs{"type" : input_type, "name" : name, "value" : value})
}

func radioTag(name, ext string) (field_html string) {
    ext_fields := strings.Split(ext, ";")
    value := ext_fields[0]
    attrs := tag_attrs{"type" : "radio", "name" : name, "value" : value}

    if len(ext_fields) == 2 {
        attrs[ext_fields[1]] = ""
    }

    return tag("input", attrs)
}

func textareaTag(name, value string) (field_html string) {
    return tag("textarea", tag_attrs{"name" : name, "value" : value})
}

func selectTag(name, ext string) (field_html string) {
    options := ""
    for _, option_meta := range strings.Split(ext, ",") {

        array1 := strings.Split(option_meta, "=")
        array2 := strings.Split(array1[1], ";")
        attrs := tag_attrs{"value" : array2[0]}

        if len(array2) == 2 && array2[1] == "selected" {
            attrs["selected"] = ""
        }

        options += contentTag("option", array1[0], attrs)
    }
    return contentTag("select", options, tag_attrs{"name" : name})
}

func inputTag(input_type, name, value, ext string) (field_html string) {
    switch input_type {
    case "textarea":
        field_html = textareaTag(name, value)
    case "radio":
        field_html = radioTag(name, ext)
    case "checkbox":
        // TODO
        field_html = ""
    case "select":
        field_html = selectTag(name, ext)
    case "text", "password", "hidden", "button":
        field_html = simpleInputTag(input_type, name, value)
    default:
        field_html = simpleInputTag("text", name, value)
    }

    return  field_html
}

func labelTag(for_attr, label string) (tag_html string) {
    return contentTag("label", label, tag_attrs{"for" : for_attr})
}

func brTag() (tag_html string) {
    return "<br>"
}

func FieldCreate(value string, options MyFieldOptions) (field_html string) {

    input := inputTag(options.Type, options.Name, value, options.Ext)
    label := labelTag(options.Name, options.Label)
    br := brTag()

    return input + label + br
}
