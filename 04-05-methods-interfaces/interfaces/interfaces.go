package interfaces

import (
	"fmt"
	"math"
)

// CartesianPoint models a point with X and Y coordinates
type CartesianPoint struct {
	X, Y float64
}

// PolarPoint models a point with radius and polar angle coordinates
type PolarPoint struct {
	R, A float64
}

// XCoord returns the X coordinate
func (p CartesianPoint) XCoord() float64 { return p.X }

// YCoord returns the Y coordinate
func (p CartesianPoint) YCoord() float64 { return p.Y }

// XCoord returns the X coordinate
func (p PolarPoint) XCoord() float64 {
	return p.R * math.Cos(p.A)
}

// YCoord returns the Y coordinate
func (p PolarPoint) YCoord() float64 {
	return p.R * math.Sin(p.A)
}
// Radius returns the polar point radius
func (p PolarPoint) Raduis() float64 {
	return p.R
}

// Angle returns the polar point angle
func (p PolarPoint) Angle() float64 {
	return p.A
}

// Print a cartesian point
func (p CartesianPoint) Print() {
	fmt.Printf("(%f, %f)", p.X, p.Y)
}

// Print a polar point
func (p PolarPoint) Print() {
	fmt.Printf("(R=%f, A=%f◦)", p.R, p.A)
}

// Point interface represents a 2D point
type Point interface {
	OtherPrinter
	Printer
	XCoord() float64
	YCoord() float64
	Print()
}

// Printer interface can print a Point
type Printer interface {
	Print()
}

type OtherPrinter interface {
	Print()
}

// Extend Cartesian point to 3D point
type Point3D struct {
	CartesianPoint
	Z float64
}

// XCoord returns the X coordinate
func (p Point3D) ZCoord() float64 { return p.Z }

// Print a cartesian point
func (p Point3D) Print() {
	fmt.Printf("(%f, %f, %f)", p.X, p.Y, p.ZCoord())
}





