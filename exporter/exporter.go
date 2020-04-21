package exporter

import (
	"errors"
	"fmt"
	"feeds"
	"os"
	"encoding/json"
	"encoding/xml"
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

	for i, fd := range u.Feed(){
		l := make(map[int]string)
		i++
		l[i] = fd

		k, _ := json.Marshal(l)
		n, err := f.Write([]byte(k))
		if err != nil  {
			return errors.New("an error occured writing to file: " + err.Error())
		}
		fmt.Printf("wrote %d bytes\n", n)
	}
	return nil
}
func Exportxml(u feeds.SocialMedia, filename string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return errors.New("an error occured opening the file: " + err.Error())
	}

	for _, fd := range u.Feed(){
		k, _ := xml.MarshalIndent(fd, " ", " ")
		n, err := f.Write([]byte(k))
		if err != nil  {
			return errors.New("an error occured writing to file: " + err.Error())
		}
		fmt.Printf("wrote %d bytes\n", n)
	}
	return nil
}


