package service

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	_ "github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type Cloudinary struct{}

func (service *Cloudinary) Credentials() (*cloudinary.Cloudinary, context.Context) {
	// Add your Cloudinary credentials, set configuration parameter
	// Secure=true to return "https" URLs, and create a context
	//===================
	// cld, _ := cloudinary.New()
	cld, _ := cloudinary.NewFromParams(os.Getenv("CLOUDINARY_NAME"), os.Getenv("CLOUDINARY_API_KEY"), os.Getenv("CLOUDINARY_API_SECRET"))
	cld.Config.URL.Secure = true
	ctx := context.Background()
	return cld, ctx
}

func (service *Cloudinary) UploadImage(cld *cloudinary.Cloudinary, ctx context.Context, folderName string, fileName interface{}) string {

	// Upload the image.
	// Set the asset's public ID and allow overwriting the asset with new versions
	resp, err := cld.Upload.Upload(ctx, fileName, uploader.UploadParams{
		Folder:         folderName,
		UniqueFilename: api.Bool(false),
		Overwrite:      api.Bool(true)})
	if err != nil {
		fmt.Println("error")
	}
	// Log the delivery URL
	fmt.Printf("****2. Upload an image****\nDelivery URL:", resp.SecureURL, "\n")
	return resp.SecureURL
}

func (service *Cloudinary) DestroyImage(cld *cloudinary.Cloudinary, ctx context.Context) {
	_, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID:     "sample",
		ResourceType: "video"})
	if err != nil {
		fmt.Println("error destroy image")
	}
}
