package creational

// BuildProcess defines the steps that are necesary to build a vehicle
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeat() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

//ManufacturingDirector is the one in charge of accepting the builders
type ManufacturingDirector struct {
	builder BuildProcess
}

//Construct uses the builder that is stored
func (f *ManufacturingDirector) Construct() {
	// implementation goes here
	f.builder.SetSeat().SetStructure().SetWheels()
}

// SetBuilder allows to change the builder
func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	// implementation goes here
	f.builder = b
}

type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeat() BuildProcess {
	c.v.Seats = 5
	return c
}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}
