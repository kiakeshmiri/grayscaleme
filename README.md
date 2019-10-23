# Grayscaleme

This project is a demo about the followings:

- Go 1.13.
- Go Image package.
- Http.
- Image color processing.

## Inital Setup

This project is develped amd tested ussing go 1.13. The follwing link contains the installation instructions:

https://golang.org/doc/install


## Running the application

Simply execure binary file and pass the url of the image (e.g. grayscaleme https://2.bp.blogspot.com/-C3L7c1Y1ck4/UZOk1ldMiBI/AAAAAAAACs4/KtymYa4tM78/s1600/Screenshot+2013-05-15+at+12.38.58+AM.jpg) 

Another alternative is to run the source as following:

go run main.go https://2.bp.blogspot.com/-C3L7c1Y1ck4/UZOk1ldMiBI/AAAAAAAACs4/KtymYa4tM78/s1600/Screenshot+2013-05-15+at+12.38.58+AM.jpg

This application will generate a file on root location with the name gray.png id the source is PNG aor gray.jpg if the source image type is JPEG.

## Build

Generate the binary by typing following command:

go build main.go -o grayscaleme

