package tracking

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/exec"
	"runtime"
	"runtime/trace"
	"strconv"
	"time"
)

var i = 0

// traceHandler serves the parsed JSON trace data
func TraceHandler(c *gin.Context) {
	traceFile := "trace_files/trace_" + strconv.Itoa(i-1) + ".out"

	data := runtime.ReadTrace()
	fmt.Println(string(data))
	// Check if trace file exists
	if _, err := os.Stat(traceFile); os.IsNotExist(err) {
		c.JSON(404, gin.H{"error": traceFile + " file not found"})
		return
	}

	// Run `go tool trace -json trace.out`
	cmd := exec.Command("go", "tool", "trace", "-json", traceFile)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Failed to run go tool trace: %v", err)})
		return
	}

	// Parse output as JSON
	var traceData map[string]interface{}
	if err := json.Unmarshal(out.Bytes(), &traceData); err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Failed to parse trace JSON: %v", err)})
		return
	}

	// Serve JSON response
	c.JSON(200, traceData)
}

func StartTracer() {
	ticker := time.Tick(5 * time.Second)

	for {
		select {
		case <-ticker:
			func() {
				f, err := os.Create("trace_files/trace_" + strconv.Itoa(i) + ".out")
				if err != nil {
					log.Fatalf("failed to create trace file: %v", err)
				}
				defer f.Close()

				if err := trace.Start(f); err != nil {
					log.Fatalf("failed to start trace: %v", err)
				}
				defer trace.Stop()
				i += 1
			}()

		}
	}

}
