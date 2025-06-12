package tools

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/guonaihong/gout"
	"github.com/mark3labs/mcp-go/mcp"
)

var CalendarTool = NewTool(
	"calendar",
	mcp.WithDescription("获取一周每日番剧信息（包含部分核心字段，减轻大模型上下文压力）"),
).WithHandler(func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	var calendar []Schedule
	if err := gout.GET("https://api.bgm.tv/calendar").BindJSON(&calendar).Do(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	bytes, err := json.Marshal(calendar)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return mcp.NewToolResultText(string(bytes)), nil
})

var CalendarFullTool = NewTool(
	"calendar_full",
	mcp.WithDescription("获取一周每日番剧信息（包含全部信息，上下文比较长，可能会导致大模型丢失有用信息）"),
).WithHandler(func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// var calendar &Schedule
	var res string
	if err := gout.GET("https://api.bgm.tv/calendar").BindBody(&res).Do(); err != nil {
		return nil, err
	}

	return mcp.NewToolResultText(res), nil
})

type Weekday struct {
	EN string `json:"en"`
	CN string `json:"cn"`
	JA string `json:"ja"`
	ID int    `json:"id"`
}

type Images struct {
	// Large  string `json:"large"`
	Common string `json:"common"`
	// Medium string `json:"medium"`
	// Small  string `json:"small"`
	// Grid   string `json:"grid"`
}

// type RatingCount struct {
// 	One   int `json:"1"`
// 	Two   int `json:"2"`
// 	Three int `json:"3"`
// 	Four  int `json:"4"`
// 	Five  int `json:"5"`
// 	Six   int `json:"6"`
// 	Seven int `json:"7"`
// 	Eight int `json:"8"`
// 	Nine  int `json:"9"`
// 	Ten   int `json:"10"`
// }

type Rating struct {
	Total int `json:"total"`
	// Count RatingCount `json:"count"`
	Score float64 `json:"score"`
}

// type Collection struct {
// 	Wish    int `json:"wish"`
// 	Collect int `json:"collect"`
// 	Doing   int `json:"doing"`
// 	OnHold  int `json:"on_hold"`
// 	Dropped int `json:"dropped"`
// }

type Item struct {
	ID      int    `json:"id"`
	URL     string `json:"url"`
	Type    int    `json:"type"`
	Name    string `json:"name"`
	NameCN  string `json:"name_cn"`
	Summary string `json:"summary"`
	AirDate string `json:"air_date"`
	// AirWeekday int        `json:"air_weekday"`
	Images   Images `json:"images"`
	Eps      int    `json:"eps"`
	EpsCount int    `json:"eps_count"`
	Rating   Rating `json:"rating"`
	Rank     int    `json:"rank"`
	// Collection Collection `json:"collection"`
}

type Schedule struct {
	Weekday Weekday `json:"weekday"`
	Items   []Item  `json:"items"`
}
