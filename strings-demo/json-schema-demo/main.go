package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

func main() {
	// 1. 定义 Schema（也可以放在文件里用 Compile 加载）
	schemaStr := `{
	  "$schema": "https://json-schema.org/draft/2020-12/schema",
	  "type": "object",
	  "properties": {
	    "username": { "type": "string", "minLength": 3, "maxLength": 20, "pattern": "^[a-zA-Z0-9_]+$" },
	    "email":    { "type": "string", "format": "email" },
	    "age":      { "type": "integer", "minimum": 18, "maximum": 120 },
	    "hobbies":  { "type": "array", "items": { "type": "string", "maxLength": 30 }, "uniqueItems": true },
	    "password": { "type": "string", "minLength": 8 }
	  },
	  "required": ["username", "email", "password"],
	  "additionalProperties": false
	}`

	// 2. 编译 Schema
	complile := jsonschema.NewCompiler()
	sch, err := complile.Compile(schemaStr)
	//sch, err := jsonschema.CompileString("user.json", schemaStr)
	if err != nil {
		log.Fatalf("Schema 编译失败: %#v", err)
	}

	// 3. 测试合法数据
	validData := `{
	  "username": "alice_123",
	  "email": "alice@example.com",
	  "age": 25,
	  "hobbies": ["reading", "coding"],
	  "password": "securePass1"
	}`
	var validInstance interface{}
	json.Unmarshal([]byte(validData), &validInstance)
	err = sch.Validate(validInstance)
	if err != nil {
		fmt.Println("✅ 合法数据 -> 不应报错，但报了:", err)
	} else {
		fmt.Println("✅ 合法数据 -> 校验通过")
	}

	// 4. 测试非法数据
	invalidData := `{
	  "username": "ab",
	  "email": "not-an-email",
	  "age": 150,
	  "hobbies": ["a very long hobby name that exceeds thirty characters!!!"],
	  "extraField": "should not exist",
	  "password": "short"
	}`
	var invalidInstance interface{}
	json.Unmarshal([]byte(invalidData), &invalidInstance)
	err = sch.Validate(invalidInstance)
	if err != nil {
		fmt.Println("\n❌ 非法数据 -> 校验失败，错误详情:")
		// 打印所有校验错误
		if ve, ok := err.(*jsonschema.ValidationError); ok {
			for _, ie := range ve.BasicOutput().Errors {
				fmt.Printf("  - 字段: %s, 错误: %s\n", ie.InstanceLocation, ie.Error)
			}
		} else {
			fmt.Printf("  %#v\n", err)
		}
	} else {
		fmt.Println("❌ 非法数据 -> 本应报错却通过了")
	}
}
