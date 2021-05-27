package main
import("testing";"io/ioutil")
func TestAnalyzeExample(t *testing.T){
  data,err:=ioutil.ReadFile("examples/input.txt");if err!=nil{t.Fatal(err)}
  output:=Analyze(string(data));if len(output)==0{t.Fatal("no report fields")}
}
