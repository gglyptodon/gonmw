package main

import (
	"fmt"
	//"greeter"
	//"bufio"
	//"os"
	"strings"
	"io/ioutil"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}
func rmEmpty(arr []string) []string{
	var res []string

	for _,v :=  range arr{
		if v != " " {
			res = append(res,v)
			fmt.Println(v, res, v!= " ")
		}
	}
	return res
}

func readSubstitutionMatrix(infile string)map[string]int{
	var m map[string]int
	m = make(map[string]int)
	dat, err := ioutil.ReadFile(infile)
    check(err)
    //fmt.Print(string(dat))
	var tmp []string = strings.Split(string(dat),"\n")
	var count int = 0
	var kh []string
	var kv []string
	var scores [][]string
	for _,v := range tmp{
		if ! strings.HasPrefix(v,"#"){
			count +=1
			if count == 1{
				//kh = strings.Split(v," ")
				kh = strings.Fields(v)
				//fmt.Println(kh)

			}else{
				tmp := strings.Fields(v)
				scores = append(scores,tmp )
				//fmt.Println(kv)
				//fmt.Println(" ")
				kv = append(kv,strings.Split(v," ")[0] )

		}
			//fmt.Println(v)
		}

	}
	for _,k := range kh{
		fmt.Print(k)
	}
	kh = rmEmpty(kh)
	kv = rmEmpty(kv)
	var matrixKeys [] string
	for _,k := range kh{
		//fmt.Print(k)
		for _,l := range kv{
			//fmt.Println(l)
			matrixKeys = append(matrixKeys, k+":"+l)
		}
	}

	for k,v := range kh{
		for j,_ := range kv {
			fmt.Println(scores[k],strings.Fields(v)[j])
		}
	}

	//for k,v := range matrixKeys {
	//	fmt.Println(v,k)
	//
	//}

	return m
}
func main() {
	readSubstitutionMatrix("/home/nin/git/gonmw/src/data/EBLOSUM62")
	//fmt.Printf("Hello world!")
	//fmt.Println(greeter.Greeting())
	var sequences []string
	var seqId map[string]string
	seqId = make(map[string]string)
	sequences = append(sequences, ">1\nMRISDQVYSLYRFIRMSFQAPPTLL", ">2\nMRISDQVYSLY")
	var mismatch int = -10
	var match int = 10


	for _, val := range sequences {
		//fmt.Println(val)
		a := strings.Split(val, "\n")
		seqId[a[0]] = a[1]
	}
	var m int = len(sequences[0])
	var n int = len(sequences[1])
	//fmt.Println(m, n)

	var mat = make([][]int, m)
	for i := 0; i < m; i++ {
		mat[i] = make([]int, n)
	}

	for k, v := range sequences[0] {
		for l, w := range sequences[1] {
			if v == w {
				mat[k][l] = match
			}else {
				mat[k][l] = mismatch
			}
		}
	}

	//for _, k := range mat {
		//for _, i := range k {
	//		fmt.Println("")
	//	}
	//}



}


	/*for k, v := range seqId {
		fmt.Println(k, v)
		var mat [len(v)][len(v)]int
		for key, val := range v:
	}*/


/*

 */


