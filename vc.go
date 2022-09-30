package main

import "fmt"

func main(){
	str:= "AOBADODB"
	x,y:=Check(str)

	fmt.Println("vowels:",x)
	fmt.Println("constants:",y)

}


func VowelsOrNot(str string) (string,string){

	var vs string
	var cs string

	for _,v:= range str {
		s := string(v)
		if s == "a" || s == "A" || s == "E" || s == "e" || s == "I" || s == "i" || s == "o" || s == "O" || s == "u" || s == "U" {
			vs = vs + s
			continue
		}
		cs = cs + s
	}

	return vs,cs

}

func Check(str string) (string,string){
	vowels,con:= VowelsOrNot(str)
	//fmt.Println(vowels,"",con)
	orderVowels:= make([]string,0)
	mv:=make(map[string]int)
	for i:=0;i<len(vowels);i++ {
		cha:= string(vowels[i])

		if _,Ok:=mv[cha];!Ok{
			mv[cha]++
			orderVowels =append(orderVowels,cha)
			continue
		}
		mv[cha] =mv[cha]+1
	}
	fmt.Println("Vowels map:", mv)
	fmt.Println("Vowels slice:",orderVowels)

	//To create order of con and find the number of con for each character
	orderCon:= make([]string,0)
	mc:=make(map[string]int)
	for j:=0;j<len(con);j++ {
		st := string(con[j])

		if _, Ok := mc[st]; !Ok {
			mc[st]++
			orderCon = append(orderCon, st)
			continue
		}
		mc[st] = mc[st] + 1
	}
	fmt.Println("con map:", mc)
	fmt.Println("con slice:",orderCon)

	var completeVowels string
	for  _,v:= range orderVowels{

		if _,ok:=mv[v];ok {
			for i:=0;i<mv[v];i++{
				completeVowels =completeVowels+v
			}

		}
	}

	var completeCon string
	for  _,c:= range orderCon{

		if _,ok:=mc[c];ok {
			for i:=0;i<mc[c];i++{
				completeCon =completeCon+c
			}

		}
	}


	return completeVowels,completeCon
}



