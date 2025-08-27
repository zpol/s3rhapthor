package main

import (
	"bufio"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type ListBucketResult struct {
	XMLName     xml.Name `xml:"ListBucketResult"`
	Name        string   `xml:"Name"`
	Prefix      string   `xml:"Prefix"`
	Marker      string   `xml:"Marker"`
	MaxKeys     int      `xml:"MaxKeys"`
	IsTruncated bool     `xml:"IsTruncated"`
	Contents    []struct {
		Key          string `xml:"Key"`
		LastModified string `xml:"LastModified"`
		ETag         string `xml:"ETag"`
		Size         int    `xml:"Size"`
		StorageClass string `xml:"StorageClass"`
	} `xml:"Contents"`
}

func banner() {
	st := "ICAgICAgICDioo7ioZEg4qKJ4qG5ICAgIOKjj+KhsSDiooDio4Ag4qOA4qGAIOKjsOKhgCDio4fioYAg4qKA4qGAIOKhgOKjgCAgICAgICAKICAg4qCJ4qCJICAg4qCi4qCcIOKgpOKgnCDioInioIkg4qCH4qCxIOKgo+KgvCDioafioJwg4qCY4qCkIOKgh+KguCDioKPioJwg4qCPICAgIOKgieKgiSAgCg=="
	decoded, err := base64.StdEncoding.DecodeString(st)
	if err != nil {
		panic(err)
	}
	fmt.Println("\n" + string(decoded) + "\t v0.0.9\n\n·\tAn AWS S3 bucket analyzer\n\n")
}

func get_data(bucket_url string) string {
	url := bucket_url
	data, err := getResponseData(url)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func getResponseData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func r_ok(msg string) {
	fmt.Printf("\033[1;32m" + msg + "\033[0m\n")
}

func r_warn(msg string) {
	fmt.Printf("\033[1;33m" + msg + "\033[0m\n")
}

func r_err(msg string) {
	fmt.Printf("\033[1;31m" + msg + "\033[0m\n")
}

func check_return_code(url string) string {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
	}
	switch resp.StatusCode {
	case 200:
		r_ok(">> [ 200 ] : " + url)
	case 404:
		r_warn(">> [ 404 Not found ] : " + url)
	case 403:
		r_err(">> [ 403 Forbidden ] : " + url)
	default:
		fmt.Println("Unexpected Status")
	}
	return ">> Status code: " + fmt.Sprint(resp.StatusCode) + "\n"
}

func save_file(content string, filename string) {
	// Create the directory if it doesn't exist
	if _, err := os.Stat("data/"); os.IsNotExist(err) {
		os.Mkdir("data/", 0755)
	}

	// Create or open the file
	f, err := os.OpenFile("data/"+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	_, err = f.Write([]byte(content))
	if err != nil {
		fmt.Println(err)
		return
	}
	err = f.Sync()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf(">> File saved at %s\n", f.Name())
}

func parse_xml(file string, bucket_url string) []string {

	xmlFile, err := os.Open("data/" + file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer xmlFile.Close()

	b, _ := ioutil.ReadAll(xmlFile)

	var q ListBucketResult
	xml.Unmarshal(b, &q)
	fmt.Println(">> Bucket Name: ", q.Name)
	fmt.Println(">> Bucket URL: " + bucket_url)

	var objects []string

	for _, v := range q.Contents {
		obj_name := bucket_url + strings.Replace(v.Key, " ", "%20", -1)
		objects = append(objects, obj_name)
	}

	return objects
}

func download_all_files(file string) {
	// Abrir el archivo de texto
	archivo, err := os.Open("data/" + file)
	if err != nil {
		panic(err)
	}
	defer archivo.Close()

	// Crear un escáner para leer el archivo línea por línea
	escaner := bufio.NewScanner(archivo)

	// Iterar sobre cada línea del archivo
	for escaner.Scan() {
		// Obtener la URL de la línea actual
		url := strings.TrimSpace(escaner.Text())

		// Descargar el archivo correspondiente
		fmt.Printf(">> DOwnloading %s...\n", url)
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		// Guardar el contenido del archivo descargado en un archivo local
		nombreArchivo := obtenerNombreArchivo(url)
		archivoLocal, err := os.Create("data/" + nombreArchivo)
		if err != nil {
			panic(err)
		}
		defer archivoLocal.Close()

		_, err = io.Copy(archivoLocal, resp.Body)
		if err != nil {
			panic(err)
		}

		// Imprimir un mensaje de éxito
		fmt.Printf(">> [OK] - Download [ %s ] completed.\n", url)
	}
}

func obtenerNombreArchivo(url string) string {
	// Obtener el nombre del archivo de la URL
	ultimoSlash := len(url) - 1
	for i := len(url) - 1; i >= 0; i-- {
		if url[i] == '/' {
			ultimoSlash = i
			break
		}
	}
	nombreArchivo := url[ultimoSlash+1:]

	// Reemplazar los caracteres especiales con guiones bajos
	for i := 0; i < len(nombreArchivo); i++ {
		if nombreArchivo[i] == '%' {
			nombreArchivo = nombreArchivo[:i] + "_" + nombreArchivo[i+3:]
		}
	}

	return nombreArchivo
}

func summary(data_file string) {
	// Abrir el archivo de texto con las URLs
	file, err := os.Open(data_file)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	// Leer las URLs del archivo y guardarlas en una lista
	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}

	// Contadores para las tareas solicitadas
	numFiles := 0
	extCounts := make(map[string]int)

	// Procesar cada URL
	for _, url := range urls {
		// Ignorar las URLs que no tienen una extensión al final
		ultimoPunto := strings.LastIndex(url, ".")
		if ultimoPunto == -1 || ultimoPunto == len(url)-1 {
			continue
		}

		if url[len(url)-1] == '/' {
			continue
		}

		// Contar el número de ficheros y las extensiones
		ext := url[ultimoPunto+1:]

		// Descartar las extensiones con más de 10 caracteres
		if len(ext) > 10 {
			continue
		}

		numFiles++
		extCounts[ext]++
	}

	// Mostrar los resultados
	fmt.Println(">> Total files: ", numFiles)

	// Contar el número de tipos de extensiones diferentes
	numExts := len(extCounts)
	fmt.Println(">> Extension filetypes: ", numExts)

	// Mostrar el número de ficheros de cada extensión y el porcentaje sobre el total
	fmt.Println(">> Extension filetypes expanded: ")
	for ext, count := range extCounts {
		percent := float64(count) / float64(numFiles) * 100
		switch ext {

		case "bak":
			fmt.Printf("\033[1;33m"+"\t.%s: %d (%.2f%%)\n"+"\033[0m", ext, count, percent)

		case "doc":
			fmt.Printf("\033[1;33m"+"\t.%s: %d (%.2f%%)\n"+"\033[0m", ext, count, percent)

		case "docx":
			fmt.Printf("\033[1;33m"+"\t.%s: %d (%.2f%%)\n"+"\033[0m", ext, count, percent)

		case "gzip":
			fmt.Printf("\033[1;33m"+"\t.%s: %d (%.2f%%)\n"+"\033[0m", ext, count, percent)

		case "iso":
			fmt.Printf("\033[1;33m"+"\t.%s: %d (%.2f%%)\n"+"\033[0m", ext, count, percent)

		case "json":
			fmt.Printf("\033[1;33m"+"\t.%s: %d (%.2f%%)\n"+"\033[0m", ext, count, percent)

		case "ova":
			fmt.Printf("\033[1;33m"+"\t.%s: %d (%.2f%%)\n"+"\033[0m", ext, count, percent)

		case "pdf":
			fmt.Printf("\033[1;33m"+"\t.%s: %d (%.2f%%)\n"+"\033[0m", ext, count, percent)

		case "php":
			fmt.Printf("\033[1;33m"+"\t.%s: %d (%.2f%%)\n"+"\033[0m", ext, count, percent)

		case "sh":
			fmt.Printf("\033[1;33m"+"\t.%s: %d (%.2f%%)\n"+"\033[0m", ext, count, percent)

		case "tgz":
			fmt.Printf("\033[1;33m"+"\t.%s: %d (%.2f%%)\n"+"\033[0m", ext, count, percent)

		case "txt":
			fmt.Printf("\033[1;33m"+"\t.%s: %d (%.2f%%)\n"+"\033[0m", ext, count, percent)

		case "vmdk":
			fmt.Printf("\033[1;33m"+"\t.%s: %d (%.2f%%)\n"+"\033[0m", ext, count, percent)

		case "xls":
			fmt.Printf("\033[1;33m"+"\t.%s: %d (%.2f%%)\n"+"\033[0m", ext, count, percent)

		case "xlsx":
			fmt.Printf("\033[1;33m"+"\t.%s: %d (%.2f%%)\n"+"\033[0m", ext, count, percent)

		case "yml":
			fmt.Printf("\033[1;33m"+"\t.%s: %d (%.2f%%)\n"+"\033[0m", ext, count, percent)

		case "zip":
			fmt.Printf("\033[1;33m"+"\t.%s: %d (%.2f%%)\n"+"\033[0m", ext, count, percent)

		default:
			fmt.Printf("\t.%s: %d (%.2f%%)\n", ext, count, percent)

		}
	}
}

func main() {
	banner()
	if len(os.Args) < 1 {
		panic(">> [ERROR]: Missing argument <s3_bucket>")
	} else {

		var s3_bucket = os.Args[1]
		// check if slash at the end
		if s3_bucket[len(s3_bucket)-1:] != "/" {
			// if not print it
			s3_bucket += "/"
		}
		re := regexp.MustCompile(`.*\/\/(.*?)\.`)

		data_file := re.FindStringSubmatch(s3_bucket)[1] + "-s3_bucket.txt"
		tmp_raw_data_file := re.FindStringSubmatch(s3_bucket)[1] + "-s3_bucket.xml"
		brief_file := re.FindStringSubmatch(s3_bucket)[1] + "-s3_bucket.brf"

		save_file("Bucket name: "+data_file+"\n", data_file)

		o := check_return_code(s3_bucket)
		save_file(o, data_file)

		data := get_data(s3_bucket)
		save_file(data, tmp_raw_data_file)

		objects := parse_xml(tmp_raw_data_file, s3_bucket)
		content := strings.Join(objects, "\n")
		save_file(content, brief_file)

		// extract_filetypes("data/" + brief_file)

		summary("data/" + brief_file)

		fmt.Println("\n>> Type: grep '\\.ext' data/" + brief_file + " to search for an specific file extension")
		fmt.Println(">> Example: grep '\\.pdf' data/" + brief_file + "\n\n")

		// download_all_files(brief_file)
	}
}
