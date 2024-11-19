package render

import (
	"bytes"
	"github.com/sean-d/tentbnb/pkg/config"
	"github.com/sean-d/tentbnb/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

// NewTemplates sets the config for the template package.
// Is called from main to establish access to the global application config struct.
func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultData returns the data provided combined with data that should be present on pages
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// TemplateRender creates a template cache with every call...will update later on.
func TemplateRender(w http.ResponseWriter, tmpl string, templateData *models.TemplateData) {
	var templateCache map[string]*template.Template
	var err error

	if app.UseCache {
		templateCache = app.TemplateCache
	} else {
		templateCache, err = CreateTemplateCache()

		if err != nil {
			log.Fatalf("Error: %s", err)
		}
	}
	// get requested template from cache and assign to newTemplate
	newTemplate, ok := templateCache[tmpl]
	if !ok {
		log.Fatalf("error: can't find template for %s", tmpl)
	}

	// using a buffer to hold bytes for finer-grained error checking. if we get an error, we know it's from a value
	// stored in the map.
	buf := new(bytes.Buffer)

	// set default data from the data supplied and combine with the data provided in AddDefaultData()
	templateData = AddDefaultData(templateData)

	err = newTemplate.Execute(buf, templateData)

	if err != nil {
		log.Fatalf("error: %s", err)
	}

	// render template...write using the buffer we created above and write to the http.ResponseWriter, w.
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
}

/*
	CreateTemplateCache:

The cycle (or loop) works like this:

Get the list of files whose names end in page.tmpl and store the result in pages.

Loop through pages and do the following.

 1. Get the base name of the page (i.e. home.page.tmpl becomes home)

 2. Create a new variable, ts, which is a *template.Template, by calling template.New and parsing the current page template.

 3. Then, we check to see if the folder contains any template layouts. Why?  Because we have no way of knowing, at the moment, whether or not the current page might use a layout template, so we have to be aware of them all (yes, an individual template can use none, one, or more than one layout).

 4. Next we call ParseGlob on the ts parsing all of the template layout files that we find, if any. Here is where you are getting confused, I think. You say that we overwrite the value of ts when we "run ParseGlob on the ts that we got before," but we do not. This is the line in question:

templateSet, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
And this is from the documentation for ParseGlob:

ParseGlob creates a new Template and parses the template definitions from the files identified by the pattern. The files are matched according to the semantics of filepath.Match, and the pattern must match at least one file. The returned template will have the (base) name and (parsed) contents of the first file matched by the pattern. ParseGlob is equivalent to calling ParseFiles with the list of files matched by the pattern.

The important thing to notice here is that we are, in the line of code above, taking the existing vlaue of ts, and then adding to it whatever we get when we call ParseGlob, since we are calling ParseGlob on ts itself.

So, another way to read that line of code is to say it this way:

ts is now equal to whatever ts was before, plus whatever we get when we parse all the templates that end in layout.tmpl.

So that this moment, ts, our template set, will include the template we want to look at (i.e. home.page.tmpl, for example) as well as all of the templates that end in layout.tmpl.

We are not overwriting anything in ts; we are adding to it.
*/
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	// get all files ending in *.page.html
	pages, err := filepath.Glob("./templates/*.page.gohtml")

	if err != nil {
		return myCache, err
	}

	// range through the pages slice and pull out the filename
	for _, page := range pages {
		pageName := filepath.Base(page)

		// create a template with the given name
		// then parse the file (the *.page.html assigned to page above) and save that as a template
		templateSet, err := template.New(pageName).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		layouts, err := filepath.Glob("./templates/*.layout.gohtml")

		if err != nil {
			return myCache, err
		}

		/*
			if there are ANY layouts, we append to the templateSet all layouts. since a page can use none or all, we
			provide all of them for use, whether used or not by all pages.

			parseGlob simply appends to templateSet with the bits found from the layouts.
		*/
		if len(layouts) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.gohtml")

			if err != nil {
				return myCache, err
			}
		}

		// once the *.page.html is turned into a templateSet and any layouts are added to it, we create
		// an entry in the myCache map for it
		myCache[pageName] = templateSet
	}
	return myCache, nil
}
