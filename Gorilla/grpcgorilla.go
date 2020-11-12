package main
import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	grpc "github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"log"
	"net/http"
	"net/rpc"
	"time"
)
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
type RpcStringRequest struct {
	String string
}
type RpcStringResult struct {
	Count int
}
type RpcStringService struct{}
func (t *RpcStringService) Count(request *RpcStringRequest, result *RpcStringResult) error {
	result.Count = len(request.String)
	return nil
}
type GorillaStringService struct{}
type GorillaStringRequest struct {
	String string `json:"string"`
}
type GorillaStringResponse struct {
	Length int `json:"length"`
}
func (h *GorillaStringService) Length(r *http.Request, args *GorillaStringRequest, reply *GorillaStringResponse) error {
	reply.Length = len(args.String)
	return nil
}
func RpcCall(url, method string, request, reply interface{}) error {
	params, err := json.EncodeClientRequest(method, request)
	if err != nil {
		return err
	}
	response, err := http.Post(url, "application/json", bytes.NewReader(params))
	if err != nil {
		return err
	}
	defer response.Body.Close()
	err = json.DecodeClientResponse(response.Body, reply)
	return err
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HelloHandler)
	r.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("/tmp"))))
	rpcServer := rpc.NewServer()
	rpcServer.Register(new(RpcStringService))
	r.Handle("/rpc", rpcServer)
	muxServer := grpc.NewServer()
	muxServer.RegisterCodec(json.NewCodec(), "application/json")
	muxServer.RegisterService(new(GorillaStringService), "")
	r.Handle("/grpc", muxServer)
	fmt.Println("Listening on 8080")
	go http.ListenAndServe(":8080", r)
	time.Sleep(1 * time.Second)
	client1, err := rpc.DialHTTPPath("tcp", "localhost:8080", "/rpc")
	if err != nil {
		log.Fatal(err)
	}
	args1 := &RpcStringRequest{"hello"}
	reply1 := new(RpcStringResult)
	fmt.Println(client1)
	fmt.Println(args1)
	fmt.Println(reply1)
	client1.Call("StringService.Count", args1, &reply1)
	fmt.Println(string(reply1.Count))
	defer client1.Close()
	args2 := GorillaStringRequest{"hello"}
	reply2 := &GorillaStringResponse{}
	err = RpcCall("http://localhost:8080/grpc", "GorillaStringService.Length", args2, reply2)
	if err != nil {
		time.Sleep(1 * time.Minute)
		log.Fatal(err)
	}
	log.Println(args2)
	log.Println(reply2)
}