package main

import "fmt"

func main() {
	testStructWrite()
}

/*
结果：
{0 {[{test  [] 0 []  0 0}] 0} }
*/
func testStructWrite() {
	myjs := MyJsonName{
		Code: 0,
		Data: struct {
			Items []struct {
				BindChannels string `json:"bind_channels"`
				Description  string `json:"description"`
				EventID      []int  `json:"event_id"`
				ID           int    `json:"id"`
				Linkage      []struct {
					AddrBookID      []int    `json:"addr_book_id"`
					Description     string   `json:"description"`
					ID              int      `json:"id"`
					LinkageChannels []string `json:"linkage_channels"`
					RecordParam     struct {
						PreRecordingTime int `json:"pre_recording_time"`
						RecordingTime    int `json:"recording_time"`
					} `json:"record_param"`
					Type int `json:"type"`
				} `json:"linkage"`
				Name   string `json:"name"`
				Status int    `json:"status"`
				Type   int    `json:"type"`
			} `json:"items"`
			ItemsTotal int `json:"items_total"`
		}{
			Items: []struct {
				BindChannels string `json:"bind_channels"`
				Description  string `json:"description"`
				EventID      []int  `json:"event_id"`
				ID           int    `json:"id"`
				Linkage      []struct {
					AddrBookID      []int    `json:"addr_book_id"`
					Description     string   `json:"description"`
					ID              int      `json:"id"`
					LinkageChannels []string `json:"linkage_channels"`
					RecordParam     struct {
						PreRecordingTime int `json:"pre_recording_time"`
						RecordingTime    int `json:"recording_time"`
					} `json:"record_param"`
					Type int `json:"type"`
				} `json:"linkage"`
				Name   string `json:"name"`
				Status int    `json:"status"`
				Type   int    `json:"type"`
			}{
				{
					BindChannels: "test",
				},
			},
		},
		Msg: "",
	}
	fmt.Println(myjs)
}

type MyJsonName struct {
	Code int `json:"code"`
	Data struct {
		Items []struct {
			BindChannels string `json:"bind_channels"`
			Description  string `json:"description"`
			EventID      []int  `json:"event_id"`
			ID           int    `json:"id"`
			Linkage      []struct {
				AddrBookID      []int    `json:"addr_book_id"`
				Description     string   `json:"description"`
				ID              int      `json:"id"`
				LinkageChannels []string `json:"linkage_channels"`
				RecordParam     struct {
					PreRecordingTime int `json:"pre_recording_time"`
					RecordingTime    int `json:"recording_time"`
				} `json:"record_param"`
				Type int `json:"type"`
			} `json:"linkage"`
			Name   string `json:"name"`
			Status int    `json:"status"`
			Type   int    `json:"type"`
		} `json:"items"`
		ItemsTotal int `json:"items_total"`
	} `json:"data"`
	Msg string `json:"msg"`
}
