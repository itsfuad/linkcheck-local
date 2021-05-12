package main
import("encoding/json";"fmt";"io/ioutil";"os")
func main(){
  if len(os.Args)<2 { fmt.Fprintln(os.Stderr,"Usage: app FILE");os.Exit(2) }
  data,err:=ioutil.ReadFile(os.Args[1]);if err!=nil{fmt.Fprintln(os.Stderr,err);os.Exit(2)}
  value:=Analyze(string(data)); output,err:=json.Marshal(value)
  if err!=nil{fmt.Fprintln(os.Stderr,err);os.Exit(2)}
  fmt.Println(string(output))
}
