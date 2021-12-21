package lib

import(
    "fmt"
    "github.com/cloudinary/cloudinary-go"
)

func Effect(cld *cloudinary.Cloudinary) string {
    // Instantiate an object for the video with public ID "dog_garden" in folder "docs/sdk/go"
    v_dogarden, err := cld.Video("docs/sdk/go/dog_garden")
    if err != nil {
        fmt.Println("error")
    }

    // Add the transformation
    v_dogarden.Transformation = "c_scale,w_400/du_10.0/e_vignette:50/e_fade:2000/e_fade:-4000/e_loop:2"

    // Generate and print the delivery URL
    myURL, err := v_dogarden.String()
    if err != nil {
        fmt.Println("error")
    }
    fmt.Println(myURL)

    // Output: https://res.cloudinary.com/demo/video/upload/c_scale,w_400/du_10.0/e_vignette:50/e_fade:2000/e_fade:-4000/e_loop:2/v1/docs/sdk/go/dog_garden 
  return myURL
}