package comp

import "fmt"

type MapperI interface {
	MapAddCompetitionData(compId, name string, teamIds []string) string
}

type Mapper struct{}

func (m Mapper) MapAddCompetitionData(compId, name string, teamIds []string) string {
	return fmt.Sprintf(PSQLAddCompetitionData, compId, name, teamIds)
}

const (
	PSQLAddCompetitionData = `insert into public.competitions (comp_id, name, teams)
			values ('%s', '%s', '{%s}') returning name;`
)
