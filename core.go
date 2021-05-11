package main
import (
    "regexp"
    "net/url"
)
// Analyze summarizes a stream of input without modifying user files.
func Analyze(input string) map[string]interface{} {

link:=regexp.MustCompile(`\[[^\]]*\]\(([^)]+)\)`)
paths:=[]string{};external:=0;invalid:=0
for _,match:=range link.FindAllStringSubmatch(input,-1){dest:=match[1]
u,err:=url.Parse(dest);if err!=nil{invalid++;continue}
if u.Scheme!=""{external++;continue};if u.Path!=""{paths=append(paths,u.Path)}}
return map[string]interface{}{"localTargets":paths,"external":external,"invalid":invalid}
}
