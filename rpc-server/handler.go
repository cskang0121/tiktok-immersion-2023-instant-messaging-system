package main

import (
	"context"
	"github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/kitex_gen/rpc"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"time"
	"strings"
)

type Message struct {
	sender 		string
	timestamp 	int64
	message 	string 
}

// IMServiceImpl implements the last service interface defined in the IDL.
type IMServiceImpl struct{}

func (s *IMServiceImpl) Send(ctx context.Context, req *rpc.SendRequest) (*rpc.SendResponse, error) {

	// Validate send requests 
	if err := validateSendRequest(req); err != nil {
		return nil, err
	}

	// Retrieve roomID, i.e., chat
	roomID := getRoomID(req.Message.GetChat())

	// Initialise connection to database "tiktok"
	db, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/tiktok")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

	// Insert message to database
	queryStatement := fmt.Sprintf("INSERT INTO messages (chat, sender, send_time, message) VALUES ('%v', '%v', %v, '%v');",
		roomID,
		req.Message.GetSender(),
		time.Now().Unix(),
		req.Message.GetText(),
	)
	
	// Set response
	resp := rpc.NewSendResponse()

	insert, err := db.Query(queryStatement)

	if err != nil {
        resp.Code, resp.Msg = 500, "failure in sending message..."
    } else {
		resp.Code, resp.Msg = 0, "success!"
	}
	defer insert.Close()

	return resp, nil
}

func (s *IMServiceImpl) Pull(ctx context.Context, req *rpc.PullRequest) (*rpc.PullResponse, error) {
	// Retrieve roomID, i.e., chat
	roomID := getRoomID(req.GetChat())

	// Retrieve range
	start := req.GetCursor()
	retr := int64(req.GetLimit())

	// Retrieve messages
	messages, hasMore, err := getDataByRoomID(roomID, start, retr, req.GetReverse())
	if err != nil {
		return nil, err
	}
	nextCursor := start + retr

	// Set response
	resp := rpc.NewPullResponse()

	respMessages := make([]*rpc.Message, 0)

	for _, msg := range messages {
		curMsg := &rpc.Message{
			Chat:     req.GetChat(),
			Text:     msg.message,
			Sender:   msg.sender,
			SendTime: msg.timestamp,
		}
		respMessages = append(respMessages, curMsg)
	}

	resp.Messages = respMessages
    resp.Code = 0
    resp.Msg = "success"
    resp.HasMore = &hasMore
    resp.NextCursor = &nextCursor

	return resp, nil
}

// ######################################### Helper Functions #########################################

// validateSendRequest
// 		1. Check if the "Chat" field is set correctly, e.g., "a1:a2"
// 		2. Check if the "Sender" exists in "Chat"
// Return an error: Error if failed the check, else nil
func validateSendRequest(req *rpc.SendRequest) error {

    senders := strings.Split(req.Message.Chat, ":")

    if len(senders) != 2 {
       err := fmt.Errorf("invalid Chat ID '%s', should be in the format of user1:user2", req.Message.GetChat())
       return err
    }
    sender1, sender2 := senders[0], senders[1]

    if req.Message.GetSender() != sender1 && req.Message.GetSender() != sender2 {
       err := fmt.Errorf("sender '%s' not in the chat room", req.Message.GetSender())
       return err
    }

    return nil
}

// getRoomID
// Return a string: A correctly formatted "roomID", i.e., "Chat"
func getRoomID (chat string) string {
	var roomID string

	lowercase := strings.ToLower(chat)
	senders := strings.Split(lowercase, ":")
	
	sender1, sender2 := senders[0], senders[1]
	
	// Compare the sender and receiver alphabetically, and sort them asc to form the room ID
	if comp := strings.Compare(sender1, sender2); comp == 1 {
		roomID = fmt.Sprintf("%s:%s", sender2, sender1)
	} else {
		roomID = fmt.Sprintf("%s:%s", sender1, sender2)
	}
 
	return roomID
}

// getDataByRoomID
// Return 
// 		1. []Messages	: Messages a roomID in the correct order
// 		2. bool 		: False if no more messages to be retrieved, else true
// 		3. error		: Error if any
func getDataByRoomID(roomID string, start int64, retr int64, reverse bool) ([]Message, bool, error) {
	var (
		messages 		[]Message
		order 			string
	)

	// Initialise connection to database "tiktok"
	db, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/tiktok")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

	if reverse { order = "DESC" } else { order = "ASC" }

	// Get messages from MySQL database
	queryStatement := fmt.Sprintf("SELECT sender, send_time, message FROM messages where chat = '%v' ORDER BY send_time %v LIMIT %v, %v;",
		roomID,
		order,
		start,
		retr,
	)	

	getMessages, err := db.Query(queryStatement)
	if err != nil {
        return nil, false, err
    } 
	defer getMessages.Close()

	for getMessages.Next() {
        var curMsg Message
        err := getMessages.Scan(&curMsg.sender, &curMsg.timestamp, &curMsg.message)
        if err != nil {
            return nil, false, err
        }
		messages = append(messages, curMsg)
   } 

	// Check if there is more messages from MySQL database
	queryStatement2 := fmt.Sprintf("SELECT COUNT(id) FROM messages WHERE chat = '%v'", roomID)

	getCount, err := db.Query(queryStatement2)
	if err != nil {
        return nil, false, err
    } 
	defer getCount.Close()

	var count int64
	getCount.Next()
	getCount.Scan(&count)

	hasMore := false
	if count > start + retr {
		hasMore = true
	}

   	return messages, hasMore, nil
}