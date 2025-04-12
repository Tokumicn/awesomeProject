package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
)

func main() {

	//LevelDebug Level = -4
	//LevelInfo  Level = 0
	//LevelWarn  Level = 4
	//LevelError Level = 8

	ctx := context.WithValue(context.Background(), "slog-key", "slog-val")
	//l := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
	//	Level: slog.LevelDebug,
	//}))

	l := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))
	// slog.SetLogLoggerLevel(slog.LevelDebug)
	slog.SetDefault(l)
	fmt.Println(l.Enabled(ctx, slog.LevelDebug))

	type User struct {
		ID   int
		Name string
		Age  int
	}

	users := []User{
		{
			1, "Tom", 18,
		},
		{
			2, "Tom2", 19,
		},
		{
			3, "Tom3", 20,
		},
	}

	// l = l.WithGroup("[TOM]")
	l.DebugContext(ctx, "[DEBUG] message")
	l.InfoContext(ctx, "[INFO] message ", slog.Any("users: ", users))
	l.WarnContext(ctx, "[WARN] message")
	l.ErrorContext(ctx, "[ERROR] message")

	// log包也被替换了
	log.Println("Println message")
	log.Fatalln("Fatalln message")
	log.Panicln("Panicln message")
}
