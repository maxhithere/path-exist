func PathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return false
	} else if err == nil {
		return true
	} else {
		return err
	}
}
