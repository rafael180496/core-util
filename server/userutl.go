package server

import (
	"fmt"
	"strings"
)

func getPlatform(comment []string) string {
	if len(comment) > 0 {
		if comment[0] != "compatible" {
			if strings.HasPrefix(comment[0], "Windows") {
				return "Windows"
			} else if strings.HasPrefix(comment[0], "Symbian") {
				return "Symbian"
			} else if strings.HasPrefix(comment[0], "webOS") {
				return "webOS"
			} else if comment[0] == "BB10" {
				return "BlackBerry"
			}
			return comment[0]
		}
	}
	return ""
}
func readUntil(ua string, index *int, delimiter byte, cat bool) []byte {
	var buffer []byte

	i := *index
	catalan := 0
	for ; i < len(ua); i = i + 1 {
		if ua[i] == delimiter {
			if catalan == 0 {
				*index = i + 1
				return buffer
			}
			catalan--
		} else if cat && ua[i] == '(' {
			catalan++
		}
		buffer = append(buffer, ua[i])
	}
	*index = i + 1
	return buffer
}
func parseProduct(product []byte) (string, string) {
	prod := strings.SplitN(string(product), "/", 2)
	if len(prod) == 2 {
		return prod[0], prod[1]
	}
	return string(product), ""
}
func parseSection(ua string, index *int) (s section) {
	var buffer []byte
	if *index < len(ua) && ua[*index] != '(' && ua[*index] != '[' {
		buffer = readUntil(ua, index, ' ', false)
		s.name, s.version = parseProduct(buffer)
	}

	if *index < len(ua) && ua[*index] == '(' {
		*index++
		buffer = readUntil(ua, index, ')', true)
		s.comment = strings.Split(string(buffer), "; ")
		*index++
	}
	if *index < len(ua) && ua[*index] == '[' {
		*index++
		buffer = readUntil(ua, index, ']', true)
		*index++
	}
	fmt.Println(buffer)
	return s
}
func gecko(p *UserAgent, comment []string) {
	if len(comment) > 1 {
		if comment[1] == "U" || comment[1] == "arm_64" {
			if len(comment) > 2 {
				p.OS = normalizeOS(comment[2])
			} else {
				p.OS = normalizeOS(comment[1])
			}
		} else {
			if strings.Contains(p.Platform, "Android") {
				p.Mobile = true
				p.Platform, p.OS = normalizeOS(comment[1]), p.Platform
			} else if comment[0] == "Mobile" || comment[0] == "Tablet" {
				p.Mobile = true
				p.OS = "FirefoxOS"
			} else {
				if p.OS == "" {
					p.OS = normalizeOS(comment[1])
				}
			}
		}
		if len(comment) > 3 && !strings.HasPrefix(comment[3], "rv:") {
			p.Localization = comment[3]
		}
	}
}
func dalvik(p *UserAgent, comment []string) {
	slen := len(comment)

	if strings.HasPrefix(comment[0], "Linux") {
		p.Platform = comment[0]
		if slen > 2 {
			p.OS = comment[2]
		}
		p.Mobile = true
	}
}
func (p *UserAgent) setSimple(name, version string, bot bool) {
	p.Bot = bot
	if !bot {
		p.Mozilla = ""
	}
	p.Browser.Name = name
	p.Browser.Version = version
	p.Browser.Engine = ""
	p.Browser.EngineVersion = ""
	p.OS = ""
	p.Localization = ""
}
func getFromSite(comment []string) string {
	if len(comment) == 0 {
		return ""
	}
	idx := 2
	if len(comment) < 3 {
		idx = 0
	} else if len(comment) == 4 {
		idx = 3
	}
	results := botFromSiteRegexp.FindStringSubmatch(comment[idx])
	if len(results) == 1 {
		if idx == 0 {
			return results[0]
		}
		return strings.TrimSpace(comment[idx-1])
	}
	return ""
}
func (p *UserAgent) checkBot(sections []section) {
	if len(sections) == 1 && sections[0].name != "Mozilla" {
		p.Mozilla = ""
		if botRegex.Match([]byte(sections[0].name)) {
			p.setSimple(sections[0].name, "", true)
			return
		}
		if name := getFromSite(sections[0].comment); name != "" {
			p.setSimple(sections[0].name, sections[0].version, true)
			return
		}

		p.setSimple(sections[0].name, sections[0].version, false)
	} else {
		for _, v := range sections {
			if name := getFromSite(v.comment); name != "" {
				results := strings.SplitN(name, "/", 2)
				version := ""
				if len(results) == 2 {
					version = results[1]
				}
				p.setSimple(results[0], version, true)
				return
			}
		}
		p.fixOther(sections)
	}
}
func (p *UserAgent) fixOther(sections []section) {
	if len(sections) > 0 {
		p.Browser.Name = sections[0].name
		p.Browser.Version = sections[0].version
		p.Mozilla = ""
	}
}
func trident(p *UserAgent, comment []string) {
	p.Platform = "Windows"
	if p.OS == "" {
		if len(comment) > 2 {
			p.OS = normalizeOS(comment[2])
		} else {
			p.OS = "Windows NT 4.0"
		}
	}
	for _, v := range comment {
		if strings.HasPrefix(v, "IEMobile") {
			p.Mobile = true
			return
		}
	}
}
func opera(p *UserAgent, comment []string) {
	slen := len(comment)

	if strings.HasPrefix(comment[0], "Windows") {
		p.Platform = "Windows"
		p.OS = normalizeOS(comment[0])
		if slen > 2 {
			if slen > 3 && strings.HasPrefix(comment[2], "MRA") {
				p.Localization = comment[3]
			} else {
				p.Localization = comment[2]
			}
		}
	} else {
		if strings.HasPrefix(comment[0], "Android") {
			p.Mobile = true
		}
		p.Platform = comment[0]
		if slen > 1 {
			p.OS = comment[1]
			if slen > 3 {
				p.Localization = comment[3]
			}
		} else {
			p.OS = comment[0]
		}
	}
}
func normalizeOS(name string) string {
	sp := strings.SplitN(name, " ", 3)
	if len(sp) != 3 || sp[1] != "NT" {
		return name
	}
	if _, ok := osWindow[sp[2]]; ok {
		return osWindow[sp[2]]
	}
	return name
}
func (p *UserAgent) iMessagePreview() bool {
	if !strings.Contains(p.UA, "facebookexternalhit") || !strings.Contains(p.UA, "Twitterbot") {
		return false
	}

	p.Bot = true
	p.Browser.Name = "iMessage-Preview"
	p.Browser.Engine = ""
	p.Browser.EngineVersion = ""
	return true
}

func (p *UserAgent) googleOrBingBot() bool {
	if !strings.Contains(p.UA, "Google") || !strings.Contains(p.UA, "bingbot") {
		p.Platform = ""
		p.Undecided = true
	}
	return p.Undecided
}

func webkit(p *UserAgent, comment []string) {
	if p.Platform == "webOS" {
		p.Browser.Name = p.Platform
		p.OS = "Palm"
		if len(comment) > 2 {
			p.Localization = comment[2]
		}
		p.Mobile = true
	} else if p.Platform == "Symbian" {
		p.Mobile = true
		p.Browser.Name = p.Platform
		p.OS = comment[0]
	} else if p.Platform == "Linux" {
		p.Mobile = true
		if p.Browser.Name == "Safari" {
			p.Browser.Name = "Android"
		}
		if len(comment) > 1 {
			if comment[1] == "U" || comment[1] == "arm_64" {
				if len(comment) > 2 {
					p.OS = comment[2]
				} else {
					p.Mobile = false
					p.OS = comment[0]
				}
			} else {
				p.OS = comment[1]
			}
		}
		if len(comment) > 3 {
			p.Localization = comment[3]
		} else if len(comment) == 3 {
			_ = p.googleOrBingBot()
		}
	} else if len(comment) > 0 {
		if len(comment) > 3 {
			p.Localization = comment[3]
		}
		if strings.HasPrefix(comment[0], "Windows NT") {
			p.OS = normalizeOS(comment[0])
		} else if len(comment) < 2 {
			p.Localization = comment[0]
		} else if len(comment) < 3 {
			if !p.googleOrBingBot() && !p.iMessagePreview() {
				p.OS = normalizeOS(comment[1])
			}
		} else {
			p.OS = normalizeOS(comment[2])
		}
		if p.Platform == "BlackBerry" {
			p.Browser.Name = p.Platform
			if p.OS == "Touch" {
				p.OS = p.Platform
			}
		}
	}
	if p.Platform == "Macintosh" && p.Browser.Engine == "AppleWebKit" && p.Browser.Name == "Firefox" {
		p.Platform = "iPad"
		p.Mobile = true
	}
}
func (p *UserAgent) detectOS(s section) {
	if s.name == "Mozilla" {
		p.Platform = getPlatform(s.comment)
		if p.Platform == "Windows" && len(s.comment) > 0 {
			p.OS = normalizeOS(s.comment[0])
		}
		switch p.Browser.Engine {
		case "":
			p.Undecided = true
		case "Gecko":
			gecko(p, s.comment)
		case "AppleWebKit":
			webkit(p, s.comment)
		case "Trident":
			trident(p, s.comment)
		}
	} else if s.name == "Opera" {
		if len(s.comment) > 0 {
			opera(p, s.comment)
		}
	} else if s.name == "Dalvik" {
		if len(s.comment) > 0 {
			dalvik(p, s.comment)
		}
	} else {
		p.Undecided = true
	}
}
