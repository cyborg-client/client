package main
import
(
	 "github.com/cyborg-client/client/datatypes"
	 "github.com/cyborg-client/client/tcphttpclient"
	"github.com/cyborg-client/client/analysis"
	"github.com/cyborg-client/client/robotserver"
	"fmt"
)

func main() {
	fmt.Print("MAIN")

	// Make channels
	 tcpDataStreamCh := make(chan tcphttpclient.TcpDataStream, 100)
	 tcpHttpClientStatusCh := make(chan tcphttpclient.Status)
	 clientRequestCh := make(chan datatypes.ClientRequest)

	 go tcphttpclient.TcpHttpClient(tcpDataStreamCh, tcpHttpClientStatusCh, clientRequestCh)

	 myReq := datatypes.ClientRequest{Request:tcphttpclient.Start}
	 myReqS := datatypes.ClientRequest{Request:tcphttpclient.Stop}
	 clientRequestCh<-myReqS
	 clientRequestCh<-myReq

	 //select{}

	 //run data parser
	timeStampChannel := make(chan []int64, 100)
	go analysis.Main(timeStampChannel, tcpDataStreamCh)
	go robotserver.Main(timeStampChannel)
	select {}

	}
