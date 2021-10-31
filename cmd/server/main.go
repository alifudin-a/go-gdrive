package main

import (
	"github.com/alifudin-a/go-gdrive/pkg/domain/helper"
	"github.com/alifudin-a/go-gdrive/pkg/router"

	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load(".env")
}

func main() {
	helper.GetDriveService()
	router.Router()
	// Step 1: Open  file
	// f, err := os.Open("Tes Upload Base64 pdf.pdf")

	// if err != nil {
	// 	panic(fmt.Sprintf("cannot open file: %v", err))
	// }

	// defer f.Close()

	// Step 2: Get the Google Drive service
	// srv, err := getDriveService()
	// if err != nil {
	// 	log.Println("An error occured! ", err)
	// }

	// GET FILE
	// file, err := getFile(srv, "17wTLYUuGMWHFnvc52tCi8zEih0nU8x0u")
	// if err != nil {
	// 	log.Println("An error occured! ", err)
	// } else {
	// 	fmt.Printf("File: %s\nFile ID: %s\n", file.Name, file.Id)
	// }

	// LIST FILE
	// files, err := listFile(srv)
	// if err != nil {
	// 	log.Println("Unable to retrieve files! ", err)
	// } else {
	// 	fmt.Printf("List of files: \n")
	// 	if len(files.Files) == 0 {
	// 		fmt.Println("No files found!")
	// 	} else {
	// 		for _, i := range files.Files {
	// 			fmt.Printf("%s (%s)\n", i.Name, i.Id)
	// 		}
	// 	}
	// }

	//DELETE
	// err = deleteFile(srv, "10Y9YgCJao0917bKfsINnlkO_dkrWLlfJ")
	// if err != nil {
	// 	log.Println("An errro occured! ", err)
	// } else {
	// 	log.Println("Successfully delete file!")
	// }

	// // Step 3: Create directory
	// // dir, err := createFolder(srv, "New Folder", "root")

	// // if err != nil {
	// // 	panic(fmt.Sprintf("Could not create dir: %v\n", err))
	// // }

	// CREATE FILE
	// //give your drive folder id here in which you want to upload or create a new directory
	// folderId := "1dpO-uaNorJ-tO3Ysh-GEC6zMfWD3NRPI"

	// // Step 4: create the file and upload
	// file, err := createFile(srv, f.Name(), "application/octet-stream", f, folderId)

	// if err != nil {
	// 	panic(fmt.Sprintf("Could not create file: %v\n", err))
	// }

	// fmt.Printf("File '%s' uploaded successfully", file.Name)
	// fmt.Printf("\nFile Id: '%s' ", file.Id)
}
