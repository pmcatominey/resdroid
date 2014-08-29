package main

import (
	"html/template"
	"io"
	"sort"
)

const drawableReportTemplate = `
<html>
	<style>
		body {
			background-color: white;
			margin: 0;
		}
		#table {
			table-layout: fixed;
		}
		.missing {
			color: red;
		}
	</style>
	<body>
	<br/>
	Press c to change color
	<br/>
	<table id="table" border="1">
		<tr>
		<th></th>
		{{ range .Dirs }}
			<th>
				{{ .Name }}
			</th>
		{{ end }}
		</tr>
		{{ range .Rows }}
		<tr>
			{{ $name := .Name }}
			<th>
				{{ .Name }}
			</th>	
			{{ range .Dirs }}
			<td>
				{{ if .HasDrawable $name }}
				{{ $drawable := .Drawable $name }}
					<img src="data:image/*;base64,{{ $drawable.Base64 }}" />
				{{ else }}
					<span class="missing">Missing</span>
				{{ end }}
			</td>
			{{ end }}
		</tr>
		{{ end }}
	</table>
	</body>
	<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.1/jquery.min.js"></script>
	<script>
		$(document).keypress(function (event) {
			if (event.which == 99) { // c - change color
				// Flip background color
				console.log($("#table").css("background-color"));
				if ($("body").css("background-color") == "rgb(255, 255, 255)") {
					$("body").css("background-color", "black");
					$("th").css("color", "white");
				} else {
					$("body").css("background-color", "white");
					$("th").css("color", "black");
				}
			}
		});
	</script>
</html>
`

type DrawableReport struct {
	Dirs []*DrawableDirectory
	Rows []*DrawableReportRow
}

type DrawableReportRow struct {
	Name string
	Dirs []*DrawableDirectory
}

func GenerateDrawableReport(r *ResDirectory, w io.Writer) error {
	// Dirs
	dirs := []*DrawableDirectory{}
	for _, dir := range r.DrawableDirectories {
		dirs = append(dirs, dir)
	}

	// Create set of drawable names
	drawableNamesSet := make(map[string]struct{})
	drawableNames := []string{}
	for _, dir := range r.DrawableDirectories {
		for _, drawable := range dir.Drawables {
			// Skip non bitmap or 9-patch files
			if drawable.Type != Bitmap && drawable.Type != NinePatch {
				continue
			}

			if _, found := drawableNamesSet[drawable.Name]; !found {
				drawableNamesSet[drawable.Name] = struct{}{}
				drawableNames = append(drawableNames, drawable.Name)
			}
		}
	}
	sort.Strings(drawableNames)

	report := DrawableReport{
		dirs,
		[]*DrawableReportRow{},
	}

	for _, name := range drawableNames {
		row := &DrawableReportRow{
			name,
			dirs,
		}

		report.Rows = append(report.Rows, row)
	}

	templ, err := template.New("report").Parse(drawableReportTemplate)
	if err != nil {
		return err
	}

	return templ.Execute(w, report)
}
