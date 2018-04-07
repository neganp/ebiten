// Copyright 2018 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package graphicsutil

func MipmapSize(w, h int, level int) (width, height, actualLevel int) {
	for i := 0; i < level; i++ {
		if w == 1 || h == 1 {
			break
		}
		w /= 2
		h /= 2
		actualLevel++
	}
	return w, h, actualLevel
}

func GenerateMipmap(pix []byte, sx, sy, sw, sh int, level int) ([]byte, int, int) {
	dw, dh, actualLevel := MipmapSize(sx, sh, level)
	scale := 1
	for i := 0; i < actualLevel; i++ {
		scale *= 2
	}
	newPix := make([]byte, 4*dw*dh)
	for j := 0; j < dh; j++ {
		for i := 0; i < dw; i++ {
			// Each values are (at least) 32bit precision and scale <= 32768,
			// then dividing values with scale^2 should be fine.
			r := 0
			g := 0
			b := 0
			a := 0
			for j2 := 0; j2 < scale; j2++ {
				y := sy + j*scale + j2
				if y > sh {
					break
				}
				for i2 := 0; i2 < scale; i2++ {
					x := sx + i*scale + i2
					if x > sw {
						break
					}
					idx := 4 * (x + y*sw)
					r += int(pix[idx])
					g += int(pix[idx+1])
					b += int(pix[idx+2])
					a += int(pix[idx+3])
				}
			}

			idx := 4 * (i + j*dw)
			newPix[idx] = byte(r / (scale * scale))
			newPix[idx+1] = byte(g / (scale * scale))
			newPix[idx+2] = byte(b / (scale * scale))
			newPix[idx+3] = byte(a / (scale * scale))
		}
	}
	return newPix, dw, dh
}
