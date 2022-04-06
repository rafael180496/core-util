package server

import (
	"strings"
)

/*detectBrowser : detecta el navegador que se esta ocupando*/
func (p *UserAgent) detectBrowser(sections []section) {
	slen := len(sections)

	switch {
	case sections[0].name == "Opera":
		p.Browser.Name = "Opera"
		p.Browser.Version = sections[0].version
		p.Browser.Engine = "Presto"
		if slen > 1 {
			p.Browser.EngineVersion = sections[1].version
		}
		return
	case sections[0].name == "Dalvik":
		p.Mozilla = "5.0"
		return
	case slen > 1:
		engine := sections[1]
		p.Browser.Engine = engine.name
		p.Browser.EngineVersion = engine.version
		if slen > 2 {
			sectionIndex := 2
			if sections[2].version == "" && slen > 3 {
				sectionIndex = 3
			}
			p.Browser.Version = sections[sectionIndex].version
			if engine.name == "AppleWebKit" {
				for _, comment := range engine.comment {
					if len(comment) > 5 &&
						(strings.HasPrefix(comment, "Googlebot") || strings.HasPrefix(comment, "bingbot")) {
						p.Undecided = true
						break
					}
				}
				switch sections[slen-1].name {
				case "Edge":
					p.Browser.Name = "Edge"
					p.Browser.Version = sections[slen-1].version
					p.Browser.Engine = "EdgeHTML"
					p.Browser.EngineVersion = ""
				case "Edg":
					if !p.Undecided {
						p.Browser.Name = "Edge"
						p.Browser.Version = sections[slen-1].version
						p.Browser.Engine = "AppleWebKit"
						p.Browser.EngineVersion = sections[slen-2].version
					}
				case "OPR":
					p.Browser.Name = "Opera"
					p.Browser.Version = sections[slen-1].version
				default:
					switch sections[slen-3].name {
					case "YaBrowser":
						p.Browser.Name = "YaBrowser"
						p.Browser.Version = sections[slen-3].version
					default:
						switch sections[slen-2].name {
						case "Electron":
							p.Browser.Name = "Electron"
							p.Browser.Version = sections[slen-2].version
						default:
							_, ok := sectionBrow[sections[sectionIndex].name]
							if ok {
								p.Browser.Name = sectionBrow[sections[sectionIndex].name]
							} else {
								p.Browser.Name = sectionBrow["Safari"]
							}
						}
					}
					for _, comment := range engine.comment {
						if len(comment) > 5 &&
							(strings.HasPrefix(comment, "Googlebot") || strings.HasPrefix(comment, "bingbot")) {
							p.Undecided = true
							break
						}
					}
				}
			} else if engine.name == "Gecko" {
				name := sections[2].name
				if name == "MRA" && slen > 4 {
					name = sections[4].name
					p.Browser.Version = sections[4].version
				}
				p.Browser.Name = name
			} else if engine.name == "like" && sections[2].name == "Gecko" {
				p.Browser.Engine = "Trident"
				p.Browser.Name = "Internet Explorer"
				for _, c := range sections[0].comment {
					version := ie11Regexp.FindStringSubmatch(c)
					if len(version) > 0 {
						p.Browser.Version = version[1]
						return
					}
				}
				p.Browser.Version = ""
			}
		}
		return
	case slen == 1 && len(sections[0].comment) > 1:
		comment := sections[0].comment
		if comment[0] == "compatible" && strings.HasPrefix(comment[1], "MSIE") {
			p.Browser.Engine = "Trident"
			p.Browser.Name = "Internet Explorer"
			for _, v := range comment {
				if strings.HasPrefix(v, "Trident/") {
					switch v[8:] {
					case "4.0":
						p.Browser.Version = "8.0"
					case "5.0":
						p.Browser.Version = "9.0"
					case "6.0":
						p.Browser.Version = "10.0"
					}
					break
				}
			}
			if p.Browser.Version == "" {
				p.Browser.Version = strings.TrimSpace(comment[1][4:])
			}
		}
		return
	default:
		return
	}
}

/*New : crea un useragent*/
func New(ua string) *UserAgent {
	o := &UserAgent{}
	o.Parse(ua)
	return o
}

/*Parse : convierte el header user-agent*/
func (p *UserAgent) Parse(ua string) {
	var sections []section
	p.UA = ua
	for index, limit := 0, len(ua); index < limit; {
		s := parseSection(ua, &index)
		if !p.Mobile && s.name == "Mobile" {
			p.Mobile = true
		}
		sections = append(sections, s)
	}

	if len(sections) > 0 {
		if sections[0].name == "Mozilla" {
			p.Mozilla = sections[0].version
		}
		p.detectBrowser(sections)
		p.detectOS(sections[0])

		if p.Undecided {
			p.checkBot(sections)
		}
	}
}
