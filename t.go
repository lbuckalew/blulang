package blulang

func() error {
	a := Adapter{}
	if err := a.Init(); err != nil {
		return err
	}
	return nil
}

