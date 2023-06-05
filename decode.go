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

// pour la date : 01 00 00 00 32
func Image_date(img []byte, ext string) (time.Time, error) {
	//Ext = file extention
	switch ext {
	case ".JPG":
		img = img[:256]
	case ".CR3":
		img = img[:256]
		//Cr3(img)
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
func Cr3(img []byte) error {
	//a := fmt.Sprintf("%X", img[4])
	motif := [5]int{1, 0, 0, 0, 32}
	i := 0
	for k, val := range img {

		//fmt.Printf("%X", val)
		if fmt.Sprintf("%X", val) == fmt.Sprintf("%d", motif[i]) {
			//fmt.Println(i)
			i += 1
			if i == len(motif) {
				break
			}
		} else {
			fmt.Printf("%X", img[k])
			i = 0
		}
		_ = k
	}

	return nil
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
