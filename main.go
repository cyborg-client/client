package main
import
(
	// "github.com/cyborg-client/client/analysis"
	// "github.com/cyborg-client/client/datatypes"
	// "github.com/cyborg-client/client/tcphttpclient"

	"github.com/cyborg-client/client/analysis"
	"github.com/cyborg-client/client/buffer"
	"time"
)

func main() {
	// Make channels
	/*
	 tcpDataStreamCh := make(chan datatypes.TcpDataStream)
	 tcpHttpClientStatusCh := make(chan datatypes.TcpHttpClientStatus)
	 clientRequestCh := make(chan datatypes.ClientRequest)

	 go tcphttpclient.TcpHttpClientMain(tcpDataStreamCh, tcpHttpClientStatusCh, clientRequestCh)
	 myReq := datatypes.ClientRequest{Request:datatypes.Start}
	 myReqStop := datatypes.ClientRequest{Request:datatypes.Stop}
	 clientRequestCh<-myReq
	 clientRequestCh<-myReqStop
	 clientRequestCh<-myReq
	 clientRequestCh<-myReqStop
	 select{}
*/
	 //run data parser
	timeStampChannel := make(chan []byte, 100)
	go analysis.Main(timeStampChannel)
	go buffer.Main(timeStampChannel)
	time.Sleep(5000*time.Millisecond)
}