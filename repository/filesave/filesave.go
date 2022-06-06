package filesave

type SaveFileResp struct {
	*responseSchema
}

func SaveFile(data []byte) (SaveFileResp, error) {
	resp, err := saveFile(&globalConfig, data)
	if err != nil {
		return SaveFileResp{}, err
	}

	return SaveFileResp{resp}, nil
}

func ReadFile(url string) ([]byte, error) {
	return getFile(&globalConfig, url)
}
