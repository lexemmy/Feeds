package exporter

import (
	"errors"
	"fmt"
	"feeds"
	"os"
	"encoding/json"
	"encoding/xml"
	"gopkg.in/yaml.v1"
)

func Exporttxt(u feeds.SocialMedia, filename string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return errors.New("an error occured opening the file: " + err.Error())
	}

	for _, fd := range u.Feed(){
		n, err := f.Write([]byte(fd + "\n"))
		if err != nil  {
			return errors.New("an error occured writing file: " + err.Error())
		}
		fmt.Printf("wrote %d bytes\n", n)
	}
	return nil
}

func Exportjson(u feeds.SocialMedia, filename string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return errors.New("an error occured opening the file: " + err.Error())
	}
	Feed := make(map[int]string)
	for i, fd := range u.Feed(){
		i++
		Feed[i] = fd
	}
		k, _ := json.MarshalIndent(Feed, " ", " ")
		n, err := f.Write([]byte(k))
		if err != nil  {
			return errors.New("an error occured writing to file: " + err.Error())
		}
		fmt.Printf("wrote %d bytes\n", n)
	
	return nil
}
func Exportxml(u feeds.SocialMedia, filename string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return errors.New("an error occured opening the file: " + err.Error())
	}

	 fd := u.Feed()
		k, _ := xml.MarshalIndent(fd, " ", " ")
		n, err := f.Write([]byte(k))
		if err != nil  {
			return errors.New("an error occured writing to file: " + err.Error())
		
		fmt.Printf("wrote %d bytes\n", n)
	}
	return nil
}


func Exportyaml(u feeds.SocialMedia, filename string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return errors.New("an error occured opening the file: " + err.Error())
	}
		k, _ := yaml.Marshal(u.Feed())
		n, err := f.Write([]byte(k))
		if err != nil  {
			return errors.New("an error occured writing to file: " + err.Error())
		}
		fmt.Printf("wrote %d bytes\n", n)
	
	return nil
}


