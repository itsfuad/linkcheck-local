package main
import("encoding/json";"fmt";"io/ioutil";"os";"strings")
func main(){
  if len(os.Args)<2 { fmt.Fprintln(os.Stderr,"Usage: app FILE");os.Exit(2) }
  var data []byte;var err error
  if os.Args[1]=="-" { data,err=ioutil.ReadAll(os.Stdin) } else { data,err=ioutil.ReadFile(os.Args[1]) };if err!=nil{fmt.Fprintln(os.Stderr,err);os.Exit(2)}
  value:=Analyze(string(data)); output,err:=json.Marshal(value)
  if err!=nil{fmt.Fprintln(os.Stderr,err);os.Exit(2)}
  if len(os.Args)>2 && os.Args[2]=="--pretty" {
    var pretty []byte
    pretty,err=json.MarshalIndent(value,"","  ");if err==nil{output=pretty}
  }
  fmt.Println(strings.TrimSpace(string(output)))
}
