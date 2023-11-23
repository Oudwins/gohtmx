package handlers

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/Oudwins/stackifyer/data"
	"github.com/Oudwins/stackifyer/pkg/apierrors"
	filestorage "github.com/Oudwins/stackifyer/pkg/fileStorage"
	"github.com/Oudwins/stackifyer/types"
	"github.com/gofiber/fiber/v2"
	"github.com/h2non/bimg"
)

func PostUploadImages(c *fiber.Ctx) error {
	// validate input
	form, err := c.MultipartForm()
	if err != nil {
		return apierrors.New(errors.New("Please submit multipart form with the files"), http.StatusBadRequest)
	}

	imgs := form.File["images"]
	if len(imgs) > types.UploadImagesMaxNumber {
		return apierrors.New(errors.New("Group image upload maximum "+string(types.UploadImagesMaxSize)), http.StatusBadRequest)
	}

	msgs := []error{}
	authDetails := c.Locals(types.LocalsKeyAuthenticatedDetails).(*data.AuthenticatedDetails)

	timestamp := time.Now().Unix()
	for i, imgHead := range imgs {
		// VALIDATION
		// 1. Is less than max upload size
		if imgHead.Size > types.UploadImagesMaxSize {
			msg := errors.New(imgHead.Filename + " is larger than 1mb")
			msgs = append(msgs, msg)
			continue
		}
		// 2. check that it is an image
		if typ := imgHead.Header.Get("Content-Type"); !strings.HasPrefix(typ, "image/") {
			msg := errors.New(imgHead.Filename + " is not an image")
			msgs = append(msgs, msg)
		}

		img, err := imgHead.Open()
		defer img.Close()
		if err != nil {
			msgs = append(msgs, err)
			continue
		}

		// this may throw an error TODO FIX IT
		extension := strings.Split(imgHead.Filename, ".")[1]
		filename := fmt.Sprintf("%s_%d_%d.%s", authDetails.User.ID, timestamp, i, extension)

		// TODO -> Move this elsewhere
		// format image
		// convert img to buf
		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, img); err != nil {
			msgs = append(msgs, err)
			continue
		}
		imgBytes := buf.Bytes()
		// generate thumbnail
		thumbnail, err := bimg.NewImage(imgBytes).Thumbnail(200)

		// save images to disk
		err = filestorage.Api.SaveFile(&img, filename)
		if err != nil {
			msgs = append(msgs, err)
			continue
		}

		// postNewFileEntries to DB
		fileobj, err := data.PostNewFileEntry(&data.FileObject{
			Ftype:    types.FileTypeImage,
			UserId:   authDetails.User.ID,
			Filename: imgHead.Filename,
			Key:      filename,
		})
		if err != nil {
			msgs = append(msgs, err)
			continue
		}

	}

	// return correct response
	return c.SendStatus(200)
}
