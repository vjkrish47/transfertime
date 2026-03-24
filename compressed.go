import (
	"fmt"
	"math"
	"math/rand"
)

const N = 8

type Block [N][N]float64

func fdct(b Block) Block {
	var out Block
	for u := 0; u < N; u++ {
		for v := 0; v < N; v++ {
			sum := 0.0
			for x := 0; x < N; x++ {
				for y := 0; y < N; y++ {
					sum += b[x][y] *
						math.Cos((2*float64(x)+1)*float64(u)*math.Pi/16) *
						math.Cos((2*float64(y)+1)*float64(v)*math.Pi/16)
				}
			}
			out[u][v] = sum
		}
	}
	return out
}

func idct(b Block) Block {
	var out Block
	for x := 0; x < N; x++ {
		for y := 0; y < N; y++ {
			sum := 0.0
			for u := 0; u < N; u++ {
				for v := 0; v < N; v++ {
					sum += b[u][v] *
						math.Cos((2*float64(x)+1)*float64(u)*math.Pi/16) *
						math.Cos((2*float64(y)+1)*float64(v)*math.Pi/16)
				}
			}
			out[x][y] = sum / 4
		}
	}
	return out
}

func quantize(b Block, q float64) Block {
	var out Block
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			out[i][j] = math.Round(b[i][j] / q)
		}
	}
	return out
}

func dequantize(b Block, q float64) Block {
	var out Block
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			out[i][j] = b[i][j] * q
		}
	}
	return out
}

func entropyEncode(b Block) []float64 {
	data := make([]float64, 0, N*N)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			data = append(data, b[i][j])
		}
	}
	return data
}

func entropyDecode(data []float64) Block {
	var out Block
	k := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			out[i][j] = data[k]
			k++
		}
	}
	return out
}

func randomBlock() Block {
	var b Block
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			b[i][j] = float64(rand.Intn(255))
		}
	}
	return b
}

func main() {
	block := randomBlock()

	dct := fdct(block)
	q := quantize(dct, 10)
	encoded := entropyEncode(q)
	decoded := entropyDecode(encoded)
	dq := dequantize(decoded, 10)
	reconstructed := idct(dq)

	fmt.Println("Original:", block[0][0])
	fmt.Println("Reconstructed:", reconstructed[0][0])
}
