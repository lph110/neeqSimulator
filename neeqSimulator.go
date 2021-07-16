package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/LindsayBradford/go-dbf/godbf"
)

/*
 * 全国股转系统交易业务模拟器，20210529
 */

var conf Config //配置文件
var wtcount, hbcount, errorcount int

func main() {
	errorcount = 0
	fmt.Println("请先确保已经清库（NQWT.DBF,NQHB.DBF）")
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(89) + 10
	n_str := strconv.Itoa(n)
	fmt.Print("请输入[ ", n_str, " ]后,开始模拟全国股转系统交易业务:")
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if in.Text() != n_str && in.Text() != n_str {
		goexit()
		return
	} else {
		fmt.Println("")
		fmt.Println("已经开始模拟全国股转系统交易业务")
		fmt.Println("")
	}
	go spinner()
	if IsExist("./conf.json") {
		JsonParse := NewJsonStruct()
		JsonParse.Load("./conf.json", &conf)
	} else {
		fmt.Println("配置文件不存在")
		goexit()
		return
	}
	if conf.NoMatchCount == 0 {
		conf.NoMatchCount = 700
	}

	var NQWT = conf.NQWT
	var NQHB = conf.NQHB
	if IsExist(NQWT) != true {
		fmt.Println("委托库不存在")
		goexit()
		return
	}
	if IsExist(NQHB) != true {
		fmt.Println("回报库不存在")
		goexit()
		return
	}
	var HBCJHM_int64 int64 //流水号
	HBCJHM_int64 = 1
	dbfTable_t, err := godbf.NewFromFileByShare(NQHB, "UTF8")
	if err != nil {
		fmt.Println(err)
		goexit()
		return
	} else {
		if dbfTable_t.NumberOfRecords() == 0 {
			HBCJHM_int64 = 1
		} else {
			HBCJHM_int64 = int64(dbfTable_t.NumberOfRecords()) + 1
		}
	}
	hbcount = dbfTable_t.NumberOfRecords()
	for {
		dbfTable, err := godbf.NewFromFileByShare(NQWT, "UTF8")
		if err != nil {
			fmt.Println(err)
			goexit()
			return
		}
		wtcount = dbfTable.NumberOfRecords()
		for i := 0; i < dbfTable.NumberOfRecords(); i++ {
			WTHTXH, err := dbfTable.FieldValueByName(i, "WTHTXH")
			if err != nil {
				fmt.Println("读取WTHTXH错误:", err)
				continue
			}

			WTZQDM, err := dbfTable.FieldValueByName(i, "WTZQDM")
			if err != nil {
				fmt.Println("读取WTZQDM错误:", err)
				continue
			}
			WTZQZH, err := dbfTable.FieldValueByName(i, "WTZQZH")
			if err != nil {
				fmt.Println("读取WTZQZH错误:", err)
				continue
			}
			WTWTSL_str, err := dbfTable.FieldValueByName(i, "WTWTSL")
			if err != nil {
				fmt.Println("读取WTWTSL错误:", err)
				continue
			}
			WTWTSL, _ := strconv.ParseFloat(WTWTSL_str, 0)
			WTWTJG_str, err := dbfTable.FieldValueByName(i, "WTWTJG")
			if err != nil {
				fmt.Println("读取WTWTJG错误:", err)
				continue
			}
			WTWTJG, _ := strconv.ParseFloat(WTWTJG_str, 3)

			WTYWLB, err := dbfTable.FieldValueByName(i, "WTYWLB")
			if err != nil {
				fmt.Println("读取WTYWLB错误:", err)
				continue
			}
			WTDFDY, err := dbfTable.FieldValueByName(i, "WTDFDY")
			if err != nil {
				fmt.Println("读取WTDFDY错误:", err)
				continue
			}
			WTDFZH, err := dbfTable.FieldValueByName(i, "WTDFZH")
			if err != nil {
				fmt.Println("读取WTDFZH错误:", err)
				continue
			}
			WTWTSL2_str, err := dbfTable.FieldValueByName(i, "WTWTSL2")
			if err != nil {
				fmt.Println("读取WTWTSL2错误:", err)
				continue
			}
			WTWTSL2, _ := strconv.ParseFloat(WTWTSL2_str, 0)
			//
			WTWTJG2_str, err := dbfTable.FieldValueByName(i, "WTWTJG2")
			if err != nil {
				fmt.Println("读取WTWTJG2错误:", err)
				continue
			}
			WTWTJG2, _ := strconv.ParseFloat(WTWTJG2_str, 3)
			WTWTSJ, err := dbfTable.FieldValueByName(i, "WTWTSJ")
			if err != nil {
				fmt.Println("读取WTWTSJ错误:", err)
				continue
			}
			WTCLBZ, err := dbfTable.FieldValueByName(i, "WTCLBZ")
			if err != nil {
				fmt.Println("读取WTLXFS错误:", err)
				continue
			}
			/*
				WTLXR, err := dbfTable.FieldValueByName(i, "WTLXR")
				if err != nil {
					fmt.Println("读取WTLXR错误:", err)
					continue
				}
				WTLXFS, err := dbfTable.FieldValueByName(i, "WTLXFS")
				if err != nil {
					fmt.Println("读取WTLXFS错误:", err)
					continue
				}
				//
				WTYDH_str, err := dbfTable.FieldValueByName(i, "WTYDH")
				if err != nil {
					fmt.Println("读取WTYDH错误:", err)
					continue
				}
				WTYDH, _ := strconv.ParseFloat(WTYDH_str, 0)

				WTBYBZ, err := dbfTable.FieldValueByName(i, "WTBYBZ")
				if err != nil {
					fmt.Println("读取WTBYBZ错误:", err)
					continue
				}

				fmt.Println(WTHTXH, WTZQDM, WTZQZH, WTWTSL, WTWTJG, WTYWLB, WTDFDY, WTDFZH, WTWTSL2, WTWTJG2, WTLXR, WTLXFS, WTYDH, WTWTSJ, WTCLBZ, WTBYBZ)
			*/
			if WTCLBZ == "z" {
				c_d := time.Now().Format("20060102")
				if WTHTXH[6:14] != c_d {
					dbfTable.SetFieldValueByName(i, "WTCLBZ", "B")
					errorcount++
					continue
				}
				switch WTYWLB {
				case "0B":
					{
						//fmt.Println("业务类别:", WTYWLB, "限价申报买入申报记录")
						if WTWTSL > 0 && WTWTJG > 0 && WTWTSL2 == 0 && WTWTJG2 == 0 {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "1")
							if int64(WTWTSL)%conf.NoMatchCount != 0 {
								now := time.Now()
								HBCJHM_str_t := strconv.FormatInt(HBCJHM_int64, 10)
								HBCJHM_str := lenstring(HBCJHM_str_t, 8)
								HBCJHM := HBCJHM_str
								HBZQDM := WTZQDM
								HBHTXH := WTHTXH
								HBZQZH := WTZQZH
								HBCJSL := WTWTSL_str
								HBCJJG := WTWTJG_str
								HBCJSL2 := WTWTSL2_str
								HBDFDY := WTDFDY
								HBDFZH := WTDFZH
								HBCJSJ := now.Format("15040500")
								HBCJRQ := now.Format("20060102")
								HBYWLB := WTYWLB
								HBCDYY := ""
								HBBYBZ := ""
								setnqhb(HBCJHM, HBZQDM, HBHTXH, HBZQZH, HBCJSL, HBCJJG, HBCJSL2, HBDFDY, HBDFZH, HBCJSJ, HBCJRQ, HBYWLB, HBCDYY, HBBYBZ)
							}
						} else {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "F")
							errorcount++
						}
						break
					}
				case "0S":
					{
						//fmt.Println("业务类别:", WTYWLB, "限价申报卖出申报记录")
						if WTWTSL > 0 && WTWTJG > 0 && WTWTSL2 == 0 && WTWTJG2 == 0 {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "1")
							if int64(WTWTSL)%conf.NoMatchCount != 0 {
								now := time.Now()
								HBCJHM_str_t := strconv.FormatInt(HBCJHM_int64, 10)
								HBCJHM_str := lenstring(HBCJHM_str_t, 8)
								HBCJHM := HBCJHM_str
								HBZQDM := WTZQDM
								HBHTXH := WTHTXH
								HBZQZH := WTZQZH
								HBCJSL := WTWTSJ
								HBCJJG := WTWTJG_str
								HBCJSL2 := WTWTSL2_str
								HBDFDY := WTDFDY
								HBDFZH := WTDFZH
								HBCJSJ := now.Format("15040500")
								HBCJRQ := now.Format("20060102")
								HBYWLB := WTYWLB
								HBCDYY := ""
								HBBYBZ := ""
								setnqhb(HBCJHM, HBZQDM, HBHTXH, HBZQZH, HBCJSL, HBCJJG, HBCJSL2, HBDFDY, HBDFZH, HBCJSJ, HBCJRQ, HBYWLB, HBCDYY, HBBYBZ)
							}
						} else {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "F")
							errorcount++
						}
						break
					}
				case "0C":
					{
						//fmt.Println("业务类别:", WTYWLB, "限价申报撤单记录")
						if WTWTSL == 0 && WTWTJG == 0 && WTWTSL2 == 0 && WTWTJG2 == 0 {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "1")
							now := time.Now()
							HBCJHM_str_t := strconv.FormatInt(HBCJHM_int64, 10)
							HBCJHM_str := lenstring(HBCJHM_str_t, 8)
							HBCJHM := HBCJHM_str
							HBZQDM := WTZQDM
							HBHTXH := WTHTXH
							HBZQZH := WTZQZH
							HBCJSL := "-" + WTWTSJ
							HBCJJG := "0"
							HBCJSL2 := WTWTSL2_str
							HBDFDY := WTDFDY
							HBDFZH := ""
							HBCJSJ := now.Format("15040500")
							HBCJRQ := now.Format("20060102")
							HBYWLB := WTYWLB
							HBCDYY := ""
							HBBYBZ := ""
							setnqhb(HBCJHM, HBZQDM, HBHTXH, HBZQZH, HBCJSL, HBCJJG, HBCJSL2, HBDFDY, HBDFZH, HBCJSJ, HBCJRQ, HBYWLB, HBCDYY, HBBYBZ)
						} else {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "F")
							errorcount++
						}
						break
					}
				case "1B":
					{
						//fmt.Println("业务类别:", WTYWLB, "成交确认申报买入申报记录")
						if WTWTSL > 0 && WTWTJG > 0 && WTWTSL2 == 0 && WTWTJG2 == 0 {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "1")
							if int64(WTWTSL)%conf.NoMatchCount != 0 {
								now := time.Now()
								HBCJHM_str_t := strconv.FormatInt(HBCJHM_int64, 10)
								HBCJHM_str := lenstring(HBCJHM_str_t, 8)
								HBCJHM := HBCJHM_str
								HBZQDM := WTZQDM
								HBHTXH := WTHTXH
								HBZQZH := WTZQZH
								HBCJSL := WTWTSL_str
								HBCJJG := WTWTJG_str
								HBCJSL2 := "0"
								HBDFDY := WTDFDY
								HBDFZH := WTDFZH
								HBCJSJ := now.Format("15040500")
								HBCJRQ := now.Format("20060102")
								HBYWLB := WTYWLB
								HBCDYY := ""
								HBBYBZ := ""
								setnqhb(HBCJHM, HBZQDM, HBHTXH, HBZQZH, HBCJSL, HBCJJG, HBCJSL2, HBDFDY, HBDFZH, HBCJSJ, HBCJRQ, HBYWLB, HBCDYY, HBBYBZ)
							}
						} else {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "F")
							errorcount++
						}
						break
					}
				case "1S":
					{
						//fmt.Println("业务类别:", WTYWLB, "成交确认申报卖出申报记录")
						if WTWTSL > 0 && WTWTJG > 0 && WTWTSL2 == 0 && WTWTJG2 == 0 {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "1")
							if int64(WTWTSL)%conf.NoMatchCount != 0 {
								now := time.Now()
								HBCJHM_str_t := strconv.FormatInt(HBCJHM_int64, 10)
								HBCJHM_str := lenstring(HBCJHM_str_t, 8)
								HBCJHM := HBCJHM_str
								HBZQDM := WTZQDM
								HBHTXH := WTHTXH
								HBZQZH := WTZQZH
								HBCJSL := WTWTSL_str
								HBCJJG := WTWTJG_str
								HBCJSL2 := "0"
								HBDFDY := WTDFDY
								HBDFZH := WTDFZH
								HBCJSJ := now.Format("15040500")
								HBCJRQ := now.Format("20060102")
								HBYWLB := WTYWLB
								HBCDYY := ""
								HBBYBZ := ""
								setnqhb(HBCJHM, HBZQDM, HBHTXH, HBZQZH, HBCJSL, HBCJJG, HBCJSL2, HBDFDY, HBDFZH, HBCJSJ, HBCJRQ, HBYWLB, HBCDYY, HBBYBZ)
							}
						} else {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "F")
							errorcount++
						}
						break
					}
				case "1C":
					{
						//fmt.Println("业务类别:", WTYWLB, "成交确认申报撤单记录")
						if WTWTSL == 0 && WTWTJG == 0 && WTWTSL2 == 0 && WTWTJG2 == 0 {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "1")

							now := time.Now()
							HBCJHM_str_t := strconv.FormatInt(HBCJHM_int64, 10)
							HBCJHM_str := lenstring(HBCJHM_str_t, 8)
							HBCJHM := HBCJHM_str
							HBZQDM := WTZQDM
							HBHTXH := WTHTXH
							HBZQZH := WTZQZH
							HBCJSL := "-" + WTWTSL_str
							HBCJJG := "0"
							HBCJSL2 := "0"
							HBDFDY := WTDFDY
							HBDFZH := WTDFZH
							HBCJSJ := now.Format("15040500")
							HBCJRQ := now.Format("20060102")
							HBYWLB := WTYWLB
							HBCDYY := ""
							HBBYBZ := ""
							setnqhb(HBCJHM, HBZQDM, HBHTXH, HBZQZH, HBCJSL, HBCJJG, HBCJSL2, HBDFDY, HBDFZH, HBCJSJ, HBCJRQ, HBYWLB, HBCDYY, HBBYBZ)

						} else {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "F")
							errorcount++
						}
						break
					}
				case "2A":
					{
						//fmt.Println("业务类别:", WTYWLB, "双向做市申报记录")
						if WTWTSL > 0 && WTWTJG > 0 && WTWTSL2 > 0 && WTWTJG2 > 0 {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "1")
							if int64(WTWTSL)%conf.NoMatchCount != 0 {
								now := time.Now()
								HBCJHM_str_t := strconv.FormatInt(HBCJHM_int64, 10)
								HBCJHM_str := lenstring(HBCJHM_str_t, 8)
								HBCJHM := HBCJHM_str
								HBZQDM := WTZQDM
								HBHTXH := WTHTXH
								HBZQZH := WTZQZH
								HBCJSL := WTWTSL_str
								HBCJJG := WTWTJG_str
								HBCJSL2 := "0"
								HBDFDY := WTDFDY
								HBDFZH := WTDFZH
								HBCJSJ := now.Format("15040500")
								HBCJRQ := now.Format("20060102")
								HBYWLB := WTYWLB
								HBCDYY := ""
								HBBYBZ := ""
								setnqhb(HBCJHM, HBZQDM, HBHTXH, HBZQZH, HBCJSL, HBCJJG, HBCJSL2, HBDFDY, HBDFZH, HBCJSJ, HBCJRQ, HBYWLB, HBCDYY, HBBYBZ)
							}
						} else if WTWTSL > 0 && WTWTJG > 0 && WTWTSL2 == 0 && WTWTJG2 == 0 { //单向做市申报买入申报记录
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "1")
							if int64(WTWTSL)%conf.NoMatchCount != 0 {
								now := time.Now()
								HBCJHM_str_t := strconv.FormatInt(HBCJHM_int64, 10)
								HBCJHM_str := lenstring(HBCJHM_str_t, 8)
								HBCJHM := HBCJHM_str
								HBZQDM := WTZQDM
								HBHTXH := WTHTXH
								HBZQZH := WTZQZH
								HBCJSL := WTWTSL_str
								HBCJJG := WTWTJG_str
								HBCJSL2 := "0"
								HBDFDY := WTDFDY
								HBDFZH := WTDFZH
								HBCJSJ := now.Format("15040500")
								HBCJRQ := now.Format("20060102")
								HBYWLB := WTYWLB
								HBCDYY := ""
								HBBYBZ := ""
								setnqhb(HBCJHM, HBZQDM, HBHTXH, HBZQZH, HBCJSL, HBCJJG, HBCJSL2, HBDFDY, HBDFZH, HBCJSJ, HBCJRQ, HBYWLB, HBCDYY, HBBYBZ)
							}
						} else if WTWTSL == 0 && WTWTJG == 0 && WTWTSL2 > 0 && WTWTJG2 > 0 { //单向做市申报卖出申报记录
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "1")
							if int64(WTWTSL)%conf.NoMatchCount != 0 {
								now := time.Now()
								HBCJHM_str_t := strconv.FormatInt(HBCJHM_int64, 10)
								HBCJHM_str := lenstring(HBCJHM_str_t, 8)
								HBCJHM := HBCJHM_str
								HBZQDM := WTZQDM
								HBHTXH := WTHTXH
								HBZQZH := WTZQZH
								HBCJSL := WTWTSL_str
								HBCJJG := WTWTJG_str
								HBCJSL2 := "0"
								HBDFDY := WTDFDY
								HBDFZH := WTDFZH
								HBCJSJ := now.Format("15040500")
								HBCJRQ := now.Format("20060102")
								HBYWLB := WTYWLB
								HBCDYY := ""
								HBBYBZ := ""
								setnqhb(HBCJHM, HBZQDM, HBHTXH, HBZQZH, HBCJSL, HBCJJG, HBCJSL2, HBDFDY, HBDFZH, HBCJSJ, HBCJRQ, HBYWLB, HBCDYY, HBBYBZ)
							}
						} else {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "F")
							errorcount++
						}
						break
					}
				case "2C":
					{
						//fmt.Println("业务类别:", WTYWLB, "做市做市申报撤单记录")
						if WTWTSL == 0 && WTWTJG == 0 && WTWTSL2 == 0 && WTWTJG2 == 0 {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "1")
							if int64(WTWTSL)%conf.NoMatchCount != 0 {
								now := time.Now()
								HBCJHM_str_t := strconv.FormatInt(HBCJHM_int64, 10)
								HBCJHM_str := lenstring(HBCJHM_str_t, 8)
								HBCJHM := HBCJHM_str
								HBZQDM := WTZQDM
								HBHTXH := WTHTXH
								HBZQZH := WTZQZH
								HBCJSL := "-" + WTWTSL_str
								HBCJJG := "0"
								HBCJSL2 := "0"
								HBDFDY := WTDFDY
								HBDFZH := WTDFZH
								HBCJSJ := now.Format("15040500")
								HBCJRQ := now.Format("20060102")
								HBYWLB := WTYWLB
								HBCDYY := ""
								HBBYBZ := ""
								setnqhb(HBCJHM, HBZQDM, HBHTXH, HBZQZH, HBCJSL, HBCJJG, HBCJSL2, HBDFDY, HBDFZH, HBCJSJ, HBCJRQ, HBYWLB, HBCDYY, HBBYBZ)
							}
						} else {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "F")
							errorcount++
						}
						break
					}
				case "3B":
					{
						//fmt.Println("业务类别:", WTYWLB, "互报成交确认申报买入申报记录")
						if WTWTSL > 0 && WTWTJG > 0 && WTWTSL2 == 0 && WTWTJG2 == 0 {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "1")
							if int64(WTWTSL)%conf.NoMatchCount != 0 {
								now := time.Now()
								HBCJHM_str_t := strconv.FormatInt(HBCJHM_int64, 10)
								HBCJHM_str := lenstring(HBCJHM_str_t, 8)
								HBCJHM := HBCJHM_str
								HBZQDM := WTZQDM
								HBHTXH := WTHTXH
								HBZQZH := WTZQZH
								HBCJSL := WTWTSL_str
								HBCJJG := WTWTJG_str
								HBCJSL2 := "0"
								HBDFDY := WTDFDY
								HBDFZH := WTDFZH
								HBCJSJ := now.Format("15040500")
								HBCJRQ := now.Format("20060102")
								HBYWLB := WTYWLB
								HBCDYY := ""
								HBBYBZ := ""
								setnqhb(HBCJHM, HBZQDM, HBHTXH, HBZQZH, HBCJSL, HBCJJG, HBCJSL2, HBDFDY, HBDFZH, HBCJSJ, HBCJRQ, HBYWLB, HBCDYY, HBBYBZ)
							}
						} else {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "F")
							errorcount++
						}
						break
					}
				case "3S":
					{
						//fmt.Println("业务类别:", WTYWLB, "互报成交确认申报卖入申报记录")
						if WTWTSL > 0 && WTWTJG > 0 && WTWTSL2 == 0 && WTWTJG2 == 0 {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "1")
							if int64(WTWTSL)%conf.NoMatchCount != 0 {
								now := time.Now()
								HBCJHM_str_t := strconv.FormatInt(HBCJHM_int64, 10)
								HBCJHM_str := lenstring(HBCJHM_str_t, 8)
								HBCJHM := HBCJHM_str
								HBZQDM := WTZQDM
								HBHTXH := WTHTXH
								HBZQZH := WTZQZH
								HBCJSL := WTWTSL_str
								HBCJJG := WTWTJG_str
								HBCJSL2 := "0"
								HBDFDY := WTDFDY
								HBDFZH := WTDFZH
								HBCJSJ := now.Format("15040500")
								HBCJRQ := now.Format("20060102")
								HBYWLB := WTYWLB
								HBCDYY := ""
								HBBYBZ := ""
								setnqhb(HBCJHM, HBZQDM, HBHTXH, HBZQZH, HBCJSL, HBCJJG, HBCJSL2, HBDFDY, HBDFZH, HBCJSJ, HBCJRQ, HBYWLB, HBCDYY, HBBYBZ)
							}
						} else {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "F")
							errorcount++
						}
						break
					}
				case "3C":
					{
						//fmt.Println("业务类别:", WTYWLB, "互报成交确认申报撤单记录")
						if WTWTSL == 0 && WTWTJG == 0 && WTWTSL2 == 0 && WTWTJG2 == 0 {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "1")
							if int64(WTWTSL)%conf.NoMatchCount != 0 {
								now := time.Now()
								HBCJHM_str_t := strconv.FormatInt(HBCJHM_int64, 10)
								HBCJHM_str := lenstring(HBCJHM_str_t, 8)
								HBCJHM := HBCJHM_str
								HBZQDM := WTZQDM
								HBHTXH := WTHTXH
								HBZQZH := WTZQZH
								HBCJSL := "-" + WTWTSL_str
								HBCJJG := "0"
								HBCJSL2 := "0"
								HBDFDY := WTDFDY
								HBDFZH := WTDFZH
								HBCJSJ := now.Format("15040500")
								HBCJRQ := now.Format("20060102")
								HBYWLB := WTYWLB
								HBCDYY := ""
								HBBYBZ := ""
								setnqhb(HBCJHM, HBZQDM, HBHTXH, HBZQZH, HBCJSL, HBCJJG, HBCJSL2, HBDFDY, HBDFZH, HBCJSJ, HBCJRQ, HBYWLB, HBCDYY, HBBYBZ)
							}
						} else {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "F")
							errorcount++
						}
						break
					}
				case "4B":
					{
						//fmt.Println("业务类别:", WTYWLB, "做市互报成交确认申报买入申报记录")
						if WTWTSL > 0 && WTWTJG > 0 && WTWTSL2 == 0 && WTWTJG2 == 0 {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "1")
							if int64(WTWTSL)%conf.NoMatchCount != 0 {
								now := time.Now()
								HBCJHM_str_t := strconv.FormatInt(HBCJHM_int64, 10)
								HBCJHM_str := lenstring(HBCJHM_str_t, 8)
								HBCJHM := HBCJHM_str
								HBZQDM := WTZQDM
								HBHTXH := WTHTXH
								HBZQZH := WTZQZH
								HBCJSL := WTWTSL_str
								HBCJJG := WTWTJG_str
								HBCJSL2 := "0"
								HBDFDY := WTDFDY
								HBDFZH := WTDFZH
								HBCJSJ := now.Format("15040500")
								HBCJRQ := now.Format("20060102")
								HBYWLB := WTYWLB
								HBCDYY := ""
								HBBYBZ := ""
								setnqhb(HBCJHM, HBZQDM, HBHTXH, HBZQZH, HBCJSL, HBCJJG, HBCJSL2, HBDFDY, HBDFZH, HBCJSJ, HBCJRQ, HBYWLB, HBCDYY, HBBYBZ)
							}
						} else {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "F")
							errorcount++
						}
						break
					}
				case "4S":
					{
						//fmt.Println("业务类别:", WTYWLB, "做市互报成交确认申报卖出申报记录")
						if WTWTSL > 0 && WTWTJG > 0 && WTWTSL2 == 0 && WTWTJG2 == 0 {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "1")
							if int64(WTWTSL)%conf.NoMatchCount != 0 {
								now := time.Now()
								HBCJHM_str_t := strconv.FormatInt(HBCJHM_int64, 10)
								HBCJHM_str := lenstring(HBCJHM_str_t, 8)
								HBCJHM := HBCJHM_str
								HBZQDM := WTZQDM
								HBHTXH := WTHTXH
								HBZQZH := WTZQZH
								HBCJSL := WTWTSL_str
								HBCJJG := WTWTJG_str
								HBCJSL2 := "0"
								HBDFDY := WTDFDY
								HBDFZH := WTDFZH
								HBCJSJ := now.Format("15040500")
								HBCJRQ := now.Format("20060102")
								HBYWLB := WTYWLB
								HBCDYY := ""
								HBBYBZ := ""
								setnqhb(HBCJHM, HBZQDM, HBHTXH, HBZQZH, HBCJSL, HBCJJG, HBCJSL2, HBDFDY, HBDFZH, HBCJSJ, HBCJRQ, HBYWLB, HBCDYY, HBBYBZ)
							}
						} else {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "F")
							errorcount++
						}
						break
					}
				case "4C":
					{
						//fmt.Println("业务类别:", WTYWLB, "做市互报成交确认申报撤单记录")
						if WTWTSL == 0 && WTWTJG == 0 && WTWTSL2 == 0 && WTWTJG2 == 0 {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "1")

							now := time.Now()
							HBCJHM_str_t := strconv.FormatInt(HBCJHM_int64, 10)
							HBCJHM_str := lenstring(HBCJHM_str_t, 8)
							HBCJHM := HBCJHM_str
							HBZQDM := WTZQDM
							HBHTXH := WTHTXH
							HBZQZH := WTZQZH
							HBCJSL := "-" + WTWTSL_str
							HBCJJG := "0"
							HBCJSL2 := "0"
							HBDFDY := WTDFDY
							HBDFZH := WTDFZH
							HBCJSJ := now.Format("15040500")
							HBCJRQ := now.Format("20060102")
							HBYWLB := WTYWLB
							HBCDYY := ""
							HBBYBZ := ""
							setnqhb(HBCJHM, HBZQDM, HBHTXH, HBZQZH, HBCJSL, HBCJJG, HBCJSL2, HBDFDY, HBDFZH, HBCJSJ, HBCJRQ, HBYWLB, HBCDYY, HBBYBZ)

						} else {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "F")
							errorcount++
						}
						break
					}
				case "6B":
					{
						//fmt.Println("业务类别:", WTYWLB, "定价申报买入申报记录")
						if WTWTSL > 0 && WTWTJG > 0 && WTWTSL2 == 0 && WTWTJG2 == 0 {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "1")
							if int64(WTWTSL)%conf.NoMatchCount != 0 {
								now := time.Now()
								HBCJHM_str_t := strconv.FormatInt(HBCJHM_int64, 10)
								HBCJHM_str := lenstring(HBCJHM_str_t, 8)
								HBCJHM := HBCJHM_str
								HBZQDM := WTZQDM
								HBHTXH := WTHTXH
								HBZQZH := WTZQZH
								HBCJSL := WTWTSL_str
								HBCJJG := WTWTJG_str
								HBCJSL2 := "0"
								HBDFDY := WTDFDY
								HBDFZH := WTDFZH
								HBCJSJ := now.Format("15040500")
								HBCJRQ := now.Format("20060102")
								HBYWLB := WTYWLB
								HBCDYY := ""
								HBBYBZ := ""
								setnqhb(HBCJHM, HBZQDM, HBHTXH, HBZQZH, HBCJSL, HBCJJG, HBCJSL2, HBDFDY, HBDFZH, HBCJSJ, HBCJRQ, HBYWLB, HBCDYY, HBBYBZ)
							}
						} else {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "F")
							errorcount++
						}
						break
					}
				case "6S":
					{
						//fmt.Println("业务类别:", WTYWLB, "定价申报卖出申报记录")
						if WTWTSL > 0 && WTWTJG > 0 && WTWTSL2 == 0 && WTWTJG2 == 0 {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "1")
							if int64(WTWTSL)%conf.NoMatchCount != 0 {
								now := time.Now()
								HBCJHM_str_t := strconv.FormatInt(HBCJHM_int64, 10)
								HBCJHM_str := lenstring(HBCJHM_str_t, 8)
								HBCJHM := HBCJHM_str
								HBZQDM := WTZQDM
								HBHTXH := WTHTXH
								HBZQZH := WTZQZH
								HBCJSL := WTWTSL_str
								HBCJJG := WTWTJG_str
								HBCJSL2 := "0"
								HBDFDY := WTDFDY
								HBDFZH := WTDFZH
								HBCJSJ := now.Format("15040500")
								HBCJRQ := now.Format("20060102")
								HBYWLB := WTYWLB
								HBCDYY := ""
								HBBYBZ := ""
								setnqhb(HBCJHM, HBZQDM, HBHTXH, HBZQZH, HBCJSL, HBCJJG, HBCJSL2, HBDFDY, HBDFZH, HBCJSJ, HBCJRQ, HBYWLB, HBCDYY, HBBYBZ)
							}
						} else {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "F")
							errorcount++
						}
						break
					}
				case "6C":
					{
						//fmt.Println("业务类别:", WTYWLB, "定价申报撤单记录")
						if WTWTSL == 0 && WTWTJG == 0 && WTWTSL2 == 0 && WTWTJG2 == 0 {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "1")

							now := time.Now()
							HBCJHM_str_t := strconv.FormatInt(HBCJHM_int64, 10)
							HBCJHM_str := lenstring(HBCJHM_str_t, 8)
							HBCJHM := HBCJHM_str
							HBZQDM := WTZQDM
							HBHTXH := WTHTXH
							HBZQZH := WTZQZH
							HBCJSL := "-" + WTWTSL_str
							HBCJJG := "0"
							HBCJSL2 := "0"
							HBDFDY := WTDFDY
							HBDFZH := WTDFZH
							HBCJSJ := now.Format("15040500")
							HBCJRQ := now.Format("20060102")
							HBYWLB := WTYWLB
							HBCDYY := ""
							HBBYBZ := ""
							setnqhb(HBCJHM, HBZQDM, HBHTXH, HBZQZH, HBCJSL, HBCJJG, HBCJSL2, HBDFDY, HBDFZH, HBCJSJ, HBCJRQ, HBYWLB, HBCDYY, HBBYBZ)

						} else {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "F")
							errorcount++
						}
						break
					}
				case "7B":
					{
						//fmt.Println("业务类别:", WTYWLB, "询价申报记录")
						if WTWTSL > 0 && WTWTJG > 0 && WTWTSL2 == 0 && WTWTJG2 == 0 {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "1")
							if int64(WTWTSL)%conf.NoMatchCount != 0 {
								now := time.Now()
								HBCJHM_str_t := strconv.FormatInt(HBCJHM_int64, 10)
								HBCJHM_str := lenstring(HBCJHM_str_t, 8)
								HBCJHM := HBCJHM_str
								HBZQDM := WTZQDM
								HBHTXH := WTHTXH
								HBZQZH := WTZQZH
								HBCJSL := WTWTSL_str
								HBCJJG := WTWTJG_str
								HBCJSL2 := "0"
								HBDFDY := WTDFDY
								HBDFZH := WTDFZH
								HBCJSJ := now.Format("15040500")
								HBCJRQ := now.Format("20060102")
								HBYWLB := WTYWLB
								HBCDYY := ""
								HBBYBZ := ""
								setnqhb(HBCJHM, HBZQDM, HBHTXH, HBZQZH, HBCJSL, HBCJJG, HBCJSL2, HBDFDY, HBDFZH, HBCJSJ, HBCJRQ, HBYWLB, HBCDYY, HBBYBZ)
							}
						} else {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "F")
							errorcount++
						}
						break
					}
				case "8B":
					{
						//fmt.Println("业务类别:", WTYWLB, "申购申报记录")
						if WTWTSL > 0 && WTWTJG > 0 && WTWTSL2 == 0 && WTWTJG2 == 0 {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "1")
							if int64(WTWTSL)%conf.NoMatchCount != 0 {
								now := time.Now()
								HBCJHM_str_t := strconv.FormatInt(HBCJHM_int64, 10)
								HBCJHM_str := lenstring(HBCJHM_str_t, 8)
								HBCJHM := HBCJHM_str
								HBZQDM := WTZQDM
								HBHTXH := WTHTXH
								HBZQZH := WTZQZH
								HBCJSL := WTWTSL_str
								HBCJJG := WTWTJG_str
								HBCJSL2 := "0"
								HBDFDY := WTDFDY
								HBDFZH := WTDFZH
								HBCJSJ := now.Format("15040500")
								HBCJRQ := now.Format("20060102")
								HBYWLB := WTYWLB
								HBCDYY := ""
								HBBYBZ := ""
								setnqhb(HBCJHM, HBZQDM, HBHTXH, HBZQZH, HBCJSL, HBCJJG, HBCJSL2, HBDFDY, HBDFZH, HBCJSJ, HBCJRQ, HBYWLB, HBCDYY, HBBYBZ)
							}
						} else {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "F")
							errorcount++
						}
						break
					}
				case "EB":
					{
						//fmt.Println("业务类别:", WTYWLB, "撤回预受要约申报记录")
						if WTWTSL > 0 && WTWTJG >= 0 && WTWTSL2 >= 0 && WTWTJG2 == 0 {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "1")
							now := time.Now()
							HBCJHM_str_t := strconv.FormatInt(HBCJHM_int64, 10)
							HBCJHM_str := lenstring(HBCJHM_str_t, 8)
							HBCJHM := HBCJHM_str
							HBZQDM := WTZQDM
							HBHTXH := WTHTXH
							HBZQZH := WTZQZH
							HBCJSL := WTWTSJ
							HBCJJG := WTWTJG2_str
							HBCJSL2 := WTWTSL2_str
							HBDFDY := WTDFDY
							HBDFZH := ""
							HBCJSJ := now.Format("15040500")
							HBCJRQ := now.Format("20060102")
							HBYWLB := WTYWLB
							HBCDYY := ""
							HBBYBZ := ""
							setnqhb(HBCJHM, HBZQDM, HBHTXH, HBZQZH, HBCJSL, HBCJJG, HBCJSL2, HBDFDY, HBDFZH, HBCJSJ, HBCJRQ, HBYWLB, HBCDYY, HBBYBZ)
						} else {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "F")
							errorcount++
						}
						break
					}
				case "ES":
					{
						//fmt.Println("业务类别:", WTYWLB, "预受要约申报记录")
						if WTWTSL > 0 && WTWTJG >= 0 && WTWTSL2 >= 0 && WTWTJG2 == 0 {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "1")
							now := time.Now()
							HBCJHM_str_t := strconv.FormatInt(HBCJHM_int64, 10)
							HBCJHM_str := lenstring(HBCJHM_str_t, 8)
							HBCJHM := HBCJHM_str
							HBZQDM := WTZQDM
							HBHTXH := WTHTXH
							HBZQZH := WTZQZH
							HBCJSL := WTWTSJ
							HBCJJG := WTWTJG2_str
							HBCJSL2 := WTWTSL2_str
							HBDFDY := WTDFDY
							HBDFZH := ""
							HBCJSJ := now.Format("15040500")
							HBCJRQ := now.Format("20060102")
							HBYWLB := WTYWLB
							HBCDYY := ""
							HBBYBZ := ""
							setnqhb(HBCJHM, HBZQDM, HBHTXH, HBZQZH, HBCJSL, HBCJJG, HBCJSL2, HBDFDY, HBDFZH, HBCJSJ, HBCJRQ, HBYWLB, HBCDYY, HBBYBZ)
						} else {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "F")
							errorcount++
						}
						break
					}
				case "EC":
					{
						//fmt.Println("业务类别:", WTYWLB, "要约撤单申报记录")
						if WTWTSL == 0 && WTWTJG == 0 && WTWTSL2 >= 0 && WTWTJG2 == 0 {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "1")
							now := time.Now()
							HBCJHM_str_t := strconv.FormatInt(HBCJHM_int64, 10)
							HBCJHM_str := lenstring(HBCJHM_str_t, 8)
							HBCJHM := HBCJHM_str
							HBZQDM := WTZQDM
							HBHTXH := WTHTXH
							HBZQZH := WTZQZH
							HBCJSL := "-" + WTWTSJ
							HBCJJG := "0"
							HBCJSL2 := "0"
							HBDFDY := WTDFDY
							HBDFZH := ""
							HBCJSJ := now.Format("15040500")
							HBCJRQ := now.Format("20060102")
							HBYWLB := WTYWLB
							HBCDYY := ""
							HBBYBZ := ""
							setnqhb(HBCJHM, HBZQDM, HBHTXH, HBZQZH, HBCJSL, HBCJJG, HBCJSL2, HBDFDY, HBDFZH, HBCJSJ, HBCJRQ, HBYWLB, HBCDYY, HBBYBZ)
						} else {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "F")
							errorcount++
						}
						break
					}
				case "VB":
					{
						//fmt.Println("业务类别:", WTYWLB, "最优五档即时成交剩余撤销买入申报记录")
						if WTWTSL > 0 && WTWTJG == 0 && WTWTSL2 == 0 && WTWTJG2 > 0 {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "1")
							now := time.Now()
							HBCJHM_str_t := strconv.FormatInt(HBCJHM_int64, 10)
							HBCJHM_str := lenstring(HBCJHM_str_t, 8)
							HBCJHM := HBCJHM_str
							HBZQDM := WTZQDM
							HBHTXH := WTHTXH
							HBZQZH := WTZQZH
							HBCJSL := WTWTSJ
							HBCJJG := WTWTJG2_str
							HBCJSL2 := WTWTSL2_str
							HBDFDY := WTDFDY
							HBDFZH := ""
							HBCJSJ := now.Format("15040500")
							HBCJRQ := now.Format("20060102")
							HBYWLB := WTYWLB
							HBCDYY := ""
							HBBYBZ := ""
							setnqhb(HBCJHM, HBZQDM, HBHTXH, HBZQZH, HBCJSL, HBCJJG, HBCJSL2, HBDFDY, HBDFZH, HBCJSJ, HBCJRQ, HBYWLB, HBCDYY, HBBYBZ)
						} else {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "F")
							errorcount++
						}
						break
					}
				case "VS":
					{
						//fmt.Println("业务类别:", WTYWLB, "最优五档即时成交剩余撤销卖出申报记录")
						if WTWTSL > 0 && WTWTJG == 0 && WTWTSL2 == 0 && WTWTJG2 > 0 {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "1")
							now := time.Now()
							HBCJHM_str_t := strconv.FormatInt(HBCJHM_int64, 10)
							HBCJHM_str := lenstring(HBCJHM_str_t, 8)
							HBCJHM := HBCJHM_str
							HBZQDM := WTZQDM
							HBHTXH := WTHTXH
							HBZQZH := WTZQZH
							HBCJSL := WTWTSJ
							HBCJJG := WTWTJG2_str
							HBCJSL2 := WTWTSL2_str
							HBDFDY := WTDFDY
							HBDFZH := ""
							HBCJSJ := now.Format("15040500")
							HBCJRQ := now.Format("20060102")
							HBYWLB := WTYWLB
							HBCDYY := ""
							HBBYBZ := ""
							setnqhb(HBCJHM, HBZQDM, HBHTXH, HBZQZH, HBCJSL, HBCJJG, HBCJSL2, HBDFDY, HBDFZH, HBCJSJ, HBCJRQ, HBYWLB, HBCDYY, HBBYBZ)
						} else {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "F")
							errorcount++
						}
						break
					}
				case "WB":
					{
						//fmt.Println("业务类别:", WTYWLB, "最优五档即时成交剩余转限价买入申报记录")
						if WTWTSL > 0 && WTWTJG == 0 && WTWTSL2 == 0 && WTWTJG2 > 0 {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "1")
							now := time.Now()
							HBCJHM_str_t := strconv.FormatInt(HBCJHM_int64, 10)
							HBCJHM_str := lenstring(HBCJHM_str_t, 8)
							HBCJHM := HBCJHM_str
							HBZQDM := WTZQDM
							HBHTXH := WTHTXH
							HBZQZH := WTZQZH
							HBCJSL := WTWTSJ
							HBCJJG := WTWTJG2_str
							HBCJSL2 := WTWTSL2_str
							HBDFDY := WTDFDY
							HBDFZH := ""
							HBCJSJ := now.Format("15040500")
							HBCJRQ := now.Format("20060102")
							HBYWLB := WTYWLB
							HBCDYY := ""
							HBBYBZ := ""
							setnqhb(HBCJHM, HBZQDM, HBHTXH, HBZQZH, HBCJSL, HBCJJG, HBCJSL2, HBDFDY, HBDFZH, HBCJSJ, HBCJRQ, HBYWLB, HBCDYY, HBBYBZ)
						} else {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "F")
							errorcount++
						}
						break
					}
				case "WS":
					{
						//fmt.Println("业务类别:", WTYWLB, "最优五档即时成交剩余转限价卖出申报记录")
						if WTWTSL > 0 && WTWTJG == 0 && WTWTSL2 == 0 && WTWTJG2 > 0 {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "1")
							now := time.Now()
							HBCJHM_str_t := strconv.FormatInt(HBCJHM_int64, 10)
							HBCJHM_str := lenstring(HBCJHM_str_t, 8)
							HBCJHM := HBCJHM_str
							HBZQDM := WTZQDM
							HBHTXH := WTHTXH
							HBZQZH := WTZQZH
							HBCJSL := WTWTSJ
							HBCJJG := WTWTJG2_str
							HBCJSL2 := WTWTSL2_str
							HBDFDY := WTDFDY
							HBDFZH := ""
							HBCJSJ := now.Format("15040500")
							HBCJRQ := now.Format("20060102")
							HBYWLB := WTYWLB
							HBCDYY := ""
							HBBYBZ := ""
							setnqhb(HBCJHM, HBZQDM, HBHTXH, HBZQZH, HBCJSL, HBCJJG, HBCJSL2, HBDFDY, HBDFZH, HBCJSJ, HBCJRQ, HBYWLB, HBCDYY, HBBYBZ)
						} else {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "F")
							errorcount++
						}
						break
					}
				case "WC":
					{
						//fmt.Println("业务类别:", WTYWLB, "最优五档即时成交剩余转限价卖出申报记录")
						if WTWTSL == 0 && WTWTJG == 0 && WTWTSL2 == 0 && WTWTJG2 == 0 {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "1")
							now := time.Now()
							HBCJHM_str_t := strconv.FormatInt(HBCJHM_int64, 10)
							HBCJHM_str := lenstring(HBCJHM_str_t, 8)
							HBCJHM := HBCJHM_str
							HBZQDM := WTZQDM
							HBHTXH := WTHTXH
							HBZQZH := WTZQZH
							HBCJSL := "-" + WTWTSJ
							HBCJJG := "0"
							HBCJSL2 := "0"
							HBDFDY := WTDFDY
							HBDFZH := ""
							HBCJSJ := now.Format("15040500")
							HBCJRQ := now.Format("20060102")
							HBYWLB := WTYWLB
							HBCDYY := ""
							HBBYBZ := ""
							setnqhb(HBCJHM, HBZQDM, HBHTXH, HBZQZH, HBCJSL, HBCJJG, HBCJSL2, HBDFDY, HBDFZH, HBCJSJ, HBCJRQ, HBYWLB, HBCDYY, HBBYBZ)
						} else {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "F")
							errorcount++
						}
						break
					}
				case "XB":
					{
						//fmt.Println("业务类别:", WTYWLB, "本方最优申报买入申报记录")
						if WTWTSL > 0 && WTWTJG == 0 && WTWTSL2 == 0 && WTWTJG2 > 0 {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "1")
							now := time.Now()
							HBCJHM_str_t := strconv.FormatInt(HBCJHM_int64, 10)
							HBCJHM_str := lenstring(HBCJHM_str_t, 8)
							HBCJHM := HBCJHM_str
							HBZQDM := WTZQDM
							HBHTXH := WTHTXH
							HBZQZH := WTZQZH
							HBCJSL := WTWTSJ
							HBCJJG := WTWTJG2_str
							HBCJSL2 := WTWTSL2_str
							HBDFDY := WTDFDY
							HBDFZH := ""
							HBCJSJ := now.Format("15040500")
							HBCJRQ := now.Format("20060102")
							HBYWLB := WTYWLB
							HBCDYY := ""
							HBBYBZ := ""
							setnqhb(HBCJHM, HBZQDM, HBHTXH, HBZQZH, HBCJSL, HBCJJG, HBCJSL2, HBDFDY, HBDFZH, HBCJSJ, HBCJRQ, HBYWLB, HBCDYY, HBBYBZ)
						} else {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "F")
							errorcount++
						}
						break
					}
				case "XS":
					{
						//fmt.Println("业务类别:", WTYWLB, "本方最优申报卖出申报记录")
						if WTWTSL > 0 && WTWTJG == 0 && WTWTSL2 == 0 && WTWTJG2 > 0 {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "1")
							now := time.Now()
							HBCJHM_str_t := strconv.FormatInt(HBCJHM_int64, 10)
							HBCJHM_str := lenstring(HBCJHM_str_t, 8)
							HBCJHM := HBCJHM_str
							HBZQDM := WTZQDM
							HBHTXH := WTHTXH
							HBZQZH := WTZQZH
							HBCJSL := WTWTSJ
							HBCJJG := WTWTJG2_str
							HBCJSL2 := WTWTSL2_str
							HBDFDY := WTDFDY
							HBDFZH := ""
							HBCJSJ := now.Format("15040500")
							HBCJRQ := now.Format("20060102")
							HBYWLB := WTYWLB
							HBCDYY := ""
							HBBYBZ := ""
							setnqhb(HBCJHM, HBZQDM, HBHTXH, HBZQZH, HBCJSL, HBCJJG, HBCJSL2, HBDFDY, HBDFZH, HBCJSJ, HBCJRQ, HBYWLB, HBCDYY, HBBYBZ)
						} else {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "F")
							errorcount++
						}
						break
					}
				case "XC":
					{
						//fmt.Println("业务类别:", WTYWLB, "本方最优申报撤单记录")
						if WTWTSL == 0 && WTWTJG == 0 && WTWTSL2 == 0 && WTWTJG2 == 0 {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "1")
							now := time.Now()
							HBCJHM_str_t := strconv.FormatInt(HBCJHM_int64, 10)
							HBCJHM_str := lenstring(HBCJHM_str_t, 8)
							HBCJHM := HBCJHM_str
							HBZQDM := WTZQDM
							HBHTXH := WTHTXH
							HBZQZH := WTZQZH
							HBCJSL := "-" + WTWTSJ
							HBCJJG := "0"
							HBCJSL2 := "0"
							HBDFDY := WTDFDY
							HBDFZH := ""
							HBCJSJ := now.Format("15040500")
							HBCJRQ := now.Format("20060102")
							HBYWLB := WTYWLB
							HBCDYY := ""
							HBBYBZ := ""
							setnqhb(HBCJHM, HBZQDM, HBHTXH, HBZQZH, HBCJSL, HBCJJG, HBCJSL2, HBDFDY, HBDFZH, HBCJSJ, HBCJRQ, HBYWLB, HBCDYY, HBBYBZ)
						} else {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "F")
							errorcount++
						}
						break
					}
				case "YB":
					{
						//fmt.Println("业务类别:", WTYWLB, "对手方最优申报买入申报记录")
						if WTWTSL > 0 && WTWTJG == 0 && WTWTSL2 == 0 && WTWTJG2 > 0 {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "1")
							now := time.Now()
							HBCJHM_str_t := strconv.FormatInt(HBCJHM_int64, 10)
							HBCJHM_str := lenstring(HBCJHM_str_t, 8)
							HBCJHM := HBCJHM_str
							HBZQDM := WTZQDM
							HBHTXH := WTHTXH
							HBZQZH := WTZQZH
							HBCJSL := WTWTSJ
							HBCJJG := WTWTJG2_str
							HBCJSL2 := WTWTSL2_str
							HBDFDY := WTDFDY
							HBDFZH := ""
							HBCJSJ := now.Format("15040500")
							HBCJRQ := now.Format("20060102")
							HBYWLB := WTYWLB
							HBCDYY := ""
							HBBYBZ := ""
							setnqhb(HBCJHM, HBZQDM, HBHTXH, HBZQZH, HBCJSL, HBCJJG, HBCJSL2, HBDFDY, HBDFZH, HBCJSJ, HBCJRQ, HBYWLB, HBCDYY, HBBYBZ)
						} else {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "F")
							errorcount++
						}
						break
					}
				case "YS":
					{
						//fmt.Println("业务类别:", WTYWLB, "对手方最优申报卖出申报记录")
						if WTWTSL > 0 && WTWTJG == 0 && WTWTSL2 == 0 && WTWTJG2 > 0 {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "1")
							now := time.Now()
							HBCJHM_str_t := strconv.FormatInt(HBCJHM_int64, 10)
							HBCJHM_str := lenstring(HBCJHM_str_t, 8)
							HBCJHM := HBCJHM_str
							HBZQDM := WTZQDM
							HBHTXH := WTHTXH
							HBZQZH := WTZQZH
							HBCJSL := WTWTSJ
							HBCJJG := WTWTJG2_str
							HBCJSL2 := WTWTSL2_str
							HBDFDY := WTDFDY
							HBDFZH := ""
							HBCJSJ := now.Format("15040500")
							HBCJRQ := now.Format("20060102")
							HBYWLB := WTYWLB
							HBCDYY := ""
							HBBYBZ := ""
							setnqhb(HBCJHM, HBZQDM, HBHTXH, HBZQZH, HBCJSL, HBCJJG, HBCJSL2, HBDFDY, HBDFZH, HBCJSJ, HBCJRQ, HBYWLB, HBCDYY, HBBYBZ)
						} else {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "F")
							errorcount++
						}
						break
					}
				case "YC":
					{
						//fmt.Println("业务类别:", WTYWLB, "对手方最优申报撤单记录")
						if WTWTSL == 0 && WTWTJG == 0 && WTWTSL2 == 0 && WTWTJG2 == 0 {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "1")
							now := time.Now()
							HBCJHM_str_t := strconv.FormatInt(HBCJHM_int64, 10)
							HBCJHM_str := lenstring(HBCJHM_str_t, 8)
							HBCJHM := HBCJHM_str
							HBZQDM := WTZQDM
							HBHTXH := WTHTXH
							HBZQZH := WTZQZH
							HBCJSL := "-" + WTWTSJ
							HBCJJG := "0"
							HBCJSL2 := "0"
							HBDFDY := WTDFDY
							HBDFZH := ""
							HBCJSJ := now.Format("15040500")
							HBCJRQ := now.Format("20060102")
							HBYWLB := WTYWLB
							HBCDYY := ""
							HBBYBZ := ""
							setnqhb(HBCJHM, HBZQDM, HBHTXH, HBZQZH, HBCJSL, HBCJJG, HBCJSL2, HBDFDY, HBDFZH, HBCJSJ, HBCJRQ, HBYWLB, HBCDYY, HBBYBZ)
						} else {
							dbfTable.SetFieldValueByName(i, "WTCLBZ", "F")
							errorcount++
						}
						break
					}
				default:
					{
						fmt.Println("不支持的业务类别:", WTYWLB)
					}
				}
				HBCJHM_int64 = HBCJHM_int64 + 1
			}
		}
		dbfTable.SaveFilenoLog(NQWT)
		time.Sleep(1 * time.Second) //等待1秒，单位秒
	}
}
func setnqhb(HBCJHM string, HBZQDM string, HBHTXH string, HBZQZH string, HBCJSL string, HBCJJG string, HBCJSL2 string, HBDFDY string, HBDFZH string, HBCJSJ string, HBCJRQ string, HBYWLB string, HBCDYY string, HBBYBZ string) {
	var NQHB = conf.NQHB
	if IsExist(NQHB) != true {
		fmt.Println("回报库不存在")
		goexit()
		return
	}

	dbfTable, err := godbf.NewFromFile(NQHB, "UTF8")
	if err != nil {
		fmt.Println(err)
		goexit()
		return
	}
	//FieldNames := dbfTable.FieldNames()
	//fmt.Println("NQHB:", FieldNames)
	//fmt.Println(HBCJHM, HBZQDM, HBHTXH, HBZQZH, HBCJSL, HBCJJG, HBCJSL2, HBDFDY, HBDFZH, HBCJSJ, HBCJRQ, HBYWLB, HBCDYY, HBBYBZ)
	//fmt.Println("NQHB Count:", dbfTable.NumberOfRecords())

	newrow := dbfTable.AddNewRecord()
	err = dbfTable.SetFieldValueByName(newrow, "HBCJHM", HBCJHM)
	if err != nil {
		fmt.Println(err)
		goexit()
		return
	}
	err = dbfTable.SetFieldValueByName(newrow, "HBZQDM", HBZQDM)
	if err != nil {
		fmt.Println(err)
		goexit()
		return
	}
	err = dbfTable.SetFieldValueByName(newrow, "HBHTXH", HBHTXH)
	if err != nil {
		fmt.Println(err)
		goexit()
		return
	}
	err = dbfTable.SetFieldValueByName(newrow, "HBZQZH", HBZQZH)
	if err != nil {
		fmt.Println(err)
		goexit()
		return
	}
	err = dbfTable.SetFieldValueByName(newrow, "HBCJSL", HBCJSL)
	if err != nil {
		fmt.Println(err)
		goexit()
		return
	}
	err = dbfTable.SetFieldValueByName(newrow, "HBCJJG", HBCJJG)
	if err != nil {
		fmt.Println(err)
		goexit()
		return
	}
	err = dbfTable.SetFieldValueByName(newrow, "HBCJSL2", HBCJSL2)
	if err != nil {
		fmt.Println(err)
		goexit()
		return
	}
	err = dbfTable.SetFieldValueByName(newrow, "HBDFDY", HBDFDY)
	if err != nil {
		fmt.Println(err)
		goexit()
		return
	}
	err = dbfTable.SetFieldValueByName(newrow, "HBDFZH", HBDFZH)
	if err != nil {
		fmt.Println(err)
		goexit()
		return
	}
	err = dbfTable.SetFieldValueByName(newrow, "HBCJSJ", HBCJSJ)
	if err != nil {
		fmt.Println(err)
		goexit()
		return
	}
	err = dbfTable.SetFieldValueByName(newrow, "HBCJRQ", HBCJRQ)
	if err != nil {
		fmt.Println(err)
		goexit()
		return
	}
	err = dbfTable.SetFieldValueByName(newrow, "HBYWLB", HBYWLB)
	if err != nil {
		fmt.Println(err)
		goexit()
		return
	}
	err = dbfTable.SetFieldValueByName(newrow, "HBCDYY", HBCDYY)
	if err != nil {
		fmt.Println(err)
		goexit()
		return
	}
	err = dbfTable.SetFieldValueByName(newrow, "HBBYBZ", HBBYBZ)
	if err != nil {
		fmt.Println(err)
		goexit()
		return
	}
	hbcount = dbfTable.NumberOfRecords()
	err = dbfTable.SaveFilenoLog(NQHB)
	if err != nil {
		fmt.Println(err)
		goexit()
		return
	}
}

//指定长度字符串
func lenstring(src string, length int) string {
	var restr string
	if len(src) == length {
		restr = src
	} else if len(src) > length {
		restr = src[:length]
	} else {
		restr = src
		for i := 0; i < length-len(src); i++ {
			restr = "0" + restr
		}
	}
	return restr
}

//
func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

//
func spinner() {
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("\r%c", r)
			fmt.Printf("\rNQWT.DBF:%d,NQHB.DBF:%d,WTERROR:%d  %c", wtcount, hbcount, errorcount, r)
			time.Sleep(1000000)
		}
	}
}

//
func goexit() {
	fmt.Printf("Press any key to exit...")
	b := make([]byte, 1)
	os.Stdin.Read(b)
}
