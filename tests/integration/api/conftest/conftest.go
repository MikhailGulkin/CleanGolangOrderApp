package conftest

import (
	"fmt"
	c "github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/config"
	"log"
	"net"
	"time"
)

func setupBaseURL(apiConfig c.APIConfig) string {
	return fmt.Sprintf("http://%s:%d%s", apiConfig.Host, apiConfig.Port, apiConfig.BaseURLPrefix)
}
func waitForServer(port string) {
	backoff := 50 * time.Millisecond

	for i := 0; i < 10; i++ {
		conn, err := net.DialTimeout("tcp", ":"+port, 1*time.Second)
		if err != nil {
			time.Sleep(backoff)
			continue
		}
		err = conn.Close()
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	log.Fatalf("Server on port %s not up after 10 attempts", port)
}
