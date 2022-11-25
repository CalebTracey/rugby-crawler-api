package comp

import (
	"fmt"
	"strings"
)

type MapperI interface {
	BuildCrawlerUrl(teamId string) string
	MapAddPSQLCompetitionData(compId, name string, teamIds []string) string
}

type Mapper struct{}

func (m Mapper) BuildCrawlerUrl(teamId string) string {
	return strings.Join([]string{CrawlBaseUrl, CrawlCompField, teamId}, "")
}

func (m Mapper) MapAddPSQLCompetitionData(compId, name string, teamIds []string) string {
	return fmt.Sprintf(PSQLAddCompetitionData, compId, name, teamIds)
}

const (
	PSQLAddCompetitionData = `insert into public.competitions (comp_id, name, teams)
			values ('%s', '%s', '{%s}') returning name;`

	CrawlBaseUrl   = "https://www.espn.co.uk/rugby/"
	CrawlCompField = "table/_/league/"
)
