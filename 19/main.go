package main

import "fmt"

var justString string

func someFunc() {
	v := "esrghkjl;dfsafdgfgjhkjuiptrsduoigfgxdcghjkl;jhgfxghjklk;hgfdzxcghvjjpjfdfzfkjhgjkjcfionidfnhguerkljfcop'diewigheriflk;ehrgvkmerjgcierl;mciero;ureigugmerilc;emgl;ejgerlk;gvmecklg;mgerkl;gcjerigber;imgerufieronmgueirl;gcmeio;erucnieo;gmeilv;gmic;ermugcioerguruighjbcuerigvuergiorugeigmergucbiegnyueriesruugcngngurijlgneruicgerubgvkieukmgvireogmueriocgprmyucgigureibgcnyerulgvyebrhulngrngvirbgrsuipgbnyugisgnurijglcbeyrnguierycbeinrutyer reui euig yerk leyhu gelrh gjer,hugl hearugiaehg e ghruhguigeru hgearu ghru iglrehtuie; uei gerhugherg iurhguerig hreui gehrgilerh uerig heriughsergh egerh gejrgherui gerhg shguerg heig rehs"
	justString = v[:100]
}

func main() {
	someFunc()
	fmt.Println(justString)
}
