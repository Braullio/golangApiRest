package enumChat

type ChatType struct {
	slug string
}

func (r ChatType) String() string {
	return r.slug
}

var (
	Success = ChatType{"🟢 *[ LOG ] SUCCESS ALERT* 🟢"}
	Warning = ChatType{"🟠 *[ LOG ] WARNING ALERT* 🟠"}
	Danger  = ChatType{"🔴 *[ LOG ] DANGER  ALERT* 🔴"}
)
