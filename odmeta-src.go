package main

import(
	"archive/zip"
	"fmt"
	"io"
	"log"
   "flag"
	"os"
	"aqwari.net/xml/xmltree"
   "encoding/json"
	)

var AppVersion string = "0.1.1 beta"

type XMLMetaTree struct {
	Title 	    string     `xml:"title"`
	Subject	    string     `xml:"subject"`
	Date	       string     `xml:"date"`
	Description  string     `xml:"description"`
	Keyword	    []string   `xml:"keyword"`
   UserDefined  []XMLUserDef   `xml:"user-defined"`
 
}

type XMLUserDef struct {
   Name  string `xml:"name,attr"`
   Value string `xml:",chardata"`
}

func usage() {
       showVersion()
       fmt.Fprintf(os.Stderr, "Usage: %s [-options] [inputfile]\n", os.Args[0])
       fmt.Println("Option list:")
       flag.PrintDefaults()
}

func showVersion() {fmt.Printf("%s version %s\n", os.Args[0], AppVersion)}

func main(){
   version := flag.Bool("v", false, "print current version and exit")
   help := flag.Bool("?",false,"print this help and exit")
   indentOutput := flag.Bool("i", false, "indent JSON output")
   flag.Usage = usage
   flag.Parse()
   if *version {showVersion(); os.Exit(0)}
   if *help {usage(); os.Exit(0)}

   args := flag.Args()
   if len(args) < 1 {fmt.Println("Input file is missing."); os.Exit(1)	}
   filename := flag.Arg(0)
   zipListing, err := zip.OpenReader(filename)
	if err != nil {log.Fatal(err)}

defer zipListing.Close()
for _, file := range zipListing.File {
 	if file.Name != "meta.xml" {continue}

  	content, err := file.Open()
  	if err != nil {log.Fatal(err)}

	bytecontent, err := io.ReadAll(content)
	if err != nil {log.Fatal(err)}
        content.Close()

	root, err := xmltree.Parse(bytecontent)
	if err != nil {log.Fatal(err)}

for _, el := range root.Search("", "meta"){

	   var tree XMLMetaTree
	   if err := xmltree.Unmarshal(el, &tree); err != nil {log.Print(err);	continue }

   var jsonTree []byte

   if *indentOutput {
   jsonTree,err = json.MarshalIndent(tree, "", "  ")
      if err != nil {log.Print(err); continue}
   } else {
   jsonTree, err = json.Marshal(tree) 
   if err != nil {log.Print(err); continue }
   }
   fmt.Println(string(jsonTree))
}

break
	} 
}
