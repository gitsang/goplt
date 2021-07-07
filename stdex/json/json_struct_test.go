package main

import (
	"encoding/json"
	"testing"
)

type AData struct {
	A string `json:"a"`
}

type BData struct {
	B string `json:"b"`
}

type Message struct {
	Name string      `json:"name"`
	Id   int         `json:"id"`
	Data interface{} `json:"data"`
}

var msgA = Message{
	Name: "msg_a",
	Id:   1,
	Data: AData{
		A: "a_data",
	},
}

var msgB = Message{
	Name: "msg_b",
	Id:   2,
	Data: BData{
		B: "b_data",
	},
}

func TestJsonStruct(t *testing.T) {
	// marshal

	msgAJ, _ := json.Marshal(msgA)
	msgBJ, _ := json.Marshal(msgB)

	// unmarshal

	msgXA := struct {
		*Message
		Data AData `json:"data"`
	}{}
	_ = json.Unmarshal(msgAJ, &msgXA)
	t.Log("msgXA", msgXA, "data", msgXA.Data.A)

	msgXB := struct {
		*Message
		Data BData `json:"data"`
	}{}
	_ = json.Unmarshal(msgBJ, &msgXB)
	t.Log("msgXB", msgXB, "data", msgXB.Data.B)
}

type ShortMessage struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

var msgAS = ShortMessage{
	Name: "msg_as",
	Id:   1,
}

var dataA = AData{
	A: "a_data",
}

var msgBS = ShortMessage{
	Name: "msg_bs",
	Id:   2,
}

var dataB = BData{
	B: "b_data",
}

func TestJsonStructSplit(t *testing.T) {
	// marshal

	msgAJ, _ := json.Marshal(struct {
		*ShortMessage
		*AData
	}{&msgAS, &dataA})

	msgBJ, _ := json.Marshal(struct {
		*ShortMessage
		*BData
	}{&msgBS, &dataB})

	// unmarshal

	var msgXA ShortMessage
	var dataXA AData
	_ = json.Unmarshal(msgAJ, &struct {
		*ShortMessage
		*AData
	}{&msgXA, &dataXA})
	t.Log("msgXA", msgXA, "data", dataXA.A)

	var msgXB ShortMessage
	var dataXB BData
	_ = json.Unmarshal(msgBJ, &struct {
		*ShortMessage
		*BData
	}{&msgXB, &dataXB})
	t.Log("msgXB", msgXB, "data", dataXB.B)
}

func TestJsonStructFull(t *testing.T) {
	// marshal

	msgAJ, _ := json.Marshal(msgA)
	msgBJ, _ := json.Marshal(msgB)

	// unmarshal

	var msgXA Message
	var dataXA AData
	_ = json.Unmarshal(msgAJ, &struct {
		*Message
		*AData `json:"data"`
	}{&msgXA, &dataXA})
	msgXA.Data = dataXA
	t.Log("msgXA", msgXA, "data", msgXA.Data.(AData).A)

	var msgXB Message
	var dataXB BData
	_ = json.Unmarshal(msgBJ, &struct {
		*Message
		*BData `json:"data"`
	}{&msgXB, &dataXB})
	msgXB.Data = dataXB
	t.Log("msgXB", msgXB, "data", msgXB.Data.(BData).B)
}
