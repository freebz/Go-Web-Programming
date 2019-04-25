// Listing 7.5  Post struct with comments struct filed

type Post struct {
	XMLName xml.Name   `xml:"post"`
	Id      string     `xml:"id,attr"`
	Content string     `xml:"content"`
	Author  Author     `xml:"author"`
	Xml     string     `xml:",interxml"`
	Comments []Comment `xml:"comments>comment"`
}
