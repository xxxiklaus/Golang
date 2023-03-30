/*
golang标准库encoding/xml xml包实现xml解析

核心的两个函数
func Marshal(v interface{}) ([]byte, error) /将struct编码成xml,可以接收任意类型

func Unmarshal(data []byte, v interface{}) error /将xml转码成struct结构体

两个核心的结构体

	type Decoder struct {
		/ Strict defaults to true, enforcing the requirements
		/ of the XML specification.
		/ If set to false, the parser allows input containing common
		/ mistakes:
		/	* If an element is missing an end tag, the parser invents
		/	  end tags as necessary to keep the return values from Token
		/	  properly balanced.
		/	* In attribute values and character data, unknown or malformed
		/	  character entities (sequences beginning with &) are left alone.
		/
		/ Setting:
		/
		/	d.Strict = false
		/	d.AutoClose = xml.HTMLAutoClose
		/	d.Entity = xml.HTMLEntity
		/
		/ creates a parser that can handle typical HTML.
		/
		/ Strict mode does not enforce the requirements of the XML name spaces TR.
		/ In particular it does not reject name space tags using undefined prefixes.
		/ Such tags are recorded with the unknown prefix as the name space URL.
		Strict bool

		/ When Strict == false, AutoClose indicates a set of elements to
		/ consider closed immediately after they are opened, regardless
		/ of whether an end element is present.
		AutoClose []string

		/ Entity can be used to map non-standard entity names to string replacements.
		/ The parser behaves as if these standard mappings are present in the map,
		/ regardless of the actual map content:
		/
		/	"lt": "<",
		/	"gt": ">",
		/	"amp": "&",
		/	"apos": "'",
		/	"quot": `"`,
		Entity map[string]string

		/ CharsetReader, if non-nil, defines a function to generate
		/ charset-conversion readers, converting from the provided
		/ non-UTF-8 charset into UTF-8. If CharsetReader is nil or
		/ returns an error, parsing stops with an error. One of the
		/ CharsetReader's result values must be non-nil.
		CharsetReader func(charset string, input io.Reader) (io.Reader, error)

		/ DefaultSpace sets the default name space used for unadorned tags,
		/ as if the entire XML stream were wrapped in an element containing
		/ the attribute xmlns="DefaultSpace".
		DefaultSpace string
		/ contains filtered or unexported fields
	}

	type Encoder struct {
		/ contains filtered or unexported fields
	}
*/
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Person struct {
	XMLName xml.Name `xml:"person"` //反引号
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	Email   string   `xml:"email"`
}

func Marshal() {
	p := Person{
		Name:  "klaus",
		Age:   22,
		Email: "pilw818@gmail.com",
	}

	b, _ := xml.MarshalIndent(p, " ", "  ") //缩进格式
	fmt.Printf("%v\n", string(b))

}

func Unmarshal() {
	b, _ := ioutil.ReadFile("a.xml")
	var p Person

	xml.Unmarshal(b, &p)
	fmt.Printf("p: %v\n", p)
	fmt.Printf("%v\n", string(b))

}

func read() {
	s := `
	<person>
    <name>klaus</name>
	<age>22</age>
	<email>pilw818@gmail.com</email>
	`

	b := []byte(s)

	var per Person

	xml.Unmarshal(b, &per)

	fmt.Printf("per: %v\n", per)
}

func write() {
	p := Person{
		Name:  "klaus",
		Age:   22,
		Email: "pilw818@gmail.com",
	}

	f, _ := os.OpenFile("a.xml", os.O_WRONLY, 0777) //writer
	defer f.Close()
	e := xml.NewEncoder(f)
	e.Encode(p)
}

func main() {
	// Marshal()
	// Unmarshal()
	// read()
	write()
}

/* <?xml version = "1.0" endcoding = "UTF-8"?>
<person>
<name>klaus</name>
<age>22</age>
<email>pilw818@gmail.com</email>
</person> */
