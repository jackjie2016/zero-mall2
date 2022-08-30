package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	proto "zero-mal/service/user/rpc/user_pb"
)

var userClient proto.UsercenterClient
var conn *grpc.ClientConn

func init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:8001", grpc.WithInsecure())
	if err != nil {
		panic(err)
	} else {
		fmt.Println("链接成功")
	}
	userClient = proto.NewUsercenterClient(conn)
}
func TestGetUserList() {

	r, err := userClient.GetUserById(context.Background(), &proto.GetUserByIdReq{Id: 4})

	if err != nil {
		panic(err)
	}

	fmt.Println(r)
}

func TestGetUserByMobile() {
	r, err := userClient.GetUserByMobile(context.Background(), &proto.MobileRequest{Mobile: "16888888886"})

	if err != nil {
		panic(err)
	}
	fmt.Println(r.NickName, r.Mobile, r.Id)
}
func TestLogin() {
	if rsp, err := userClient.GetUserByMobile(context.Background(), &proto.MobileRequest{
		Mobile: "15958615780",
	}); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(rsp)
	}

}
func TestCreateUser() {

	//for i := 0; i < 9; i++ {
	//	r, err := userClient.CreateUser(context.Background(), &proto.CreateUserInfo{
	//		Mobile:   fmt.Sprintf("1595861578%d", i),
	//		NickName: fmt.Sprintf("zifeng_%d", i),
	//		Password: "zifeng234",
	//	})
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(r.Id)
	//}

}
func main() {
	//TestLogin()
	//TestCreateUser()
	TestGetUserList()
	//TestGetUserByMobile()
	//TestCreateUser()
	conn.Close()
}
