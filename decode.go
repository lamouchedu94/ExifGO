package decode

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"regexp"
	"time"
)

var regDate = regexp.MustCompile(`\d{4}:\d{2}:\d{2} \d{2}:\d{2}:\d{2}`)
var MissingDate = errors.New("no date")
var MissingOwner = errors.New("no owner")

// pour la date : 01 00 00 00 32
func Image_date(img []byte, ext string) (time.Time, error) {
	//Ext = file extension
	switch ext {
	case ".JPG":
		return Jpg_Date(img)
		img = img[:256]
	case ".CR3":
		return Cr3_Date(img[:1024])
	default:
		img = img[:1024]
	}

	m := regDate.Find(img)
	if m == nil {
		return time.Time{}, errors.New("file not supported")
	}

	date, err := time.Parse("2006:01:02 15:04:05", string(m))
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}

func Jpg_Date(img []byte) (time.Time, error) {
	start := bytes.Index(img, []byte{0x48, 0, 0, 0, 1, 0, 0, 0, 0x48, 0, 0, 0, 1, 0, 0, 0}) + 16
	if start < 0 {
		return time.Time{}, MissingDate
	}
	date, err := time.Parse("2006:01:02 15:04:05", string(img[start:start+19]))
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}

func Cr3_Date(img []byte) (time.Time, error) {
	start := bytes.Index(img, []byte{1, 0, 0, 0, 0x32}) + 4
	if start < 0 {
		return time.Time{}, MissingDate
	}

	date, err := time.Parse("2006:01:02 15:04:05", string(img[start:start+19]))
	if err != nil {
		return time.Time{}, err
	}

	return date, nil
}

func Camera_name(img []byte, ext string) (string, error) {
	switch ext {
	case ".JPG":
		img = img[:256]
	case ".CR3":
		//img = img[:1024]
		s, err := Cr3_Name(img[:1024])
		if err != nil {
			return "", err
		}
		return s, nil
	default:
		img = img[:1024]
	}
	i_model := bytes.Index(img, []byte("Canon E"))
	if i_model == -1 {
		return "", errors.New("file not supported")
	}

	count := 0
	res := ""
	for _, car := range img[i_model:] {
		if 'A' <= car && car <= 'z' || '0' <= car && car <= '9' || car == ':' || car == ' ' {
			res = res + string(car)
			count = 0
		} else {
			count += 1
		}
		if count > 1 {
			break
		}
	}

	return res, nil
}

func Cr3_Name(img []byte) (string, error) {
	start := bytes.Index(img, []byte{1, 0, 0, 0, 0x32}) + 4
	if start < 0 {
		return "", MissingOwner
	}
	end := 0
	for i, val := range img[start+20:] {
		if fmt.Sprintf("%X", val) == "0" {
			end = i
			break
		}
	}
	return string(img[start+20 : start+20+end]), nil
}

func Read_img(img_path string) ([]byte, error) {
	b := make([]byte, 1024)
	f, err := os.Open(img_path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	_, err = f.Read(b)

	return b, err
}
