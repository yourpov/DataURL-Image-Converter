# DataURL-Image-Converter

<p align="center">
  <img alt="Top language" src="https://img.shields.io/github/languages/top/yourpov/DataURL-Image-Converter?color=56BEB8">
  <img alt="Language count" src="https://img.shields.io/github/languages/count/yourpov/DataURL-Image-Converter?color=56BEB8">
  <img alt="Repository size" src="https://img.shields.io/github/repo-size/yourpov/DataURL-Image-Converter?color=56BEB8">
  <img alt="License" src="https://img.shields.io/github/license/yourpov/DataURL-Image-Converter?color=56BEB8">
</p>

---

## About

**DataURL-Image-Converter** is a tool for encoding images to base64 data URLs and decoding base64 data URLs back to image files.

| Feature         | Description                                      |
|-----------------|--------------------------------------------------|
| Encode image    | Convert PNG image to base64 data URL             |
| Decode data URL | Convert base64 data URL to PNG image             |

---

## Tech Stack

- [Go](https://go.dev/) (1.18+)

---

## Setup

```bash
# Clone & enter project
 git clone https://github.com/yourpov/DataURL-Image-Converter
 cd DataURL-Image-Converter

# Install deps
 go mod tidy

# Build
 go build
```

---

## Usage

### Encode an image to a data URL

```bash
imgtool.exe encode image.png encoded.txt
```

### Decode a data URL to an image

```bash
imgtool.exe decode encoded.txt out.png
```

---

## Examples

```sh
# Encode an image
imgtool.exe encode image.png encoded.txt

# Decode a data URL (from file)
imgtool.exe decode encoded.txt out.png
```

---

## Troubleshooting

- **"invalid data URL"**
  - Make sure the input is a valid base64 data URL (starts with `data:image/png;base64,`)
- **File not found**
  - Check your input/output file paths

---

## Showcase

<!-- You can add a screenshot or gif here -->
<!-- ![Preview](preview.gif) -->
