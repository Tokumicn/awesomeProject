package handler

import (
	"context"
	"i_intent/data/scene_data"
	"i_intent/service"
	"io"
	"net/http"
	"strings"
	"time"

	"i_intent/data/model"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/frame/g"
)

func RegisterRoutes(ctx context.Context, router *gin.Engine) {
	apiPrefix := g.Cfg().MustGet(ctx, "server.APIPrefix").String()

	//router.POST("/multi_question", APIMultiQuestion)
	router.POST(apiPrefix+"/llm_chat", APILLMChat)
	//router.GET(apiPrefix+"/mock_slots", APIMockSlots)
	router.POST(apiPrefix+"/reset_session", APIResetSession)
	router.GET(apiPrefix+"/health", APIHealth)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func APIMultiQuestion(c *gin.Context) {
	var req struct {
		Question string `json:"question" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No question provided"})
		return
	}

	response := service.NewChatbotService(scene_data.SceneTemplateMap).ProcessMultiQuestion(c, req.Question)
	c.JSON(http.StatusOK, gin.H{"answer": response})
}

func APILLMChat(c *gin.Context) {
	var req struct {
		Messages  []models.ChatMessage `json:"messages"`
		UserInput string               `json:"user_input" binding:"required"`
		SessionID string               `json:"session_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No user_input provided"})
		return
	}

	sessionID, session := service.NewSessionManager().GetOrCreate(req.SessionID)

	acceptHeader := c.GetHeader("Accept")
	isStream := strings.Contains(acceptHeader, "text/event-stream")

	if isStream {
		c.Header("Content-Type", "text/event-stream")
		c.Header("Cache-Control", "no-cache")
		c.Header("Connection", "keep-alive")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		c.Header("X-Session-ID", sessionID)

		handleStreamResponse(c, req.UserInput, session, sessionID)
	} else {
		response := service.NewChatbotService(scene_data.SceneTemplateMap).ProcessMultiQuestion(c, req.UserInput)
		c.JSON(http.StatusOK, gin.H{
			"response":   response,
			"session_id": sessionID,
		})
	}
}

func handleStreamResponse(c *gin.Context, userInput string, session *service.Session, sessionID string) {
	response := service.NewChatbotService(scene_data.SceneTemplateMap).ProcessMultiQuestion(c, userInput)

	c.Stream(func(w io.Writer) bool {
		buffer := ""
		punctuation := "。！？，、；："

		for _, char := range response {
			buffer += string(char)

			if len(buffer) >= 3 || contains(punctuation, string(char)) {
				c.SSEvent("message", buffer)
				c.Writer.Flush()
				buffer = ""
				time.Sleep(50 * time.Millisecond)
			}
		}

		if buffer != "" {
			c.SSEvent("message", buffer)
			c.Writer.Flush()
		}

		c.SSEvent("done", "[DONE]")
		return false
	})
}

func contains(s, substr string) bool {
	return len(s) > 0 && len(substr) > 0 && indexOf(s, substr) >= 0
}

func indexOf(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

func APIMockSlots(c *gin.Context) {
	mockData := gin.H{
		"slots": gin.H{
			"phone_number": "13812345678",
			"user_name":    "张三",
			"service_type": "流量套餐",
			"package_type": "月套餐",
		},
		"available_services": []gin.H{
			{"id": 1, "name": "流量套餐", "description": "包月流量服务"},
			{"id": 2, "name": "通话套餐", "description": "包月通话服务"},
			{"id": 3, "name": "短信套餐", "description": "包月短信服务"},
		},
	}
	c.JSON(http.StatusOK, mockData)
}

func APIResetSession(c *gin.Context) {
	var req struct {
		SessionID string `json:"session_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No session_id provided"})
		return
	}

	service.NewSessionManager().Reset(req.SessionID)
	c.JSON(http.StatusOK, gin.H{
		"message":    "Session reset successfully",
		"session_id": req.SessionID,
	})
}

func APIHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":      "healthy",
		"backend_url": g.Cfg().MustGet(c, "server.BackendURL").String(),
		"environment": g.Cfg().MustGet(c, "server.ENV").String(),
	})
}
