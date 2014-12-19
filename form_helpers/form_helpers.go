package form_helpers

import (
    "fmt"
)

type MyFieldOptions struct {
    Type string
    Name string
    Label string
    Required bool
    Default bool
}

func simpleInputTag(input_type, name, value string) (field_html string) {
    // TODO remove empty attributes
    return fmt.Sprintf(`<input type="%s" name="%s" value="%s"></input>`, input_type, name, value)
}

func textareaTag(name, value string) (field_html string) {
    return fmt.Sprintf(`<textarea name="%s" value="%s"></textarea>`, name, value)
}

func inputTag(input_type, name, value string) (field_html string) {
    switch input_type {
    case "textarea":
        field_html = textareaTag(name, value)
    case "raido":
        // TODO
        field_html = ""
    case "checkbox":
        // TODO
        field_html = ""
    case "select":
        // TODO
        field_html = ""
    case "text", "password", "hidden", "button":
        field_html = simpleInputTag(input_type, name, value)
    default:
        field_html = simpleInputTag("text", name, value)
    }

    return  field_html
}

func labelTag(for_attr, label string) (html string) {
    return fmt.Sprintf(`<label for="%s">%s</label>`, for_attr, label)
}

func brTag() (html string) {
    return "<br>"
}

func FieldCreate(value string, options MyFieldOptions) (field_html string) {

    input := inputTag(options.Type, options.Name, value)
    label := labelTag(options.Name, options.Label)
    br := brTag()

    return input + label + br
}
