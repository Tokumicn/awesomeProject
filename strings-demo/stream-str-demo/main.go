package main

import "fmt"

func main() {
	checker := &CardChecker{}

	// 测试字符串
	testStr := "【第一个】文本【第二个】"

	// 逐个字符处理
	for i := 0; i < len(testStr); {
		remainingStr := testStr[i:]
		isCountinue, isSendCard, cardIndex := checker.checkNeedSendCard(remainingStr)

		fmt.Printf("位置 %d: isCountinue=%t, isSendCard=%t, cardIndex=%d\n",
			i, isCountinue, isSendCard, cardIndex)

		i += checker.processedChars

		// 如果是发送卡片的情况，下一次调用会自动重置状态
		if isSendCard {
			fmt.Println("↑ 发送了卡片，下次调用会重置isSendCard和isCountinue")
		}
	}
}

type CardChecker struct {
	bracketStack     []int
	currentCardIndex int
	processedChars   int
	hasSentCard      bool // 标记是否已经发送了卡片
}

func (cc *CardChecker) checkNeedSendCard(str string) (isCountinue bool, isSendCard bool, cardIndex int) {
	// 如果上次已经发送了卡片，这次重置发送状态
	if cc.hasSentCard {
		cc.hasSentCard = false
		isCountinue = false
		isSendCard = false
		cardIndex = cc.currentCardIndex - 1 // 使用上一个cardIndex
		return
	}

	for i, char := range str {
		if char == '【' {
			// 遇到【，入栈并记录当前cardIndex
			cc.bracketStack = append(cc.bracketStack, cc.currentCardIndex)
			isCountinue = true
			isSendCard = false
			cardIndex = cc.currentCardIndex
			cc.processedChars = i + 1
			return
		} else if char == '】' {
			// 遇到】，检查是否有匹配的【
			if len(cc.bracketStack) > 0 {
				// 弹出栈顶的【
				matchedIndex := cc.bracketStack[len(cc.bracketStack)-1]
				cc.bracketStack = cc.bracketStack[:len(cc.bracketStack)-1]
				isCountinue = false
				isSendCard = true
				cardIndex = matchedIndex
				cc.currentCardIndex++ // 完成一个完整的【】对，索引加1
				cc.processedChars = i + 1
				cc.hasSentCard = true // 标记已经发送了卡片
				return
			} else {
				// 没有匹配的【，忽略这个】
				continue
			}
		}
	}

	// 处理完所有字符，没有遇到【或】
	cc.processedChars = len(str)
	isCountinue = false
	isSendCard = false
	cardIndex = cc.currentCardIndex
	return
}

// 重置所有状态
func (cc *CardChecker) reset() {
	cc.bracketStack = nil
	cc.currentCardIndex = 0
	cc.processedChars = 0
	cc.hasSentCard = false
}
