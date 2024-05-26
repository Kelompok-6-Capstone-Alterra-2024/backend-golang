package utilities

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"time"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func UploadImage(file multipart.File, path string) (string, error) {
	// Mendapatkan URL Cloudinary dari environment variables
	cloudinaryURL := os.Getenv("CLOUDINARY_URL")
	if cloudinaryURL == "" {
		return "", fmt.Errorf("CLOUDINARY_URL is not set")
	}

	// Inisialisasi Cloudinary
	cld, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		return "", fmt.Errorf("failed to initialize Cloudinary: %v", err)
	}

	// Membuat context dengan timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Mengunggah file ke Cloudinary
	uploadResult, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: path,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload image: %v", err)
	}

	// Mengembalikan URL gambar yang diunggah
	return uploadResult.SecureURL, nil
}

func DeleteImage(path string) error {
	// Mendapatkan URL Cloudinary dari environment variables
	cloudinaryURL := os.Getenv("CLOUDINARY_URL")
	if cloudinaryURL == "" {
		return fmt.Errorf("CLOUDINARY_URL is not set")
	}

	// Inisialisasi Cloudinary
	cld, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		return fmt.Errorf("failed to initialize Cloudinary: %v", err)
	}

	// Membuat context dengan timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Menghapus file dari Cloudinary
	_, err = cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: path,
	})
	if err != nil {
		return fmt.Errorf("failed to delete image: %v", err)
	}

	return nil
}
