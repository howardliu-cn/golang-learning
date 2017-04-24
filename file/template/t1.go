package main

import (
	"fmt"
	"html/template"
	"os"
)

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func main() {
	t := template.New("fieldname example")
	t, _ = t.Parse("Hello, {{.UserName}}!")
	p := Person{UserName: "Howard Liu"}
	t.Execute(os.Stdout, p)

	fmt.Println()

	f1 := Friend{Fname: "u1"}
	f2 := Friend{Fname: "u2"}
	t = template.New("fieldname example")
	t, _ = t.Parse(
		`Hello, {{.UserName}}!
		{{range .Emails}}
		an email {{.}}
		{{end}}
		{{with .Friends}}
		{{range .}}
			my friend name is {{.Fname}}
		{{end}}
		{{end}}`)
	p = Person{UserName: "Howard Liu", Emails: []string{"howardliu1988@163.com"}, Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)

	fmt.Println()

	t = template.New("template test")
	t = template.Must(t.Parse("空 pipeline if demo: {{if ``}} 不会输出. {{end}}\n"))
	t.Execute(os.Stdout, nil)

	t = template.New("template test")
	t = template.Must(t.Parse("不为空的 pipeline if demo: {{if `anything`}} 我有内容，我会输出. {{end}}\n"))
	t.Execute(os.Stdout, nil)

	t=template.New("template test")
	t=template.Must(t.Parse("if-else demo: {{if `anything`}} if部分 {{else}} else部分.{{end}}\n"))
	t.Execute(os.Stdout, nil)
}
