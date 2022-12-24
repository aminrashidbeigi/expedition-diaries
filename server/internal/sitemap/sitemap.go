package sitemap

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aminrashidbeigi/expedition-diaries/storage/queries"
	"github.com/sabloger/sitemap-generator/smg"
)

type SitemapGenerator struct {
	Storage    *queries.Queries
	HostName   string
	OutputPath string
}

func (sg *SitemapGenerator) Generate() {
	now := time.Now().UTC()

	sm := smg.NewSitemap(true)
	sm.SetName("sitemap")
	sm.SetHostname(sg.HostName)
	sm.SetOutputPath(sg.OutputPath)
	sm.SetLastMod(&now)
	sm.SetCompress(false)

	err := sm.Add(&smg.SitemapLoc{
		Loc:        "/",
		LastMod:    &now,
		ChangeFreq: smg.Daily,
		Priority:   1.0,
	})
	if err != nil {
		log.Fatal("Unable to add SitemapLoc:", err)
	}

	err = sm.Add(&smg.SitemapLoc{
		Loc:        "/map",
		LastMod:    &now,
		ChangeFreq: smg.Monthly,
		Priority:   1.0,
	})
	if err != nil {
		log.Fatal("Unable to add SitemapLoc:", err)
	}

	err = sm.Add(&smg.SitemapLoc{
		Loc:        "/add-expedition",
		LastMod:    &now,
		ChangeFreq: smg.Monthly,
		Priority:   0.3,
	})
	if err != nil {
		log.Fatal("Unable to add SitemapLoc:", err)
	}

	ctx := context.Background()
	travels, err := sg.Storage.GetTravels(ctx, queries.GetTravelsParams{
		Offset: 0,
		Limit:  1000,
	})
	if err != nil {
		log.Fatal("can not get travels: ", err)
	}
	for _, travel := range travels {
		err := sm.Add(&smg.SitemapLoc{
			Loc:        "/travels/" + travel.Slug.String,
			LastMod:    &now,
			ChangeFreq: smg.Weekly,
			Priority:   0.8,
		})
		if err != nil {
			log.Fatal("Unable to add SitemapLoc:", err)
		}
	}

	countries, err := sg.Storage.GetCountries(ctx)
	if err != nil {
		log.Fatal("can not get countries: ", err)
	}

	for _, country := range countries {
		err := sm.Add(&smg.SitemapLoc{
			Loc:        "/countries/" + country.Code,
			LastMod:    &now,
			ChangeFreq: smg.Daily,
			Priority:   0.7,
		})
		if err != nil {
			log.Fatal("Unable to add SitemapLoc:", err)
		}
	}

	filenames, err := sm.Save()
	if err != nil {
		log.Fatal("Unable to Save Sitemap:", err)
	}
	for i, filename := range filenames {
		fmt.Println("file no.", i+1, filename)
	}

	fmt.Println("Sitemap generated")
}