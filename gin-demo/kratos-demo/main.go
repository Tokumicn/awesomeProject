package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {

	//3c011db1-7403-4f9a-a2f1-617fb2947813
	//92ccae26-b0b3-4ab4-a3c6-d92af564f6b4
	//40157bbd-8a71-40a9-a4ee-ae86c2a55da2
	//c7c76fac-765e-495a-a394-3cc62ce91f58
	//305c9417-234e-4804-ac95-6dd89dd149a1
	//2491e96c-41e7-4f7b-b0b0-2fa6d1c8b363
	//95371c52-a2d8-41f1-83cc-4ae0455fa31b
	//018e3ac7-5408-47a5-b440-51b32fcafb6a
	//82f53afb-9186-4113-945e-33d365d40291
	//79c6160a-801d-47ee-b722-78d475993326
	for i := 0; i < 10; i++ {
		fmt.Println(uuid.NewString())
	}
}
