package models

type SceneConfig struct {
	SceneName   string `json:"scene_name"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Parameters  []Slot `json:"parameters"`
	Enabled     bool   `json:"enabled"`
	Example     string `json:"example"`
}

type SceneTemplates map[string]SceneConfig

type SceneListItem struct {
	SceneName   string `json:"scene_name"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Parameters  []Slot `json:"parameters"`
	Enabled     bool   `json:"enabled"`
	Example     string `json:"example"`
}

type SceneTemplateFile struct {
	CommonFields []Slot          `json:"common_fields"`
	SceneList    []SceneListItem `json:"scene_list"`
}

func GetDynamicExample(sceneConfig SceneConfig) string {
	if sceneConfig.Example != "" {
		return sceneConfig.Example
	}
	return "JSON：[{'name': 'phone', 'desc': '需要查询的手机号', 'value': ''}, {'name': 'month', 'desc': '查询的月份，格式为yyyy-MM', 'value': ''} ]\n输入：帮我查一下18724011022在2024年7月的流量\n答：{ 'phone': '18724011022', 'month': '2024-07' }"
}
