package lib

import(
    "fmt"
    "github.com/cloudinary/cloudinary-go"
)

func Concat(cld *cloudinary.Cloudinary) string {
    // Instantiate an object that stores information for asset with public ID "horse_race" in folder "docs/sdk/go"
    v_races, err := cld.Video("docs/sdk/go/horse_race")
    if err != nil {
        fmt.Println("error")
    }

    // Add the transformation
    v_races.Transformation = "c_fill,h_300,w_450/du_5/fl_splice,l_video:swimming_race/c_fill,h_300,w_450/du_5/fl_layer_apply"

    // Generate and print the delivery URL
    myURL, err := v_races.String()
    fmt.Println(myURL)
    if err != nil {
        fmt.Println("error")
    }
    return myURL
}