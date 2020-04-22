package main

import (
	"fmt"
	"feeds"
	"feeds/facebook"
	"feeds/twitter"
	"feeds/linkedin"
	"feeds/exporter"

)

func main()  {
	fb := new(facebook.Facebook)
	twt := new(twitter.Twitter)
	lnkd := new(linkedin.Linkedin)
	err := exporter.Exporttxt(fb,"fb.txt")
	err = exporter.Exporttxt(twt, "twt.txt")
	err = exporter.Exporttxt(lnkd, "lnkd.txt")

	err = exporter.Exportjson(fb,"fb.json")
	err = exporter.Exportjson(twt, "twt.json")
	err = exporter.Exportjson(lnkd, "lnkd.json")

	err = exporter.Exportxml(fb,"fb.xml")
	err = exporter.Exportxml(twt, "twt.xml")
	err = exporter.Exportxml(lnkd, "lnkd.xml")

	err = exporter.Exportyaml(fb,"fb.yaml")
	err = exporter.Exportyaml(twt, "twt.yaml")
	err = exporter.Exportyaml(lnkd, "lnkd.yaml")
	

	if err != nil {
		panic(err)
	}

	//twt := new(twitter.Twitter)
	//lnkd := new(linkedin.Linkedin)
		
	//ScrollFeeds(fb, twt, lnkd)
}

func ScrollFeeds(platforms ...feeds.SocialMedia){
	for _, sm := range platforms{
		for _, fd := range sm.Feed(){
			fmt.Println(fd)
		}
		fmt.Println("===============================")
	}
}




