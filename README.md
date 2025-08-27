# 🚀 S3Rapthor - AWS S3 Bucket Analyzer

[![Go Version](https://img.shields.io/badge/Go-1.18+-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

> **S3Rapthor** es una herramienta de análisis de buckets de AWS S3 que permite explorar, analizar y descargar contenido de buckets S3 públicos de manera eficiente.

## 🎯 Características

- 🔍 **Análisis de Buckets**: Explora buckets S3 públicos y extrae información detallada
- 📊 **Estadísticas de Archivos**: Genera estadísticas completas sobre tipos de archivos y extensiones
- 💾 **Descarga Masiva**: Descarga automática de todos los archivos del bucket
- 📁 **Organización**: Guarda todos los datos en una estructura organizada
- 🎨 **Interfaz Colorida**: Salida con colores para mejor legibilidad
- ⚡ **Rápido y Eficiente**: Escrito en Go para máximo rendimiento

## 🛠️ Instalación

### Prerrequisitos

- Go 1.18 o superior
- Git

### Instalación desde el repositorio

```bash
# Clonar el repositorio
git clone https://github.com/tu-usuario/s3Rapthor.git
cd s3Rapthor

# Instalar dependencias
go mod tidy

# Compilar la herramienta
go build -o s3Rapthor main.go
```

### Instalación directa con Go

```bash
go install github.com/tu-usuario/s3Rapthor@latest
```

## 📖 Uso

### Uso Básico

```bash
./s3Rapthor <URL_DEL_BUCKET_S3>
```

### Ejemplos

```bash
# Analizar un bucket S3 público
./s3Rapthor https://mi-bucket.s3.amazonaws.com/

# Analizar bucket con contenido específico
./s3Rapthor https://datos-publicos.s3.us-east-1.amazonaws.com/
```

### Parámetros

- `<URL_DEL_BUCKET_S3>`: URL completa del bucket S3 a analizar (requerido)

## 📁 Estructura de Salida

La herramienta crea automáticamente un directorio `data/` con los siguientes archivos:

```
data/
├── nombre-bucket-s3_bucket.txt    # Información básica del bucket
├── nombre-bucket-s3_bucket.xml    # Datos XML raw del bucket
├── nombre-bucket-s3_bucket.brf    # Lista de URLs de archivos
└── archivos_descargados/          # Archivos descargados (si se activa)
```

## 📊 Funcionalidades

### 1. Análisis de Bucket
- Verifica la accesibilidad del bucket
- Extrae metadatos del bucket
- Lista todos los objetos contenidos

### 2. Estadísticas de Archivos
- Cuenta total de archivos
- Análisis por extensiones de archivo
- Porcentajes de cada tipo de archivo
- Destaca archivos potencialmente interesantes (PDF, DOC, etc.)

### 3. Búsqueda de Archivos
```bash
# Buscar archivos por extensión
grep '\.pdf' data/nombre-bucket-s3_bucket.brf
grep '\.doc' data/nombre-bucket-s3_bucket.brf
grep '\.zip' data/nombre-bucket-s3_bucket.brf
```

### 4. Descarga de Archivos
La funcionalidad de descarga está comentada por defecto. Para activarla, descomenta la línea:
```go
download_all_files(brief_file)
```

## 🎨 Salida de Ejemplo

```
🚀 S3Rapthor v0.0.9

·	An AWS S3 bucket analyzer

>> [ 200 ] : https://mi-bucket.s3.amazonaws.com/
>> Bucket Name: mi-bucket
>> Bucket URL: https://mi-bucket.s3.amazonaws.com/
>> Total files: 1250
>> Extension filetypes: 15
>> Extension filetypes expanded:
	.pdf: 45 (3.60%)
	.jpg: 234 (18.72%)
	.png: 156 (12.48%)
	.zip: 23 (1.84%)
	.doc: 12 (0.96%)
	...

>> Type: grep '\.ext' data/mi-bucket-s3_bucket.brf to search for an specific file extension
>> Example: grep '\.pdf' data/mi-bucket-s3_bucket.brf
```

## 🔧 Configuración

### Variables de Entorno (Opcional)

```bash
export AWS_REGION=us-east-1
export AWS_PROFILE=default
```

## 🚨 Consideraciones de Seguridad

- ⚠️ **Solo para buckets públicos**: Esta herramienta está diseñada para analizar buckets S3 públicos
- 🔒 **Respeto por los datos**: Usa esta herramienta de manera ética y responsable
- 📋 **Cumplimiento**: Asegúrate de cumplir con las políticas de uso de los datos
- 🛡️ **Permisos**: Verifica que tienes autorización para acceder a los buckets

## 🤝 Contribuciones

Las contribuciones son bienvenidas. Por favor:

1. Fork el proyecto
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

## 📝 Licencia

Este proyecto está bajo la Licencia MIT. Ver el archivo `LICENSE` para más detalles.

## ⚠️ Disclaimer

Esta herramienta está diseñada únicamente para fines educativos y de investigación. Los usuarios son responsables de cumplir con todas las leyes y regulaciones aplicables al usar esta herramienta.

## 📞 Soporte

Si tienes problemas o sugerencias:

- 📧 Abre un issue en GitHub
- 🐛 Reporta bugs con detalles del problema
- 💡 Sugiere nuevas características

---

**Desarrollado con ❤️ en Go**
