package simple

import (
	"sync"

	"google.golang.org/grpc"

	userpb "github.com/ymw0407/GoLang-Studying/protos/v1/user"
)

var (
	once sync.Once
	cli  userpb.UserClient
)

func GetUserClient(serviceHost string) userpb.UserClient {
	once.Do(func() { // sync의 Once.Do 메서드를 사용하여 한 번만 실행하도록 보장함
		conn, _ := grpc.Dial(serviceHost, // grpc dial 함수를 사용하여 연결 생성
			grpc.WithInsecure(), // 보안 연결이 아니라 일반 연결
			grpc.WithBlock(), // 서버와의 연결이 성립될때까지 호출을 차단
		)

		cli = userpb.NewUserClient(conn)
	})

	return cli
}

/*
	이 코드는 동시성을 고려하여 gRPC 클라이언트를 초기화하고 반환하는 패턴을 구현합니다.
	
	GetUserClient 함수를 호출할 때마다 매번 새로운 클라이언트를 생성하는 것이 아니라, 
	한 번만 초기화하고 공유된 클라이언트를 사용하여 효율적으로 리소스를 관리할 수 있습니다.
*/