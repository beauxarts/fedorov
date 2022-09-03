package cli

import "time"

func Sync() error {

	start := time.Now()

	//if err := GetMyBooksFresh(); err != nil {
	//	return err
	//}

	if err := ExtractMyBooksIds(); err != nil {
		return err
	}

	if err := GetDetailedData(); err != nil {
		return err
	}

	if err := ReduceDetailedData(start.Unix()); err != nil {
		return err
	}

	return nil
}
