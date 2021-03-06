package test

import (
	"github.com/qst-project/backend.git/pkg"
	"google.golang.org/grpc/test/bufconn"
)

var lis *bufconn.Listener

const bufSize = 1024 * 1024

func NewTestConfig() pkg.Config {
	return pkg.Config{
		PostgresDsn: "host=localhost port=5432 user=qst password=qst_pwd dbname=qst_db",
	}
}

func init() {
	//app := fx.New(
	//	fx.Provide(pkg.NewLogger()),
	//	fx.Provide(NewTestConfig),
	//	fx.Provide(gateway.NewPostgresClient),
	//	fx.Provide(gateway.Setup),
	//	fx.Provide(api.Setup),
	//	fx.Provide(api.NewGrpcHandler),
	//	fx.Invoke(api.RegisterGrpcServer),
	//)
	lis = bufconn.Listen(bufSize)
	//server := grpc.NewServer()
	//logger := pkg.NewLogger()
	//config, _ := pkg.NewConfig()
	//postgresClient, _ := gateway.NewPostgresClient(config, logger)
	//repos := gateway.Setup(postgresClient)
	//serv := delegate.NewDelegateModule(repos)
	//handler, _ := grpc_handler.NewGrpcHandler(serv)
	//api.RegisterQuestionnaireServer(server, handler)
	//go func() {
	//	if err := server.Serve(lis); err != nil {
	//		log.Printf("%v%v", repos, app)
	//		log.Fatalf("Server exited with error: %v", err)
	//	}
	//}()
}
