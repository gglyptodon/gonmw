package main

import (
"fmt"
"strconv"
"strings"
"os"
"io/ioutil"
"runtime"
"sync"
"math"
)
/*
 SUBSTITUTION MATRIX
*/
type SubstitutionMatrix struct{
	name string
	data map[string]int
}
func (sm SubstitutionMatrix) GetName() string {
	return sm.name
}
func (sm SubstitutionMatrix) GetVal(keyA string, keyB string) int {
	if val,ok := sm.data[keyA+keyB]; ok {
    	return val
	}else if val,ok := sm.data[keyB+keyA]; ok{
		return val
	}
	//not in map
	panic("Aaah!")
}

func (sm *SubstitutionMatrix) setMap(mapName string){
	if mapName == "EBLOSUM62"{
		sm.data = nil
		sm.data = map[string]int{"**":1,"*A":-4,"*C":-4,"*B":-4,"*E":-4,"*D":-4,"*G":-4,"*F":-4,"*I":-4,
"*H":-4,"*K":-4,"*M":-4,"*L":-4,"*N":-4,"*Q":-4,"*P":-4,"*S":-4,
"*R":-4,"*T":-4,"*W":-4,"*V":-4,"*Y":-4,"*X":-4,"*Z":-4,"A*":-4,
"AA":4,"AC":0,"AB":-2,"AE":-1,"AD":-2,"AG":0,"AF":-2,"AI":-1,
"AH":-2,"AK":-1,"AM":-1,"AL":-1,"AN":-2,"AQ":-1,"AP":-1,"AS":1,
"AR":-1,"AT":0,"AW":-3,"AV":0,"AY":-2,"AX":0,"AZ":-1,"C*":-4,
"CA":0,"CC":9,"CB":-3,"CE":-4,"CD":-3,"CG":-3,"CF":-2,"CI":-1,
"CH":-3,"CK":-3,"CM":-1,"CL":-1,"CN":-3,"CQ":-3,"CP":-3,"CS":-1,
"CR":-3,"CT":-1,"CW":-2,"CV":-1,"CY":-2,"CX":-2,"CZ":-3,"B*":-4,
"BA":-2,"BC":-3,"BB":4,"BE":1,"BD":4,"BG":-1,"BF":-3,"BI":-3,
"BH":0,"BK":0,"BM":-3,"BL":-4,"BN":3,"BQ":0,"BP":-2,"BS":0,
"BR":-1,"BT":-1,"BW":-4,"BV":-3,"BY":-3,"BX":-1,"BZ":1,"E*":-4,
"EA":-1,"EC":-4,"EB":1,"EE":5,"ED":2,"EG":-2,"EF":-3,"EI":-3,
"EH":0,"EK":1,"EM":-2,"EL":-3,"EN":0,"EQ":2,"EP":-1,"ES":0,
"ER":0,"ET":-1,"EW":-3,"EV":-2,"EY":-2,"EX":-1,"EZ":4,"D*":-4,
"DA":-2,"DC":-3,"DB":4,"DE":2,"DD":6,"DG":-1,"DF":-3,"DI":-3,
"DH":-1,"DK":-1,"DM":-3,"DL":-4,"DN":1,"DQ":0,"DP":-1,"DS":0,
"DR":-2,"DT":-1,"DW":-4,"DV":-3,"DY":-3,"DX":-1,"DZ":1,"G*":-4,
"GA":0,"GC":-3,"GB":-1,"GE":-2,"GD":-1,"GG":6,"GF":-3,"GI":-4,
"GH":-2,"GK":-2,"GM":-3,"GL":-4,"GN":0,"GQ":-2,"GP":-2,"GS":0,
"GR":-2,"GT":-2,"GW":-2,"GV":-3,"GY":-3,"GX":-1,"GZ":-2,"F*":-4,
"FA":-2,"FC":-2,"FB":-3,"FE":-3,"FD":-3,"FG":-3,"FF":6,"FI":0,
"FH":-1,"FK":-3,"FM":0,"FL":0,"FN":-3,"FQ":-3,"FP":-4,"FS":-2,
"FR":-3,"FT":-2,"FW":1,"FV":-1,"FY":3,"FX":-1,"FZ":-3,"I*":-4,
"IA":-1,"IC":-1,"IB":-3,"IE":-3,"ID":-3,"IG":-4,"IF":0,"II":4,
"IH":-3,"IK":-3,"IM":1,"IL":2,"IN":-3,"IQ":-3,"IP":-3,"IS":-2,
"IR":-3,"IT":-1,"IW":-3,"IV":3,"IY":-1,"IX":-1,"IZ":-3,"H*":-4,
"HA":-2,"HC":-3,"HB":0,"HE":0,"HD":-1,"HG":-2,"HF":-1,"HI":-3,
"HH":8,"HK":-1,"HM":-2,"HL":-3,"HN":1,"HQ":0,"HP":-2,"HS":-1,
"HR":0,"HT":-2,"HW":-2,"HV":-3,"HY":2,"HX":-1,"HZ":0,"K*":-4,
"KA":-1,"KC":-3,"KB":0,"KE":1,"KD":-1,"KG":-2,"KF":-3,"KI":-3,
"KH":-1,"KK":5,"KM":-1,"KL":-2,"KN":0,"KQ":1,"KP":-1,"KS":0,
"KR":2,"KT":-1,"KW":-3,"KV":-2,"KY":-2,"KX":-1,"KZ":1,"M*":-4,
"MA":-1,"MC":-1,"MB":-3,"ME":-2,"MD":-3,"MG":-3,"MF":0,"MI":1,
"MH":-2,"MK":-1,"MM":5,"ML":2,"MN":-2,"MQ":0,"MP":-2,"MS":-1,
"MR":-1,"MT":-1,"MW":-1,"MV":1,"MY":-1,"MX":-1,"MZ":-1,"L*":-4,
"LA":-1,"LC":-1,"LB":-4,"LE":-3,"LD":-4,"LG":-4,"LF":0,"LI":2,
"LH":-3,"LK":-2,"LM":2,"LL":4,"LN":-3,"LQ":-2,"LP":-3,"LS":-2,
"LR":-2,"LT":-1,"LW":-2,"LV":1,"LY":-1,"LX":-1,"LZ":-3,"N*":-4,
"NA":-2,"NC":-3,"NB":3,"NE":0,"ND":1,"NG":0,"NF":-3,"NI":-3,
"NH":1,"NK":0,"NM":-2,"NL":-3,"NN":6,"NQ":0,"NP":-2,"NS":1,
"NR":0,"NT":0,"NW":-4,"NV":-3,"NY":-2,"NX":-1,"NZ":0,"Q*":-4,
"QA":-1,"QC":-3,"QB":0,"QE":2,"QD":0,"QG":-2,"QF":-3,"QI":-3,
"QH":0,"QK":1,"QM":0,"QL":-2,"QN":0,"QQ":5,"QP":-1,"QS":0,
"QR":1,"QT":-1,"QW":-2,"QV":-2,"QY":-1,"QX":-1,"QZ":3,"P*":-4,
"PA":-1,"PC":-3,"PB":-2,"PE":-1,"PD":-1,"PG":-2,"PF":-4,"PI":-3,
"PH":-2,"PK":-1,"PM":-2,"PL":-3,"PN":-2,"PQ":-1,"PP":7,"PS":-1,
"PR":-2,"PT":-1,"PW":-4,"PV":-2,"PY":-3,"PX":-2,"PZ":-1,"S*":-4,
"SA":1,"SC":-1,"SB":0,"SE":0,"SD":0,"SG":0,"SF":-2,"SI":-2,
"SH":-1,"SK":0,"SM":-1,"SL":-2,"SN":1,"SQ":0,"SP":-1,"SS":4,
"SR":-1,"ST":1,"SW":-3,"SV":-2,"SY":-2,"SX":0,"SZ":0,"R*":-4,
"RA":-1,"RC":-3,"RB":-1,"RE":0,"RD":-2,"RG":-2,"RF":-3,"RI":-3,
"RH":0,"RK":2,"RM":-1,"RL":-2,"RN":0,"RQ":1,"RP":-2,"RS":-1,
"RR":5,"RT":-1,"RW":-3,"RV":-3,"RY":-2,"RX":-1,"RZ":0,"T*":-4,
"TA":0,"TC":-1,"TB":-1,"TE":-1,"TD":-1,"TG":-2,"TF":-2,"TI":-1,
"TH":-2,"TK":-1,"TM":-1,"TL":-1,"TN":0,"TQ":-1,"TP":-1,"TS":1,
"TR":-1,"TT":5,"TW":-2,"TV":0,"TY":-2,"TX":0,"TZ":-1,"W*":-4,
"WA":-3,"WC":-2,"WB":-4,"WE":-3,"WD":-4,"WG":-2,"WF":1,"WI":-3,
"WH":-2,"WK":-3,"WM":-1,"WL":-2,"WN":-4,"WQ":-2,"WP":-4,"WS":-3,
"WR":-3,"WT":-2,"WW":11,"WV":-3,"WY":2,"WX":-2,"WZ":-3,"V*":-4,
"VA":0,"VC":-1,"VB":-3,"VE":-2,"VD":-3,"VG":-3,"VF":-1,"VI":3,
"VH":-3,"VK":-2,"VM":1,"VL":1,"VN":-3,"VQ":-2,"VP":-2,"VS":-2,
"VR":-3,"VT":0,"VW":-3,"VV":4,"VY":-1,"VX":-1,"VZ":-2,"Y*":-4,
"YA":-2,"YC":-2,"YB":-3,"YE":-2,"YD":-3,"YG":-3,"YF":3,"YI":-1,
"YH":2,"YK":-2,"YM":-1,"YL":-1,"YN":-2,"YQ":-1,"YP":-3,"YS":-2,
"YR":-2,"YT":-2,"YW":2,"YV":-1,"YY":7,"YX":-1,"YZ":-2,"X*":-4,
"XA":0,"XC":-2,"XB":-1,"XE":-1,"XD":-1,"XG":-1,"XF":-1,"XI":-1,
"XH":-1,"XK":-1,"XM":-1,"XL":-1,"XN":-1,"XQ":-1,"XP":-2,"XS":0,
"XR":-1,"XT":0,"XW":-2,"XV":-1,"XY":-1,"XX":-1,"XZ":-1,"Z*":-4,
"ZA":-1,"ZC":-3,"ZB":1,"ZE":4,"ZD":1,"ZG":-2,"ZF":-3,"ZI":-3,
"ZH":0,"ZK":1,"ZM":-1,"ZL":-3,"ZN":0,"ZQ":3,"ZP":-1,"ZS":0,
"ZR":0,"ZT":-1,"ZW":-3,"ZV":-2,"ZY":-2,"ZX":-1,"ZZ":4}
	}else if mapName=="EBLOSUM50"{
		sm.data = map[string]int{"**":1,"*A":-5,"*C":-5,"*B":-5,"*E":-5,"*D":-5,"*G":-5,"*F":-5,"*I":-5,
"*H":-5,"*K":-5,"*M":-5,"*L":-5,"*N":-5,"*Q":-5,"*P":-5,"*S":-5,
"*R":-5,"*T":-5,"*W":-5,"*V":-5,"*Y":-5,"*X":-5,"*Z":-5,"A*":-5,
"AA":5,"AC":-1,"AB":-2,"AE":-1,"AD":-2,"AG":0,"AF":-3,"AI":-1,
"AH":-2,"AK":-1,"AM":-1,"AL":-2,"AN":-1,"AQ":-1,"AP":-1,"AS":1,
"AR":-2,"AT":0,"AW":-3,"AV":0,"AY":-2,"AX":-1,"AZ":-1,"C*":-5,
"CA":-1,"CC":13,"CB":-3,"CE":-3,"CD":-4,"CG":-3,"CF":-2,"CI":-2,
"CH":-3,"CK":-3,"CM":-2,"CL":-2,"CN":-2,"CQ":-3,"CP":-4,"CS":-1,
"CR":-4,"CT":-1,"CW":-5,"CV":-1,"CY":-3,"CX":-2,"CZ":-3,"B*":-5,
"BA":-2,"BC":-3,"BB":5,"BE":1,"BD":5,"BG":-1,"BF":-4,"BI":-4,
"BH":0,"BK":0,"BM":-3,"BL":-4,"BN":4,"BQ":0,"BP":-2,"BS":0,
"BR":-1,"BT":0,"BW":-5,"BV":-4,"BY":-3,"BX":-1,"BZ":2,"E*":-5,
"EA":-1,"EC":-3,"EB":1,"EE":6,"ED":2,"EG":-3,"EF":-3,"EI":-4,
"EH":0,"EK":1,"EM":-2,"EL":-3,"EN":0,"EQ":2,"EP":-1,"ES":-1,
"ER":0,"ET":-1,"EW":-3,"EV":-3,"EY":-2,"EX":-1,"EZ":5,"D*":-5,
"DA":-2,"DC":-4,"DB":5,"DE":2,"DD":8,"DG":-1,"DF":-5,"DI":-4,
"DH":-1,"DK":-1,"DM":-4,"DL":-4,"DN":2,"DQ":0,"DP":-1,"DS":0,
"DR":-2,"DT":-1,"DW":-5,"DV":-4,"DY":-3,"DX":-1,"DZ":1,"G*":-5,
"GA":0,"GC":-3,"GB":-1,"GE":-3,"GD":-1,"GG":8,"GF":-4,"GI":-4,
"GH":-2,"GK":-2,"GM":-3,"GL":-4,"GN":0,"GQ":-2,"GP":-2,"GS":0,
"GR":-3,"GT":-2,"GW":-3,"GV":-4,"GY":-3,"GX":-2,"GZ":-2,"F*":-5,
"FA":-3,"FC":-2,"FB":-4,"FE":-3,"FD":-5,"FG":-4,"FF":8,"FI":0,
"FH":-1,"FK":-4,"FM":0,"FL":1,"FN":-4,"FQ":-4,"FP":-4,"FS":-3,
"FR":-3,"FT":-2,"FW":1,"FV":-1,"FY":4,"FX":-2,"FZ":-4,"I*":-5,
"IA":-1,"IC":-2,"IB":-4,"IE":-4,"ID":-4,"IG":-4,"IF":0,"II":5,
"IH":-4,"IK":-3,"IM":2,"IL":2,"IN":-3,"IQ":-3,"IP":-3,"IS":-3,
"IR":-4,"IT":-1,"IW":-3,"IV":4,"IY":-1,"IX":-1,"IZ":-3,"H*":-5,
"HA":-2,"HC":-3,"HB":0,"HE":0,"HD":-1,"HG":-2,"HF":-1,"HI":-4,
"HH":10,"HK":0,"HM":-1,"HL":-3,"HN":1,"HQ":1,"HP":-2,"HS":-1,
"HR":0,"HT":-2,"HW":-3,"HV":-4,"HY":2,"HX":-1,"HZ":0,"K*":-5,
"KA":-1,"KC":-3,"KB":0,"KE":1,"KD":-1,"KG":-2,"KF":-4,"KI":-3,
"KH":0,"KK":6,"KM":-2,"KL":-3,"KN":0,"KQ":2,"KP":-1,"KS":0,
"KR":3,"KT":-1,"KW":-3,"KV":-3,"KY":-2,"KX":-1,"KZ":1,"M*":-5,
"MA":-1,"MC":-2,"MB":-3,"ME":-2,"MD":-4,"MG":-3,"MF":0,"MI":2,
"MH":-1,"MK":-2,"MM":7,"ML":3,"MN":-2,"MQ":0,"MP":-3,"MS":-2,
"MR":-2,"MT":-1,"MW":-1,"MV":1,"MY":0,"MX":-1,"MZ":-1,"L*":-5,
"LA":-2,"LC":-2,"LB":-4,"LE":-3,"LD":-4,"LG":-4,"LF":1,"LI":2,
"LH":-3,"LK":-3,"LM":3,"LL":5,"LN":-4,"LQ":-2,"LP":-4,"LS":-3,
"LR":-3,"LT":-1,"LW":-2,"LV":1,"LY":-1,"LX":-1,"LZ":-3,"N*":-5,
"NA":-1,"NC":-2,"NB":4,"NE":0,"ND":2,"NG":0,"NF":-4,"NI":-3,
"NH":1,"NK":0,"NM":-2,"NL":-4,"NN":7,"NQ":0,"NP":-2,"NS":1,
"NR":-1,"NT":0,"NW":-4,"NV":-3,"NY":-2,"NX":-1,"NZ":0,"Q*":-5,
"QA":-1,"QC":-3,"QB":0,"QE":2,"QD":0,"QG":-2,"QF":-4,"QI":-3,
"QH":1,"QK":2,"QM":0,"QL":-2,"QN":0,"QQ":7,"QP":-1,"QS":0,
"QR":1,"QT":-1,"QW":-1,"QV":-3,"QY":-1,"QX":-1,"QZ":4,"P*":-5,
"PA":-1,"PC":-4,"PB":-2,"PE":-1,"PD":-1,"PG":-2,"PF":-4,"PI":-3,
"PH":-2,"PK":-1,"PM":-3,"PL":-4,"PN":-2,"PQ":-1,"PP":10,"PS":-1,
"PR":-3,"PT":-1,"PW":-4,"PV":-3,"PY":-3,"PX":-2,"PZ":-1,"S*":-5,
"SA":1,"SC":-1,"SB":0,"SE":-1,"SD":0,"SG":0,"SF":-3,"SI":-3,
"SH":-1,"SK":0,"SM":-2,"SL":-3,"SN":1,"SQ":0,"SP":-1,"SS":5,
"SR":-1,"ST":2,"SW":-4,"SV":-2,"SY":-2,"SX":-1,"SZ":0,"R*":-5,
"RA":-2,"RC":-4,"RB":-1,"RE":0,"RD":-2,"RG":-3,"RF":-3,"RI":-4,
"RH":0,"RK":3,"RM":-2,"RL":-3,"RN":-1,"RQ":1,"RP":-3,"RS":-1,
"RR":7,"RT":-1,"RW":-3,"RV":-3,"RY":-1,"RX":-1,"RZ":0,"T*":-5,
"TA":0,"TC":-1,"TB":0,"TE":-1,"TD":-1,"TG":-2,"TF":-2,"TI":-1,
"TH":-2,"TK":-1,"TM":-1,"TL":-1,"TN":0,"TQ":-1,"TP":-1,"TS":2,
"TR":-1,"TT":5,"TW":-3,"TV":0,"TY":-2,"TX":0,"TZ":-1,"W*":-5,
"WA":-3,"WC":-5,"WB":-5,"WE":-3,"WD":-5,"WG":-3,"WF":1,"WI":-3,
"WH":-3,"WK":-3,"WM":-1,"WL":-2,"WN":-4,"WQ":-1,"WP":-4,"WS":-4,
"WR":-3,"WT":-3,"WW":15,"WV":-3,"WY":2,"WX":-3,"WZ":-2,"V*":-5,
"VA":0,"VC":-1,"VB":-4,"VE":-3,"VD":-4,"VG":-4,"VF":-1,"VI":4,
"VH":-4,"VK":-3,"VM":1,"VL":1,"VN":-3,"VQ":-3,"VP":-3,"VS":-2,
"VR":-3,"VT":0,"VW":-3,"VV":5,"VY":-1,"VX":-1,"VZ":-3,"Y*":-5,
"YA":-2,"YC":-3,"YB":-3,"YE":-2,"YD":-3,"YG":-3,"YF":4,"YI":-1,
"YH":2,"YK":-2,"YM":0,"YL":-1,"YN":-2,"YQ":-1,"YP":-3,"YS":-2,
"YR":-1,"YT":-2,"YW":2,"YV":-1,"YY":8,"YX":-1,"YZ":-2,"X*":-5,
"XA":-1,"XC":-2,"XB":-1,"XE":-1,"XD":-1,"XG":-2,"XF":-2,"XI":-1,
"XH":-1,"XK":-1,"XM":-1,"XL":-1,"XN":-1,"XQ":-1,"XP":-2,"XS":-1,
"XR":-1,"XT":0,"XW":-3,"XV":-1,"XY":-1,"XX":-1,"XZ":-1,"Z*":-5,
"ZA":-1,"ZC":-3,"ZB":2,"ZE":5,"ZD":1,"ZG":-2,"ZF":-4,"ZI":-3,
"ZH":0,"ZK":1,"ZM":-1,"ZL":-3,"ZN":0,"ZQ":4,"ZP":-1,"ZS":0,
"ZR":0,"ZT":-1,"ZW":-2,"ZV":-3,"ZY":-2,"ZX":-1,"ZZ":5}


	}
}





/*
SEQUENCE
*/
type Sequence struct{
	header string
	sequence string
}
func toFasta(seq Sequence) string{
	var res string
	res = seq.header+"\n"+seq.sequence
	return res
}

/*
SCORE
*/
type Score struct {
	seqA Sequence
	seqB Sequence
	sm SubstitutionMatrix
	res float64
}
func (s *Score ) setScore(newVal float64){
	s.res=newVal
}
func prettyPrint(score Score) string{
	var res string
	//print(score.res)
	res = score.seqA.header+"\t"+score.seqB.header+"\t"+strconv.FormatFloat(score.res,'f', 2, 32)+"\n"
	return res
}
/*
NMW
*/
func nmw(seqA Sequence, seqB Sequence, substMat SubstitutionMatrix )Score{
	//twoD array to iterate over
	m := len(seqA.sequence)+1
	n := len(seqB.sequence)+1
	//fmt.Println("m,n",m,n)
	var s Score
	s = Score{seqA:seqA, seqB:seqB, sm: substMat, res:0}
	//2d Scoring Matrix
	var mat = make([][]float64, m)
	for i := 0; i < m; i++ {
		mat[i] = make([]float64, n)
	}
	// Scoring matrix done

	//Boolean matrix to allow for gap extension
	var isGap = make([][]bool, m)
	for i := 0; i < m; i++ {
		isGap[i] = make([]bool, n)
	}

	//String matrix for path
	var pathMat = make([][]string,m)
	// ↑　top, ←　left, 　diag, × not used
	for i := 0; i < m; i++ {
		pathMat[i] = make([]string, n)
	}
	//fmt.Print(pathMat)




	var seqA_arr []string
	seqA_arr = strings.Split(seqA.sequence,"")

	var seqB_arr []string
	seqB_arr = strings.Split(seqB.sequence,"")

	//var gapOpenPenalty float64
	gapOpenPenalty,err := strconv.ParseFloat(os.Args[3],64)
	if err !=nil{panic(err)}
	//var gapExtendPenalty float64
	gapExtendPenalty,err := strconv.ParseFloat(os.Args[4],64)
	if err !=nil{panic(err)}

	var left float64
	var top float64
	var diag float64

	//var OP bool

	//gap open only if noy j==0..i==0
	for i:=1; i<=len(seqA_arr);i++{
		v := seqA_arr[i-1]
		for  j:= 1;j <= len(seqB_arr);j++ {
			isGap[i][j] = false
			//fmt.Println(seqA_arr[i-1], seqB_arr[j-1])
			if isGap [i][j - 1] {
				left = mat[i][j-1]+gapExtendPenalty
				//OP = false
			}else {
				left = mat[i][j-1]+gapOpenPenalty
				//OP = true
			}

			if isGap [i - 1][j] {
				top = mat[i-1][j]+gapExtendPenalty
				//OP = false
			}else {
				top = mat[i-1][j]+gapOpenPenalty
				//OP=true
			}
			w := seqB_arr[j - 1]

			diag = mat[i-1][j-1]+float64(s.sm.GetVal(v, w))
			mat[i][j] = maxOfThree(top, left, diag)

			if mat[i][j] == top{
				pathMat[i][j] = "↑"
			}else if mat[i][j] == left{
				pathMat[i][j] = "←"
			}else if mat[i][j] == diag{
				pathMat[i][j] = "↖"
				//pathMat[i-1][j-1] = "×"
			}else{
				pathMat[i][j] = "-"
			}
			if mat[i][j] == top || mat[i][j] == left {
				isGap[i][j] = true
				//if OP {fmt.Println("GAPOPEN", i,j,v,w)
				//}else{fmt.Println("GAPExtend", i,j,v,w)}
			}

		}
	}
	//print gap matrix
	/*for  i,_ := range isGap{
		for j,_ :=range isGap[i]{
			fmt.Print(isGap[i][j], " ")
		}
		fmt.Println("")
	}*/

	/* print matrix
	for i,_ := range pathMat{
		print(i," ")
		for j,_ :=range pathMat[i]{
			fmt.Print(pathMat[i][j], " ")
		}
		fmt.Println("")
	}
	fmt.Println("\n\n")
	*/

	fmt.Println("##########\n")
	//print matrix
	for i,_ := range mat{
		//print(i," ")
		for j,_ :=range mat[i]{
			if mat[i][j] >= 0&& mat[i][j] <10  {
				fmt.Print(" ")
			fmt.Print(mat[i][j],"  ")
			}else{
				fmt.Print(mat[i][j],"  ")
			}
		}
		fmt.Println("")
	}
	//fmt.Println("##########")
	//
	maxCol,maxRow, maxScore :=getMax3(mat)
	//fmt.Println("TEST", len(seqA_arr), len(seqB_arr), maxRow, maxCol, "\n\n")
	if pathMat[maxCol][maxRow] == "←"{
		 pathMat[maxCol][maxRow] ="⇐"
	}else if pathMat[maxCol][maxRow] == "↑"{
	pathMat[maxCol][maxRow] ="⇑"
	}else{
	pathMat[maxCol][maxRow] ="⇖"
	}
	//pathMat[maxRow][maxCol] = "当"
	//print matrix
	for i,_ := range pathMat{
		//print(i," ")
		for j,_ :=range pathMat[i]{
			fmt.Print(pathMat[i][j], " ")
		}
		fmt.Println("")
	}
	fmt.Println("\n")

	//fmt.Println(maxRow,maxCol,pathMat[maxCol][maxRow]  )
	s.setScore(maxScore)
	return s
}

func getMax3(twod [][]float64) (int,int,float64){
	var max float64
	var n int
	var m int
	n = len(twod)
	m = len(twod[n-1])
	//fmt.Println("LENGTHS",n,m)
	max = -1

	//last col max
	lastCol := n-1
	lastRow := twod[lastCol]
	lastRowMaxIndex,max := maximum(lastRow)
	//fmt.Println(lastRow)
	//fmt.Println("MAX from last Row", max)
	//last row max
	lastRowIndex := m-1
	var lastColArr []float64
	for i:=0;i<n;i++{
		lastColArr = append(lastColArr,  twod[i][lastRowIndex])
	}
	//fmt.Println("#####")
	//fmt.Println("lastCOl:" ,lastColArr)
	lastColMaxIndex,lastColMax := maximum(lastColArr)
	//fmt.Println("\n",lastColMax,"lcm")
	if max < lastColMax {
		return lastColMaxIndex,lastRowIndex, lastColMax
		//return  lastColMax
	}else{return lastCol,lastRowMaxIndex ,max	}
}

func maxOfThree(x float64, y float64, z float64)float64{
	if x >= y && x >= z{
		return x
	}else if y >= x && y >= z{
		return(y)
	}
	return z
}

/*
MAIN
*/

func main(){
	var fastaFileA string
	var fastaFileB string
	fastaFileA = os.Args[1]
	fastaFileB = os.Args[2]

	var allB []Sequence
	var allA []Sequence

	allA = FastaReader{file:fastaFileA}.getSequences()
	allB = FastaReader{file:fastaFileB}.getSequences()


	eblosum62 := SubstitutionMatrix{name:"EBLOSUM62", data: map[string]int{"_":-1}}
	eblosum62.setMap("EBLOSUM62")

	var maxCPU int = runtime.NumCPU()
	var wg sync.WaitGroup
	tasks :=make(chan Task,10000)
	var resultStr []string

	if (len(allA)*len(allB)<maxCPU){
		maxCPU = len(allA)*len(allB)
	}

	for i := 0; i < maxCPU; i++ {
        wg.Add(1)
        go func() {
            for t := range tasks {
				resultStr = append(resultStr,prettyPrint(nmw(t.a,t.b,t.sm)))
				//resultStr = resultStr +prettyPrint(nmw(t.a,t.b,t.sm))
            }
            wg.Done()
        }()
    }
	var seen map[string]bool
	seen = make(map[string]bool)
	var ok bool
	var tmpkey string
	for _,a := range allA{
		if a.sequence ==""{continue}
		for _,b:=range allB{
			if b.sequence ==""{continue}
			tmpkey = b.header+a.header
			ok = seen[tmpkey]
			if ok{continue}
			tmpkey = a.header+b.header
			ok = seen[tmpkey]
			if ok{continue}
			seen[a.header+b.header]=true
			seen[b.header+a.header]=true
			tasks <-  Task{a:a,b:b,sm:eblosum62}

		}
	}

	close(tasks)
	wg.Wait()

	for _,v := range resultStr{
		fmt.Print(v)
	}
}

func Consumer(limit int, inChan <-chan Task){
	for i := 0; i< limit;i++{
		for s := range inChan {
			fmt.Println(prettyPrint(nmw(s.a, s.b, s.sm)))
    	}
	}
}

func Producer(allA []Sequence, allB []Sequence, sm SubstitutionMatrix) <-chan Task {
    ch := make(chan Task)
    go func() {
		for _,a :=range allA{
			for _,b:= range allB{
				ch <- Task{a,b,sm}
			}
        }
        close(ch)
    }()
    return ch
}


func check(e error) {
    if e != nil {
        panic(e)
    }
}

type FastaReader struct{
	file string
}
func(fr FastaReader) getSequences()[]Sequence{
	var res []Sequence
	dat, err := ioutil.ReadFile(fr.file)
    check(err)
	var seqarr []string
	seqarr = strings.Split(string(dat),">") // now everything should start with a header
	for _,v := range(seqarr){
		var tmp []string
		tmp = strings.Split(v,"\n")
		shortHeader := strings.Split(tmp[0]," ")[0]
		res = append(res, Sequence{header:shortHeader, sequence:strings.Join(tmp[1:],"")})
	}
	return res

}
type Task struct{
	a Sequence
	b Sequence
	sm SubstitutionMatrix
}

//return max from slice
func maximum(slice []float64) (int,float64, ){
	maxIndex := -1
	currentMax := math.Inf(-1)
	for i,v := range slice{
		if v > currentMax{
			maxIndex = i
		}
		currentMax = math.Max(currentMax,v)

	}
	return maxIndex, currentMax
}
