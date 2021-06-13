package main
import("testing")
func TestEmptyInput(t *testing.T){
    got:=Analyze("")
    if got==nil{t.Fatal("nil report")}
}
