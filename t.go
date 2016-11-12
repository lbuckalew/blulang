package blulang

func main() error {
	a := Adapter{}
	if err := a.Init(); err != nil {
		return err
	}
	return nil
}
