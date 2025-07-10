package tp

import (
	"sync"
)

type Results struct {
	a                    [3]float64
	maxPoints            int
	convolution1DResults [][]float64
	originalSignal       []float64
	corruptedSignal      []float64
	filters              [][]float64
}

func NewResults() *Results {
	results := &Results{
		a:               [3]float64{0.1, 0.6, 0.9},
		maxPoints:       200,
		originalSignal:  make([]float64, 200),
		corruptedSignal: make([]float64, 200),
	}
	convolution := make([][]float64, 3)
	for i := range convolution {
		convolution[i] = make([]float64, 200)
	}
	results.convolution1DResults = convolution

	filter := make([][]float64, 3)
	for i := range filter {
		filter[i] = make([]float64, 200)
	}
	results.filters = filter
	return results
}

func (c *Results) GetA() [3]float64 {
	return c.a
}
func (c *Results) GetMaxPoints() int {
	return c.maxPoints
}
func (c *Results) GetOriginalSignal() []float64 {
	return c.originalSignal
}
func (c *Results) GetCorruptedSignal() []float64 {
	return c.corruptedSignal
}
func (c *Results) GetConvolution1DResults() [][]float64 {
	return c.convolution1DResults
}

func (c *Results) GetFilter() [][]float64 {
	return c.filters
}

func (c *Results) SetOriginalSignal(f func(float64) float64) {
	for i := 0; i < c.maxPoints; i++ {
		c.originalSignal[i] = f(float64(i))
	}
}

func (c *Results) SetCorruptedSignal(f func(float64) float64) {
	for i := 0; i < c.maxPoints; i++ {
		c.corruptedSignal[i] = f(float64(i))
	}
}

func (c *Results) SetFilters(f func(float64, float64) float64) {
	for i := range c.filters {
		for j := 0; j < c.maxPoints; j++ {
			c.filters[i][j] = f(float64(j), float64(c.a[i]))
		}
	}

}

func (c *Results) Convolution1D(index int) {
	wg := sync.WaitGroup{}
	filter := c.filters[index]
	signal := c.corruptedSignal
	convolutionResult := c.convolution1DResults[index]

	for i := range convolutionResult {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			var temp float64
			for j := 0; j < c.maxPoints; j++ {
				if n-j < 0 {
					continue
				}
				temp += signal[j] * filter[n-j]
			}
			convolutionResult[n] = temp
		}(i)
	}
	wg.Wait()
}
