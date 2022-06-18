package helpers

func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
