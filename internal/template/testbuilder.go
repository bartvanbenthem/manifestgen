package template

// TestBuilder of type Builder interface for testing
// functions that require a builder interface for input
type TestBuilder struct{}

func (r *TestBuilder) ParseToStdout([]byte, string) error {
	var err error
	return err
}
func (r *TestBuilder) ParseToFile([]byte, string, string) error {
	var err error
	return err
}
