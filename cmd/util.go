/*
Copyright © 2021 Zoraiz Hassan <hzoraiz8@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"
	"path"

	"github.com/TheZoraiz/ascii-image-converter/aic_package/winsize"
)

// Check input and flag values for detecting errors or invalid inputs
func checkInputAndFlags(args []string) bool {

	gifCount := 0
	gifPresent := false
	nonGifPresent := false
	for _, arg := range args {
		extension := path.Ext(arg)

		if extension == ".gif" {
			gifPresent = true
			gifCount++
		} else {
			nonGifPresent = true
		}
	}

	if gifPresent && nonGifPresent {
		fmt.Printf("Error: There are other inputs along with GIFs\nDue to the potential looping nature of GIFs, non-GIFs must not be supplied alongside\n\n")
		return true
	}

	if gifCount > 1 {
		fmt.Printf("Error: There are multiple GIFs supplied\nDue to the potential looping nature of GIFs, only one GIF per command is supported\n\n")
		return true
	}

	if formatsTrue {
		fmt.Printf("Supported input formats:\n\n" +
			"JPEG/JPG\n" +
			"PNG\n" +
			"WEBP\n" +
			"BMP\n" +
			"TIFF/TIF\n" +
			"GIF\n\n")
		return true
	}

	if len(args) < 1 {
		fmt.Printf("Error: Need at least 1 input path/url\nUse the -h flag for more info\n\n")
		return true
	}

	if customMap != "" && len(customMap) < 2 {
		fmt.Printf("Need at least 2 characters for --map flag\n\n")
		return true
	}

	if dimensions != nil {

		numberOfDimensions := len(dimensions)
		if numberOfDimensions != 2 {
			fmt.Printf("Error: requires 2 dimensions, got %v\n\n", numberOfDimensions)
			return true
		}

		if dimensions[0] < 1 || dimensions[1] < 1 {
			fmt.Printf("Error: invalid values for dimensions\n\n")
			return true
		}

		defaultTermWidth, _, err := winsize.GetTerminalSize()
		if err != nil {
			fmt.Printf("Error: %v\n\n", err)
			return true
		}

		defaultTermWidth -= 1
		if dimensions[0] > defaultTermWidth {
			fmt.Printf("Error: set width must be lower than terminal width\n\n")
			return true
		}
	}

	if width != 0 || height != 0 {

		if width != 0 && height != 0 {
			fmt.Printf("Error: both --width and --height can't be set. Use --dimensions instead\n\n")
			return true
		} else {

			defaultTermWidth, _, err := winsize.GetTerminalSize()
			if err != nil {
				fmt.Printf("Error: %v\n\n", err)
				return true
			}

			// Check if set width exceeds terminal
			defaultTermWidth -= 1
			if width > defaultTermWidth {
				fmt.Printf("Error: set width must be lower than terminal width\n\n")
				return true
			}

			if width < 0 {
				fmt.Printf("Error: invalid value for width\n\n")
				return true
			}

			if height < 0 {
				fmt.Printf("Error: invalid value for height\n\n")
				return true
			}

		}

	}

	if saveBgColor == nil {
		saveBgColor = []int{0, 0, 0}
	} else {
		bgValues := len(saveBgColor)
		if bgValues != 3 {
			fmt.Printf("Error: --save-bg requires 3 values for RGB, got %v\n\n", bgValues)
			return true
		}

		if saveBgColor[0] < 0 || saveBgColor[1] < 0 || saveBgColor[2] < 0 {
			fmt.Printf("Error: RBG values must be between 0 and 255\n\n")
			return true
		}

		if saveBgColor[0] > 255 || saveBgColor[1] > 255 || saveBgColor[2] > 255 {
			fmt.Printf("Error: RBG values must be between 0 and 255\n\n")
			return true
		}
	}

	if fontColor == nil {
		fontColor = []int{255, 255, 255}
	} else {
		fontColorValues := len(fontColor)
		if fontColorValues != 3 {
			fmt.Printf("Error: --font-color requires 3 values for RGB, got %v\n\n", fontColorValues)
			return true
		}

		if fontColor[0] < 0 || fontColor[1] < 0 || fontColor[2] < 0 {
			fmt.Printf("Error: RBG values must be between 0 and 255\n\n")
			return true
		}

		if fontColor[0] > 255 || fontColor[1] > 255 || fontColor[2] > 255 {
			fmt.Printf("Error: RBG values must be between 0 and 255\n\n")
			return true
		}
	}

	if threshold == 0 {
		threshold = 128
	}

	if threshold < 0 || threshold > 255 {
		fmt.Printf("Error: threshold must be between 0 and 255\n\n")
		return true
	}

	return false
}
