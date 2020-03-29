package configurations

// Configurations : Represent configurations
type Configurations struct {
	port string
}

// New : Instantiates the configurations
func New() Configurations {
	c := Configurations{
		port: ":3000",
	}
	return c
}

// Port : Getter for port
func (c Configurations) Port() string {
	return c.port
}
