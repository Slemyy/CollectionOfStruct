package dbms

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

type FileInfo struct {
	FileName    string `json:"file_name"`
	FileContent string `json:"file_content"`
}

func SaveFileToDB(dataFile string, structFile string, structure string) error {
	// Создание файла с расширением .json
	dataFileName := dataFile
	if !strings.HasSuffix(dataFileName, ".json") {
		dataFileName += ".json"
	}

	// Загрузка существующих данных из JSON файла, если он существует.
	fileData, err := ioutil.ReadFile(dataFileName)
	var fileInfos []FileInfo

	if err == nil {
		err = json.Unmarshal(fileData, &fileInfos)
		if err != nil {
			return err
		}
	}

	// Загрузка существующих данных из STRUCT файла, если он существует.
	fileContent, err := os.ReadFile(structFile)
	if err != nil {
		return err
	}

	// Инициализируем новый файл, добавляем в него имя и контент.
	newFileInfo := FileInfo{
		FileName:    structFile + structure,
		FileContent: string(fileContent),
	}

	// Проверка на уникальность новых данных в существующем срезе.
	isUnique := true
	for i, existingFileInfo := range fileInfos {
		if newFileInfo.FileName == existingFileInfo.FileName {
			isUnique = false
			// Если данные не уникальны, перезапишем их в существующем срезе.
			fileInfos[i] = newFileInfo
			break
		}
	}

	// Если данные уникальны, добавьте их в существующий срез.
	if isUnique {
		fileInfos = append(fileInfos, newFileInfo)
	}

	// Сериализация обновленных данных в JSON.
	jsonData, err := json.Marshal(fileInfos)
	if err != nil {
		return err
	}

	// Сохранение JSON-данных в файл, перезаписывая существующий файл.
	err = ioutil.WriteFile(dataFileName, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
