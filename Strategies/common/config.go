package common

// func ReadConfig() (File, error) {
// 	jsonFile, err := os.Open("config.json")
// 	if err != nil {
// 		return File{}, nil
// 	}
// 	defer jsonFile.Close()

// 	byteValue, err := ioutil.ReadAll(jsonFile)
// 	if err != nil {
// 		return File{}, err
// 	}

// 	var f File
// 	if err := json.Unmarshal(byteValue, &f); err != nil {
// 		return File{}, err
// 	}

// 	return f, nil
// }

// func WriteConfig(f File) error {
// 	file, err := json.MarshalIndent(f, "", " ")
// 	if err != nil {
// 		return err
// 	}
// 	if err := ioutil.WriteFile("config.json", file, 0644); err != nil {
// 		return err
// 	}
// 	return nil
// }
