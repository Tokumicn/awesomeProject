package prompts

import "testing"

func TestBasePrompt(t *testing.T) {
	prompt, err := BasePrompt()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(prompt)
}

func TestGoTemplateFormat(t *testing.T) {
	promptStr, err := GoTemplateFormat("", nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(promptStr)
}
