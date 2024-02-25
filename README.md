## odmeta - OpenDocument meta reader

### by Guido Cauli <guido.cauli@gmail.com>

Odmeta is a metadata reader for OpenDocument format written in GO. 
It get the job done by parsing the 'meta.xml' file found into the OpenDocument (odt, odc, odg...)
archive, parsing and outputting it in standard JSON format.

Using the -i option it can output indented JSON.

## Compiling and installing

In order to compile the source, you need the GO compiler. 
You can compile the source with the command:

make

and install into your system with:

sudo make install

You can also uninstall from the system with:

sudo make uninstall

NOTICE: this software is licensed by GNU General Public License v3.0 and is provided WITHOUT ANY WARRANTY.
