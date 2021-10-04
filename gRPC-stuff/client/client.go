package main

import (
	"context"
	pb "exercise/gRPC-stuff"
	"fmt"

	"google.golang.org/grpc"
)

func main() {
	var opts []grpc.DialOption
	conn, err := grpc.Dial("http://localhost:8080/v1", opts...)
	if err != nil {
		return
	}

	defer conn.Close()
	client := pb.NewCourseServiceClient(conn)

	course, err := client.GetCourse(context.Background(), &pb.CourseID{Id: 1})
	if err != nil {
		fmt.Printf(course.Name)
	}
}
