package core

type Id uint

type Questionnaire struct {
	Id
	Title string
	Ref   string

	Questions               []Question
	RadioPossibleAnswers    []RadioPossibleAnswer
	CheckboxPossibleAnswers []CheckboxPossibleAnswer
	TextPossibleAnswers     []TextPossibleAnswer
	//RadioAnswers            []*gateway.RadioAnswer
	//CheckboxAnswers         []*gateway.CheckboxAnswer
	//TextAnswers             []*gateway.TextAnswer
}

type Question struct {
	Id            int
	Questionnaire Questionnaire
	Order         string
	Kind          string
	Text          string
}

type RadioPossibleAnswer struct {
	Question Question
	Text     string
}

type CheckboxPossibleAnswer struct {
	Question Question
	Text     string
}

type TextPossibleAnswer struct {
	Question    Question
	Text        string
	Placeholder string
}

type RadioAnswer struct {
	Questionnaire       Questionnaire
	Question            Question
	RadioPossibleAnswer RadioPossibleAnswer
}

type CheckboxAnswer struct {
	Questionnaire          Questionnaire
	Question               Question
	CheckboxPossibleAnswer CheckboxPossibleAnswer
}

type TextAnswer struct {
	Questionnaire      Questionnaire
	Question           Question
	TextPossibleAnswer TextPossibleAnswer
	Answer             string
}