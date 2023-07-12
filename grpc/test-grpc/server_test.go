package main

import (
	"context"
	"errors"
	"io"
	"log"
	"net"
	"testing"

	"github.com/renatospaka/grpc/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func server(ctx context.Context) (pb.TelephoneClient, func()) {
	buffer := 101024 * 1024
	lis := bufconn.Listen(buffer)

	baseServer := grpc.NewServer()
	pb.RegisterTelephoneServer(baseServer, NewServer())
	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Printf("error serving server :%v\n", err)
		}
	}()

	conn, err := grpc.DialContext(ctx, "",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error connecting to server :%v\n", err)
	}

	closer := func() {
		err := lis.Close()
		if err != nil {
			log.Printf("error closing listener :%v\n", err)
		}
	}

	client := pb.NewTelephoneClient(conn)
	return client, closer
}

func TestTelephoneServer_GetContact(t *testing.T) {
	ctx := context.Background()
	client, closer := server(ctx)
	defer closer()

	type expectation struct {
		out *pb.GetContactReply
		err error
	}

	tests := map[string]struct {
		in       *pb.GetContactRequest
		expected expectation
	}{
		"Must_Success": {
			in: &pb.GetContactRequest{
				Number: "33333333333",
			},
			expected: expectation{
				out: &pb.GetContactReply{
					Name:     "Sebnem",
					Lastname: "Ferah",
					Number:   "33333333333",
				},
				err: nil,
			},
		},
		"Not_Found_Number": {
			in: &pb.GetContactRequest{
				Number: "44444444444",
			},
			expected: expectation{
				out: &pb.GetContactReply{},
				err: errors.New("rpc error: code = Unknown desc = no contact found"),
			},
		},
		"Empty_Number": {
			in: &pb.GetContactRequest{
				Number: "",
			},
			expected: expectation{
				out: &pb.GetContactReply{},
				err: errors.New("rpc error: code = Unknown desc = invalid number"),
			},
		},
		"Invalid_Number": {
			in: &pb.GetContactRequest{
				Number: "test",
			},
			expected: expectation{
				out: &pb.GetContactReply{},
				err: errors.New("rpc error: code = Unknown desc = invalid number"),
			},
		},
		"Short_Number": {
			in: &pb.GetContactRequest{
				Number: "333333333",
			},
			expected: expectation{
				out: &pb.GetContactReply{},
				err: errors.New("rpc error: code = Unknown desc = invalid number"),
			},
		},
		"Long_Number": {
			in: &pb.GetContactRequest{
				Number: "3333333333333",
			},
			expected: expectation{
				out: &pb.GetContactReply{},
				err: errors.New("rpc error: code = Unknown desc = invalid number"),
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			out, err := client.GetContact(ctx, tt.in)
			if err != nil {
				if tt.expected.err.Error() != err.Error() {
					t.Errorf("Err -> \nWant: %q\nGot: %q\n", tt.expected.err, err)
				}
			} else {
				if tt.expected.out.Name != out.Name ||
				tt.expected.out.Number != out.Number ||
				tt.expected.out.Lastname != out.Lastname {
					t.Errorf("Out -> \nWant: %q\nGot : %q", tt.expected.out, out)
				}
			}
		})
	}
}

func TestTelephoneServer_ListContacts(t *testing.T) {
	ctx := context.Background()
	client, closer := server(ctx)
	defer closer()

	type expectation struct {
		out []*pb.ListContactsReply
		err error
	}

	tests := map[string]struct {
		in *pb.ListContactsRequest
		expected expectation
	}{
		"Must_Success": {
			in: &pb.ListContactsRequest{},
			expected: expectation{
				out: []*pb.ListContactsReply{
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
				},
				err: nil,
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			out, err := client.ListContacts(ctx, tt.in)
			var outs []*pb.ListContactsReply

			for {
				o, err := out.Recv()
				if errors.Is(err, io.EOF) {
					break
				}
				outs = append(outs, o)
			}

			if err != nil {
				if tt.expected.err.Error() != err.Error() {
					t.Errorf("Err -> \nWant: %q\nGot: %q\n", tt.expected.err, err)
				}
			} else {
				if len(outs) != len(tt.expected.out) {
					t.Errorf("Out -> \nWant: %q\nGot : %q", tt.expected.out, outs)
				} else {
					for i, o := range outs {
						if o.Name != tt.expected.out[i].Name ||
						o.Lastname != tt.expected.out[i].Lastname ||
						o.Number != tt.expected.out[i].Number {
							t.Errorf("Out -> \nWant: %q\nGot : %q", tt.expected.out, outs)
						}
					}
				}
			}
		})
	}
}
