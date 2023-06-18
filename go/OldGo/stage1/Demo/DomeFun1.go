package main

//你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
//分配规则如下：
//a. 名字中每包含1个'e'或'E'分1枚金币
//b. 名字中每包含1个'i'或'I'分2枚金币
//c. 名字中每包含1个'o'或'O'分3枚金币
//d: 名字中每包含1个'u'或'U'分4枚金币
//写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
import "fmt"

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	// 返回剩余金币数量  // map中 存储 名字对应的分配金币数量
	distribution = make(map[string]int, len(users))
)

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
	fmt.Println(distribution)

}

//  按规则 分金币  返回剩余金币数量
func dispatchCoin() int {
	// 1.依次给给每人分金币(拿到名字)
	// 2. 按照规则分配金币(判断规则)

	for _, v := range users {
		userNum := 0
		for _, name := range v {
			// 取出名的的字母
			switch name {
			case 'e', 'E':
				userNum = userNum + 1
			case 'i', 'I':
				userNum = userNum + 2
			case 'o', 'O':
				userNum = userNum + 3
			case 'u', 'U':
				userNum = userNum + 4
			}
		}
		// 3. 登记每人分配了多少
		distribution[v] = userNum
		// 4. 计算剩余金币
		coins = coins - userNum
	}

	return coins
}
