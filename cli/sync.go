package cli

func Sync() error {

	//if err := GetMyBooksFresh(); err != nil {
	//	return err
	//}
	//if err := ReduceMyBooksFresh(); err != nil {
	//	return err
	//}

	if err := GetMyBooksDetails(); err != nil {
		return err
	}
	if err := ReduceMyBooksDetails(); err != nil {
		return err
	}

	//if err := GetDetailedData(); err != nil {
	//	return err
	//}
	//
	//if err := ReduceDetailedData(start.Unix()); err != nil {
	//	return err
	//}

	return nil
}
