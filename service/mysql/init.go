package mysql

var (
	scpDB  = &DB{Name: "scp"}
	scpDAO = &SCPDAO{db: scpDB}
)

func init() {

}
