package ghutil

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func mustDo(err error, callback func(args ...interface{}), args ...interface{}) {
	if err != nil {
		if callback != nil {
			if args != nil {
				callback(args...)
			} else {
				callback()
			}
		}
		panic(err)
	}
}
