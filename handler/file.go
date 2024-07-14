package filesHandler

import "strconv"

type File struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func files() []File {
	return []File{
		{
			ID:   1,
			Name: "asd",
			Type: ".jpg",
		},
		{
			ID:   2,
			Name: "dfg",
			Type: ".png",
		},
		{
			ID:   3,
			Name: "ghr",
			Type: ".pdf",
		},
	}
}

func getFiles() map[string]File {
	files := files()
	response := make(map[string]File, len(files))

	for _, x := range files {
		response[strconv.Itoa(int(x.ID))] = x
	}

	return response
}
