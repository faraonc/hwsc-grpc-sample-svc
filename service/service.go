package service

import (
	"fmt"
	"golang.org/x/net/context"
	"time"

	pb "github.com/hwsc-org/hwsc-api-blocks/int/hwsc-grpc-sample-svc/proto"
)

// Service implements the service SayHello defined in the proto.
type Service struct{}

// SayHello matches the proto in https://github.com/hwsc-org/hwsc-api-blocks/blob/master/int/hwsc-grpc-sample-svc/proto/hwsc-grpc-sample-svc.proto
func (s Service) SayHello(ctx context.Context, in *pb.SampleServiceRequest) (*pb.SampleServiceResponse, error) {
	fmt.Printf("Request { %v}\n", in)
	fmt.Printf("Receiving UTC date: %s\n", time.Unix(in.GetCreateTimestamp(), 0).UTC())

	now := time.Now().UTC()
	fmt.Printf("Sending response with UTC date: %s\n", now)
	fmt.Printf("Sending response with UTC unix: %d\n", now.Unix())
	response := &pb.SampleServiceResponse{
		Message:           "Hello " + in.FirstName,
		ResponseTimestamp: now.Unix(),
	}
	return response, nil
}
