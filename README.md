## odmeta - OpenDocument meta reader

### by Guido Cauli <guido.cauli@gmail.com>
NOTICE: this software is licensed by GNU General Public License v3.0 and is provided WITHOUT ANY WARRANTY.

'odmeta' is a metadata reader for OpenDocument format and is written in GO. 
You can download binaries for your system visiting https://go.dev/dl/ .

It gets the job done by retrieving metadata from 'meta.xml' found into the OpenDocument
(odt, odc, odg...) archive, then parsing and outputting them in standard JSON format.

As of now, 'odmeta' can retrieve all main metadata, such as Title, Subject, Date, Description, the Keyword list, 
and all the custom field that can be specified into the document. It does not retrieve numeric metadata,
like the word count, the page count and whitespace count (just because I actually don't feel the need to retrieve them).

Using the -i option 'odmeta' it can output indented JSON.

## Compile and Install


## Install XML required modules

You will need GO compiler correctly installed onto your system.

First and foremost, you need to get the XML module:
From your console (Windows CLI/PowerShell or Linux console) type:

~~~
go mod init aqwari.net/xml
go mod download aqwari.net/xml
go mod tidy
~~~

## Compile and install

### MS Windows

In order to compile the source in MS Windows, from CLI (or PowerShell) you have to issue the commmand:

~~~
go build odmeta-src.go -o odmeta
~~~

Then you can copy 'odmeta.exe' anywhere you need.
It's recommended to add the folder containing 'odmeta.exe' to your PATH:

~~~
set PATH=%PATH%;C:\add_your_path_here\
~~~

### GNU/Linux

In GNU/Linux you can compile the source with the command:
~~~
make
~~~
and install into your system with:
~~~
sudo make install
~~~
You can also uninstall from the system with:
~~~
sudo make uninstall
~~~
