package markdown

type Options func(r *renderer)

// Use a custom collection of ANSI colors for the headings
func WithHeadingShades(shades []shadeFmt) Options {
	return func(r *renderer) {
		r.headingShade = shade(shades)
	}
}

// Use a custom collection of ANSI colors for the blockquotes
func WithBlockquoteShades(shades []shadeFmt) Options {
	return func(r *renderer) {
		r.blockQuoteShade = shade(shades)
	}
}
