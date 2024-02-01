package main

// CAN IGNORE GRAPHING


import (
	"MapReduce/MultiThreaded"
	"MapReduce/SingleThreaded"
	"time"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

)

func main() {
	numIterations := 10
	singleThreadedDurations := make([]time.Duration, numIterations)
	multiThreadedDurations := make([]time.Duration, numIterations)

	// Run the tests for multiple iterations
	for i := 0; i < numIterations; i++ {
		runSingleThreaded(&singleThreadedDurations[i])
		runMultiThreaded(&multiThreadedDurations[i])
	}

	// Convert durations to float64
	var singleThreadedData, multiThreadedData plotter.XYs
	for i := 0; i < numIterations; i++ {
		singleThreadedData = append(singleThreadedData, plotter.XY{X: float64(i), Y: singleThreadedDurations[i].Seconds() * 1000})
		multiThreadedData = append(multiThreadedData, plotter.XY{X: float64(i), Y: multiThreadedDurations[i].Seconds() * 1000})
	}

	// Plotting
	plotResults(singleThreadedData, multiThreadedData)

}

func runSingleThreaded(duration *time.Duration) {
	startTimeSingle := time.Now()
	files := []string{
		"alien3a.txt", "alien3a.txt", "alien3a.txt", "alien3a.txt",
		"alien3a.txt", "alien3a.txt", "alien3a.txt", "alien3a.txt",
		"alien3a.txt", "alien3a.txt", "alien3a.txt", "alien3a.txt",
		"alien3a.txt", "alien3a.txt", "alien3a.txt", "alien3a.txt",
		"alien3a.txt", "alien3a.txt", "alien3a.txt", "alien3a.txt",
		"alien3a.txt", "alien3a.txt", "alien3a.txt", "alien3a.txt",
		"alien3a.txt", "alien3a.txt", "alien3a.txt", "alien3a.txt",
		"alien3a.txt", "alien3a.txt", "alien3a.txt", "alien3a.txt",
		"alien3a.txt", "alien3a.txt", "alien3a.txt", "alien3a.txt",
		"alien3a.txt", "alien3a.txt", "alien3a.txt", "alien3a.txt",
		"alien3a.txt", "alien3a.txt", "alien3a.txt", "alien3a.txt",
		"alien3a.txt", "alien3a.txt", "alien3a.txt", "alien3a.txt",
		"alien3a.txt", "alien3a.txt", "alien3a.txt", "alien3a.txt",
		"alien3a.txt", "alien3a.txt", "alien3a.txt", "alien3a.txt",
		"alien3a.txt", "alien3a.txt", "alien3a.txt", "alien3a.txt",
		"alien3a.txt", "alien3a.txt", "alien3a.txt", "alien3a.txt",
		"alien3a.txt", "alien3a.txt", "alien3a.txt", "alien3a.txt",
		"alien3a.txt", "alien3a.txt", "alien3a.txt", "alien3a.txt",
	}
	singleThreadedMR := SingleThreaded.NewSingleThreadedMR(files)
	_ = singleThreadedMR.Process()
	*duration = time.Since(startTimeSingle)
}

func runMultiThreaded(duration *time.Duration) {
	numReducers := 4
	startTimeMulti := time.Now()
	files := []string{
		"alien3a.txt", "alien3a.txt", "alien3a.txt", "alien3a.txt",
		"alien3a.txt", "alien3a.txt", "alien3a.txt", "alien3a.txt",
		"alien3a.txt", "alien3a.txt", "alien3a.txt", "alien3a.txt",
		"alien3a.txt", "alien3a.txt", "alien3a.txt", "alien3a.txt",
		"alien3a.txt", "alien3a.txt", "alien3a.txt", "alien3a.txt",
		"alien3a.txt", "alien3a.txt", "alien3a.txt", "alien3a.txt",
		"alien3a.txt", "alien3a.txt", "alien3a.txt", "alien3a.txt",
		"alien3a.txt", "alien3a.txt", "alien3a.txt", "alien3a.txt",
		"alien3a.txt", "alien3a.txt", "alien3a.txt", "alien3a.txt",
		"alien3a.txt", "alien3a.txt", "alien3a.txt", "alien3a.txt",
		"alien3a.txt", "alien3a.txt", "alien3a.txt", "alien3a.txt",
		"alien3a.txt", "alien3a.txt", "alien3a.txt", "alien3a.txt",
		"alien3a.txt", "alien3a.txt", "alien3a.txt", "alien3a.txt",
		"alien3a.txt", "alien3a.txt", "alien3a.txt", "alien3a.txt",
		"alien3a.txt", "alien3a.txt", "alien3a.txt", "alien3a.txt",
		"alien3a.txt", "alien3a.txt", "alien3a.txt", "alien3a.txt",
		"alien3a.txt", "alien3a.txt", "alien3a.txt", "alien3a.txt",
		"alien3a.txt", "alien3a.txt", "alien3a.txt", "alien3a.txt",
	}
	mr := MultiThreaded.NewMultiThreadedMR(numReducers)
	_ = mr.Process(files)
	*duration = time.Since(startTimeMulti)
}

func plotResults(singleThreadedDurations, multiThreadedDurations plotter.XYs) {
	p := plot.New()


	// Create scatter plots for individual iteration times
	singleThreadedScatter, err := plotter.NewScatter(singleThreadedDurations)
	if err != nil {
		panic(err)
	}
	multiThreadedScatter, err := plotter.NewScatter(multiThreadedDurations)
	if err != nil {
		panic(err)
	}

	// Set up the plot
	p.Add(singleThreadedScatter, multiThreadedScatter)
	p.Title.Text = "Individual Iteration Processing Time Comparison"
	p.X.Label.Text = "Iteration"
	p.Y.Label.Text = "Time (ms)"

	// Save the plot to a file or display it
	if err := p.Save(6*vg.Inch, 4*vg.Inch, "chart.png"); err != nil {
		panic(err)
	}
}


