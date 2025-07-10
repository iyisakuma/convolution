package tp

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"math"
	"math/rand"
)

func s(n float64) float64 { return math.Cos(n*0.08) + 0.3*math.Sin(0.2*n) }
func r() float64          { return (rand.Float64() * 0.6) - 0.3 }
func h(n, a float64) float64 {
	if n < 0 {
		return 0
	}
	return (1 - a) * math.Pow(a, float64(n))
}

func convertToXYs(vect []float64, maxPoints int) plotter.XYs {
	pts := make(plotter.XYs, maxPoints)
	for i := 0; i < maxPoints; i++ {
		pts[i].X = float64(i)
		pts[i].Y = float64(vect[i])
	}
	return pts
}

func Part1() {
	var results = NewResults()
	results.SetOriginalSignal(s)
	results.SetFilters(h)
	results.SetCorruptedSignal(func(n float64) float64 {
		return s(n) + r()
	})
	p := plot.New()
	p.X.Label.Text = "amostras"
	p.Y.Label.Text = "sinal"
	err := plotutil.AddLines(p, "sinal original", convertToXYs(results.GetOriginalSignal(), results.GetMaxPoints()), "sinal com ruido", convertToXYs(results.GetCorruptedSignal(), results.GetMaxPoints()))
	if err != nil {
		panic(err)
	}
	if err := p.Save(8*vg.Inch, 8*vg.Inch, "imagens-sinais.png"); err != nil {
		panic(err)
	}

	results.Convolution1D(0)
	results.Convolution1D(1)
	results.Convolution1D(2)

	ploterConv := plot.New()
	ploterConv.Title.Text = "Sinal após filtro"
	p.X.Label.Text = "amostras"
	p.Y.Label.Text = "sinal filtrado"

	err = plotutil.AddLines(ploterConv, "Sinal pós filtro a: 0.1", convertToXYs(results.GetConvolution1DResults()[0], results.GetMaxPoints()), "sinal com ruido", convertToXYs(results.GetCorruptedSignal(), results.GetMaxPoints()))
	if err != nil {
		panic(err)
	}
	if err := ploterConv.Save(8*vg.Inch, 8*vg.Inch, "imagens-sinal-filtrado00.png"); err != nil {
		panic(err)
	}

	ploterConv = plot.New()
	ploterConv.Title.Text = "Sinal após filtro"
	p.X.Label.Text = "amostras"
	p.Y.Label.Text = "sinal filtrado"
	err = plotutil.AddLines(ploterConv, "Sinal pós filtro a: 0.1", convertToXYs(results.GetConvolution1DResults()[0], results.GetMaxPoints()), "sinal original", convertToXYs(results.GetOriginalSignal(), results.GetMaxPoints()))
	if err != nil {
		panic(err)
	}
	if err := ploterConv.Save(8*vg.Inch, 8*vg.Inch, "imagens-sinal-filtrado-original0.png"); err != nil {
		panic(err)
	}

	ploterConv = plot.New()
	ploterConv.Title.Text = "Sinal após filtro"
	p.X.Label.Text = "amostras"
	p.Y.Label.Text = "sinal filtrado"
	err = plotutil.AddLines(ploterConv, "Sinal pós filtro a: 0.6", convertToXYs(results.GetConvolution1DResults()[1], results.GetMaxPoints()), "sinal com ruido", convertToXYs(results.GetCorruptedSignal(), results.GetMaxPoints()))
	if err != nil {
		panic(err)
	}
	if err := ploterConv.Save(8*vg.Inch, 8*vg.Inch, "imagens-sinal-filtrado05.png"); err != nil {
		panic(err)
	}
	ploterConv = plot.New()
	ploterConv.Title.Text = "Sinal após filtro"
	p.X.Label.Text = "amostras"
	p.Y.Label.Text = "sinal filtrado"
	err = plotutil.AddLines(ploterConv, "Sinal pós filtro a: 0.6", convertToXYs(results.GetConvolution1DResults()[1], results.GetMaxPoints()), "sinal original", convertToXYs(results.GetOriginalSignal(), results.GetMaxPoints()))
	if err != nil {
		panic(err)
	}
	if err := ploterConv.Save(8*vg.Inch, 8*vg.Inch, "imagens-sinal-filtrado-original2.png"); err != nil {
		panic(err)
	}

	ploterConv = plot.New()
	ploterConv.Title.Text = "Sinal após filtro"
	p.X.Label.Text = "amostras"
	p.Y.Label.Text = "sinal filtrado"
	err = plotutil.AddLines(ploterConv, "Sinal pós filtro a: 0.9", convertToXYs(results.GetConvolution1DResults()[2], results.GetMaxPoints()), "sinal com ruido", convertToXYs(results.GetCorruptedSignal(), results.GetMaxPoints()))
	if err != nil {
		panic(err)
	}
	if err := ploterConv.Save(8*vg.Inch, 8*vg.Inch, "imagens-sinal-filtrado09.png"); err != nil {
		panic(err)
	}

	ploterConv = plot.New()
	ploterConv.Title.Text = "Sinal após filtro"
	p.X.Label.Text = "amostras"
	p.Y.Label.Text = "sinal filtrado"
	err = plotutil.AddLines(ploterConv, "Sinal pós filtro a: 0.9", convertToXYs(results.GetConvolution1DResults()[2], results.GetMaxPoints()), "sinal original", convertToXYs(results.GetOriginalSignal(), results.GetMaxPoints()))
	if err != nil {
		panic(err)
	}
	if err := ploterConv.Save(8*vg.Inch, 8*vg.Inch, "imagens-sinal-filtrado-original3.png"); err != nil {
		panic(err)
	}
}
