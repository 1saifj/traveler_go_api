package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"image/png"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"taveler/infrastructure/model"
	"taveler/infrastructure/utils"
	"taveler/usecase/interactor"
)

type fileController struct {
	interactor interactor.FileInteractor
}

type FileController interface {
	UploadImage(ctx *fiber.Ctx) error
	//TODO: CHANGE NAME TO GET IMAGE BY ID
	GetFileByID(ctx *fiber.Ctx) error
}

func NewFileController(i interactor.FileInteractor) FileController {
	return &fileController{interactor: i}
}

func (f *fileController) UploadImage(ctx *fiber.Ctx) error {
	fileModel := model.File{}
	header, err := ctx.FormFile("image")
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to generate uuid",
		})
	}

	file, err := header.Open()
	if err != nil {
		fmt.Println(err)
	}
	defer func(file multipart.File) {
		_ = file.Close()
	}(file)

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to read file",
		})
	}
	//list of string to string

	fileName := strings.Join(strings.Split(header.Filename, " "), "-")
	err = ioutil.WriteFile("./public/"+fileName, fileBytes, os.ModePerm)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to write file",
		})
	}
	//fileModel.Path = ctx.Protocol() + "://" + ctx.Hostname() + "/public/" + fileName
	fileModel.Path = "./public/" + fileName
	res, err := f.interactor.UploadImage(&fileModel)
	return ctx.Status(http.StatusOK).JSON(res)
}

func (f *fileController) GetFileByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	imagePath, err := f.interactor.GetFileByID(id)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to generate uuid111",
		})
	}
	img, err := utils.GetImageInStorage(imagePath.Path)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to generate uuid",
		})
	}
	size := ctx.Query("size")
	if size == "" {
		return ctx.Status(http.StatusOK).JSON(imagePath)
	} else {
		nImg := utils.ImageResizerBySize(img, size)
		err = png.Encode(ctx, nImg)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"message": "failed to generate uuid",
			})
		}
	}
	return nil

}
