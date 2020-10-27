package blockchain

//var TimeTotals []int64

/*		system given difficulty: how many zeros are in front of the []byte
//[]byte 256位   系统给定的难度:  前多少位数字为为零?       难度与总矿工的数量有关. :目前自己有多少个协程在运行so
与上一个周期矿工挖出的一个平均时间有关			golang don't know how many goroutines his has running right now
//    系统给定的值 [1,0,0,0,0 ....  0,0,0]		我们矿工挖到的[]byte要小于系统给定的这个值

2016个区块出块总用时
*/
//func Minner(sysBytes,block []byte) (int64, error) { //sysBytes 为 系统给的[]byte 返回一个挖矿的时间和一个error
//	timeStart := time.Now().UnixNano()
//	fmt.Println(timeStart)
//	var i int64 = 0
//	for {
//		i++
//		byteI, err := utils.Int64ToByte(i)
//		if err != nil {
//			fmt.Println(err.Error())
//			return -1, err
//		}
//		hashByteI, err := utils.SHA256HashByte(byteI)
//		if err != nil {
//			fmt.Println(err.Error())
//			return -1, err
//		}
//		hashByteI = bytes.Join([][]byte{
//			block,hashByteI,
//		},[]byte{})
//		if CompareBytes(hashByteI, sysBytes) { //当hashByteI < sysBytess 时 返回true
//			timeEnd := time.Now().UnixNano()
//			fmt.Println(timeEnd)
//
//			timeTotal := timeEnd - timeStart
//			//fmt.Println("花费了我",timeTotal,"秒")
//			fmt.Println("我找到这个数学题的解决方案啦! ")
//			return timeTotal, nil
//		}
//		fmt.Printf("第 %v 次寻找hash值\n", i)
//	}
//}
//func GetSySBytes(n int)[]byte { /*判断条件来更改sysBytes的255索引值下标   //
//				*		思路 ：把每个区块创建的时间 存到一个数据库的表中 当 这个表中的数据行数 % 2016 == 0  难度更迭
//	*						上一个2016个区块的挖掘时间 < 预定的时间  难度提升 最多提升4倍  ---->
//				*			上一个2016个区块的挖掘时间 = 预定的时间  难度不变
//				*		 	上一个2016个区块的挖掘时间 > 预定的时间  难度下降 最多下降4倍 ---->
//						总：我们每个区块挖掘的难度是多少？ 即找到平均这个区块要用多长的时间？ 秒还是分钟为单位？
//							我们是多少个区块调整一次难度？  个位？十位百位还是用千位？
//
//	*/
//	var sysBytes = []byte{255}
//	for i := 1; i <= 31; i++ {
//		sysBytes = append(sysBytes, 0)
//	}
//	fmt.Println(sysBytes)
//	fmt.Println(hex.EncodeToString(sysBytes))
//	index1 := bytes.Index(sysBytes, []byte{255})
//	sysBytes[index1], sysBytes[index1+n] = sysBytes[index1+n], sysBytes[index1]
//	fmt.Println("更改索引值后")
//	fmt.Println(sysBytes)
//	fmt.Println(hex.EncodeToString(sysBytes))
//	return sysBytes
//
//}
//func CompareBytes(a, b []byte) bool {
//	if bytes.Compare(a, b) < 0 { //a小于b
//		//fmt.Println("a less than b")
//		return true
//	} else if bytes.Compare(a, b) > 0 { //a大于b
//		//fmt.Println("a greater than b")
//		return false
//
//	} else if bytes.Compare(a, b) == 0 { //a与b相等
//		//fmt.Println("a equals b")
//		return false
//
//	}
//	return false
//}
func TestTime(){
	pow := NewPow(nil,30)
	pow.RunTest(0)//48898039 25级 找4889万次
	//						112092883 27级  找11209万次
	//						178171688   29级    找17817万次
}
/*
//前个零是 1293653  23级 129万次
 */