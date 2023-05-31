package decode

import (
	"bytes"
	"errors"
	"os"
	"time"
)

func Image_date(img []byte, ext string) (time.Time, error) {
	//
	switch ext {
	case ".JPG":
		img = img[:256]
	case ".CR3":
		img = img[:1024]
	default:
		img = img[:1024]
	}
	i_date := bytes.Index(img, []byte(":"))
	if i_date == -1 {
		return time.Time{}, errors.New("File Not supported.")
	}

	datestr := ""
	for _, car := range img[i_date-4 : i_date+15] {
		datestr += string(car)
	}
	date, err := time.Parse("2006:01:02 15:04:05", datestr)
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
		img = img[:1024]
	default:
		img = img[:1024]
	}
	i_model := bytes.Index(img, []byte("Canon E"))
	if i_model == -1 {
		return "", errors.New("File not supported.")
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

func Read_img(img_path string) ([]byte, error) {
	b, err := os.ReadFile(img_path)
	if err != nil {
		return nil, err
	}
	return b, nil
}