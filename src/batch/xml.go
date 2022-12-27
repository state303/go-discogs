package batch

import (
	"github.com/reactivex/rxgo/v2"
	"github.com/state303/go-discogs/model"
	"github.com/state303/go-discogs/src/cache"
	"github.com/state303/go-discogs/src/dateparser"
	"github.com/state303/go-discogs/src/helper"
	"github.com/state303/go-discogs/src/unique"
	"strconv"
	"strings"
)

type XmlRef struct {
	ID   int32  `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type XmlArtist struct {
	ID          int32   `xml:"id"`
	Name        *string `xml:"name"`
	DataQuality *string `xml:"data_quality"`
	Profile     *string `xml:"profile"`
	RealName    *string `xml:"realname"`
}

func (a *XmlArtist) Transform() rxgo.Observable {
	return rxgo.Just(&model.Artist{
		ID:          a.ID,
		DataQuality: a.DataQuality,
		Name:        a.Name,
		Profile:     a.Profile,
		RealName:    a.RealName,
	})()
}

type XmlArtistRelation struct {
	ID       int32    `xml:"id" gorm:"column:id"`
	Urls     []string `xml:"urls>url"`
	NameVars []string `xml:"namevariations>name"`
	Aliases  []XmlRef `xml:"aliases>name"`
	Groups   []XmlRef `xml:"groups>name"`
}

func (a *XmlArtistRelation) GetUrls() []*model.ArtistURL {
	slice := make([]*model.ArtistURL, 0)
	for _, url := range a.Urls {
		slice = append(slice, &model.ArtistURL{ArtistID: a.ID, URLHash: int64(helper.Fnv32Str(url)), URL: url})
	}
	return unique.Slice(slice)
}

func (a *XmlArtistRelation) GetNameVars() []*model.ArtistNameVariation {
	slice := make([]*model.ArtistNameVariation, 0)
	for _, nameVar := range a.NameVars {
		slice = append(slice, &model.ArtistNameVariation{ArtistID: a.ID, NameVariation: nameVar, NameVariationHash: int64(helper.Fnv32Str(nameVar))})
	}
	return unique.Slice(slice)
}

func (a *XmlArtistRelation) GetAliases() []*model.ArtistAlias {
	slice := make([]*model.ArtistAlias, 0)
	for _, alias := range a.Aliases {
		if _, ok := cache.ArtistIDCache.Load(alias.ID); ok {
			slice = append(slice, &model.ArtistAlias{ArtistID: a.ID, AliasID: alias.ID})
		}
	}
	return unique.Slice(slice)
}

func (a *XmlArtistRelation) GetGroups() []*model.ArtistGroup {
	slice := make([]*model.ArtistGroup, 0)
	for _, group := range a.Groups {
		if _, ok := cache.ArtistIDCache.Load(group.ID); ok {
			slice = append(slice, &model.ArtistGroup{ArtistID: a.ID, GroupID: group.ID})
		}
	}
	return unique.Slice(slice)
}

type XmlLabel struct {
	ID          int32   `xml:"id"`
	Name        *string `xml:"name"`
	ContactInfo *string `xml:"contactinfo"`
	Profile     *string `xml:"profile"`
	DataQuality *string `xml:"data_quality"`
}

func (l *XmlLabel) Transform() rxgo.Observable {
	return rxgo.Just(&model.Label{
		ID:          l.ID,
		Name:        l.Name,
		ContactInfo: l.ContactInfo,
		Profile:     l.Profile,
		DataQuality: l.DataQuality,
	})()
}

type XmlLabelRelation struct {
	ID          int32    `xml:"id"`
	Urls        []string `xml:"urls>url"`
	ParentLabel *XmlRef  `xml:"parentLabel"`
}

func (l *XmlLabelRelation) GetUrls() []*model.LabelURL {
	r := make([]*model.LabelURL, 0)
	for _, url := range l.Urls {
		r = append(r, &model.LabelURL{
			LabelID: l.ID,
			URLHash: int64(helper.Fnv32Str(url)),
			URL:     url,
		})
	}
	return unique.Slice(r)
}

func (l *XmlLabelRelation) GetParentID() *int32 {
	if l.ParentLabel == nil {
		return nil
	}
	return &l.ParentLabel.ID
}

type XmlMaster struct {
	ID          int32   `xml:"id,attr"`
	Title       *string `xml:"title"`
	DataQuality *string `xml:"data_quality"`
	Year        *int16  `xml:"year"`
}

type XmlGenreStyle struct {
	Styles []string `xml:"styles>style"`
	Genres []string `xml:"genres>genre"`
}

func (m *XmlMaster) Transform() rxgo.Observable {
	return rxgo.Just(&model.Master{
		ID:           m.ID,
		Title:        m.Title,
		DataQuality:  m.DataQuality,
		ReleasedYear: m.Year,
	})()
}

type XmlMasterRelation struct {
	ID          int32      `xml:"id,attr"`
	Title       *string    `xml:"title"`
	DataQuality *string    `xml:"data_quality"`
	Year        *int16     `xml:"year"`
	Styles      []string   `xml:"styles>style"`
	Genres      []string   `xml:"genres>genre"`
	Artists     []int32    `xml:"artists>artist>id"`
	Videos      []XmlVideo `xml:"videos>video"`
}

type XmlVideo struct {
	URL         string  `xml:"src,attr"`
	Title       *string `xml:"title"`
	Description *string `xml:"description"`
}

func (m *XmlMasterRelation) GetStyles() []*model.Style {
	s := make([]*model.Style, 0)
	for _, style := range unique.Slice(m.Styles) {
		s = append(s, &model.Style{Name: style})
	}
	return s
}

func (m *XmlMasterRelation) GetGenres() []*model.Genre {
	g := make([]*model.Genre, 0)
	for _, genre := range unique.Slice(m.Genres) {
		g = append(g, &model.Genre{Name: genre})
	}
	return g
}

func (m *XmlMasterRelation) GetMaster() *model.Master {
	return &model.Master{
		ID:           m.ID,
		Title:        m.Title,
		DataQuality:  m.DataQuality,
		ReleasedYear: m.Year,
	}
}

func (m *XmlMasterRelation) GetMasterStyles() []*model.MasterStyle {
	filteredMasterStyles := make([]*model.MasterStyle, 0)
	for _, style := range m.Styles {
		if id, ok := cache.StyleCache.Load(style); ok {
			ms := &model.MasterStyle{MasterID: m.ID, StyleID: id.(int32)}
			filteredMasterStyles = append(filteredMasterStyles, ms)
		}
	}
	return unique.Slice(filteredMasterStyles)
}

func (m *XmlMasterRelation) GetMasterGenres() []*model.MasterGenre {
	filteredMasterGenres := make([]*model.MasterGenre, 0)
	for _, genre := range m.Genres {
		if id, ok := cache.GenreCache.Load(genre); ok {
			mg := &model.MasterGenre{MasterID: m.ID, GenreID: id.(int32)}
			filteredMasterGenres = append(filteredMasterGenres, mg)
		}
	}
	return unique.Slice(filteredMasterGenres)
}

func (m *XmlMasterRelation) GetMasterVideos() []*model.MasterVideo {
	items := make([]*model.MasterVideo, 0)
	for _, vid := range m.Videos {
		items = append(items, &model.MasterVideo{
			MasterID:    m.ID,
			URLHash:     int64(helper.Fnv32Str(vid.URL)),
			URL:         vid.URL,
			Description: vid.Description,
			Title:       vid.Title,
		})
	}
	return unique.Slice(items)
}

func (m *XmlMasterRelation) GetMasterArtists() []*model.MasterArtist {
	items := make([]*model.MasterArtist, 0)
	for _, id := range m.Artists {
		if _, ok := cache.ArtistIDCache.Load(id); ok {
			items = append(items, &model.MasterArtist{
				ArtistID: id,
				MasterID: m.ID,
			})
		}
	}
	return unique.Slice(items)
}

type XmlRelease struct {
	ID                int32                `xml:"id,attr"`
	Title             *string              `xml:"title"`
	Country           *string              `xml:"country"`
	DataQuality       *string              `xml:"data_quality"`
	ListedReleaseDate *string              `xml:"released"`
	Notes             *string              `xml:"notes"`
	MasterInfo        XmlReleaseMasterInfo `xml:"master_id"`
	Status            *string              `xml:"status,attr"`
}

type XmlReleaseMasterInfo struct {
	MasterID *int32 `xml:",cdata"`
	IsMaster bool   `xml:"is_main_release,attr"`
}

func (m *XmlRelease) Transform() rxgo.Observable {
	var year, month, day *int16
	if ymd := m.ListedReleaseDate; ymd != nil {
		year, month, day = dateparser.ParseYMD(*ymd)
	}

	var masterID *int32

	if m.MasterInfo.MasterID != nil {
		if _, ok := cache.MasterIDCache.Load(*m.MasterInfo.MasterID); ok {
			masterID = m.MasterInfo.MasterID
		}
	}

	return rxgo.Just(&model.Release{
		ID:                m.ID,
		Title:             m.Title,
		Country:           m.Country,
		DataQuality:       m.DataQuality,
		ReleasedYear:      year,
		ReleasedMonth:     month,
		ReleasedDay:       day,
		ListedReleaseDate: m.ListedReleaseDate,
		IsMaster:          &m.MasterInfo.IsMaster,
		MasterID:          masterID,
		Notes:             m.Notes,
		Status:            m.Status,
	})()
}

type XmlLabelRelease struct {
	LabelID          int32  `xml:"id,attr"`
	CategoryNotation string `xml:"catno,attr"`
}

type XmlCreditedArtist struct {
	ArtistID int32  `xml:"id"`
	Role     string `xml:"role"`
}

type XmlFormat struct {
	Name         *string  `xml:"name,attr"`
	Quantity     *int32   `xml:"qty,attr"`
	Text         *string  `xml:"text,attr"`
	Descriptions []string `xml:"descriptions>description"`
}

type XmlTrack struct {
	Position string `xml:"position"`
	Title    string `xml:"title"`
	Duration string `xml:"duration"`
}

type XmlIdentifier struct {
	Typ   string `xml:"type,attr"`
	Desc  string `xml:"description,attr"`
	Value string `xml:"value,attr"`
}

type XmlContract struct {
	ResourceUrl string `xml:"resource_url"`
	Content     string `xml:"entity_type_name"`
}

type XmlReleaseRelation struct {
	ID                int32                `xml:"id,attr"`
	Title             *string              `xml:"title"`
	Country           *string              `xml:"country"`
	DataQuality       *string              `xml:"data_quality"`
	ListedReleaseDate *string              `xml:"released"`
	Notes             *string              `xml:"notes"`
	MasterInfo        XmlReleaseMasterInfo `xml:"master_id"`
	Status            *string              `xml:"status,attr"`
	Artists           []int32              `xml:"artists>artist>id"`
	Labels            []XmlLabelRelease    `xml:"labels>label"`
	CreditedArtists   []XmlCreditedArtist  `xml:"extraartists>artist"`
	Formats           []XmlFormat          `xml:"formats>format"`
	Genres            []string             `xml:"genres>genre"`
	Styles            []string             `xml:"styles>style"`
	Tracks            []XmlTrack           `xml:"tracklist>track"`
	Identifiers       []XmlIdentifier      `xml:"identifiers>identifier"`
	Videos            []XmlVideo           `xml:"videos>video"`
	Contracts         []XmlContract        `xml:"companies>company"`
}

func (r *XmlReleaseRelation) GetGenres() []*model.Genre {
	genres := make([]*model.Genre, 0)
	for _, v := range unique.Slice(r.Genres) {
		genres = append(genres, &model.Genre{Name: v})
	}
	return genres
}

func (r *XmlReleaseRelation) GetStyles() []*model.Style {
	styles := make([]*model.Style, 0)
	for _, v := range unique.Slice(r.Styles) {
		styles = append(styles, &model.Style{Name: v})
	}
	return styles
}

func (r *XmlReleaseRelation) GetRelease() *model.Release {
	var year, month, day *int16
	if ymd := r.ListedReleaseDate; ymd != nil {
		year, month, day = dateparser.ParseYMD(*ymd)
	}
	var masterID *int32
	if r.MasterInfo.MasterID != nil {
		if _, ok := cache.MasterIDCache.Load(*r.MasterInfo.MasterID); ok {
			masterID = r.MasterInfo.MasterID
		}
	}
	return &model.Release{
		ID:                r.ID,
		Title:             r.Title,
		Country:           r.Country,
		DataQuality:       r.DataQuality,
		ReleasedYear:      year,
		ReleasedMonth:     month,
		ReleasedDay:       day,
		ListedReleaseDate: r.ListedReleaseDate,
		IsMaster:          &r.MasterInfo.IsMaster,
		MasterID:          masterID,
		Notes:             r.Notes,
		Status:            r.Status,
	}
}

func (r *XmlReleaseRelation) GetContracts() []*model.ReleaseContract {
	items := make([]*model.ReleaseContract, 0)
	for _, rc := range r.Contracts {
		labelID, err := strconv.Atoi(helper.GetLastUriSegment(rc.ResourceUrl))
		if err != nil {
			continue
		}

		lid32 := int32(labelID)
		if _, ok := cache.LabelIDCache.Load(lid32); !ok {
			continue
		}
		items = append(items, &model.ReleaseContract{
			ReleaseID:    r.ID,
			LabelID:      int32(labelID),
			ContractHash: int64(helper.Fnv32Str(rc.Content)),
			Contract:     rc.Content,
		})
	}
	return unique.Slice(items)
}

func (r *XmlReleaseRelation) GetVideos() []*model.ReleaseVideo {
	items := make([]*model.ReleaseVideo, 0)
	for _, vid := range r.Videos {
		items = append(items, &model.ReleaseVideo{
			ReleaseID:   r.ID,
			Description: vid.Description,
			Title:       vid.Title,
			URL:         vid.URL,
			URLHash:     int64(helper.Fnv32Str(vid.URL)),
		})
	}
	return unique.Slice(items)
}

func (r *XmlReleaseRelation) GetIdentifiers() []*model.ReleaseIdentifier {
	items := make([]*model.ReleaseIdentifier, 0)
	for _, identifier := range r.Identifiers {
		items = append(items, &model.ReleaseIdentifier{
			ReleaseID:      r.ID,
			Description:    &identifier.Desc,
			Type:           &identifier.Typ,
			Value:          &identifier.Value,
			IdentifierHash: int64(helper.Fnv32Str(identifier.Desc + identifier.Typ + identifier.Value)),
		})
	}
	return unique.Slice(items)
}

func (r *XmlReleaseRelation) GetTracks() []*model.ReleaseTrack {
	items := make([]*model.ReleaseTrack, 0)
	for _, track := range r.Tracks {
		items = append(items, &model.ReleaseTrack{
			ReleaseID: r.ID,
			Duration:  &track.Duration,
			Position:  &track.Position,
			Title:     &track.Title,
			TitleHash: int64(helper.Fnv32Str(track.Title)),
		})
	}
	return unique.Slice(items)
}

func (r *XmlReleaseRelation) GetFormats() []*model.ReleaseFormat {
	items := make([]*model.ReleaseFormat, 0)
	for _, format := range r.Formats {
		desc := strings.Join(format.Descriptions, ",")
		hashSrc := desc
		if format.Name != nil {
			hashSrc += *format.Name
		}
		if format.Quantity != nil {
			hashSrc += string(*format.Quantity)
		}
		if format.Text != nil {
			hashSrc += *format.Text
		}
		items = append(items, &model.ReleaseFormat{
			ReleaseID:   r.ID,
			Description: &desc,
			Name:        format.Name,
			Quantity:    format.Quantity,
			Text:        format.Text,
			FormatHash:  int64(helper.Fnv32Str(hashSrc)),
		})
	}
	return unique.Slice(items)
}

func (r *XmlReleaseRelation) GetCreditedArtists() []*model.ReleaseCreditedArtist {
	items := make([]*model.ReleaseCreditedArtist, 0)
	for _, ca := range r.CreditedArtists {
		if _, ok := cache.ArtistIDCache.Load(ca.ArtistID); ok && len(ca.Role) > 0 {
			parts := strings.Split(ca.Role, ",")
			for i := range parts {
				parts[i] = strings.TrimSpace(parts[i])
			}
			ca.Role = strings.Join(parts, ",")
			items = append(items, &model.ReleaseCreditedArtist{
				ReleaseID: r.ID,
				ArtistID:  ca.ArtistID,
				RoleHash:  int64(helper.Fnv32Str(ca.Role)),
				Role:      &ca.Role,
			})
		}
	}
	return unique.Slice(items)
}

func (r *XmlReleaseRelation) GetReleaseArtists() []*model.ReleaseArtist {
	items := make([]*model.ReleaseArtist, 0)
	for _, artistID := range r.Artists {
		if _, ok := cache.ArtistIDCache.Load(artistID); ok {
			items = append(items, &model.ReleaseArtist{
				ReleaseID: r.ID,
				ArtistID:  artistID,
			})
		}
	}
	return unique.Slice(items)
}

func (r *XmlReleaseRelation) GetLabels() []*model.LabelRelease {
	items := make([]*model.LabelRelease, 0)
	for _, label := range r.Labels {
		if _, ok := cache.LabelIDCache.Load(label.LabelID); ok {
			items = append(items, &model.LabelRelease{
				LabelID:          label.LabelID,
				ReleaseID:        r.ID,
				CategoryNotation: &label.CategoryNotation,
			})
		}
	}
	return unique.Slice(items)
}

func (r *XmlReleaseRelation) GetReleaseStyles() []*model.ReleaseStyle {
	filteredReleaseStyles := make([]*model.ReleaseStyle, 0)
	for _, style := range r.Styles {
		if styleID, ok := cache.StyleCache.Load(style); ok {
			rs := &model.ReleaseStyle{ReleaseID: r.ID, StyleID: styleID.(int32)}
			filteredReleaseStyles = append(filteredReleaseStyles, rs)
		}
	}
	return unique.Slice(filteredReleaseStyles)
}

func (r *XmlReleaseRelation) GetReleaseGenres() []*model.ReleaseGenre {
	filteredReleaseGenres := make([]*model.ReleaseGenre, 0)
	for _, genre := range r.Genres {
		if genreID, ok := cache.GenreCache.Load(genre); ok {
			rg := &model.ReleaseGenre{ReleaseID: r.ID, GenreID: genreID.(int32)}
			filteredReleaseGenres = append(filteredReleaseGenres, rg)
		}
	}
	return unique.Slice(filteredReleaseGenres)
}
