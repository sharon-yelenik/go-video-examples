package main


import (
    
    "html/template"
    "net/http"
    "github.com/cloudinary/cloudinary-go"
    ec "lib"
    "os"
)

type Transformation struct {
    Description string
    Asset string
    Link string
    Text string
    Done  bool
}

type TransformationPageData struct {
    PageTitle string
    Transformations     []Transformation
}

func main() {
    // Add your Cloudinary credentials.
    cld, _ := cloudinary.NewFromParams(os.Getenv("CLOUD_NAME"), os.Getenv("API_KEY"), os.Getenv("API_SECRET"))

    url_resize_crop:=ec.Resize_crop(cld)
    url_concat := ec.Concat(cld) 
    url_trim := ec.Trim(cld)
    url_overlay := ec.Overlay(cld)
    url_effect := ec.Effect(cld)

  
    tmpl := template.Must(template.ParseFiles("transformations.html"))
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := TransformationPageData{
          PageTitle: "Cloudinary Transformations: Go SDK",
            Transformations: []Transformation{
              {Description: "Resizing and cropping", Asset: url_resize_crop, Done: true, Link: "https://cloudinary.com/documentation/go_media_transformations#resizing_and_cropping_videos", Text: "Resizing and cropping videos"},  
              {Description: "Concatenating", Asset: url_concat, Done: true, Link: "https://cloudinary.com/documentation/go_media_transformations#concatenating_videos", Text: "Concatenating videos"},
              {Description: "Trimming", Asset: url_trim, Done: true, Link: "https://cloudinary.com/documentation/go_media_transformations#trimming_videos", Text: "Trimming videos"},
              {Description: "Overlaying", Asset: url_overlay, Done: true, Link: "https://cloudinary.com/documentation/go_media_transformations#adding_video_overlays", Text: "Adding video overlays"},
              {Description: "Applying effects", Asset: url_effect, Done: true, Link: "https://cloudinary.com/documentation/go_media_transformations#adding_video_effects", Text: "Adding video effects"},
            },
        }
        tmpl.Execute(w, data)
    })
    http.ListenAndServe(":8000", nil)
}
