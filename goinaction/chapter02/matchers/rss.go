package matchers

import "encoding/xml"

type (
	//item根据item字段的标签，将定义的字段与rss文档的字段关联起来
	item struct {
		XMLName      	xml.Name	`xml:"item"`
		PubDate      	string   	`xml:"pubDate"`
		Title        	string   	`xml:"title"`
		Description  	string   	`xml:"description"`
		Link		 	string   	`xml:"link"`
		GUID		 	string   	`xml:"guid"`
		GeoRssPoint  	string	  	`xml:"georss:point"`
	}

	//image根据image字段的标签，将定义的字段与rss文档关联起来
	//与rss文档的字段关联起来
	image struct {
		XMLName      	xml.Name		`xml:"image"`
		URL			 	string			`xml:"url"`
		Title		 	string			`xml:"title"`
		Link		 	string			`xml:"link"`
	}

	//channel根据channel字段的标签，将定义的字段与rss文档关联起来
	channel struct {
		XMLName			xml.Name		`xml:"channel"`
		Title			string			`xml:"title"`
		Description 	string			`xml:"description"`
		Link			string			`xml:"link"`
		PubDate			string			`xml:"pubDate"`
		LastBuildDate 	string			`xml:"lastBuildDate"`
		TTL				string			`xml:"ttl"`
		Language		string			`xml:"language"`
		ManagingEditor	string			`xml:"managingEditor"`
		WebMaster		string			`xml:"webMaster"`
		Image			string			`xml:"image"`
		Item			[]item			`xml:"item"`
	}
	//rssDocument定义了与rss文档关联的字段
	rssDocument struct {
		XMLName 		xml.Name		`xml:"rss"`
		Channel			channel			`xml:"channel"`
	}
)

