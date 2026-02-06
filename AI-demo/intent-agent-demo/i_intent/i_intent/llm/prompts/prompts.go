package prompts

import (
	"fmt"
	"strings"
	"time"
)

func GetCurrentDate() string {
	return time.Now().Format("2006-01-02")
}

func GetSlotUpdateMessage(sceneName string, dynamicExample string, slotTemplate []string, userInput string) string {
	currentDate := GetCurrentDate()
	//template := "你是一个信息抽取机器人。\n" +
	//	"当前问答场景是：【" + sceneName + "】\n" +
	//	"当前日期是：" + currentDate + "\n\n" +
	//	"JSON中每个元素代表一个参数信息：\n" +
	//	"'''\n" +
	//	"name是参数名称\n" +
	//	"desc是参数注释，可以做为参数信息的补充\n" +
	//	"'''\n\n" +
	//	"需求：\n" +
	//	"#01 根据用户输入内容提取有用的信息到value值，严格提取，没有提及就丢弃该元素，禁止将\"未提及\"写入value\n" +
	//	"#02 返回JSON结果，只需要name和value收到\n\n" +
	//	"返回样例：\n" +
	//	"```\n" +
	//	dynamicExample + "\n" +
	//	"```\n\n" +
	//	"JSON：" + strings.Join(slotTemplate, "\n") + "\n" +
	//	"输入：" + userInput + "\n" +
	//	"答："

	template := fmt.Sprintf(slot_update, sceneName, currentDate, dynamicExample, strings.Join(slotTemplate, "\n"), userInput)
	return template
}

func GetSlotQueryUserMessage(sceneName string, slotTemplate []string, userInput string) string {
	//template := "你是一个专业的客服。\n" +
	//	"当前问答场景是：【" + sceneName + "】\n\n" +
	//	"JSON中每个元素代表一个参数信息：\n" +
	//	"'''\n" +
	//	"name表示参数名称\n" +
	//	"desc表示参数的描述，你要根据描述引导用户补充参数value值\n" +
	//	"'''\n\n" +
	//	"需求：\n" +
	//	"#01 一次最多只向用户问两个参数\n" +
	//	"#02 回答以\"请问\"开头\n\n" +
	//	"JSON：" + strings.Join(slotTemplate, "\n") + "\n" +
	//	"向用户提问："

	template := fmt.Sprintf(slot_query_user, sceneName, strings.Join(slotTemplate, "\n"))
	return template
}

func GetNoSceneResponsePrompt(userInput string, options string) string {
	//template := "你是一个专业的电信客服助手。\n" +
	//	"你可以处理的场景有：\n" +
	//	options + "\n\n" +
	//	"首先请礼貌拒绝用户的要求（如有），并说明这在你能力之外。\n" +
	//	"然后引导用户明确表达他们的需求。\n" +
	//	"用户输入：" + userInput

	template := fmt.Sprintf(no_scene_response, options, userInput)
	return template
}

func GetSceneSwitchDetectionPrompt(currentSceneName string, userInput string) string {
	//template := "你是一个场景意图判断助手。\n" +
	//	"当前正在处理的场景是：【" + currentSceneName + "】\n\n" +
	//	"请判断用户当前的输入是否表达了切换场景的意图。\n" +
	//	"- 如果用户想要切换到其他场景或开始新的任务，回复：1\n" +
	//	"- 如果用户没有明确表达切换场景的意图，回复：0\n\n" +
	//	"用户输入：" + userInput + "\n" +
	//	"请回复（只回复0或1）："

	template := fmt.Sprintf(scene_switch_detection, currentSceneName, userInput)
	return template
}

var slot_update = `你是一个信息抽取机器人。
当前问答场景是：【%s】
当前日期是：%s

JSON中每个元素代表一个参数信息：

'''
name是参数名称
desc是参数注释，可以做为参数信息的补充
'''

需求：
#01 根据用户输入内容提取有用的信息到value值，严格提取，没有提及就丢弃该元素，禁止将“未提及”写入value
#02 返回JSON结果，只需要name和value收到

返回样例：
%s

JSON：%s
输入：%s
答：
`

var slot_query_user = `你是一个专业的客服。
当前问答场景是：【%s】

JSON中每个元素代表一个参数信息：
'''
name表示参数名称
desc表示参数的描述，你要根据描述引导用户补充参数value值
'''

需求：
#01 一次最多只向用户问两个参数
#02 回答以"请问"开头

JSON：%s
向用户提问：`

var no_scene_response = `你是一个专业的电信客服助手。
你可以处理的场景有：
%s

首先请礼貌拒绝用户的要求（如有），并说明这在你能力之外。
然后引导用户明确表达他们的需求。
用户输入：%s`

var scene_switch_detection = `你是一个场景意图判断助手。
当前正在处理的场景是：【%s】

请判断用户当前的输入是否表达了切换场景的意图。
- 如果用户想要切换到其他场景或开始新的任务，回复：1
- 如果用户没有明确表达切换场景的意图，回复：0

用户输入：%s
请回复（只回复0或1）：`
