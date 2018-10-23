package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

// rotate rotates the (x, y) point around origo rot degrees
func rotate(x, y, rot float64) (float64, float64) {
	rad := rot / (2 * math.Pi)
	xt := x*math.Cos(rad) - y*math.Sin(rad)
	yt := x*math.Sin(rad) + y*math.Cos(rad)
	return xt, yt
}

func main() {
	rotation := flag.Int("rotate", 0, "Rotation, degrees (0-360)")
	offsetX := flag.Float64("xoffset", 0, "Offset X-axis (after rotation)")
	offsetY := flag.Float64("yoffset", 0, "Offset Y-axis (after rotation)")
	inputFile := flag.String("input_file", "", "Input file")
	outputFile := flag.String("output_file", "", "Output file")
	columnX := flag.Int("x_column", 4, "Column for X values")
	columnY := flag.Int("y_column", 5, "Column for Y values")
	columnRot := flag.Int("rot_column", 6, "Rotation column")
	flag.Parse()

	if *inputFile == "" {
		panic("Must specify input file")
	}
	if *outputFile == "" {
		panic("Must specify output file")
	}
	if *rotation == 0 && *offsetX == 0.0 && *offsetY == 0.0 {
		panic("Nothing to do")
	}

	// Read the file
	buf, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		panic(err)
	}
	var output []string
	lines := strings.Split(string(buf), "\n")

	for _, v := range lines {
		columns := strings.Split(v, ",")
		if len(columns) < *columnRot {
			output = append(output, v)
			continue
		}
		x, err := strconv.ParseFloat(strings.Replace(columns[*columnX], "\"", "", -1), 64)
		if err != nil {
			output = append(output, v)
			continue
		}
		y, err := strconv.ParseFloat(strings.Replace(columns[*columnY], "\"", "", -1), 64)
		if err != nil {
			panic(err)
		}
		rot, err := strconv.ParseInt(strings.Replace(columns[*columnRot], "\"", "", -1), 10, 32)
		if err != nil {
			panic(err)
		}

		rot += int64(*rotation)
		x, y = rotate(x, y, float64(*rotation))
		x += *offsetX
		y += *offsetY
		columns[*columnX] = fmt.Sprintf("\"%6.2f\"", x)
		columns[*columnY] = fmt.Sprintf("\"%6.2f\"", y)
		columns[*columnRot] = fmt.Sprintf("\"%d\"", rot)

		output = append(output, strings.Join(columns, ","))
	}
	if err := ioutil.WriteFile(*outputFile, []byte(strings.Join(output, "\n")), 0666); err != nil {
		panic(err)
	}
}
