package main

import (
	"context"
	"errors"
	"io"
	"strconv"
	"time"

	"github.com/renatospaka/grpc/grpc/pb"
)

type telephoneServer struct {
	pb.UnimplementedTelephoneServer
	contacts []*pb.ListContactsReply
	calls    []string
}

func NewServer() pb.telephoneServer {
	t := &telephoneServer{}
	t.loadContact()
	return t
}

func (t *telephoneServer) GetContact(ctx context.Context, telNum *pb.GetContactRequest) {
	_, err := strconv.ParseInt(telNum.Number, 10, 64)
	if len(telNum.Number) != 11 || telNum.Number == "" || err != nil {
		return &pb.GetContactRequest{}, errors.New("invalid number")
	}

	for _, c := range t.contacts {
		if telNum.Number == c.Number {
			return &pb.GetContactReply{
				Name: c.Name,
				LastName: c.LastName,
				Number: c.Number,
			}, nil
		}
	}
	return &pb.GetContactReply{}, errors.New("no contact found")
}

func (t *telephoneServer) RecordCallistory(stream pb.Telephone_RecordCallHistoryServer) {
	var callCount int32
	start := time.Now()

	for {
		_, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return stream.SendAndClose(&pb.RecordCallHistoryReply{
					CallCount: callCount,
					ElapsedTime: int32(time.Since(start)),
				})
			}
			return err
		}
		callCount++
	}
}

func (t *telephoneServer) SendMessage(stream pb.Telephone_SendMessageServer) error {
	for {
		rec, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}
	}
	t.calls = append(t.calls, string(rec.Msg))

	if t. calls[len(t.calls) - 1] == "." {
		for _, m := range t.calls {
			switch m {
			case ".":
			case "Hi!":
				stream.Send(&pb.SendMessageReply{
					Msg: []byte("Hello!"),
				})
			case "How are you?":
				stream.Send(&pb.SendMessageReply{
					Msg: []byte("Fine, you?"),
				})
			case "See you later":
				stream.Send(&pb.SendMessageReply{
					Msg: []byte("See you!"),
				})
			default:
				stream.Send(&pb.SendMessageReply{
					Msg: []byte("Sorry, I don't understand :/"),
				})
			}
		}
		t.calls = t.calls[:0]
	}
}


func (t *telephoneServer) loadContact() {
	t.contacts = []*pb.ListContactsReply{
		{
			Name:     "Nukhet",
			Lastname: "Duru",
			Number:   "11111111111",
		},
		{
			Name:     "Zeki",
			Lastname: "Muren",
			Number:   "22222222222",
		},
		{
			Name:     "Sebnem",
			Lastname: "Ferah",
			Number:   "33333333333",
		},
	}
}
