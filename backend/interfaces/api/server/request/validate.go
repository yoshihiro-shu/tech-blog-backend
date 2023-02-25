package request

func (c Context) Validate(s interface{}) error {
	err := c.validate.Struct(s)
	if err != nil {
		return err
	}

	return nil
}
