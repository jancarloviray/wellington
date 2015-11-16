### Mixins

#### @include sprite-dimensions($map, "file")

Sprite-dimensions outputs the height and width properties of the specified image.

```
div {
  @include sprite-dimensions($spritemap, "file");
}
```

*Output*

```css
div {
	width: 100px;
	height: 50px;
}
```

### Functions

Don't see a function you want?  Check out [handlers](http://godoc.org/github.com/wellington/wellington/handlers) and submit a pull request!

#### sprite-map("glob/pattern"[, $spacing: 10px])

sprite-map generates a sprite from the matched images optinally with spacing between the images.  No output is generated by this function, instead the return is used in other functions.

```
$spritemap: sprite-map("*.png");
```

*Output*

```css

```

#### sprite($map, $name: "image"[, $offsetX: 0px, $offsetY: 0px])|

sprite generates a background url with background position to the position of the specified `"image"` in the spritesheet.  Optionally, offsets can be used to slightly modify the background position.

```
div {
	background: sprite($spritemap, "image");
}
```

*Output*

```css
div {
	background: url("spritegen.png") -0px -149px;
}
```

#### sprite-file($map, $name: "image")

Sprite-file returns an encoded string only useful for passing to image-width or image-height.

```
div {
	background: sprite-file($spritemap, "image");
}
```

*Output*

```css
div {
	background: {encodedstring};
}
```

#### image-height($path)

image-height returns the height of the image specified by `$path`.

```
div {
	height: image-height(sprite-file($spritemap, "image"));
}
div {
	height: image-height("path/to/image.png");
}
```

*Output*

```css
div {
	height: 50px;
}
div {
	height: 50px;
}
```

#### image-width($path)

image-width returns the width of the image specified by `$path`.

```
.first {
	width: image-width(sprite-file($spritemap, "image"));
}
.second {
	width: image-width("path/to/image.png");
}
```

*Output*

```css
.first {
	width: 50px;
}
.second {
	width: 50px;
}
```

#### inline-image($path[, $encode: false])

inline-image base64 encodes binary images (png, jpg, gif are currently supported). SVG images are by default url escaped. Optionally SVG can be base64 encoded by specifying `$encode: true`. Base64 encoding incurs a (10-30%) file size penalty.

```
.png {
	background: inline-image("path/to/image.png");
}
.svg {
	background: inline-image("path/to/image.svg", $encode: false);
}
```

*Output*

```css
.png {
	background: inline-image("data:image/png;base64,iVBOR...");
}
.svg {
	background: inline-image("data:image/svg+xml;utf8,%3C%3F...");
}
```

#### image-url($path)

image-url returns a relative path to an image in the image directory from the built css directory.

```
div {
	background: image-url("path/to/image.png");
}
```

*Output*

```css
div {
	background: url('../imgdirectory/path/to/image.png");
}
```

### font-url($path, [$raw:false])

font-url returns a relative path to fonts in your font directory.  You must set the font path to use this function.  By default, font-url will return `url('path/to/font')`, set `$raw: true` to only return the path

```
div {
	$path: font-url("arial.eot", true);
	@font-face {
		src: font-url("arial.eot");
		src: url("#{$path}");
	}
}
```

*Output*

```css
div {
	@font-face {
		src: url("../font/arial.eot");
		src: url("../font/arial.eot");
	}
}
```