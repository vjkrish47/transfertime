import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
const (
	nodes        = 5
	filesPerNode = 5
	baseBW       = 120.0
)
var fileSizes = []int{100, 300, 500, 700, 900}
type Result struct {
	size int
	time time.Duration
}
func networkBandwidth(active int) float64 {
	return baseBW / float64(active)
}
func send(size int, active *int, m *sync.Mutex, wg *sync.WaitGroup, ch chan<- Result) {
	m.Lock()
	*active++
	bw := networkBandwidth(*active)
	m.Unlock()
	start := time.Now()
	latency := rand.Float64()*8 + 2
	transfer := float64(size) / bw
	total := latency + transfer
	time.Sleep(time.Duration(total*10) * time.Millisecond)
	elapsed := time.Since(start)
	m.Lock()
	*active--
	m.Unlock()
	ch <- Result{size: size, time: elapsed}
	wg.Done()
}
func simulate(size int) Result {
	var wg sync.WaitGroup
	var mu sync.Mutex
	active := 0
	results := make(chan Result, nodes*filesPerNode)
	start := time.Now()
	for i := 0; i < nodes; i++ {
		for j := 0; j < filesPerNode; j++ {
			wg.Add(1)
			go send(size, &active, &mu, &wg, results)
		}
	}
	wg.Wait()
	close(results)
	total := time.Since(start)
	return Result{size: size, time: total}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Uncompressed Transfer Simulation")
	var totalTime time.Duration
	for _, size := range fileSizes {
		r := simulate(size)
		fmt.Println("Size:", size, "MB  Time(ms):", r.time.Milliseconds())
		totalTime += r.time
	}
	avg := totalTime / time.Duration(len(fileSizes))
	fmt.Println("Average Time(ms):", avg.Milliseconds())
}
