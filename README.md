# resdroid

Resdroid is a tool for getting an overview of the res directory within an android project.

I plan on adding more reports to it but currently it only generates an overview of Bitmap and 9-Patch drawables.

## Usage

```
	# With a working go 1.3 install and workspace
	go get github.com/mcatominey/resdroid

	resdroid <flags> /Path/To/Res
```

### Flags

#### -drawable output.html

This will generate a drawable report in HTML format in output.html

![alt tag](https://raw.githubusercontent.com/mcatominey/resdroid/master/screenshots/drawable.png)