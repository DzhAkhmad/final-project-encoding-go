package encoding

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hive-bootcamp/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	// 1 читаем файл, который был создан в папке проекта
	dataJ, err := os.ReadFile(j.FileInput)
	if err != nil {
		fmt.Print(err)
		return err
	}
	// 2 преобразуем json из слайса байтов в данные структуры DockerCompose
	err = json.Unmarshal(dataJ, &j.DockerCompose)
	if err != nil {
		fmt.Print(err)
		return err
	}

	// 3 данные структуры DockerCompose преобразуем в yaml []byte
	dataY, err := yaml.Marshal(j.DockerCompose)
	if err != nil {
		fmt.Print(err)
		return err
	}

	// 4 создаем файл для записи данных
	yamlFile, err := os.Create(j.FileOutput)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer yamlFile.Close()

	// 5 записываем данные в выходной файл
	_, err = yamlFile.Write(dataY)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	// Ниже реализуйте метод
	// 1 читаем файл, который был создан в папке проекта
	dataY, err := os.ReadFile(y.FileInput)
	if err != nil {
		fmt.Print(err)
		return err
	}

	// 2 преобразуем yaml из слайса байтов в данные структуры DockerCompose
	err = yaml.Unmarshal(dataY, &y.DockerCompose)
	if err != nil {
		fmt.Print(err)
		return err
	}

	// 3 данные структуры DockerCompose преобразуем в json []byte
	dataJ, err := json.Marshal(y.DockerCompose)
	if err != nil {
		fmt.Print(err)
		return err
	}

	// 4 создаем файл для записи данных
	jsonFile, err := os.Create(y.FileOutput)
	if err != nil {
		fmt.Print(err)
		return err
	}
	defer jsonFile.Close()

	// 5 записываем данные в выходной файл
	_, err = jsonFile.Write(dataJ)
	if err != nil {
		fmt.Print(err)
		return err
	}
	return nil
}
