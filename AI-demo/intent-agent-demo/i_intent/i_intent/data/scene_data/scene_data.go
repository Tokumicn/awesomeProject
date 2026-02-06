package scene_data

import (
	"encoding/json"
	"i_intent/data/model"
	"os"
)

var SceneTemplateMap = map[string]models.SceneConfig{}

func LoadAllSceneConfigs() {
	filePath := "config/scene_templates.json"
	templates, err := loadSceneTemplates(filePath)
	if err != nil {
		panic(err)
	}
	SceneTemplateMap = templates
}

func loadSceneTemplates(filePath string) (map[string]models.SceneConfig, error) {
	data, err := readFile(filePath)
	if err != nil {
		return nil, err
	}

	var fileData models.SceneTemplateFile
	if err := json.Unmarshal(data, &fileData); err != nil {
		return nil, err
	}

	allSceneConfigs := make(map[string]models.SceneConfig)

	for _, scene := range fileData.SceneList {
		sceneName := scene.SceneName
		if sceneName != "" && allSceneConfigs[sceneName].SceneName == "" {
			sceneParams := make([]models.Slot, len(scene.Parameters))
			copy(sceneParams, scene.Parameters)

			mergedParams := append([]models.Slot{}, fileData.CommonFields...)
			mergedParams = append(mergedParams, sceneParams...)

			allSceneConfigs[sceneName] = models.SceneConfig{
				SceneName:   sceneName,
				Name:        scene.Name,
				Description: scene.Description,
				Parameters:  mergedParams,
				Enabled:     scene.Enabled,
				Example:     scene.Example,
			}
		}
	}

	return allSceneConfigs, nil
}

func readFile(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return data, nil
}
