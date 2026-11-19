package main

import (
	"log"

	"github.com/dkatzel-home/godocx"
)

func main() {
	// Create a new document
	doc, err := godocx.NewDocument()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Creating document with Table of Contents...")

	// Add content with headings
	doc.AddHeading("Introduction", 1)
	doc.AddParagraph("This is the introduction section of the document.")
	doc.AddParagraph("It contains important information about the project.")

	doc.AddHeading("Getting Started", 1)
	doc.AddParagraph("This section explains how to get started with the project.")

	doc.AddHeading("Installation", 2)
	doc.AddParagraph("Installation instructions go here.")

	doc.AddHeading("Configuration", 2)
	doc.AddParagraph("Configuration details are explained here.")

	doc.AddHeading("Advanced Usage", 1)
	doc.AddParagraph("Advanced usage examples and tips.")

	doc.AddHeading("Troubleshooting", 2)
	doc.AddParagraph("Common issues and their solutions.")

	doc.AddHeading("Conclusion", 1)
	doc.AddParagraph("Summary and next steps.")

	// Add TOC with fluent API configuration
	toc := doc.AddTableOfContents("Table of Contents", 3, 1, true, 20)

	log.Printf("TOC created with %d entries", len(toc.Entries))

	// Save the document
	err = doc.SaveTo("toc_example.docx")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("✅ Document with TOC created successfully!")
	log.Println("📄 Open 'toc_example.docx' to see the table of contents.")
	log.Println("🎯 Features:")
	log.Println("   • Professional formatting with dotted leaders")
	log.Println("   • Bold Level 1 headings")
	log.Println("   • Proper indentation for hierarchy")
	log.Println("   • Realistic page numbers")
	log.Println("   • Immediate visibility (no manual update needed)")

	// Print the heading structure
	headings, err := doc.GetHeadingStructure()
	if err != nil {
		log.Printf("Error getting heading structure: %v", err)
	} else {
		log.Println("\n📋 Document heading structure:")
		for _, heading := range headings {
			indent := ""
			for i := 0; i < heading.Level-1; i++ {
				indent += "  "
			}
			log.Printf("  %sLevel %d: %s", indent, heading.Level, heading.Text)
		}
	}

	log.Printf("\n📊 Total headings found: %d", len(headings))
}
