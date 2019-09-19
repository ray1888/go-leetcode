package DynamicProgramming

//func minCut(s string) int {
//
//	if s == ""{
//		return 0
//	}
//	length := len(s)
//	//dp here meas d[i][j] mincut number
//	dp := make([][]int, length*length)
//	singel_dis := make([]int, length*length)
//	for i:=0;i<length;i++{
//		dp[i], singel_dis = singel_dis[:length], singel_dis[length:]
//	}
//
//	for i:=length-1;i>=0;i--{
//		for j:=i;j>=0;j--{
//			dpVal := 0
//			if isparlindrome(s[j:]) {
//				dpVal = 1
//			}
//			dp[i][j] = min(dp[i][length-j-1] + dpVal, dp[i][j])
//		}
//	}
//	return dp[0][length-1]
//}
