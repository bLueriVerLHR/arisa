package plugins

import (
	"arisa/tools"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type logonByQRCode struct {
	Code   int `json:"code"`
	Result struct {
		CodeURL string `json:"codeUrl"`
		PollKey string `json:"pollKey"`
	} `json:"result"`
	Message string `json:"message"`
	TraceID string `json:"traceId"`
	Sampled bool   `json:"sampled"`
}

type pollkey struct {
	Code   int `json:"code"`
	Result struct {
		CodeStatus int    `json:"codeStatus"`
		Token      string `json:"token"`
	} `json:"result"`
	Message string `json:"message"`
	TraceID string `json:"traceId"`
	Sampled bool   `json:"sampled"`
}

type LessonList struct {
	Code   int `json:"code"`
	Result struct {
		Result []struct {
			ID        int    `json:"id"`
			ShortName string `json:"shortName"`
			Name      string `json:"name"`
			ImgURL    string `json:"imgUrl"`
			TermPanel struct {
				ID                        int         `json:"id"`
				CourseID                  int         `json:"courseId"`
				CourseName                interface{} `json:"courseName"`
				StartTime                 int64       `json:"startTime"`
				EndTime                   int64       `json:"endTime"`
				Duration                  string      `json:"duration"`
				PublishStatus             int         `json:"publishStatus"`
				CloseVisableStatus        int         `json:"closeVisableStatus"`
				BigPhotoURL               string      `json:"bigPhotoUrl"`
				LectorPanels              interface{} `json:"lectorPanels"`
				ScoreCardDto              interface{} `json:"scoreCardDto"`
				SchoolPanel               interface{} `json:"schoolPanel"`
				JSONContent               interface{} `json:"jsonContent"`
				HasEnroll                 bool        `json:"hasEnroll"`
				SpecialChargeableTerm     bool        `json:"specialChargeableTerm"`
				AchievementStatus         int         `json:"achievementStatus"`
				CertStatus                interface{} `json:"certStatus"`
				ChargeCertStatus          interface{} `json:"chargeCertStatus"`
				OrdinaryEditors           interface{} `json:"ordinaryEditors"`
				CertNo                    interface{} `json:"certNo"`
				ChargeableCert            int         `json:"chargeableCert"`
				CopyRight                 interface{} `json:"copyRight"`
				Mode                      int         `json:"mode"`
				SelfMocTermCopyright      interface{} `json:"selfMocTermCopyright"`
				OriginMocTermCopyRight    interface{} `json:"originMocTermCopyRight"`
				FromTermID                int         `json:"fromTermId"`
				FromTermMode              int         `json:"fromTermMode"`
				OriginalCourseChannel     interface{} `json:"originalCourseChannel"`
				ApplyMoocStatus           int         `json:"applyMoocStatus"`
				SpocToOocStatus           int         `json:"spocToOocStatus"`
				ApplyConvertChannelStatus interface{} `json:"applyConvertChannelStatus"`
				ApplyPassedTermID         interface{} `json:"applyPassedTermId"`
				SyncPrice                 interface{} `json:"syncPrice"`
				AsynPrice                 interface{} `json:"asynPrice"`
				Copied                    int         `json:"copied"`
				CopyTime                  interface{} `json:"copyTime"`
				OrderPrice                interface{} `json:"orderPrice"`
				EnrollCount               int         `json:"enrollCount"`
				Price                     float64     `json:"price"`
				OriginalPrice             float64     `json:"originalPrice"`
				CertApplyStartTime        interface{} `json:"certApplyStartTime"`
				CertApplyEndTime          interface{} `json:"certApplyEndTime"`
				LessonsCount              int         `json:"lessonsCount"`
				ProductType               int         `json:"productType"`
				SchoolID                  int         `json:"schoolId"`
			} `json:"termPanel"`
			SchoolPanel struct {
				ID                  int         `json:"id"`
				Name                string      `json:"name"`
				ShortName           string      `json:"shortName"`
				ImgURL              interface{} `json:"imgUrl"`
				SupportMooc         interface{} `json:"supportMooc"`
				SupportSpoc         interface{} `json:"supportSpoc"`
				BgPhoto             interface{} `json:"bgPhoto"`
				SmallLogo           interface{} `json:"smallLogo"`
				SupportCommonMooc   interface{} `json:"supportCommonMooc"`
				SupportPostgradexam interface{} `json:"supportPostgradexam"`
				ClassroomSupport    interface{} `json:"classroomSupport"`
			} `json:"schoolPanel"`
			Mode       int `json:"mode"`
			Channel    int `json:"channel"`
			MocTagDtos []struct {
				ID            int         `json:"id"`
				TagFamilyID   int         `json:"tagFamilyId"`
				TagFamilyName interface{} `json:"tagFamilyName"`
				TargetID      interface{} `json:"targetId"`
				TargetType    interface{} `json:"targetType"`
				Name          string      `json:"name"`
				Colour        string      `json:"colour"`
				Weight        int         `json:"weight"`
				Comment       string      `json:"comment"`
				Link          string      `json:"link"`
				IsTop         interface{} `json:"isTop"`
			} `json:"mocTagDtos"`
			LearnedCount int         `json:"learnedCount"`
			ReleaseCount interface{} `json:"releaseCount"`
			ProductType  int         `json:"productType"`
			CourseType   int         `json:"courseType"`
		} `json:"result"`
		Pagination struct {
			SortCriterial         interface{} `json:"sortCriterial"`
			DefaultPageSize       int         `json:"DEFAULT_PAGE_SIZE"`
			DefaultPageIndex      int         `json:"DEFAULT_PAGE_INDEX"`
			DefaultTotlePageCount int         `json:"DEFAULT_TOTLE_PAGE_COUNT"`
			DefaultTotleCount     int         `json:"DEFAULT_TOTLE_COUNT"`
			DefaultOffset         int         `json:"DEFAULT_OFFSET"`
			PageSize              int         `json:"pageSize"`
			PageIndex             int         `json:"pageIndex"`
			TotlePageCount        int         `json:"totlePageCount"`
			TotleCount            int         `json:"totleCount"`
			Offset                int         `json:"offset"`
			Limit                 int         `json:"limit"`
		} `json:"pagination"`
	} `json:"result"`
	Message string `json:"message"`
	TraceID string `json:"traceId"`
	Sampled bool   `json:"sampled"`
}

type Lesson struct {
	Code   int `json:"code"`
	Result struct {
		LastLearnUnitID interface{} `json:"lastLearnUnitId"`
		MocTermDto      struct {
			Times              int         `json:"times"`
			ID                 int         `json:"id"`
			GmtCreate          interface{} `json:"gmtCreate"`
			GmtModified        interface{} `json:"gmtModified"`
			CourseID           int         `json:"courseId"`
			CloseVisableStatus int         `json:"closeVisableStatus"`
			StartTime          int64       `json:"startTime"`
			Duration           interface{} `json:"duration"`
			EndTime            int64       `json:"endTime"`
			PublishStatus      interface{} `json:"publishStatus"`
			CourseLoad         interface{} `json:"courseLoad"`
			SmallPhoto         interface{} `json:"smallPhoto"`
			BigPhoto           interface{} `json:"bigPhoto"`
			FirstPublishTime   interface{} `json:"firstPublishTime"`
			EnrollCount        interface{} `json:"enrollCount"`
			LessonsCount       interface{} `json:"lessonsCount"`
			CourseName         string      `json:"courseName"`
			CoverPhoto         string      `json:"coverPhoto"`
			Chapters           []struct {
				ID          int         `json:"id"`
				GmtCreate   int64       `json:"gmtCreate"`
				GmtModified int64       `json:"gmtModified"`
				Name        string      `json:"name"`
				Position    int         `json:"position"`
				TermID      int         `json:"termId"`
				ContentType int         `json:"contentType"`
				ContentID   interface{} `json:"contentId"`
				ReleaseTime int64       `json:"releaseTime"`
				Published   bool        `json:"published"`
				Lessons     []struct {
					ID            int         `json:"id"`
					GmtCreate     int64       `json:"gmtCreate"`
					GmtModified   int64       `json:"gmtModified"`
					Name          string      `json:"name"`
					Position      int         `json:"position"`
					TermID        int         `json:"termId"`
					ChapterID     int         `json:"chapterId"`
					ContentType   int         `json:"contentType"`
					ContentID     interface{} `json:"contentId"`
					IsTestChecked bool        `json:"isTestChecked"`
					Units         []struct {
						ID                 int         `json:"id"`
						GmtCreate          int64       `json:"gmtCreate"`
						GmtModified        int64       `json:"gmtModified"`
						Name               string      `json:"name"`
						Position           int         `json:"position"`
						LessonID           int         `json:"lessonId"`
						ChapterID          int         `json:"chapterId"`
						TermID             int         `json:"termId"`
						ContentType        int         `json:"contentType"`
						ContentID          int         `json:"contentId"`
						UnitID             interface{} `json:"unitId"`
						Live               interface{} `json:"live"`
						FreePreview        interface{} `json:"freePreview"`
						DurationInSeconds  interface{} `json:"durationInSeconds"`
						LearnCount         interface{} `json:"learnCount"`
						ViewStatus         int         `json:"viewStatus"`
						NameAlias          interface{} `json:"nameAlias"`
						PhotoURL           interface{} `json:"photoUrl"`
						ResourceInfo       interface{} `json:"resourceInfo"`
						Attachments        interface{} `json:"attachments"`
						AnchorQuestions    interface{} `json:"anchorQuestions"`
						JSONContent        interface{} `json:"jsonContent"`
						LiveInfoDto        interface{} `json:"liveInfoDto"`
						YktRelatedLiveInfo interface{} `json:"yktRelatedLiveInfo"`
					} `json:"units"`
					ReleaseTime     int64       `json:"releaseTime"`
					ViewStatus      int         `json:"viewStatus"`
					TestDraftStatus int         `json:"testDraftStatus"`
					Test            interface{} `json:"test"`
				} `json:"lessons"`
				Homeworks []struct {
					ID              int         `json:"id"`
					GmtCreate       int64       `json:"gmtCreate"`
					GmtModified     int64       `json:"gmtModified"`
					Name            string      `json:"name"`
					Position        int         `json:"position"`
					TermID          int         `json:"termId"`
					ChapterID       int         `json:"chapterId"`
					ContentType     int         `json:"contentType"`
					ContentID       int         `json:"contentId"`
					IsTestChecked   bool        `json:"isTestChecked"`
					Units           interface{} `json:"units"`
					ReleaseTime     int64       `json:"releaseTime"`
					ViewStatus      int         `json:"viewStatus"`
					TestDraftStatus int         `json:"testDraftStatus"`
					Test            struct {
						ID                       int         `json:"id"`
						ReleaseTime              int64       `json:"releaseTime"`
						Type                     int         `json:"type"`
						Name                     string      `json:"name"`
						Deadline                 int64       `json:"deadline"`
						TestTime                 interface{} `json:"testTime"`
						Trytime                  interface{} `json:"trytime"`
						UsedTryCount             int         `json:"usedTryCount"`
						EvaluateJudgeType        int         `json:"evaluateJudgeType"`
						EvaluateNeedTrain        int         `json:"evaluateNeedTrain"`
						EvaluateStart            int64       `json:"evaluateStart"`
						EvaluateEnd              int64       `json:"evaluateEnd"`
						EvaluateScoreReleaseTime int64       `json:"evaluateScoreReleaseTime"`
						ScorePubStatus           int         `json:"scorePubStatus"`
						EnableEvaluation         bool        `json:"enableEvaluation"`
						UserScore                interface{} `json:"userScore"`
						TotalScore               float64     `json:"totalScore"`
						BonusScore               interface{} `json:"bonusScore"`
						ExamID                   int         `json:"examId"`
					} `json:"test"`
				} `json:"homeworks"`
				Quizs []struct {
					ID              int         `json:"id"`
					GmtCreate       int64       `json:"gmtCreate"`
					GmtModified     int64       `json:"gmtModified"`
					Name            string      `json:"name"`
					Position        int         `json:"position"`
					TermID          int         `json:"termId"`
					ChapterID       int         `json:"chapterId"`
					ContentType     int         `json:"contentType"`
					ContentID       int         `json:"contentId"`
					IsTestChecked   bool        `json:"isTestChecked"`
					Units           interface{} `json:"units"`
					ReleaseTime     int64       `json:"releaseTime"`
					ViewStatus      int         `json:"viewStatus"`
					TestDraftStatus int         `json:"testDraftStatus"`
					Test            struct {
						ID                       int         `json:"id"`
						ReleaseTime              int64       `json:"releaseTime"`
						Type                     int         `json:"type"`
						Name                     string      `json:"name"`
						Deadline                 int64       `json:"deadline"`
						TestTime                 int         `json:"testTime"`
						Trytime                  int         `json:"trytime"`
						UsedTryCount             int         `json:"usedTryCount"`
						EvaluateJudgeType        interface{} `json:"evaluateJudgeType"`
						EvaluateNeedTrain        interface{} `json:"evaluateNeedTrain"`
						EvaluateStart            interface{} `json:"evaluateStart"`
						EvaluateEnd              interface{} `json:"evaluateEnd"`
						EvaluateScoreReleaseTime interface{} `json:"evaluateScoreReleaseTime"`
						ScorePubStatus           int         `json:"scorePubStatus"`
						EnableEvaluation         bool        `json:"enableEvaluation"`
						UserScore                interface{} `json:"userScore"`
						TotalScore               float64     `json:"totalScore"`
						BonusScore               interface{} `json:"bonusScore"`
						ExamID                   int         `json:"examId"`
					} `json:"test"`
				} `json:"quizs"`
				HasFreePreviewVideo bool        `json:"hasFreePreviewVideo"`
				Exam                interface{} `json:"exam"`
				DraftStatus         int         `json:"draftStatus"`
			} `json:"chapters"`
			Exams                 []interface{} `json:"exams"`
			Mode                  int           `json:"mode"`
			FromTermID            interface{}   `json:"fromTermId"`
			SchoolID              int           `json:"schoolId"`
			HasFreePreviewVideo   bool          `json:"hasFreePreviewVideo"`
			VideoID               interface{}   `json:"videoId"`
			Description           interface{}   `json:"description"`
			BgKnowledge           interface{}   `json:"bgKnowledge"`
			Outline               interface{}   `json:"outline"`
			OutlineStructure      interface{}   `json:"outlineStructure"`
			ReommendRead          interface{}   `json:"reommendRead"`
			CourseStyle           interface{}   `json:"courseStyle"`
			Faq                   interface{}   `json:"faq"`
			JSONContent           interface{}   `json:"jsonContent"`
			Requirements          interface{}   `json:"requirements"`
			RequirementsForCert   interface{}   `json:"requirementsForCert"`
			DescriptionForCert    interface{}   `json:"descriptionForCert"`
			Target                interface{}   `json:"target"`
			MobDescription        interface{}   `json:"mobDescription"`
			ChiefLectorDto        interface{}   `json:"chiefLectorDto"`
			StaffLectorDtos       interface{}   `json:"staffLectorDtos"`
			StaffAssistDtos       interface{}   `json:"staffAssistDtos"`
			ChiefLector           interface{}   `json:"chiefLector"`
			StaffLectors          interface{}   `json:"staffLectors"`
			StaffAssists          interface{}   `json:"staffAssists"`
			ChargeableCert        interface{}   `json:"chargeableCert"`
			SpecialChargeableTerm bool          `json:"specialChargeableTerm"`
			TimeToFreeze          interface{}   `json:"timeToFreeze"`
			PreviousCourseDtos    interface{}   `json:"previousCourseDtos"`
			AnnouncementDtos      interface{}   `json:"announcementDtos"`
			Enrolled              interface{}   `json:"enrolled"`
			HasPaid               interface{}   `json:"hasPaid"`
			AchievementStatus     interface{}   `json:"achievementStatus"`
			FromTermMode          interface{}   `json:"fromTermMode"`
			ApplyMoocStatus       interface{}   `json:"applyMoocStatus"`
			OriginCopyRightTermID interface{}   `json:"originCopyRightTermId"`
			ExtraInfo             interface{}   `json:"extraInfo"`
			NeedPassword          bool          `json:"needPassword"`
			Copied                interface{}   `json:"copied"`
			CopyTime              interface{}   `json:"copyTime"`
			Price                 interface{}   `json:"price"`
			IsStart               bool          `json:"isStart"`
			IsEnd                 bool          `json:"isEnd"`
			OriginalPrice         interface{}   `json:"originalPrice"`
			OutLineStructureDtos  interface{}   `json:"outLineStructureDtos"`
			ProductType           interface{}   `json:"productType"`
			Channel               int           `json:"channel"`
			PositionStatus        interface{}   `json:"positionStatus"`
			DetailDraftStatus     int           `json:"detailDraftStatus"`
			WebVisible            interface{}   `json:"webVisible"`
		} `json:"mocTermDto"`
	} `json:"result"`
	Message string `json:"message"`
	TraceID string `json:"traceId"`
	Sampled bool   `json:"sampled"`
}

func GetLessonList(cookie string) LessonList {
	client := &http.Client{}
	csrfkey := string(regexp.MustCompile(`NTESSTUDYSI=[0-9a-z]{1,}`).Find([]byte(cookie)))[12:]
	uid := string(regexp.MustCompile(`NETEASE_WDA_UID=[0-9]{1,}`).Find([]byte(cookie)))[16:]
	req, err := http.NewRequest("POST", "https://www.icourse163.org/web/j/learnerCourseRpcBean.getMyLearnedCoursePanelList.rpc?csrfKey="+csrfkey, strings.NewReader("type=30&p=1&psize=8&courseType=1"))
	tools.Check(err)
	req.Header.Set("Host", "www.icourse163.org")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,en-US;q=0.7,ja;q=0.3")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Length", "32")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", cookie)
	req.Header.Set("DNT", "1")
	req.Header.Set("edu-script-token", csrfkey)
	req.Header.Set("Origin", "https://www.icourse163.org")
	req.Header.Set("Referer", "https://www.icourse163.org/home.htm?userId="+uid)
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("TE", "trailers")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:92.0) Gecko/20100101 Firefox/92.0")
	resp, err := client.Do(req)
	tools.Check(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	tools.Check(err)
	var leslist LessonList
	err = json.Unmarshal(body, &leslist)
	tools.Check(err)
	return leslist
}

func GetLessonInfo(cookie string, termId string, lessonId string) Lesson {
	client := &http.Client{}
	csrfkey := string(regexp.MustCompile(`NTESSTUDYSI=[0-9a-z]{1,}`).Find([]byte(cookie)))[12:]
	req, err := http.NewRequest("POST", "https://www.icourse163.org/web/j/courseBean.getLastLearnedMocTermDto.rpc?csrfKey="+csrfkey, strings.NewReader("termId="+termId))
	tools.Check(err)
	req.Header.Set("Host", "www.icourse163.org")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:92.0) Gecko/20100101 Firefox/92.0")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,en-US;q=0.7,ja;q=0.3")
	req.Header.Set("edu-script-token", csrfkey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Length", "17")
	req.Header.Set("Origin", "https://www.icourse163.org")
	req.Header.Set("DNT", "1")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", "https://www.icourse163.org/learn/NEU-"+lessonId+"?tid="+termId)
	req.Header.Set("Cookie", cookie)
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("TE", "trailers")
	resp, err := client.Do(req)
	tools.Check(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	tools.Check(err)
	var les Lesson
	err = json.Unmarshal(body, &les)
	tools.Check(err)
	return les
}

func MOOCtime2China(input int64) time.Time {
	return tools.Time2China(time.Unix(input/1000, 0))
}

func MOOCLoginQRcode() logonByQRCode {
	response, err := http.Get("https://www.icourse163.org/logonByQRCode/code.do?width=128&height=128")
	tools.Check(err)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	tools.Check(err)
	var lbq logonByQRCode
	err = json.Unmarshal(body, &lbq)
	tools.Check(err)
	return lbq
}

func CheckingQRcode(Pollkey string) (string, bool, []*http.Cookie) {
	var pk pollkey
	for {
		response, err := http.Get("https://www.icourse163.org/logonByQRCode/poll.do?pollKey=" + Pollkey)
		tools.Check(err)
		body, err := ioutil.ReadAll(response.Body)
		tools.Check(err)
		err = json.Unmarshal(body, &pk)
		tools.Check(err)
		if pk.Result.CodeStatus == 2 {
			return pk.Result.Token, true, response.Cookies()
		} else if pk.Result.CodeStatus == 3 {
			return "", false, nil
		}
	}
}

func MocMobChangeCookie(Token string, STUDY_WTR string) ([]*http.Cookie, string) {
	client := &http.Client{}
	url := "https://www.icourse163.org/passport/logingate/mocMobChangeCookie.htm?returnUrl=aHR0cHM6Ly93d3cuaWNvdXJzZTE2My5vcmcvaW5kZXguaHRt&token=" + Token
	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	tools.Check(err)
	req.Header.Set("Host", "www.icourse163.org")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:92.0) Gecko/20100101 Firefox/92.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "zh-CN,en-US;q=0.7,ja;q=0.3")
	req.Header.Set("Cookie", `STUDY_WTR="`+STUDY_WTR+`"`)
	resp, err := client.Do(req)
	tools.Check(err)
	return resp.Cookies(), url
}

func GetNeteaseWDAuid(cookie string, ref string) []*http.Cookie {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.icourse163.org/", strings.NewReader(""))
	tools.Check(err)
	req.Header.Set("Host", "www.icourse163.org")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:92.0) Gecko/20100101 Firefox/92.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "zh-CN,en-US;q=0.7,ja;q=0.3")
	req.Header.Set("Referer", ref)
	req.Header.Set("Cookie", cookie)
	resp, err := client.Do(req)
	tools.Check(err)
	return resp.Cookies()
}
